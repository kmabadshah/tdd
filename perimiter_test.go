package main

import (
	"math"
	"testing"
)

/*
return slice of shape with w1 as the wanted value for the first shape and w2 as the
wanted value for the second shape and so on

shapes := GetShapes(2 * (5 + 4), 2 * math.Pi * 5, (10 * 5) / 2)
fmt.Println(shapes)

Output: {
	{Rectangle{5, 4}, 2 * (5 + 4)},
	{Circle{5}, 2 * math.Pi * 5},
	{Triangle{10, 5}, (10 * 5) / 2},
}
*/
func GetShapes(w1, w2, w3 float64) []struct {
	name  string
	shape Shape
	want  float64
} {
	return []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{5, 4}, want: w1},
		{name: "Circle", shape: Circle{5}, want: w2},
		{name: "Triangle", shape: Triangle{10, 5}, want: w3},
	}
}

func TestPerimeter(t *testing.T) {
	shapes := GetShapes(
		2*(5+4),
		2*math.Pi*5,
		10+5+math.Sqrt(math.Pow(10, 2)+math.Pow(5, 2)),
	)

	for _, st := range shapes {
		t.Run(st.name, func(t *testing.T) {
			got := st.shape.Perimeter()
			if got != st.want {
				t.Errorf("%#v GOT %v WANT %v", st.shape, got, st.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	shapes := GetShapes(5*4, math.Pi*math.Pow(5, 2), (10*5)/2)

	for _, st := range shapes {
		t.Run(st.name, func(t *testing.T) {
			got := st.shape.Area()
			if got != st.want {
				t.Errorf("%#v GOT %v WANT %v", st.shape, got, st.want)
			}
		})
	}
}
