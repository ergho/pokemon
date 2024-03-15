package pokemon

import (
	"encoding/json"
	"os"
)

type FileIOHandler interface {
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

type OSFileIOHandler struct {}

func (h *OSFileIOHandler) ReadFile(filename string) ([]byte, error) {
    return os.ReadFile(filename)
}

func (h *OSFileIOHandler) WriteFile(filename string, data []byte, perm os.FileMode) error {
    return os.WriteFile(filename, data, perm)
}

type MockFileIOHandler struct {
	FileData map[string][]byte
}

func (m *MockFileIOHandler) ReadFile(filename string) ([]byte, error) {
	data, exists := m.FileData[filename]

	if !exists {
		return nil, os.ErrNotExist
	}
	return data, nil
}

func (m *MockFileIOHandler) WriteFile(filename string, data[]byte, perm os.FileMode) error {
    m.FileData[filename] = data
    return nil
}

func SaveTOJSON(handler FileIOHandler, data interface{}, filename string) error {
    jsonData, err := json.Marshal(data)
    if err != nil {
        return err
    }
    return handler.WriteFile(filename, jsonData, 0644)
}

func LoadFromJSON(handler FileIOHandler, filename string, data interface{}) error {
    jsonData, err := handler.ReadFile(filename)

    if err != nil {
        return err
    }

    return json.Unmarshal(jsonData, data)
}
