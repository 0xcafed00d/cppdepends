package main

import "github.com/simulatedsimian/cppdepends/cppdep"
import "fmt"

func main() {
	fmt.Println(cppdep.ListFiles("..", []string{".h", ".md"}))
}
