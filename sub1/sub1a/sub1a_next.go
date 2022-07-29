package sub1a

import (
	"fmt"
	"net/http"
)

func Sub1a_3(x int) int {
	http.ListenAndServe("", nil)
	res := x * x
	fmt.Println("sub1a_3 says :", res)
	return res
}

func Sub1a_4(x int) int {
	res := x * x
	fmt.Println("sub1a_4 says :", res)
	return res
}
