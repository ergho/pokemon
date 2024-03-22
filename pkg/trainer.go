package pokemon

type Trainer struct {
	ID           string
	Name         string
	Team         [6]*Pokemon
	Items        []Item
	Location     string
	Pokedex      PokedexRepository
	Achievements []Achievement
}

func NewTrainer(name string, team [6]*Pokemon) *Trainer {
	return &Trainer{
		Name: name,
		Team: team,
	}
}

func (t *Trainer) AddPokemon(newPokemon *Pokemon) bool {
    for i, pokemon := range t.Team {
        if pokemon == nil {
            t.Team[i] = newPokemon
            return true
        }

    }
    return false
}

func (t *Trainer) SwapActivePokemon(swapIndex int) bool {
	if swapIndex > 0 && swapIndex < len(t.Team) && t.Team[swapIndex] != nil {
		t.Team[0], t.Team[swapIndex] = t.Team[swapIndex], t.Team[0]
		return true
	}
	return false

}

func (t *Trainer) RemovePokemon(pokemon Pokemon) error {
	// implement the logic to remove a Pokemon from the trainer's team
	// return an error if the Pokemon is not found in the team
	return nil
}

// More methods related to the Trainer can be added here

type Achievement interface {
	Name() string
	Description() string
	Earned() bool
}

type Badge struct {
	name        string
	description string
	earned      bool
}

func (b *Badge) Name() string {
	return b.name
}

func (b *Badge) Description() string {
	return b.description
}

func (b *Badge) Earned() bool {
	return b.earned
}

// func(t *Trainer) Trade(myPokemon *Pokemon, otherTrainer *Trainer, theirPokemon *Pokemon)

// Quests and or Challenges ?
