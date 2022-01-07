package eveapi

import (
	"os"
	"testing"
)

func TestGetUniverseNames(t *testing.T) {
	cfg, err := ReadConfig(os.Getenv("CONFIG"))
	if err != nil {
		t.Error(err)
	}
	orders, _ := GetOrders(cfg)
	length := len(orders)
	if length < 100 {
		t.Errorf("GetOrders = %d, wanted > 100", length)
	}
	names, _ := GetUniverseNames(cfg, orders.UniqIds())
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
	cfg, err := ReadConfig(os.Getenv("CONFIG"))
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
