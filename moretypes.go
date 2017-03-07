package main

import (
	"fmt"
	//"strings"
	"math"
)

type Vertex struct {
	X int
	Y int
}

func pointers() {
	fmt.Println("\n\tpointers")
	i, j := 42, 2701
	fmt.Println("i", i, "j", j)
	p := &i         // point to i
	fmt.Println("p : = &i")
	fmt.Println("*p", *p) // read i through the pointer
	fmt.Println("p", p)
	*p = 21         // set i through the pointer
	fmt.Println("*p = 21")
	fmt.Println("i", i)  // see the new value of i

	p = &j         // point to j
	fmt.Println("p = &j")
	*p = *p / 37   // divide j through the pointer
	fmt.Println("*p = *p / 37")
	fmt.Println("j", j) // see the new value of j
}

func structs() {
	fmt.Println("\n\tstructs")
	fmt.Println(Vertex{1, 2})
	v := Vertex{3, 4}
	v.X = 5
	fmt.Println(v)
	p := &v
	p.X = 1e9
	fmt.Println(v)
	v2 := Vertex{X: 6}  // Y:0 is implicit
	v3 := Vertex{}      // X:0 and Y:0
	p2 := &Vertex{7, 8} // has type *Vertex
	fmt.Println(p2, v2, v3)
}

func arrays() {
	fmt.Println("\n\tarrays")
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	//var s []int = primes[1:4]
	s := primes[1:4]
	fmt.Println(s)
}

func slices() {
	fmt.Println("\n\tslices")
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
func slicesliterals() {
	fmt.Println("\n\tslices literals")
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	fmt.Println(s[0].i)
}
func slicesdefaults() {
	fmt.Println("\n\tslices defaults")
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)
	s = s[1:4]
	fmt.Println("[1:4]", s)
	s = s[:2]
	fmt.Println("[:2]", s)
	s = s[1:]
	fmt.Println("[1:]", s)
	s = s[:]
	fmt.Println("[:]", s)
	/*
	For the array
		var a [10]int
	these slice expressions are equivalent:
		a[0:10]
		a[:10]
		a[0:]
		a[:]
	*/
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func sliceslengthcapacity() {
	fmt.Println("\n\tslices length capacity")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	s = s[:0]//Slice the slice to give it zero length.
	fmt.Println("[:0]")
	printSlice(s)
	s = s[:4]//Extend its length.
	fmt.Println("[:4]")
	printSlice(s)
	s = s[2:]//Drop its first two values.
	fmt.Println("[2:]")
	printSlice(s)
}
func slicesnil() {
	fmt.Println("\n\tslices nil")
	var snil []int
	fmt.Println(snil, len(snil), cap(snil))
	if snil == nil {
		fmt.Println("nil!")
	}
}

func slicesmake() {
	fmt.Println("\n\tslices make")
	a := make([]int, 5)
	fmt.Println("make([]int, 5)")
	printSlice(a)
	b := make([]int, 0, 5)
	fmt.Println("make([]int, 0, 5)")
	printSlice(b)
	c := b[:2]
	fmt.Println("[:2]")
	printSlice(c)
	d := c[2:5]
	fmt.Println("[2:5]")
	printSlice(d)
}
func slicesofslices() {
	fmt.Println("\n\tslices of slices")
	// Create a tic-tac-toe board.
	/*
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	*/
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		//fmt.Printf("%s\n", strings.Join(board[i], " "))
		fmt.Println(board[i])
	}
}
func slicesappend() {
	fmt.Println("\n\tslices append")
	var s []int
	fmt.Println("[]int")
	printSlice(s)
	// append works on nil slices.
	s = append(s, 0)
	fmt.Println("append(s, 0)")
	printSlice(s)
	// The slice grows as needed.
	s = append(s, 1)
	fmt.Println("append(s, 1)")
	printSlice(s)
	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	fmt.Println("append(s, 2, 3, 4)")
	printSlice(s)
}
func slicesrange() {
	fmt.Println("\n\tslices range")
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2^i
	}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}

func maps() {
	fmt.Println("\n\tmaps")
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["a"] = Vertex{
		40, -74,
	}
	m["b"] = Vertex{
		-40, 74,
	}
	fmt.Println("m[a]", m["a"])
	fmt.Println("m[b]", m["b"])
	fmt.Println("m[c]", m["c"])
}
func mapsliterals() {
	fmt.Println("\n\tmaps literals")
	/*
	//If the top-level type is just a type name, you can omit it from the elements of the literal.
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40, -74,
		},
		"Google": Vertex{
			37, -122,
		},
	}
	*/
	m := map[string]Vertex{
		"Bell Labs": {40, -74},
		"Google":    {37, -122},
	}
	fmt.Println(m)
}
func mapsmutating() {
	fmt.Println("\n\tmaps mutating")
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func functionvalues() {
	fmt.Println("\n\tfunction values")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func functionclosures() {
	fmt.Println("\n\tfunction closures")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			"pos", pos(i),
			"neg", neg(-2*i),
		)
	}
}

func fibonacci() func() int {
    x, y := 0, 1
    return func() (r int) {
		r = x
        x, y = y, x + y
        return
    }
}

func main() {
	fmt.Println("\n\t--- begin ---\n")

	pointers()
	structs()
	arrays()
	slices()
	slicesliterals()
	slicesdefaults()
	sliceslengthcapacity()
	slicesnil()
	slicesmake()
	slicesofslices()
	slicesappend()
	slicesrange()
	maps()
	mapsliterals()
	mapsmutating()
	functionvalues()
	functionclosures()

	fmt.Println("\n\tfibonacci")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	fmt.Println("\n\t--- end ---")
}
