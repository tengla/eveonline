package main

import (
	"fmt"
	"log"

	"github.com/tengla/drdropin/eveapi"
)

func main() {
	cfg, err := eveapi.ReadConfig("./config.yml")
	if err != nil {
		log.Fatal(err)
	}

	orders, err := eveapi.GetOrders(cfg)
	if err != nil {
		log.Fatal(err)
	}

	universeNames, err := eveapi.GetUniverseNames(
		cfg, orders.UniqIds())

	if err != nil {
		log.Fatal(err)
	}

	sorted := universeNames.LexSortByName()
	for _, u := range sorted {
		fmt.Println(u.PrettyPrint())
	}
}
