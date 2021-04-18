package generator

type NewGeneratorFunc = func(int) Generator

type Generator interface {
	Name() string
	Next() string
}
