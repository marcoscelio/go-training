package main

import (
	"fmt"
)

type Skill struct {
	Name string
}

type Profile struct {
	Name   string
	Skills []Skill
}

func (profile *Profile) AddSkill(skill Skill) {
	profile.Skills = append(profile.Skills, skill)
}

func (profile *Profile) showSkills() {
	for i, s := range profile.Skills {
		fmt.Println(i, s.Name)
	}
}

func changeNameNoPointer(p Profile, newName string) Profile {
	p.Name = newName
	return p
}

func changeName(p *Profile, newName string) *Profile {
	p.Name = newName
	return p
}

func main() {
	fmt.Printf("############# No pointer no change##############\n")
	p1 := Profile{Name: "Marcos Celio", Skills: []Skill{Skill{Name: "Golang"}}}
	fmt.Printf("Name: %s\n", p1.Name)
	changeNameNoPointer(p1, "Leonardo Celio")
	fmt.Printf("Name not changed: %s\n", p1.Name)
	fmt.Printf("#### Using pointer we can change mem attribute ##\n")
	p2 := &Profile{Name: "Marcos Celio", Skills: []Skill{Skill{Name: "Golang"}}}
	fmt.Printf("Name: %s\n", p2.Name)
	changeName(p2, "Leonardo Celio")
	fmt.Printf("Name changed: to %s\n", p2.Name)
	fmt.Printf("################################################\n")

	fmt.Println("End")
}
