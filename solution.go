package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type vs int

func CalcSquare(sideLen float64, sidesNum vs) (res float64) {

	const SidesTriangle vs = 3
	const SidesSquare vs = 4
	const SidesCircle vs = 0

	switch sidesNum {
	case SidesCircle:
		res = math.Pi * math.Pow(sideLen, 2.0)
	case SidesTriangle:
		res = (math.Pow(sideLen, 2.0) * math.Sqrt(3.0)) / 4
	case SidesSquare:
		res = math.Pow(sideLen, 2.0)
	default:
		res = 0
	}

	return
}
