package pokemon

import "math/rand"

type Effect func(*Pokemon, *Pokemon)

type MoveCategory string

const (
	Physical MoveCategory = "Physical"
	Special  MoveCategory = "Special"
)

type Move struct {
	Name         string
	Type         Type
	Category     MoveCategory
	Power        int
	PP           int
	Accuracy     int
	Effects      []Effect
	StatusEffect StatusEffect
	Priority     int
}

func (m *Move) Execute(user *Pokemon, target *Pokemon) {
	if rand.Intn(100) < m.Accuracy {
		for _, effect := range m.Effects {
			effect(user, target)
		}

		if m.StatusEffect != nil {
			m.StatusEffect.Apply(target)
		}
	}
}

// Example of defining a stat-boosting effect
var attackBoost = func(user *Pokemon, _ *Pokemon) {
	user.Modifiers.AttackMultiplier *= 2.0
	if user.Modifiers.AttackMultiplier > 4.0 {
		user.Modifiers.AttackMultiplier = 4.0
	}
}

// Example move with multiple effects
var buffBuff = Move{
	Name:     "Buff buff",
	Type:     Normal,
	Power:    0,
	Accuracy: 100,
	PP:       8,
	Effects:  []Effect{attackBoost, attackBoost}, // should apply both example effects, will probably move to test code or something
}
