package main

import (
	"fmt"
	"reflect"
	// "github.com/fatih/structs"
)

type Model interface {
	Default() string
}

type El struct {
	Message string
}

func (e El) Default() string {
	return "default"
}

type Al struct {
	val int
}

func main() {
	// element := El{"mon", 
	// 	&El{
	// 		Message:"yo", 
	// 	},
	// }

	// name := structs.Name(element)
	// mapped := structs.Map(element)
	// names := structs.Names(element)
	// values := structs.Values(element)

	// fmt.Print("name ")
	// fmt.Println(name)

	// fmt.Print("names ")
	// fmt.Println(names)

	// fmt.Print("values ")
	// fmt.Println(values)

	// fmt.Print("mapped ")
	// fmt.Println(mapped)

	// refVal := reflect.ValueOf(element)
	// refKind := refVal.Kind()
	// refFBN := refVal.FieldByName("Time")

	// fmt.Print("refVal ")
	// fmt.Println(refVal)

	// fmt.Print("refKind ")
	// fmt.Println(refKind)

	// fmt.Print("refFBN ")
	// fmt.Println(refFBN)
	a := "coucou"       // default
	b := 3              // default
	c := []int{3,5,8}   // list 
	d := El{"blabla"}   // model
	e := Al{42}         // struct
	f := []El{d}        // list
	switchPrint(a)
	switchPrint(b)
	switchPrint(c)
	switchPrint(d)
	switchPrint(e)
	switchPrint(f)
}

func switchPrint(el any) {
	v := reflect.ValueOf(el)
	switch v.Kind() {	
	case reflect.Slice :
		fmt.Println(v, v.Kind(), "list")
	case reflect.Map :
		fmt.Println(v, v.Kind(), "map")
	case reflect.Struct:
		switch el.(type) {
		case Model:
			fmt.Println(v, v.Kind(), "model")
		default:
			fmt.Println(v, v.Kind(), "struct")
		}
	default:
		fmt.Println(v, v.Kind(), "default")
	}
}