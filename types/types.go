//+build generate
package types
//go:generate codecgen -o types.gen.go types.go

type Class int
const (
	WARIOR Class = iota
	MAGE   Class = iota
)

type Spell int
const (
	FIREBALL Spell = iota
	THUNDERBOLT Spell = iota
)

type Weapon int
const (
	SWORD Weapon = iota
	BOW Weapon = iota
)

type WariorInfo struct {
	Weapon Weapon `codec:"w"`
	ArrowsNumber int `codec:"a"`
}

type MageInfo struct {
	Spellbook []Spell `codec:"s"`
	Mana int `codec:"m"`
}

type Hero struct {
	Name string `codec:"n"`
	HP int `codec:"h"`
	XP int `codec:"x"`
	WariorInfo *WariorInfo `codec:"w"`
	MageInfo *MageInfo `codec:"m"`
}
