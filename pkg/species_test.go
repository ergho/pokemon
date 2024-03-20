package pokemon

import (
	"testing"
)

type MockCoat struct {
	Item
}

var (
	Mockdata  = []Species{
		{
			ID:   5,
			Name: "Charmeleon",
			EvolutionStages: []EvolutionStage{
				{
					EvolvesInto: 6, // Charizard's ID
					Method:      LevelEvolution{RequiredLevel: 36},
				},
			},
		},
		{ID: 6, Name: "Charizard"},
		{
			ID:   95,
			Name: "Onix",
			EvolutionStages: []EvolutionStage{
				{
					EvolvesInto: 208, // Steelix's ID
					Method:      ItemEvolution{RequiredItem: MockCoat{}},
				},
			},
		},
		{
			ID:   133,
			Name: "Eevee",
			EvolutionStages: []EvolutionStage{
				{
					EvolvesInto: 196, // Espeon's ID
					Method: FriendshipEvolution{
						RequiredFriendship: 220,
						RequiredTime:       "Day",
					},
				},
				{
					EvolvesInto: 197, // Umbreon's ID
					Method: FriendshipEvolution{
						RequiredFriendship: 220,
						RequiredTime:       "Night",
					},
				},
			},
		},
		{ID: 196, Name: "Espeon"},

		{ID: 197, Name: "Umbreon"},
		{ID: 208, Name: "Steelix"},
	}
)

func TestEvolution(t *testing.T) {
	pokedex := NewPokedex(Mockdata)
	tests := []struct {
		name            string
		pokemon         *Pokemon
		currentTime     Time
		currentWeather  Weather
		currentLocation string
		wantSpeciesID   int
	}{
		{"Eevee to Espeon", &Pokemon{Species: pokedex.GetSpeciesByID(133), Level: 20, Friendship: 220}, "Day", "", "", 196},
		{"Eevee to Umbreon", &Pokemon{Species: pokedex.GetSpeciesByID(133), Level: 20, Friendship: 220}, "Night", "", "", 197},
		{"Charmeleon to Charizard", &Pokemon{Species: pokedex.GetSpeciesByID(5), Level: 36}, "Anytime", "", "", 6},
		{"Onix to Steelix with Metal Coat", &Pokemon{Species: pokedex.GetSpeciesByID(95), Level: 25, HeldItem: MockCoat{}}, "Anytime", "", "", 208},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.pokemon.Evolve(pokedex, tt.currentTime, tt.currentWeather, tt.currentLocation)
			if err != nil {
				t.Fatalf("Evolve() error %v", err)
			}

			if got := tt.pokemon.Species.ID; got != tt.wantSpeciesID {
				t.Errorf("Evolve() =  %v, want %v", got, tt.wantSpeciesID)
			}
		})
	}
}
