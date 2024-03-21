package pokemon

import "fmt"

type ActionType int

const (
	Attack ActionType = iota
	UseItem
	SwitchPokemon
	Flee
)

type BattleAction struct {
	Type     ActionType
	Move     Move
	Item     Item
	SwitchTo int
}

type Battler interface {
	ChooseAction() BattleAction
	ExecuteAction(action BattleAction, opponent *Pokemon)
}

type WildPokemon struct {
	Pokemon *Pokemon
}

func (w *WildPokemon) ChooseAction() BattleAction {
	//We can add more logic on these later but just jam first move always for tests
	return BattleAction{Type: Attack, Move: w.Pokemon.Moves[0]}
}

func (w *WildPokemon) ExecuteAction(action BattleAction, opponent *Pokemon) {
	//doing minimal logic for now...
	if action.Type == Attack {
		fmt.Printf("Wild %s used %s!\n", w.Pokemon.Species.Name, action.Move.Name)
		opponent.Health.Current -= action.Move.Power // will add damage calculating at a later time
	}
}

func (t *Trainer) ChooseAction() BattleAction {
	// we need logic here somehow to choose an action...
	return BattleAction{Type: Attack, Move: t.Team[0].Moves[0]}
}

func (t *Trainer) ExecuteAction(action BattleAction, opponent *Pokemon) {
	switch action.Type {
	case Attack:
		opponent.Health.decrease(10)
	case UseItem:
		// later
	case SwitchPokemon:
		// later
	case Flee:
		// later
	}
}

type Battle struct {
	Battler1 Battler
	Battler2 Battler
	Running  bool
}

func NewBattle(battler1, battler2 Battler) *Battle {
	return &Battle{
		Battler1: battler1,
		Battler2: battler2,
		Running:  true,
	}
}

func (b *Battle) Run() {
	for b.Running {
		//we keep this design for now, maybe we will track rounds or something later...
		action1 := b.Battler1.ChooseAction()
		action2 := b.Battler2.ChooseAction()

		b.Battler1.ExecuteAction(action1, b.getBattlerPokemon(b.Battler2))
		b.Battler2.ExecuteAction(action2, b.getBattlerPokemon(b.Battler1))

        if b.checkEndConditions() {
            b.Running = false
        }
	}
}

func (b *Battle) getBattlerPokemon(battler Battler) *Pokemon {
	switch v := battler.(type) {
	case *Trainer:
		return &v.Team[0]
	case *WildPokemon:
		return v.Pokemon
	default:
		// shouldnt be reachable... if implemented correctly
		return nil
	}
}

func (b *Battle) checkEndConditions() bool {
    return false
}
