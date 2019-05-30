package main

import (
	"fmt"
	"math"

	"github.com/abhi2109/go_crash_course/03_Package/ourpackage"
)

func main() {
	fmt.Println(math.Floor(2.7)) //2
	fmt.Println(math.Ceil(2.71)) //3
	fmt.Println(math.Sqrt(16))   //4
	ourpackage.Hello()
}
