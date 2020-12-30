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
	profile.Skills  = append(profile.Skills, skill)
}

func (profile *Profile) showSkills() {
	for i, s := range profile.Skills {
		fmt.Println(i, s.Name)
	}
}

func main() {
	profile := &Profile{Name: "Marcos Celio", Skills: []Skill{Skill{Name: "Golang"}}}
	fmt.Printf("Name: %s\n", profile.Name)
	fmt.Println("Skills:")
	profile.showSkills()
	profile.AddSkill(Skill{"Java"})
	fmt.Println("Added Skills")
	profile.showSkills()
	fmt.Println("End")
}
