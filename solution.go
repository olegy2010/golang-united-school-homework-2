package square

import (
	"math"
)

type SidesNum int

const (
	SidesTriangle SidesNum = 3
	SidesSquare   SidesNum = 4
	SidesCircle   SidesNum = 0
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum SidesNum) float64 {
	switch sidesNum {
	case SidesTriangle:
		return math.Sqrt(3.0) * sideLen * sideLen / 4.0
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	default:
		return 0
	}
}
