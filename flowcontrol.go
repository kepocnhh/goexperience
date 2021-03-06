package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func newtonsqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++{
		z -= (z*z - x)/(2*z)
	}
	return z
}

func defercall() {
	defer fmt.Println("call")
	fmt.Print("defer ")
}
func defercounting() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main() {
	fmt.Println("\n\t--- begin ---\n")

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1//The init and post statement are optional.
	for ; sum2 < 1000; {
		sum2 += sum
	}
	fmt.Println(sum2)

	sum3 := 1
	for sum3 < 10000 {
		sum3 += sum2
	}
	fmt.Println(sum3)

	sum4 := 1
	for {
		if sum4 > 100000 {
			break
		}
		sum4 += sum3
	}
	fmt.Println(sum4)

	fmt.Println("sqrt(2)", sqrt(2), "sqrt(-4)", sqrt(-4))

	fmt.Println(
		"pow(3, 2, 10)", pow(3, 2, 10),
		"pow(3, 3, 20)", pow(3, 3, 20),
	)

	fmt.Println("newtonsqrt of 576", newtonsqrt(576))

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	timeNow := time.Now()
	switch {
	case timeNow.Hour() < 12:
		fmt.Println("Good morning!")
	case timeNow.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	defercall()
	defercounting()

	fmt.Println("\n\t--- end ---")
}
