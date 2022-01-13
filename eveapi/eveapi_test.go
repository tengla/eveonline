package eveapi

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestGetUniverseNames(t *testing.T) {
	err := ReadConfig(os.Getenv("CONFIG"))
	if err != nil {
		t.Error(err)
	}
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

func TestConfig(t *testing.T) {
	err := ReadConfig(os.Getenv("CONFIG"))
	if err != nil {
		t.Error(err)
	}
	if len(cfg.Urls.OrdersUrl) < 1 {
		t.Errorf("Expected cfg.Urls.OrdersUrl, was '%s'", cfg.Urls.OrdersUrl)
	}
	if len(cfg.Urls.LookupUrl) < 1 {
		t.Errorf("Expected cfg.Urls.OrdersUrl, was '%s'", cfg.Urls.LookupUrl)
	}
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func BenchmarkLexSortByName(b *testing.B) {
	list := UniverseNameList{}
	for i := 0; i < 2000; i++ {
		list = append(list, UniverseName{
			Name:     randomString(64),
			ID:       i,
			Category: "",
		})
	}
	b.ResetTimer()
	list.LexSortByName()
}
