package main

type Animal interface{
  Sound() string
  Move() string
}

type Dog struct{
  Name string
}

func (d *Dog) Sound() string {
  d.Name = "Rex"
  return d.Name + " say Au Au"
}

func (d *Dog) Move() string{
  return d.Name + " is running"
}

func MakeActions (a Animal){
  println(a.Move())
  println(a.Sound())
}

func main(){
  dog := Dog{
    Name: "Buddy",
  }
  MakeActions(&dog)
  println(dog.Name)
}