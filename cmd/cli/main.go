package main

import (
	"fmt"
	"log"

	"github.com/tengla/drdropin/eveapi"
)

func main() {

	orders, err := eveapi.GetOrders()
	if err != nil {
		log.Fatal(err)
	}

	universeNames, err := eveapi.GetUniverseNames(
		orders.UniqIds())

	if err != nil {
		log.Fatal(err)
	}

	sorted := universeNames.LexSortByName()
	for _, u := range sorted {
		fmt.Println(u.PrettyPrint())
	}
}
