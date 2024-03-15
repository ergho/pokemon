package pokemon

import (
	"math/rand"
	"reflect"
	"testing"
)

var (
	BulbasaurSpecies = &Species{
		Name:  "Bulbasaur",
		Types: []Type{Grass, Poison},
	}
	CharmanderSpecies = &Species{
		Name:  "Charmander",
		Types: []Type{Fire},
	}
)

type MockRand struct {
	IntnFunc func(n int) int
}

func (m *MockRand) Intn(n int) int {
	return m.IntnFunc(n)
}

// HelperFunc
func newPoisonMove() Move {
	return Move{
		Name:         "Poison Sting",
		Type:         Poison,
		Category:     Physical,
		Power:        15,
		Accuracy:     100,
		PP:           35,
		StatusEffect: &PoisonEffect{Chance: 30},
	}
}

// Helper function to create a move that causes sleep
func NewSleepMove() Move {
	return Move{
		Name:         "Sleep Powder",
		Type:         Grass,
		Category:     Special,
		Power:        0,
		Accuracy:     75,
		PP:           15,
		StatusEffect: &SleepStatus{Duration: rand.Intn(3) + 1},
	}
}

func TestNewPokemon(t *testing.T) {
	species := &Species{
		ID:        25,
		Name:      "Pikachu",
		Types:     []Type{Electric},
		BaseStats: Stats{HP: 35, Attack: 55, Defense: 40, SpecialAttack: 50, SpecialDefense: 50, Speed: 90},
	}

	pikachu := NewPokemon(species, 5)

	if pikachu.Species.Name != "Pikachu" {
		t.Errorf("NewPokemon() species = %v, want %v", pikachu.Species.Name, "Pikachu")
	}

	if pikachu.Level != 5 {
		t.Errorf("NewPokemon() level = %v, want %v", pikachu.Level, 5)
	}

	// We could check ivs here i guess, or maybe enouhg in TestGenerateRandomIVs function hmm
	for _, iv := range []int{pikachu.ivs.HP, pikachu.ivs.Attack, pikachu.ivs.Defense,
		pikachu.ivs.SpecialAttack, pikachu.ivs.SpecialDefense, pikachu.ivs.Speed} {
		if iv < 0 || iv > 31 {
			t.Errorf("NewPokemon() IVs not in range 0-31, got %v", iv)
		}
	}
}

func TestGenerateRandomIVs(t *testing.T) {
	for i := 0; i < 100; i++ {
		ivs := GenerateRandomIVs()
		rv := reflect.ValueOf(ivs)
		for k := 0; k < rv.NumField(); k++ {
			if iv := rv.Field(k).Int(); iv < 0 || iv > 31 {
				t.Errorf("GenerateRandomIVs() generated IV out of range, got %v", iv)
			}
		}
	}
}

func TestCalculateStats(t *testing.T) {
	stats := Stats{HP: 35, Attack: 55, Defense: 40, SpecialAttack: 50, SpecialDefense: 50, Speed: 90}
	ivs := Stats{HP: 31, Attack: 31, Defense: 31, SpecialAttack: 31, SpecialDefense: 31, Speed: 31}
	level := 50

	expectedStats := Stats{HP: 110, Attack: 75, Defense: 60, SpecialAttack: 70, SpecialDefense: 70, Speed: 110}

	calculatedStats := CalculateStats(stats, level, ivs)

	if !reflect.DeepEqual(calculatedStats, expectedStats) {
		t.Errorf("Calculatestats() stats = %v, expected = %v", calculatedStats, expectedStats)
	}
}

// Test the execution of a move with a chance to poison
func TestMoveExecuteWithChanceToPoison(t *testing.T) {
	user := NewPokemon(BulbasaurSpecies, 50)
	target := NewPokemon(CharmanderSpecies, 50)
	poisonMove := newPoisonMove()

	poisonMove.Execute(user, target)

	// Check if the target is poisoned (30% chance)
	if target.StatusManager.Primary != nil && target.StatusManager.Primary.Name() == "Poison" {
		t.Log("Target successfully poisoned.")
	} else {
		t.Log("Target not poisoned.")
	}
}

// Test the execution of a move that causes sleep
func TestMoveExecuteWithSleep(t *testing.T) {
	user := NewPokemon(BulbasaurSpecies, 50)
	target := NewPokemon(CharmanderSpecies, 50)
	sleepMove := NewSleepMove()

	sleepMove.Execute(user, target)

	// Check if the target is asleep
	if target.StatusManager.Primary != nil && target.StatusManager.Primary.Name() == "Sleep" {
		t.Log("Target successfully put to sleep.")
	} else {
		t.Log("Target not put to sleep.")
	}
}

// Test the update of status effects at the start of each turn
func TestUpdateStatusEffects(t *testing.T) {
	pokemon := NewPokemon(CharmanderSpecies, 50)
	pokemon.StatusManager.Primary = &SleepStatus{Duration: 2}

	// Simulate turns
	for i := 0; i < 3; i++ {
		pokemon.StatusManager.UpdateStatusEffects(pokemon)
		if pokemon.StatusManager.Primary != nil && pokemon.StatusManager.Primary.Name() == "Sleep" {
			t.Logf("Charmander is asleep. Turns left: %d", 2-i)
		} else {
			t.Log("Charmander woke up!")
			break
		}
	}
}
