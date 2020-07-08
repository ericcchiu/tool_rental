package psql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ericcchiu/tool_rental/tools"
	_ "github.com/lib/pq"
)

type toolDataStore struct {
	db *sql.DB
}

func NewPostgresToolDataStore(db *sql.DB) *toolDataStore {
	return &toolDataStore{
		db,
	}
}

// Create method takes in a Tool instance and places the new tool into the database for the toolDataStore struct
func (t toolDataStore) Create(tool *tools.Tool) (err error) {
	sqlStatement := `INSERT INTO tools(name, description, price, quantity, created_at, update_at)
					 VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	row := t.db.QueryRow(sqlStatement, tool.Name, tool.Description, tool.Price, tool.Quantity, time.Now(), time.Now())
	err = row.Scan(&tool.ID)

	if err == nil {
		fmt.Println("New record ID is:", tool.ID)
	}

	return
}

// FindByID method takes in a Tool ID and returns the associated tool, if it exists, for the toolDataStore struct.
func (t toolDataStore) FindByID(id string) (tool *tools.Tool, err error) {
	tool = new(tools.Tool)
	sqlStatement := `SELECT * FROM tools WHERE id=$1`
	row := t.db.QueryRow(sqlStatement, id)
	err = row.Scan(&tool.ID, &tool.Name, &tool.Description, &tool.Price, &tool.Quantity, &tool.Created, &tool.Updated)

	switch err {
	case sql.ErrNoRows:
		return tool, fmt.Errorf("no rows were returned")
	case nil:
		return
	default:
		return
	}
}

// FindAll method returns all the number of rows for the toolDataStore struct
func (t toolDataStore) FindAll() (allTools []*tools.Tool, err error) {
	sqlStatement := `SELECT * FROM tools`
	rows, err := t.db.Query(sqlStatement)
	defer rows.Close()
	for rows.Next() {
		currentTool := new(tools.Tool)
		err = rows.Scan(&currentTool.ID, &currentTool.Name, &currentTool.Description, &currentTool.Price, &currentTool.Quantity, &currentTool.Created, &currentTool.Updated)
		if err == nil {
			allTools = append(allTools, currentTool) //append only if there was not an error
		}
	}

	err = rows.Err()
	return
}

// Update method takes in a Tool instance and updates the tool with the corresponding ID in the database for the toolDataStore struct.
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

// Delete method takes in a Tool ID and deletes the tool with the corresponding ID in the database for the toolDataStore struct
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
