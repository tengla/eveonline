package main

import (
	"fmt"
	"log"

	"github.com/tengla/drdropin/eveapi"
)

func main() {
	err := eveapi.ReadConfig("./config.yml")
	if err != nil {
		log.Fatal(err)
	}
	orders, err := eveapi.GetOrders()
	if err != nil {
		log.Fatal(err)
	}

	universeNames, err := eveapi.GetUniverseNames(orders.UniqIds())

	if err != nil {
		log.Fatal(err)
	}

	for _, sorted := range universeNames.LexSortByName() {
		fmt.Println(sorted.PrettyPrint())
	}
}
