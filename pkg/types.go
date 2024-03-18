package pokemon

import (
	"errors"
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
