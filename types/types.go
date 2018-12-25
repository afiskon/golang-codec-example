//+build !test generate

package types

//go:generate codecgen -o types.gen.go types.go

type Spell int

const (
	FIREBALL    Spell = iota
	THUNDERBOLT Spell = iota
)

type Weapon int

const (
	SWORD Weapon = iota
	BOW   Weapon = iota
)

type WarriorInfo struct {
	Weapon       Weapon `codec:"w"`
	ArrowsNumber int    `codec:"a"`
}

type MageInfo struct {
	Spellbook []Spell `codec:"s"`
	Mana      int     `codec:"m"`
}

type Hero struct {
	Name        string       `codec:"n"`
	HP          int          `codec:"h"`
	XP          int          `codec:"x"`
	WarriorInfo *WarriorInfo `codec:"w"`
	MageInfo    *MageInfo    `codec:"m"`
}
