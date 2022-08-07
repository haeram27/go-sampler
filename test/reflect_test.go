package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCheckTypeOfVariable(t *testing.T) {

	varString := "varString"

	fmt.Printf("%T\n", varString) // string

	// reflect.TypeOf() returns reflect.Type of interface represent go types
	// reflect.TypeOf().Kind() returns reflect.Kind of const int of type

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

func TestDeepEqual(t *testing.T) {
	map_1 := map[int]string{

		200: "Anita",
		201: "Neha",
		203: "Suman",
		204: "Robin",
		205: "Rohit",
	}

	map_2 := map[int]string{
		200: "Anita",
		201: "Neha",
		203: "Suman",
		204: "Robin",
		205: "Rohit",
		206: "Sumit",
	}

	map_3 := map[int]string{ // same with map_1
		200: "Anita",
		201: "Neha",
		203: "Suman",
		204: "Robin",
		205: "Rohit",
	}

	map_4 := map[string]int{
		"Anita": 200,
		"Neha":  201,
		"Suman": 203,
		"Robin": 204,
		"Rohit": 205,
	}

	// Comparing maps
	// Using DeepEqual() function
	t.Log(reflect.DeepEqual(map_1, map_2)) // false
	t.Log(reflect.DeepEqual(map_1, map_3)) // true
	t.Log(reflect.DeepEqual(map_1, map_4)) // false
	t.Log(reflect.DeepEqual(map_2, map_3)) // false
	t.Log(reflect.DeepEqual(map_3, map_4)) // false
	t.Log(reflect.DeepEqual(map_4, map_4)) // true
	t.Log(reflect.DeepEqual(map_2, map_4)) // false
}
