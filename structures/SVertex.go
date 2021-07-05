package structures

import "fmt"

type SVertex struct {
	X int
	Y int
}

func (vertex SVertex) String() string {
	res := fmt.Sprint("X :")
	res += fmt.Sprint(vertex.X)
	res += fmt.Sprint(" Y :")
	res += fmt.Sprint(vertex.Y)
	res += fmt.Sprint(" ")
	return res
}
