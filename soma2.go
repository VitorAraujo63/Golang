package main

func Soma() func () int {
	count := 0
	incrementa := func() int{
		count++
		return count 
	}
	return incrementa
}

func main()  {
	contador := Soma()
	print(contador())
	print(contador())
	print(contador())
	print(contador())
}