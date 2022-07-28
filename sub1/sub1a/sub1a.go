package sub1a

import "fmt"

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
	res := x * x + PublicGeneric(x) + privateGeneric(x)
	fmt.Println("sub1a_2 says :", res)
	return res
}
