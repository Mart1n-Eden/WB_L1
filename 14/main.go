package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{}
	x = 1
	typeDefinition(x)

	x = "qwerty"
	typeDefinition(x)

	x = true
	typeDefinition(x)

	x = make(chan any)
	typeDefinition(x)
}

func typeDefinition(x interface{}) {
	fmt.Println(reflect.TypeOf(x))
}


// func main() {
// 	var x interface{}
// 	x = 1
// 	typeDefinition(x)

// 	x = "qwerty"
// 	typeDefinition(x)

// 	x = true
// 	typeDefinition(x)

// 	x = make(chan any)
// 	typeDefinition(x)
// }

// func typeDefinition(x interface{}) {
// 	switch x.(type) {
// 	case int:
// 		fmt.Println("int")
// 	case string:
// 		fmt.Println("string")
// 	case bool:
// 		fmt.Println("bool")
// 	case chan any:
// 		fmt.Println("chan")
// 	}
// }
