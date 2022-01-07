package eveapi

import (
	"testing"
)

func TestGetUniverseNames(t *testing.T) {
	orders, _ := GetOrders()
	length := len(orders)
	if length < 100 {
		t.Errorf("GetOrders = %d, wanted > 100", length)
	}
	names, _ := GetUniverseNames(orders.UniqIds())
	sorted := names.LexSortByName()
	length = len(names)
	if length < 100 {
		t.Errorf("GetUniverseNames = %d, wanted > 100", length)
	}
	lenSorted := len(sorted)
	if lenSorted != length {
		t.Errorf("sort = %d, wanted = %d", lenSorted, length)
	}
}
