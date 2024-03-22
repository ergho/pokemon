package pokemon

import (
	"fmt"
	"testing"
)

// MockBattler is a mock implementation of the Battler interface for testing
type MockBattler struct {
	Pokemon *Pokemon
	Action  BattleAction
}

func (mb *MockBattler) ChooseAction() BattleAction {
	return mb.Action
}

func (mb *MockBattler) ExecuteAction(action BattleAction, opponent *Pokemon) {

	if action.Type == Attack {
        fmt.Println("hello world im here")
        action.Move.Execute(mb.GetPokemon(), opponent)
	// Simplified logic for testing: just reduce the opponent's health by the move's power
    }
}

func (mb *MockBattler) GetPokemon() *Pokemon {
	return mb.Pokemon
}

// TestBattleOrder tests the order of actions in a battle based on move priority and speed
func TestBattleOrder(t *testing.T) {
	// Setup mock Pokémon with different speeds and moves
	quickAttack := Move{Name: "Quick Attack", Power: 40, Priority: 1}
	tackle := Move{Name: "Tackle", Power: 50, Priority: 0}

	fastPokemon := &Pokemon{Species: CharmanderSpecies, Health: Health{Current: 100, Max: 100}, Stats: Stats{Speed: 90}, Moves: [4]Move{quickAttack}}
	slowPokemon := &Pokemon{Species: BulbasaurSpecies, Health: Health{Current: 100, Max: 100}, Stats: Stats{Speed: 45}, Moves: [4]Move{tackle}}

	fastBattler := &MockBattler{Pokemon: fastPokemon, Action: BattleAction{Type: Attack, Move: quickAttack}}
	slowBattler := &MockBattler{Pokemon: slowPokemon, Action: BattleAction{Type: Attack, Move: tackle}}

	// Run the battle
	battle := NewBattle(fastBattler, slowBattler)
	battle.Run()

	// Verify that the fast Pokémon with the priority move attacked first
	if slowPokemon.Health.Current != 60 {
		t.Errorf("Expected slow Pokémon to have 60 health after Quick Attack, got %d", slowPokemon.Health.Current)
	}

	// Verify that the slow Pokémon attacked second
	if fastPokemon.Health.Current != 50 {
		t.Errorf("Expected fast Pokémon to have 50 health after Tackle, got %d", fastPokemon.Health.Current)
	}
}

func TestStatusEffects(t *testing.T) {
	poisonSting := newPoisonMove()
	sleepingPowder := NewSleepMove()

	sleepPokemon := &Pokemon{Species: CharmanderSpecies, Health: Health{Current: 100, Max: 100}, Stats: Stats{Speed: 90}, Moves: [4]Move{sleepingPowder}}
	poisonPokemon := &Pokemon{Species: BulbasaurSpecies, Health: Health{Current: 100, Max: 100}, Stats: Stats{Speed: 50}, Moves: [4]Move{poisonSting}}
	sleepBattler := &MockBattler{Pokemon: sleepPokemon, Action: BattleAction{Type: Attack, Move: sleepingPowder}}
	poisonBattler := &MockBattler{Pokemon: poisonPokemon, Action: BattleAction{Type: Attack, Move: poisonSting}}
	battle := NewBattle(sleepBattler, poisonBattler)
	battle.Run()

	fmt.Printf("%v, %v", battle.Battler1.GetPokemon().Health, battle.Battler2.GetPokemon().Health)

}
