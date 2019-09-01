package main

type Character struct {
	Name           string
	WeaponBehavior WeaponBehavior
}

func (c *Character) setWeaponBehavior(w WeaponBehavior) {
	c.WeaponBehavior = w
}

type Wizzard struct {
	Character
}

type Wolf struct {
	Character
}

type Vampire struct {
	Character
}
