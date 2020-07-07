package tools

type ToolDataStore interface {
	Create(tool *Tool) error
	FindByID(id string) (*Tool, error)
	FindAll() ([]*Tool, error)
	Update(tool *Tool) error
	Delete(id string) error
}
