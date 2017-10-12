package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

//go:generate bash -c "go run tools/get-codes.go tools/airport-codes.csv codes.go && goimports -w codes.go"

func main() {
	for _, code := range os.Args[1:] {
		code = strings.ToUpper(code)

		i := sort.Search(len(airports), func(i int) bool {
			return airports[i].iata >= code
		})

		if i < len(airports) && airports[i].iata == code {
			fmt.Printf("%s - %s (%s, %s)\n", code, airports[i].name, airports[i].city, airports[i].country)
		} else {
			fmt.Printf("%s - not found\n", code)
		}
	}
}
