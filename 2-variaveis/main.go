// main entry point
package main

const grettings = "Hello, "
var (
	b bool
	c int
	d string
	e float64
	f = "Kleber"
)

func main() {
	var nome = "Kleber"
	b = true
	shorthand := "shorthand"
	println(grettings + nome)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
	println(shorthand)
}