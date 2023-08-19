package main

type animal interface{
	Sound() string
	Move() string
}

type Cachorro struct {
	Name string
}

func (d Cachorro) Sound() string {
	return d.Name + " Say Au, Au"
}

func (d Cachorro) Move() string {
	return d.Name + " Is running"
}

func MakeMove(a animal) {
	print(a.Move())
	print(a.Sound())
}

func main(){
  cachorro := Cachorro{
    Name: "Cachorro",
  }

  MakeMove(cachorro)
}