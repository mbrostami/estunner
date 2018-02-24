package debug

import (
	"fmt"
	"reflect"
)

func ListMethods(test interface{}) {
	fooType := reflect.TypeOf(test)
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		fmt.Println(method.Name)
	}
}
