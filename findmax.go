package main

import "fmt"

type Number interface{
  ~int8 | ~int16 | ~float32 | ~float64
}

func FindMax [T Number] (arr []T) T {
  max := arr[0] 
  for _,value := range arr{
    if value > max {
      max = value
    }
  }
  return max
}

func main() {
  maxInt := FindMax([]int8{1,2,9,56,0,3,7})
  maxFloat := FindMax([]float32{0.8, 5.2, 5.3, 1.0, 7.9})

  fmt.Println(maxInt)
  fmt.Printf("%.2f", maxFloat)
}