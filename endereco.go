package main

func main() {
	x := 10
	b := &x

	print (&x)
	print ("X:", x)

	println (&b)
	println ("B:", *b)

  *b = 55
}