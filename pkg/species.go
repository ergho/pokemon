package pokemon

import "fmt"

type Item string
type Time string
type Weather string

type EvolutionMethod interface {
	CanEvolve(p *Pokemon, currentTime Time, currentWeather Weather, currentLocation string) bool
}

type LevelEvolution struct {
	RequiredLevel int
}

func (l LevelEvolution) CanEvolve(p *Pokemon, _ Time, _ Weather, _ string) bool {
	return p.Level >= l.RequiredLevel
}

type FriendshipEvolution struct {
	RequiredFriendship int
	RequiredTime       Time
}

func (f FriendshipEvolution) CanEvolve(p *Pokemon, currentTime Time, _ Weather, _ string) bool {
	return p.Friendship >= f.RequiredFriendship && currentTime == f.RequiredTime
}

type ItemEvolution struct {
	RequiredItem Item
}

func (i ItemEvolution) CanEvolve(p *Pokemon, _ Time, _ Weather, _ string) bool {
	return p.HasItem(i.RequiredItem)
}

type EvolutionStage struct {
	EvolvesInto int
	Method      EvolutionMethod
}

type Species struct {
	ID              int
	Name            string
	Types           []Type
	BaseStats       Stats
	BaseExpYield    int
	EvolutionStages []EvolutionStage
	Learnset        map[int]Move
}

//func GetSpeciesByID(id int) *Species {
//	return nil
//}

func (p *Pokemon) Evolve(pokedex Pokedex, currentTime Time, currentWeather Weather, currentLocation string) error {
	for _, stage := range p.Species.EvolutionStages {
		if stage.Method.CanEvolve(p, currentTime, currentWeather, currentLocation) {
			newSpecies := pokedex.GetSpeciesByID(stage.EvolvesInto)

			if newSpecies == nil {
				return fmt.Errorf("Species with ID %d does not exist", stage.EvolvesInto)
			}

			fmt.Printf("%s has evolved into %s!\n", p.Species.Name, newSpecies.Name)
			p.Species = newSpecies
			return nil
		}
	}
	return fmt.Errorf("no evolution conditions met")
}
