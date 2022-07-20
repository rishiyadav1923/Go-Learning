/*
	Go has built-in support for JSON encoding and decoding. it also supports custom datatypes.
	The Marshal function is used to convert go data types into JSON format.

	Marshal function syntax is:

    "func Marshal(v interface{}) ([]byte, error)"

	Marshal returns the JSON encoding of v.
	Boolean is converted to JSON booleans. Floating point, integer, and Number are converted to JSON numbers. The map's key type must either be a string, an integer type, or implement encoding.TextMarshaler.
	The decoding of JSON is done using Unmarshal function.

	Unmarshal function syntax is:

	"func Unmarshal(data []byte, v interface{}) error"

	Unmarshal decodes JSON-encoded value and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
	We can also customize the fields stored under the "json" key in the struct field's tag. we can have name of the field, followed by a comma-separated list of options. Like

    Field int 'json:"myName"' // The appears in JSON as key "myName".
    Field int 'json:"myName,omitempty?'// The field is omitted from the object if its value is empty,
    Field int 'json:"-"' // Field is ignored by this package.
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	bolType, _ := json.Marshal(false) //boolean Value
	fmt.Println(string(bolType))
	intType, _ := json.Marshal(10) // integer value
	fmt.Println(string(intType))
	fltType, _ := json.Marshal(3.14) //float value
	fmt.Println(string(fltType))
	strType, _ := json.Marshal("JavaTpoint") // string value
	fmt.Println(string(strType))
	slcA := []string{"sun", "moon", "star"} //slice value
	slcB, _ := json.Marshal(slcA)
	fmt.Println(string(slcB))
	mapA := map[string]int{"sun": 1, "moon": 2} //map value
	mapB, _ := json.Marshal(mapA)
	fmt.Println(string(mapB))
}
