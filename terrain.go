package main

import (
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	draw "github.com/fogleman/gg"
)

// Land is a range of x y and z coordinates
type Land struct {
	x, y, z int
}

var (
	seed   time.Time
	random rand.Rand
	sqSize int
	xSize  int
	ySize  int
	zStart int
	img    gg.Context
)

func init() {

	seed = time.Now()

	random = *rand.New(rand.NewSource(seed.Unix()))

	sqSize = 5

	ySize = 180

	xSize = 320

	zStart = random.Intn(256) + 1

	img = *draw.NewContext(xSize*sqSize, ySize*sqSize)

}

func main() {
	// define the landmass
	landMass := make([][]Land, xSize)

	for lands := range landMass {
		landMass[lands] = make([]Land, ySize)
	}
	// define the heights
	for a, arr := range landMass {
		for b, land := range arr {
			land = assignEasy(a, b, land)
			land.z = assignHard(landMass, land)
			//			fmt.Println("x: ", land.x, "y: ", land.y)
			landMass[land.x][land.y] = land
			img.DrawRectangle(float64(land.x*sqSize), float64(land.y*sqSize), float64(sqSize), float64(sqSize))
			img.SetRGB(float64(land.x), float64(land.y), float64(land.z))
			img.Fill()
		}
	}

	img.SavePNG("terrain.png")
}

func assignEasy(x int, y int, land Land) Land {
	land.x = x
	land.y = y

	return land
}

func assignHard(mass [][]Land, land Land) int {
	var calc int
	top, left := false, false

	//	fmt.Println(land.x, land.y)
	if land.y == 0 {
		left = true
	}
	if land.x == 0 {
		top = true
	}

	//	fmt.Println("assHard: ", left, top)

	switch {
	case top && left:
		calc = zStart
	case top && !left:
		calc = determineHeight(mass[0][land.y-1].z)
	case !top && left:
		calc = determineHeight(mass[land.x-1][0].z)
	case !top && !left:
		calc = determineHeight(mass[land.x-1][land.y-1].z)
	default:
		calc = 0
	}

	return calc
}

func determineHeight(a int) int {
	var calc int

	//	fmt.Print("detHigt: ")
	if halfsies() {
		calc = a
		//		fmt.Print("true")
	} else {
		calc = a + plusMinus(a)
		//		fmt.Print("false")
	}
	//	fmt.Print("\n")

	return calc
}

func plusMinus(a int) int {
	b := random.Intn(3) + 1
	//	fmt.Println("pM:", b)
	switch b {
	case 1:
		return -1
	case 2:
		return 0
	}
	return 1
}

func halfsies() bool {
	if (random.Intn(2)+1)%2 == 0 {
		return true
	}
	return false
}
