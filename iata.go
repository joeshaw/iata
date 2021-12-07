package main

import (
	"fmt"
	"os"
	"strings"
)

//go:generate bash -c "go run tools/get-codes.go tools/airports.csv codes.go && goimports -w codes.go"

func (a airport) String() string {
	var code string
	if a.iata == "" {
		code = a.icao
	} else {
		code = fmt.Sprintf("%s (%s)", a.iata, a.icao)
	}

	return fmt.Sprintf(
		"%s - %s (%s, %s)",
		code,
		a.name,
		a.city,
		a.country,
	)
}

func main() {
	for _, code := range os.Args[1:] {
		code = strings.ToUpper(code)

		var match bool
		if a, ok := iatas[code]; ok {
			fmt.Println(a)
			match = true
		}
		if a, ok := icaos[code]; ok {
			fmt.Println(a)
			match = true
		}
		if !match {
			fmt.Printf("%s - not found\n", code)
		}
	}
}
