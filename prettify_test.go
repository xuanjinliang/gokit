package gotool

import (
	"testing"
)

type Person struct {
	Name   *string
	Age    int
	School map[string]*string
	Hobbit []string
}

func TestPrettify(t *testing.T) {
	school := make(map[string]*string)
	school["junior"] = StringPtr("junior school")
	school["high"] = StringPtr("high school")
	person := Person{
		Age:    1,
		Name:   StringPtr("jess"),
		School: school,
		Hobbit: []string{"swim", "game"},
	}

	t.Log(Prettify(person))
}

func TestPrettifyJson(t *testing.T) {
	school := make(map[string]*string)
	school["junior"] = StringPtr("junior school")
	school["high"] = StringPtr("high school")
	person := Person{
		Age:    1,
		Name:   StringPtr("jess"),
		School: school,
		Hobbit: []string{"swim", "game"},
	}
	t.Log(PrettifyJson(person, false))

	t.Log(PrettifyJson([]*string{StringPtr("book"), StringPtr("bed")}, true))
}
