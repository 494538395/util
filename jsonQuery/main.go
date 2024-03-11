package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/jsonquery"
	"github.com/antchfx/xpath"
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
	fmt.Println("fist age-->", age)

	// or
	age = jsonquery.FindOne(doc, "person/age")
	fmt.Printf("second age-->%#v[%T]\n", age.Value(), age.Value()) // prints 31[float64]

	hobbies := jsonquery.FindOne(doc, "//hobbies")
	fmt.Printf("%#v\n", hobbies.Value()) // prints []interface {}{"coding", "eating", "football"}
	for _, element := range hobbies.Value().([]interface{}) {
		fmt.Println("hobby element-->", element)
	}

	n := jsonquery.QuerySelector(doc, xpath.MustCompile("//a"))
	fmt.Println("n-->", n)

	firstHobby := jsonquery.FindOne(doc, "//hobbies/*[2]")
	fmt.Printf("%#v\n", firstHobby.Value()) // "coding"

}
