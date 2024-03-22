package pokemon

//import "fmt"

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

//type BattlerAction struct {
//	Action  BattleAction
//	Battler Battler
//}
type Battler interface {
	ChooseAction() BattleAction
	ExecuteAction(action BattleAction, opponent *Pokemon)
	GetPokemon() *Pokemon
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
        action.Move.Execute(w.GetPokemon(), opponent)
		//fmt.Printf("Wild %s used %s!\n", w.Pokemon.Species.Name, action.Move.Name)
		//opponent.Health.Current -= action.Move.Power // will add damage calculating at a later time
	}
}

func (w *WildPokemon) GetPokemon() *Pokemon {
	return w.Pokemon
}

func (t *Trainer) ChooseAction() BattleAction {
	// we need logic here somehow to choose an action...
	return BattleAction{Type: Attack, Move: t.Team[0].Moves[0]}
}

func (t *Trainer) GetPokemon() *Pokemon {
	if len(t.Team) > 0 && t.Team[0] != nil {
		return t.Team[0]
	}
	return nil
}

func (t *Trainer) ExecuteAction(action BattleAction, opponent *Pokemon) {
	switch action.Type {
	case Attack:
		action.Move.Execute(t.GetPokemon(), opponent)
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
		action1 := b.Battler1.ChooseAction()
		action2 := b.Battler2.ChooseAction()

		// Determine the order of execution based on priority and speed
		firstBattler, secondBattler := b.determineActionOrder(b.Battler1, action1, b.Battler2, action2)

		// Execute actions in the determined order
		firstBattler.ExecuteAction(action1, secondBattler.GetPokemon())
		if secondBattler.GetPokemon().Health.Current > 0 {
			secondBattler.ExecuteAction(action2, firstBattler.GetPokemon())
		}

		// Check for end conditions
		if b.checkEndConditions() {
			b.Running = false
		}
	}
}

// determineActionOrder determines which battler should act first based on action priority and speed
func (b *Battle) determineActionOrder(battler1 Battler, action1 BattleAction, battler2 Battler, action2 BattleAction) (firstBattler, secondBattler Battler) {
	switch {
	case action1.Move.Priority > action2.Move.Priority:
		return battler1, battler2
	case action1.Move.Priority < action2.Move.Priority:
		return battler2, battler1
	case b.Battler1.GetPokemon().Stats.Speed > b.Battler2.GetPokemon().Stats.Speed:
		return battler1, battler2
	case b.Battler1.GetPokemon().Stats.Speed < b.Battler2.GetPokemon().Stats.Speed:
		return battler2, battler1
	default:
		// If priority and speed are the same, choose randomly or by another logic
		// Placeholder for additional logic
		return battler1, battler2
	}
}

func (b *Battle) checkEndConditions() bool {
	return true
}
