package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

//go:generate bash -c "go run tools/get-codes.go tools/airports.csv codes.go && goimports -w codes.go"

func (a airport) String() string {
	return fmt.Sprintf(
		"%s (%s) - %s (%s, %s)",
		a.iata,
		a.icao,
		a.name,
		a.city,
		a.country,
	)
}

func main() {
	for _, code := range os.Args[1:] {
		code = strings.ToUpper(code)

		i := sort.Search(len(airports), func(i int) bool {
			return airports[i].iata >= code
		})

		var match airport

		if i < len(airports) && airports[i].iata == code {
			match = airports[i]
		} else {
			for _, a := range airports {
				if a.icao == code {
					match = a
					break
				}
			}
		}

		if match.iata != "" {
			fmt.Println(match)
		} else {
			fmt.Printf("%s - not found\n", code)
		}
	}
}
