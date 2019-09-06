package main

import "fmt"

type WeaponBehavior interface {
	Shot()
}

type Gun struct {
	munition int32
}

func (g *Gun) Shot() {
	fmt.Println("Shot a gun!")
}

type Bazooka struct {
	munition int32
}

func (b *Bazooka) launchRocket() {
	fmt.Println("Shot a Bazooka!")
}

func (b *Bazooka) Shot() {
	b.launchRocket()
}

type Sniper struct {
	munition int32
}

func (s *Sniper) headShot() {
	fmt.Println("Shot a headshot with sniper!")
}

func (s *Sniper) Shot() {
	s.headShot()
}
