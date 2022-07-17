package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCheckTypeOfVariable(t *testing.T) {

	varString := "varString"

	fmt.Printf("%T\n", varString) // string

	if reflect.TypeOf(varString).Kind() == reflect.String {
		fmt.Println("Equal")                                   // Equal
		fmt.Println(reflect.TypeOf(varString).Name())          // string
		fmt.Println(reflect.TypeOf(varString).String())        // string
		fmt.Println(reflect.TypeOf(varString).Kind().String()) // string
	}

	if reflect.ValueOf(varString).Kind() == reflect.String {
		fmt.Println("Equal") // Equal
	}

	var itfs interface{}
	itfs = varString
	switch t := itfs.(type) {
	case string:
		fmt.Println("string") // string
	default:
		fmt.Printf("type unknown %T\n", t)
	}
}
