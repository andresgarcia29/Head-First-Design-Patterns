package main

func main() {
	vampire := new(Vampire)
	vampire.Name = "I'm a vampire"
	vampire.setWeaponBehavior(new(Sniper))
	vampire.WeaponBehavior.Shot()
	vampire.setWeaponBehavior(new(Bazooka))
	vampire.WeaponBehavior.Shot()
}
