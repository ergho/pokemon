package pokemon

import (
	"testing"
)

func TestHeldItem(t *testing.T) {
	species := &Species{BaseStats: Stats{HP: 100}}
	pokemon := NewPokemon(species, 10, OranBerry, nil, [4]Move{})

	pokemon.Health.decrease(20)

	pokemon.UseItem()

	if pokemon.Health.Current != pokemon.Health.Max-10 {
		t.Errorf("Expected health to be 90, not %d", pokemon.Health.Current)
	}
}
