package structures

import "fmt"

type Tnumber uint8
type Tgoban [][]Tnumber

type SContext struct {
	Goban         Tgoban
	NSize         uint8
	SizeCase      int32
	Size          int32
	CurrentPlayer uint8
	NbVictoryP1   int
	NbVictoryP2   int
	NbCaptureP1   int
	NbCaptureP2   int
}

func (goban Tgoban) String() string {
	res := "Goban: {\n"
	for _, tab := range goban {
		res += "\t"
		for _, cases := range tab {
			res += fmt.Sprint(cases)
			res += " "
		}
		res += "\n"
	}
	return res + "}"
}

func (ctx SContext) String() string {
	res := "---------------SContext---------------\n"
	res += fmt.Sprint(ctx.Goban)
	res += "\n--------------------------------------"
	return res
}
