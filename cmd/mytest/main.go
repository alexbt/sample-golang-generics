package main

import (
	"fmt"
	"github.com/alexbt/sample-golang/cmd/mytest/types"
)

func main() {
	fmt.Println("First's Get with int:", Get[int](types.First{10}))
	fmt.Println("Second's Get with string:", Get[string](types.Second{"value"}))
	fmt.Println()

	fmt.Println("First's GetAny with int:", GetAny[int](types.First{666}))
	fmt.Println("Second's GetAny with string:", GetAny[string](types.Second{"any value"}))
	fmt.Println()

	fmt.Println("First's PrintIntOrStringElements with int:")
	PrintIntOrStringElements[int]([]types.First{types.First{10}, types.First{11}})
	fmt.Println()

	fmt.Println("Second's PrintIntOrStringElements with string:")
	PrintIntOrStringElements[string]([]types.Second{types.Second{"value1"}, types.Second{"value2"}})
	fmt.Println()
	fmt.Println("First's PrintAnyElements with int:")
	PrintAnyElements[int]([]types.First{types.First{10}, types.First{11}})
	fmt.Println()
	fmt.Println("Second's PrintAnyElements with string:")
	PrintAnyElements[string]([]types.Second{types.Second{"valueAny1"}, types.Second{"valueAny2"}})
}

func Get[V int | string, K types.MyTypeReturnsIntOrString[V]](obj K) V {
	return obj.GetIntOrString()
}

func GetAny[V any, K types.MyTypeReturnsAny[V]](obj K) V {
	return obj.GetAny()
}

func PrintAnyElements[V any, K types.MyTypeReturnsAny[V]](list []K) string {
	for k, v := range list {
		value := v.GetAny()
		fmt.Println(" - elem", k+1, ":", value)
	}
	return "ok"
}

func PrintIntOrStringElements[V int | string, K types.MyTypeReturnsIntOrString[V]](list []K) string {
	for k, v := range list {
		value := v.GetIntOrString()
		fmt.Println(" - elem", k+1, ":", value)
	}
	return "ok"
}
