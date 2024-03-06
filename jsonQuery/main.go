package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/jsonquery"
)

func main() {
	s := `{
            "person":{
               "name":"John",
               "age":31,
               "female":false,
               "city":null,
               "hobbies":[
                  "coding",
                  "eating",
                  "football"
               ]
            }
         }`
	doc, err := jsonquery.Parse(strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	// xpath query
	age := jsonquery.FindOne(doc, "age")
	// or
	age = jsonquery.FindOne(doc, "person/age")
	fmt.Printf("%#v[%T]\n", age.Value(), age.Value()) // prints 31[float64]

	hobbies := jsonquery.FindOne(doc, "//hobbies")
	fmt.Printf("%#v\n", hobbies.Value()) // prints []interface {}{"coding", "eating", "football"}
	firstHobby := jsonquery.FindOne(doc, "//hobbies/*[2]")
	fmt.Printf("%#v\n", firstHobby.Value()) // "coding"

}
