package pokemon

type ItemEffect func(*Pokemon)

type Item interface {
	Use(p *Pokemon)
	Name() string
}

type Berry struct {
	name       string
	isConsumed bool
	effect     func(p *Pokemon)
}

func (b *Berry) Use(p *Pokemon) {
	b.effect(p)
	if b.isConsumed {
		p.HeldItem = nil
	}
}

func (b *Berry) Name() string {
	return b.name
}

// example item
var OranBerry = &Berry{
	name: "Oran Berry",
	effect: func(p *Pokemon) {
		if p.Health.Current < p.Health.Max {
			p.Health.increase(10)
		}
	},
	isConsumed: true,
}

type ReviveItem interface {
	Item
	Revive(p *Pokemon)
}

type Revive struct {
	Item
}

// example of revive being implemented
func (r *Revive) Revive(p *Pokemon) {
	if p.StatusManager.Primary != nil && p.StatusManager.Primary.Name() == "Fainted" {
		p.StatusManager.Primary = nil
		p.Health.increase(int(p.Health.Max / 2))
	}
}
