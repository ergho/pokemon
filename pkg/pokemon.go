package pokemon

import (
	"fmt"
	"math/rand"
)

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

type Health struct {
	Current int
	Max     int
}

func NewHealth(maxHP int) *Health {
	return &Health{
		Current: maxHP,
		Max:     maxHP,
	}
}

func (h *Health) increase(amount int) {
	h.Current += amount
	if h.Current > h.Max {
		h.Current = h.Max
	}
}

func (h *Health) decrease(amount int) {
	h.Current -= amount

	if h.Current <= 0 {
		h.Current = 0
	}
}

func (h *Health) IsFainted() bool {
	return h.Current <= 0
}

type Pokemon struct {
	Species       *Species
	Health        Health
	Level         int
	Experience    int
	HeldItem      Item
	StatusManager StatusEffectManager
	Stats         Stats
	ivs           Stats
	Friendship    int
	Nature        *Nature
	Moves         [4]Move
	Modifiers     StatModifiers
}

func (p *Pokemon) LevelUp() {
	oldMaxHealth := p.Health.Max
	p.Level++
	fmt.Println(p.Species.Name, "leveled up to ", p.Level)
	p.Stats = CalculateStats(p.Stats, p.Level, p.ivs)

	healthIncrease := p.Stats.HP - oldMaxHealth

	p.Health.increase(healthIncrease)

	if p.Health.Current > p.Health.Max {
		p.Health.Current = p.Health.Max
	}

	// Should probably check for new moves to learn or if evolution happens.
}

func (p *Pokemon) Heal(item Item) error {
	// check if pokemon is fainted, return early with error if so, dont consume item then.
	if p.Health.IsFainted() {
		if reviveItem, ok := item.(ReviveItem); ok {
			reviveItem.Revive(p)
			return nil
		} else {
			return fmt.Errorf("Cannot heal a fainted pokemon with %s", item.Name())

		}
	}
	if p.Health.Current >= p.Health.Max {
		return fmt.Errorf("Cannot heal a pokemon with full health")
	}
	item.Use(p)

	return nil
}

func (p *Pokemon) TakeDamage(amount int) {
	p.Health.decrease(amount)

	if p.Health.IsFainted() {
		p.StatusManager.Primary = &FaintedStatus{}
	}
}

func (p *Pokemon) HasItem(item Item) bool {
	return true
}

func (p *Pokemon) UseItem() {
	p.HeldItem.Use(p)
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

type StatusEffect interface {
	Apply(p *Pokemon) bool // true if its still active
	Name() string
}

type StatusEffectManager struct {
	Primary   StatusEffect
	Secondary []StatusEffect
}

func (m *StatusEffectManager) PrimaryStatus() string {
	return m.Primary.Name()
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

type FaintedStatus struct{}

func (f *FaintedStatus) Apply(p *Pokemon) bool {
	// Fainted status remains until removed by special item or being healed by a healer.
	return true
}

func (f *FaintedStatus) Name() string {
	return "Fainted"
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

func NewPokemon(species *Species, level int, heldItem Item, nature *Nature, moves [4]Move) *Pokemon {
	ivs := GenerateRandomIVs()
	stats := CalculateStats(species.BaseStats, level, ivs)
	pokemon := Pokemon{
		Species:   species,
		Health:    *NewHealth(stats.HP),
		Level:     level,
		HeldItem:  heldItem,
		Nature:    nature,
		Moves:     moves,
		ivs:       ivs,
		Stats:     stats,
		Modifiers: StatModifiers{AttackMultiplier: 1.0, DefenseMultplier: 1.0, SpecialAttackMultiplier: 1.0, SpecialDefenseMultiplier: 1.0, SpeedMultiplier: 1.0},
	}
	return &pokemon
}
