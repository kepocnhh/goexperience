package main

import (
	"fmt"
	"math"
	"time"
	"io"
	//"os"
	"strings"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v Vertex) Smethod() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vertex) Spointermethod() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func Sfunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Reverse() float64 {
	return float64(-f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func Scalefunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Squarer interface {
	Square() (float64, error)
}

func (f MyFloat) Square() (float64, error) {
	return float64(f*f), nil
}
func (v *Vertex) Square() (float64, error) {
	if v == nil {
		//fmt.Println("Vertex Square <nil>")
		return -1, &MyError{
			time.Now(),
			"Vertex Square <nil>",
		}
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y), nil
}

func describe(s Squarer) {
	fmt.Printf("(%v, %T)\n", s, s)
}

func interfaces() {
	fmt.Println("\n\tinterfaces")
	var i interface{}
	fmt.Printf("interface (%v, %T)\n", i, i)
	i = 42
	fmt.Printf("interface (%v, %T)\n", i, i)
	i = "hello"
	fmt.Printf("interface (%v, %T)\n", i, i)
}

func typeassertions() {
	fmt.Println("\n\ttype assertions")
	var i interface{} = "hello"
	s := i.(string)
	//s := string(i)//cannot convert i (type interface {}) to type string: need type assertion
	fmt.Println(s)
	s, ok := i.(string)//If i holds a T, then t will be the underlying value and ok will be true.
	fmt.Println(s, ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)
	//f = i.(float64)//panic
	//fmt.Println(f)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func typeswitches() {
	fmt.Println("\n\ttype switches")
	do(21)
	do("hello")
	do(true)
}

type Person struct {
	Name string
	Age  int
}
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type MyError struct {
	When time.Time
	What string
}
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func readers() {
	fmt.Println("\n\treaders")
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// (v*) p:= v -> &p
// (v*) p:=&v -> *p
func main() {
	fmt.Println("\n\t--- begin ---\n")

	v := Vertex{3, 4}
	fmt.Println(v.Smethod())
	fmt.Println(Sfunc(v))

	flt := MyFloat(-5)
	fmt.Println(flt.Reverse())

	v.Scale(10)
	fmt.Println(v.Smethod())
	Scalefunc(&v, 10)
	fmt.Println(v.Smethod())

	p := &Vertex{7, 9}
	//p := &v
	Scalefunc(p, 8)
	fmt.Println("p", p.Smethod())
	fmt.Println(Sfunc(*p))
	fmt.Println("v", v.Smethod())
	fmt.Println("v pointer", v.Spointermethod())

	var squarer Squarer
	describe(squarer)
	var n MyFloat
	squarer = n
	describe(squarer)
	var vn *Vertex
	squarer = vn
	describe(squarer)
	tmp1, err1 := squarer.Square()
	if err1 != nil {
		fmt.Println("Square nil", err1)
	} else {
		fmt.Println("Square nil", tmp1)
	}
	squarer = flt//MyFloat implements Squarer
	describe(squarer)
	squarer = &v//*Vertex implements Squarer
	describe(squarer)
	//squarer = v//v is a Vertex (not *Vertex) and does NOT implement Squarer
	tmp2, err2 := squarer.Square()
	if err2 != nil {
		fmt.Println("Square", err2)
	} else {
		fmt.Println("Square", tmp2)
	}
	interfaces()
	typeassertions()
	typeswitches()

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println("\n\tStringer:\n", a, "\n", z)

	fmt.Println("\n\tError check")
	if err := run(); err != nil {
		fmt.Println(err)
	}

	readers()

	fmt.Println("\n\t--- end ---")
}
