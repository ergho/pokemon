package pokemon

import (
	"testing"
)

var (
	mockRepo = MockPokedex{
		Mockdata: map[int]*Species{
			5: {
				ID:   5,
				Name: "Charmeleon",
				EvolutionStages: []EvolutionStage{
					{
						EvolvesInto: 6, // Charizard's ID
						Method:      LevelEvolution{RequiredLevel: 36},
					},
				},
			},
			6: {ID: 6, Name: "Charizard"},
			95: {
				ID:   95,
				Name: "Onix",
				EvolutionStages: []EvolutionStage{
					{
						EvolvesInto: 208, // Steelix's ID
						Method:      ItemEvolution{RequiredItem: "Metal Coat"},
					},
				},
			},
			133: {
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
			196: {ID: 196, Name: "Espeon"},
			197: {ID: 197, Name: "Umbreon"},
			208: {ID: 208, Name: "Steelix"},
		},
	}
)

func TestEvolution(t *testing.T) {
	tests := []struct {
		name            string
		pokemon         *Pokemon
		currentTime     Time
		currentWeather  Weather
		currentLocation string
		wantSpeciesID   int
	}{
		{"Eevee to Espeon", &Pokemon{Species: mockRepo.GetSpeciesByID(133), Level: 20, Friendship: 220}, "Day", "", "", 196},
		{"Eevee to Umbreon", &Pokemon{Species: mockRepo.GetSpeciesByID(133), Level: 20, Friendship: 220}, "Night", "", "", 197},
		{"Charmeleon to Charizard", &Pokemon{Species: mockRepo.GetSpeciesByID(5), Level: 36}, "Anytime", "", "", 6},
		{"Onix to Steelix with Metal Coat", &Pokemon{Species: mockRepo.GetSpeciesByID(95), Level: 25, HeldItem: "Metal Coat"}, "Anytime", "", "", 208},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.pokemon.Evolve(&mockRepo, tt.currentTime, tt.currentWeather, tt.currentLocation)
			if err != nil {
				t.Fatalf("Evolve() error %v", err)
			}

			if got := tt.pokemon.Species.ID; got != tt.wantSpeciesID {
				t.Errorf("Evolve() =  %v, want %v", got, tt.wantSpeciesID)
			}
		})
	}
}
