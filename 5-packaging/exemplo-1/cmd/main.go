package main

import (
	"fmt"
	"github.com/anap7/goexpert-toolbox/math"
)

func main() {
	m := math.NewMath(30, 50)
	fmt.Println(m.Add())
	fmt.Println(math.X)
}