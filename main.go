package main

import (
	"fmt"

)


func main()  {                           // for å komme i gang fikk jeg informasjon fra: https://go.dev/doc/tutorial/getting-started
	fmt.Println(myquote.First())        // for å printe alt til skjermen sjekket jeg: https://pkg.go.dev/rsc.io/quote/v4
	fmt.Println(myquote.Second())
	fmt.Println(myquote.Third())
	fmt.Println(myquote.Fourth())
}