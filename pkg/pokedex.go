package pokemon

type Pokedex interface {
	GetSpeciesByID(id int) *Species
}

type PokedexDatabase struct{}

type PokedexJson struct {}

type MockPokedex struct {
    Mockdata map[int]*Species
}

func (m *MockPokedex) GetSpeciesByID(id int) *Species {
    if species, exists := m.Mockdata[id]; exists {
        return species
    }

    return nil
}
