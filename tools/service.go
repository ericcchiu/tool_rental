package tools

import (
	"time"

	"github.com/google/uuid"
)

type ToolService interface {
	CreateTool(tool *Tool) error
	FindToolByID(id string) (*Tool, error)
	FindAllTools() ([]*Tool, error)
	UpdateTool(tool *Tool) error
	DeleteTool(id string) error
}

type toolService struct {
	dataStore ToolDataStore
}

func NewToolService(dataStore ToolDataStore) ToolService {
	return &toolService{
		dataStore,
	}
}

func (s *toolService) CreateTool(tool *Tool) error {
	tool.ID = uuid.New().String()
	tool.Created = time.Now()
	tool.Updated = time.Now()

	return s.dataStore.Create(tool)
}

func (s *toolService) FindToolByID(id string) (*Tool, error) {
	tool, err := s.dataStore.FindByID(id)
	return tool, err
}

func (s *toolService) FindAllTools() ([]*Tool, error) {
	tools, err := s.dataStore.FindAll()
	return tools, err
}

func (s *toolService) UpdateTool(tool *Tool) error {
	tool.Updated = time.Now()
	return s.dataStore.Update(tool)
}

func (s *toolService) DeleteTool(id string) error {
	return s.dataStore.Delete(id)
}
