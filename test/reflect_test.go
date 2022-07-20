package test

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

	var itf interface{}
	itf = varString
	switch t := itf.(type) {
	case string:
		fmt.Println("string") // string
	default:
		fmt.Printf("type unknown %T\n", t)
	}
}
