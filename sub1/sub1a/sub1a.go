package sub1a

import (
	"fmt"
	"net/http"
)

func PublicGeneric(y int) int {
	return y + 1
}

func privateGeneric(y int) int {
	return y + 10
}

func Sub1a_1(x int) int {
	res := x * x + PublicGeneric(x) + privateGeneric(x)
	fmt.Println("sub1a_1 says :", res)
	return res
}

func Sub1a_2(x int) int {
	http.ListenAndServe("", nil)
	res := x * x + PublicGeneric(x) + privateGeneric(x)
	fmt.Println("sub1a_2 says :", res)
	return res
}

type Sub1a_Type struct {
	x int
	y int
}

func (a Sub1a_Type) Method1(x int) int {
	return x * (a.x * a.y) + 1
}

func (a Sub1a_Type) Method2(x int) int {
	return x * (a.x * a.y) + 2
}

func (a Sub1a_Type) Method3(x int) int {
	return x * (a.x * a.y) + 3
}

func (a Sub1a_Type) Method4(x int) int {
	return x * (a.x * a.y) + 4
}
