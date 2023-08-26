package main

func main() {

	a := 1
	//b := 2
	//c := 3

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("d")
	}

	//if a > b && c > a {
	//	println("a > b && c > a")
	//}

	// go env | go env GOOS GOARCH
	// go build main.go
	// GOOS=windows GOARCH=amd64 go build main.go
	// go mod init meu-modulo -> go build

}
