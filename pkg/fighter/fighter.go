package fighter

import "fmt"

type Fighter struct {
	name        string
	nationality string
	age         int32
	height      float32
	weight      float32
	category    string
	victories   int32
	defeats     int32
	draws       int32
}

func (f *Fighter) SetName(name string) {
	f.name = name
}

func (f Fighter) Name() string {
	return f.name
}

func (f *Fighter) SetNationality(nationality string) {
	f.nationality = nationality
}

func (f Fighter) Nationality() string {
	return f.nationality
}

func (f *Fighter) SetAge(age int32) {
	f.age = age
}

func (f Fighter) Age() int32 {
	return f.age
}

func (f *Fighter) SetHeight(height float32) {
	f.height = height
}

func (f Fighter) Height() float32 {
	return f.height
}

func (f *Fighter) SetWeight(weight float32) {
	f.weight = weight
}

func (f Fighter) Weight() float32 {
	return f.weight
}
func (f *Fighter) SetCategory(category string) {
	f.category = category
}

func (f Fighter) Category() string {
	return f.category
}

func (f *Fighter) SetVictories(victories int32) {
	f.victories = victories
}

func (f Fighter) Victories() int32 {
	return f.victories
}
func (f *Fighter) SetDefeats(defeats int32) {
	f.defeats = defeats
}

func (f Fighter) Defeats() int32 {
	return f.defeats
}
func (f *Fighter) SetDraws(draws int32) {
	f.draws = draws
}

func (f Fighter) Draws() int32 {
	return f.draws
}

func (f Fighter) Present() string {
	return fmt.Sprintf("Name: %s \nNationality: %s \nAge: %d \nHeight: %f \nWeight: %f \nCategory: %s", f.name, f.nationality, f.age, f.height, f.weight, f.category)
}

func (f Fighter) Status() string {
	return fmt.Sprintf(
		"Name: %s \nNum of Fights: %d \nVictories: %d \nDefeats: %d \nDraws: %d",
		f.Name(),
		f.Victories()+f.Defeats()+f.Draws(),
		f.Victories(),
		f.Defeats(),
		f.Draws(),
	)
}

func (f *Fighter) WinFight() {
	f.SetVictories(f.Victories() + 1)
}

func (f *Fighter) LoseFight() {
	f.SetDefeats(f.Defeats() + 1)
}

func (f *Fighter) DrawFight() {
	f.SetDraws(f.Draws() + 1)
}
