package main

import (
	"fmt"
	"time"
	"math/cmplx"
	"math/rand"
	"math"
)

func add(x int, y int) int {
	return x + y
}
func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var c2, python2, java2 = true, false, "no!"
var c3, python3, java3 int = 1, 2, 4

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("\n\t--- begin ---\n")

	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)

	fmt.Println("add", add(42, 13))
	fmt.Println("add2", add2(42, 13))

	a, b := swap("hello", "world")
	fmt.Println("swap", a, b)

	fmt.Println(split(17))
	//fmt.Println("split", split(17))//multiple-value split() in single-value context

	var cfunclevel int
	fmt.Println(cfunclevel, c, python, java)
	fmt.Println(c2, python2, java2)
	fmt.Println(c3, python3, java3)
	//only function level
	c4, python4, java4 := true, false, "no!"
	fmt.Println(c4, python4, java4)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var intzero int
	var float64zero float64
	var boolzero bool
	var stringzero string
	fmt.Printf("%v %v %v %q\n", intzero, float64zero, boolzero, stringzero)

	/*
	var intjust int = 42
	var floatfromint float64 = float64(intjust)
	var uintfromfloat uint = uint(floatfromint)
	*/
	intjust := 42
	floatfromint := float64(intjust)
	uintfromfloat := uint(floatfromint)
	fmt.Println(intjust, floatfromint, uintfromfloat)

	//value_without_specifying_type := 42
	//value_without_specifying_type := 3.142
	value_without_specifying_type := 0.867 + 0.5
	fmt.Printf("value is of type %T\n", value_without_specifying_type)

	const Truth = true
	fmt.Println("Happy", Pi, "Day")
	fmt.Println("Go rules?", Truth)

	fmt.Println("int small", needInt(Small))
	fmt.Println("float small", needFloat(Small))
	fmt.Println("float big", needFloat(Big))

	fmt.Println("\n\t--- end ---")
}
