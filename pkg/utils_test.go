package pokemon

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSaveToJSON(t *testing.T) {
	mockHandler := &MockFileIOHandler{FileData: make(map[string][]byte)}
	testData := []Species{{Name: "Bulbasaur"}, {Name: "Charmander"}}
	err := SaveTOJSON(mockHandler, testData, "test.json")

	if err != nil {
		t.Fatalf("SaveTOJSON() error = %v", err)
	}

	if _, exists := mockHandler.FileData["test.json"]; !exists {
		t.Errorf("SaveTOJSON() did not save data to mock file system")
	}
}

func TestLoadFromJSON(t *testing.T) {
	mockHandler := &MockFileIOHandler{FileData: make(map[string][]byte)}
	testData := []Species{{Name: "Bulbasaur"}, {Name: "Charmander"}}
    jsonData, _ := json.Marshal(testData)
    mockHandler.FileData["test.json"] = jsonData
    var loadedData []Species
    err := LoadFromJSON(mockHandler, "test.json", &loadedData)

    if err != nil {
        t.Fatalf("LoadFromJSON() error = %v", err)
    }

    if !reflect.DeepEqual(loadedData, testData) {
        t.Errorf("LoadFromJSON() loaded data = %v, want %v", loadedData, testData)
    }
}
