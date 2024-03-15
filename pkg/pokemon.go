package pokemon

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type Type int

const (
	Normal Type = iota
	Fire
	Water
	Electric
	Grass
	Ice
	Fighting
	Poison
	Ground
	Flying
	Psychic
	Bug
	Rock
	Ghost
	Dragon
	Dark
	Steel
)

func (t Type) String() string {
	return [...]string{"Normal", "Fire", "Water", "Electric", "Grass", "Ice",
		"Fighting", "Poison", "Ground", "Flying", "Psychic", "Bug", "Rock",
		"Ghost", "Dragon", "Dark", "Steel"}[t]
}

func StringToPokemonType(typeStr string) (Type, error) {
	switch strings.ToLower(typeStr) {
	case "normal":
		return Normal, nil
	case "fire":
		return Fire, nil
	case "water":
		return Water, nil
	case "electric":
		return Electric, nil
	case "grass":
		return Grass, nil
	case "ice":
		return Ice, nil
	case "fighting":
		return Fighting, nil
	case "poison":
		return Poison, nil
	case "ground":
		return Ground, nil
	case "flying":
		return Flying, nil
	case "psychic":
		return Psychic, nil
	case "bug":
		return Bug, nil
	case "rock":
		return Rock, nil
	case "ghost":
		return Ghost, nil
	case "dragon":
		return Dragon, nil
	case "dark":
		return Dark, nil
	case "steel":
		return Steel, nil
	default:
		return 0, errors.New("invalid Pokemon type")
	}
}

type Nature int

const (
	Adamant Nature = iota
	Bashful
	Bold
	Brave
	Calm
	Careful
	Docile
	Gentle
	Hardy
	Hasty
	Impish
	Jolly
	Lax
	Lonely
	Mild
	Modest
	Naive
	Naughty
	Quiet
	Quirky
	Rash
	Relaxed
	Sassy
	Serious
	Speed
)

var NatureInfo = map[Nature][2]string{
	Adamant: {"Attack", "SpecialAttack"},
	Bashful: {"", ""},
	Bold:    {"Defense", "Attack"},
	Brave:   {"Attack", "Speed"},
	Calm:    {"SpecialDefense", "Attack"},
	Careful: {"SpecialDefense", "SpecialAttack"},
	Docile:  {"", ""},
	Gentle:  {"SpecialDefense", "Defense"},
	Hardy:   {"", ""},
	Hasty:   {"Speed", "Defense"},
	Impish:  {"Defense", "SpecialAttack"},
	Jolly:   {"Speed", "SpecialAttack"},
	Lax:     {"Defense", "SpecialDefense"},
	Lonely:  {"Attack", "Defense"},
	Mild:    {"SpecialAttack", "Defense"},
	Modest:  {"SpecialAttack", "Attack"},
	Naive:   {"Speed", "SpecialDefense"},
	Naughty: {"Attack", "SpecialDefense"},
	Quiet:   {"SpecialAttack", "Speed"},
	Quirky:  {"", ""},
	Rash:    {"SpecialAttack", "SpecialDefense"},
	Relaxed: {"Defense", "Speed"},
	Sassy:   {"SpecialDefense", "Speed"},
	Serious: {"", ""},
	Speed:   {"Speed", "Attack"},
}

type StatusCondition int

const (
	Healthy StatusCondition = iota
	Sleep
)

type Status struct {
	Condition StatusCondition
	Duration  int
	Damage    int
}

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

// These should reset at end of battle or if pokemon is switched out
type StatModifiers struct {
	AttackMultiplier         float32
	DefenseMultplier         float32
	SpecialAttackMultiplier  float32
	SpecialDefenseMultiplier float32
	SpeedMultiplier          float32
}

type Stats struct {
	HP             int
	Attack         int
	Defense        int
	SpecialAttack  int
	SpecialDefense int
	Speed          int
}

type Species struct {
	ID             int
	Name           string
	Types          []Type
	BaseStats      Stats
	BaseExpYield   int
	EvolutionChain []string
	Learnset       map[int]Move
}

type Pokemon struct {
	Species       *Species
	CurrentHealth int
	Level         int
	Experience    int
	StatusManager StatusEffectManager
	Stats         Stats
	ivs           Stats
	Nature        *Nature
	Moves         [4]Move
	Modifiers     StatModifiers
}

type StatusEffect interface {
	Apply(p *Pokemon) bool // true if its still active
	Name() string
}

type StatusEffectManager struct {
	Primary   StatusEffect
	Secondary []StatusEffect
}

func (m *StatusEffectManager) UpdateStatusEffects(p *Pokemon) {
	if m.Primary != nil && !m.Primary.Apply(p) {
		m.Primary = nil
	}

	for i := len(m.Secondary) - 1; i >= 0; i-- {
		if !m.Secondary[i].Apply(p) {
			m.Secondary = append(m.Secondary[:i], m.Secondary[i+1:]...) // remove inactive secondary status effects
		}
	}
}

type SleepStatus struct {
	Duration int
}

func (s *SleepStatus) Apply(p *Pokemon) bool {
	s.Duration--
	return s.Duration > 0
}

func (s *SleepStatus) Name() string {
	return "Sleep"
}

// PoisonEffect to implement the StatusEffect interface for poison
type PoisonEffect struct {
	Chance int // Probability of poisoning the target, represented as a percentage
}

func (p *PoisonEffect) Apply(target *Pokemon) bool {
	// Generate a random number to determine if the poison effect is applied
	if rand.Intn(100) < p.Chance {
		// Inflict poison status if not already poisoned
		if target.StatusManager.Primary == nil || target.StatusManager.Primary.Name() != "Poison" {
			target.StatusManager.Primary = p
		}
	}
	// The poison effect persists until cured
	return true
}

func (p *PoisonEffect) Name() string {
	return "Poison"
}

// Redo this with IVS accounted for.
func CalculateStats(stats Stats, level int, ivs Stats) Stats {
	calc_stat := func(base int, iv int) int {
		return (2*base+iv)*level/100 + 5

	}
	stats.HP = (2*stats.HP+ivs.HP)*level/100 + level + 10
	stats.Attack = calc_stat(stats.Attack, ivs.Attack)
	stats.Defense = calc_stat(stats.Defense, ivs.Defense)
	stats.SpecialAttack = calc_stat(stats.SpecialAttack, ivs.SpecialAttack)
	stats.SpecialDefense = calc_stat(stats.SpecialDefense, ivs.SpecialDefense)
	stats.Speed = calc_stat(stats.Speed, ivs.Speed)

	return stats
}

func (p *Pokemon) LevelUp() {
	p.Level++
	fmt.Println(p.Species.Name, "leveled up to ", p.Level)
	p.Stats = CalculateStats(p.Stats, p.Level, p.ivs)

	// Should probably check for new moves to learn or if evolution happens.
}

func GenerateRandomIVs() Stats {
	return Stats{
		HP:             rand.Intn(32),
		Attack:         rand.Intn(32),
		Defense:        rand.Intn(32),
		SpecialAttack:  rand.Intn(32),
		SpecialDefense: rand.Intn(32),
		Speed:          rand.Intn(32),
	}
}

func NewPokemon(species *Species, level int) *Pokemon {
	ivs := GenerateRandomIVs()
	pokemon := Pokemon{
		Species:    species,
		Level:      level,
		Experience: 0, //We should calculate the actual exprience the pokemon has.
		ivs:        ivs,
		Stats:      CalculateStats(species.BaseStats, level, ivs),
	}
	return &pokemon
}

type Trainer struct {
	Name string
	Team []Pokemon
}

//implement constructors for these and probably to be able to store it as json too

//Apply
