/*
	Go has good support for url parsing. URL contains a scheme, authentication info, host, port,
    path, query params, and query fragment. we can parse URL and deduce what are the parameter is
	coming to the server and then process the request accordingly.
	The net/url package has the required functions like Scheme, User, Host, Path, RawQuery etc.
*/

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	p := fmt.Println

	s := "Mysql://admin:password@serverhost.com:8081/server/page1?key=value&key2=value2#X"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	p(u.Scheme)          //prints Schema of the URL
	p(u.User)            // prints the parsed user and password
	p(u.User.Username()) //prints user's name
	pass, _ := u.User.Password()
	p(pass)                                    //prints user password
	p(u.Host)                                  //prints host and port
	host, port, _ := net.SplitHostPort(u.Host) //splits host name and port
	p(host)                                    //prints host
	p(port)                                    //prints port
	p(u.Path)                                  //prints the path
	p(u.Fragment)                              //prints fragment path value
	p(u.RawQuery)                              //prints query param name and value as provided
	m, _ := url.ParseQuery(u.RawQuery)         //parse query param into map
	p(m)                                       //prints param map
	p(m["key2"][0])                            //prints specific key value
}
