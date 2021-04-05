package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},

		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},

		{
			"Struct with mixed fields",
			struct {
				Name string
				Age  int
				X    []int
			}{"Chris", 20, []int{1, 2, 3}},
			[]string{"Chris"},
		},

		{
			"Nested fields",
			Person{
				"Chris",
				Profile{20, "London"},
			},
			[]string{"Chris", "London"},
		},

		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{20, "London"},
			},
			[]string{"Chris", "London"},
		},

		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Mars"},
			},
			[]string{"London", "Mars"},
		},

		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Mars"},
			},
			[]string{"London", "Mars"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
