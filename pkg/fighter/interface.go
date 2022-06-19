package fighter

type FighterInterface interface {
	Name() string
	SetName(name string)
	Nationality() string
	SetNationality(nationality string)
	Age() int32
	SetAge(age int32)
	Height() float32
	SetHeight(height float32)
	Weight() float32
	SetWeight(weight float32)
	Category() string
	SetCategory(category string)
	Victories() int32
	SetVictories(victories int32)
	Defeats() int32
	SetDefeats(defeats int32)
	Draws() int32
	SetDraws(draws int32)

	Present() string
	Status() string
	WinFight()
	LoseFight()
	DrawFight()
}
