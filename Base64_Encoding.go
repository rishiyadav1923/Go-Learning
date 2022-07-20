/*
	we can encode String and url in Go. Go has Encoder which takes byte array and convert into string encoding.
	The Decoder takes the encoded value and covert it to the original string.
*/

package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "JavaTpoint@12345!@#$%^&*()"
	strEncode := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("value to be encode  " + data)
	fmt.Println("Encoden value:  " + strEncode)

	fmt.Println()

	fmt.Print("Value to be decode  " + strEncode)
	strDecode, _ := b64.StdEncoding.DecodeString(strEncode)
	fmt.Println("Decoded value  " + string(strDecode))
	fmt.Println()

	url := "https://golang.org/ref/spec"

	fmt.Println("url to be encode  " + url)
	urlEncode := b64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println("Encoded url  " + urlEncode)

	fmt.Println("value to be decode  " + urlEncode)
	strDecode2, _ := b64.URLEncoding.DecodeString(urlEncode)

	fmt.Println("Decoded value  " + string(strDecode2))
}
