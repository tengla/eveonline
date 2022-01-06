package main

import (
	"fmt"
	"log"
	"sort"

	drdropin "github.com/tengla/drdropin/eveapi"
)

func uniqIds(ids []int) []int {
	found := map[int]bool{}
	result := []int{}
	for _, id := range ids {
		if !found[id] {
			found[id] = true
			result = append(result, id)
		}
	}
	return result
}
func main() {
	orders, err := drdropin.GetOrders()
	if err != nil {
		log.Fatal(err)
	}
	type_ids := make([]int, 0)
	for _, order := range orders {
		type_ids = append(type_ids, order.TypeID)
	}
	universeNames, err := drdropin.GetUniverseNames(uniqIds(type_ids))
	if err != nil {
		log.Fatal(err)
	}
	names := []string{}
	for _, universeName := range universeNames {
		names = append(names, universeName.Name)
	}

	sort.Strings(names)

	for _, name := range names {
		universeName, err := universeNames.FindByName(name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s Id: %d Category: %s\n", universeName.Name, universeName.ID,
			universeName.Category)
	}
}
