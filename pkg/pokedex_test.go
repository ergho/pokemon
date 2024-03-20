package pokemon

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func createTestPokedex() *Pokedex {
	species := []Species{
		{ID: 4, Name: "Charmander", Types: []Type{Fire}},
		{ID: 26, Name: "Raichu", Types: []Type{Electric}},
	}

	pokedex := NewPokedex(species)
	return pokedex
}

func TestSaveAndLoadPokedex(t *testing.T) {
	originalPokdex := createTestPokedex()

	filename := "test_pokedex.json"

	if err := originalPokdex.SavePokedexToJSON(filename); err != nil {
		t.Fatalf("Failed to save Pokedex to json: %v", err)
	}

	loadedPokedex, err := LoadPokedexFromJSON(filename)

	if err != nil {
		t.Fatalf("Failed to load Pokedex from json: %v", err)
	}

	if !reflect.DeepEqual(originalPokdex, loadedPokedex) {
		fmt.Printf("%v, %v", originalPokdex, loadedPokedex)
		t.Errorf("Original and loaded pokdex are not the same")
	}
}

func TestAddNewSpecies(t *testing.T) {
	pokedex := createTestPokedex()

	newSpecies := Species{ID: 152, Name: "Chikorita", Types: []Type{Grass}}

	err := pokedex.AddNewSpecies(newSpecies)

	if err != nil {
		t.Fatalf("Failed to add new species %v", err)
	}

	results := pokedex.SearchByNamePrefix("Chik")
	if len(results) != 1 || results[0].Name != "Chikorita" {
		fmt.Printf("%v", pokedex.nameTrie)
		t.Errorf("SearchByNamePrefix(Chiko) returned incorrect results")
	}
}

func TestAddNewSpecies_DuplicateID(t *testing.T) {
	pokedex := createTestPokedex()

	newSpecies := Species{ID: 4, Name: "Chikorita", Types: []Type{Grass}}

	err := pokedex.AddNewSpecies(newSpecies)

	if err == nil {
		t.Errorf("Expected error when adding duplicate ID, got nil")
	}
}

func TestSearchByType(t *testing.T) {
	pokedex := createTestPokedex()
	electricSpecies := pokedex.SearchByType(Electric)

	if len(electricSpecies) != 1 || electricSpecies[0].Name != "Raichu" {
		t.Errorf("SearchByType(Electric) returned incorrect results")
	}
}

func TestSearchByType_EmptyType(t *testing.T) {
	pokedex := createTestPokedex()
    results := pokedex.SearchByType(Water)
    if len(results) != 0 {
        t.Errorf("Expected no results for type with no members, got %d", len(results))
    }
}

func TestSearchByNamePrefix_EmptyString(t *testing.T) {
	pokedex := createTestPokedex()
	results := pokedex.SearchByNamePrefix("")
	if len(results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(results))
	}
}

func TestGetSpeciesByID_NegativeID(t *testing.T) {
	pokedex := createTestPokedex()
	species := pokedex.GetSpeciesByID(-1)
	if species != nil {
		t.Errorf("Expected nil, got %v", species)
	}
}

func TestEmptyPokedex(t *testing.T) {
	pokedex := NewPokedex([]Species{})
	if len(pokedex.Species) != 0 {
		t.Errorf("Expected Pokedex to be empty, got %d species", len(pokedex.Species))
	}
}

func TestLargePokedex(t *testing.T) {
	species := make([]Species, 10000)
	for i := range species {
		species[i] = Species{ID: i, Name: fmt.Sprintf("Species%d", i), Types: []Type{Normal}}
	}
	pokedex := NewPokedex(species)
	if len(pokedex.Species) != len(species) {
		t.Errorf("Expected Pokedex to have %d species, got %d", len(species), len(pokedex.Species))
	}

}

func TestLoadPokedexFromJSON_NonExistentFile(t *testing.T) {
	_, err := LoadPokedexFromJSON("non_existent_file.json")
	if err == nil {
		t.Errorf("Expected error when loading from non-existent file, got nil")
	}
}

func TestLoadPokedexFromJSON_InvalidJSON(t *testing.T) {
	filename := "invalid.json"
	err := os.WriteFile(filename, []byte("{invalid json}"), 0644)
	if err != nil {
		t.Fatalf("Failed to create invalid json file: %v", err)
	}
	_, err = LoadPokedexFromJSON(filename)
	if err == nil {
		t.Errorf("Expected error when loading from invalid json file, got nil")
	}
	os.Remove(filename)
}
