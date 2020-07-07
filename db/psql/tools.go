package psql
​
import (
	"database/sql"
	"fmt"
	"time"
​
	"github.homedepot.com/emc4jq2/tool_rental/tools" //replace with yourldap with your ldap
​
	_ "github.com/lib/pq" //gives access to the PSQL driver
)
​
type toolDataStore struct { //1
	db *sql.DB
}
​
func NewPostgresToolDataStore(db *sql.DB) *toolDataStore { //2
	return &toolDataStore{
		db,
	}
}
​
func (t toolDataStore) Create(tool *tools.Tool) (err error) {
	sqlStatement := `INSERT INTO tools(name, description, price, quantity, created_at, update_at)
					 VALUES ($1, $2, $3, $4, $5, $6) RETURNING id` //1
	row := t.db.QueryRow(sqlStatement, tool.Name, tool.Description, tool.Price, tool.Quantity, time.Now(), time.Now()) //2
	err = row.Scan(&tool.ID)
	if err == nil {
		fmt.Println("New record ID is:", tool.ID) //3
	}
​
	return
}
​
func (t toolDataStore) FindByID(id string) (tool *tools.Tool, err error) {
	tool = new(tools.Tool)
	sqlStatement := `SELECT * FROM tools WHERE id=$1`
	row := t.db.QueryRow(sqlStatement, id)                                                                             //1
	err = row.Scan(&tool.ID, &tool.Name, &tool.Description, &tool.Price, &tool.Quantity, &tool.Created, &tool.Updated) //2
​
	switch err {
	case sql.ErrNoRows: //3
		return tool, fmt.Errorf("no rows were returned")
	case nil:
		return
	default:
		return
	}
}
​
func (t toolDataStore) FindAll() (allTools []*tools.Tool, err error) {
	sqlStatement := `SELECT * FROM tools`
	rows, err := t.db.Query(sqlStatement) //1
	defer rows.Close()                    //2
	for rows.Next() {                     //3
		currentTool := new(tools.Tool)
		err = rows.Scan(&currentTool.ID, &currentTool.Name, &currentTool.Description, &currentTool.Price, &currentTool.Quantity, &currentTool.Created, &currentTool.Updated)
		if err == nil {
			allTools = append(allTools, currentTool) //append only if there was not an error
		}
	}
​
	err = rows.Err() //4
	return
}
​
func (t toolDataStore) Update(tool *tools.Tool) (err error) {
	_, err = t.FindByID(tool.ID) //Checks to see if ID is provided/valid
	if err != nil {
		return
	}
	sqlStatement := `UPDATE tools
					SET name = $2, description = $3, price = $4, quantity = $5, update_at = $6
					WHERE id = $1;`
	result, err := t.db.Exec(sqlStatement, tool.ID, tool.Name, tool.Description, tool.Price, tool.Quantity, time.Now())
	fmt.Printf("%v row affected\n", result)
	return err
}
​
func (t toolDataStore) Delete(id string) (err error) {
	_, err = t.FindByID(id)
	if err != nil {
		return fmt.Errorf("delete: id not found")
	}
	sqlStatement := `DELETE FROM tools
					WHERE id = $1;`
	result, err := t.db.Exec(sqlStatement, id)
	numRows, _ := result.RowsAffected()
	fmt.Printf("%v rows deleted", numRows)
	return err
}




