package main

// facing is +/- for each of x, y, z.
type facing int

const (
	forward  facing = iota // -z
	backward               // +z
	north                  // +y
	south                  // -y
	east                   // +x
	west                   // -x
)

var allFacings = [...]facing{
	forward, backward, north, south, east, west,
}

var facingString = [...]string{
	"forward", "backward", "north", "south", "east", "west",
}

func (f facing) String() string {
	return facingString[f]
}

// rotation is around the z axis.
type rotation int

const (
	zero         rotation = iota // 0 degrees
	oneQuarter                   // 90 degrees
	twoQuarter                   // 180 degrees
	threeQuarter                 // 270 degrees
)

var allRotations = [...]rotation{
	zero, oneQuarter, twoQuarter, threeQuarter,
}

var rotationStrings = [...]string{
	"zero", "oneQuarter", "twoQuarter", "threeQuarter",
}

func (r rotation) String() string {
	return rotationStrings[r]
}
