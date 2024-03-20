package pokemon

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dghubble/trie"
)

type PokedexRepository interface {
	GetSpeciesByID(id int) *Species
	SearchByNamePrefix(prefix string) []*Species
	SearchByType(pokemonType Type) []*Species
	buildIndices()
	AddNewSpecies(newSpecies Species) error
}

type Pokedex struct {
	Species     []Species           `json:"species"`
	indexByID   map[int]*Species    `json:"-"`
	nameTrie    *trie.RuneTrie      `json:"-"`
	typeHashMap map[Type][]*Species `json:"-"`
}

func NewPokedex(species []Species) *Pokedex {
	p := &Pokedex{
		Species:     species,
		indexByID:   make(map[int]*Species),
		nameTrie:    trie.NewRuneTrie(),
		typeHashMap: make(map[Type][]*Species),
	}

	p.buildIndices()
	return p
}

func LoadPokedexFromJSON(filename string) (*Pokedex, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var pokedex Pokedex
	if err := json.Unmarshal(bytes, &pokedex); err != nil {
		return nil, err
	}

	pokedex.buildIndices()
	os.Remove(filename)
	return &pokedex, nil
}

func (p *Pokedex) SavePokedexToJSON(filename string) error {
	//only Species will get serialized to json
	bytes, err := json.Marshal(p)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, bytes, 0644)
}

func (p *Pokedex) buildIndices() {
	if p.indexByID == nil {
		p.indexByID = make(map[int]*Species)
	}

	if p.typeHashMap == nil {
		p.typeHashMap = make(map[Type][]*Species)
	}

	if p.nameTrie == nil {
		p.nameTrie = trie.NewRuneTrie()

	}

	for i := range p.Species {
		p.indexByID[p.Species[i].ID] = &p.Species[i]
		p.nameTrie.Put(p.Species[i].Name, &p.Species[i])
		//p.indexByName[p.Species[i].Name] = &p.Species[i]
		for _, t := range p.Species[i].Types {
			p.typeHashMap[t] = append(p.typeHashMap[t], &p.Species[i])
		}
	}
}

func (p *Pokedex) GetSpeciesByID(id int) *Species {
	if species, exists := p.indexByID[id]; exists {
		return species
	}
	return nil
}

func (p *Pokedex) SearchByType(pokemonType Type) []*Species {
	return p.typeHashMap[pokemonType]
}

func (p *Pokedex) SearchByNamePrefix(prefix string) []*Species {
	var matchingSpecies []*Species
	err := p.nameTrie.Walk(func(word string, specie interface{}) error {
		if len(word) >= len(prefix) && word[:len(prefix)] == prefix {
			s, ok := specie.(*Species)
			if ok {
				matchingSpecies = append(matchingSpecies, s)
			}
		}
		return nil
	})

	if err != nil {
		return nil
	}

	return matchingSpecies
}

func (p *Pokedex) AddNewSpecies(newSpecies Species) error {
	if _, exists := p.indexByID[newSpecies.ID]; exists {
		return fmt.Errorf("species with ID: %d  already exists", newSpecies.ID)
	}

	// implement check to ensure duplicate names are not possible while allowing for partial name similarity

	p.Species = append(p.Species, newSpecies)

	p.buildIndices()
	return nil
}
