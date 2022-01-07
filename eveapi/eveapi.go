package eveapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"gopkg.in/yaml.v2"
)

// UniverseNames
type UniverseName struct {
	Category string `json:"category"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
}

// PrettyPrint UniverseName instance
func (u UniverseName) PrettyPrint() string {
	return fmt.Sprintf("%s, Id: %d, Category: %s", u.Name, u.ID,
		u.Category)
}

// Order
type Order struct {
	Duration     int       `json:"duration"`
	IsBuyOrder   bool      `json:"is_buy_order"`
	Issued       time.Time `json:"issued"`
	LocationID   int       `json:"location_id"`
	MinVolume    int       `json:"min_volume"`
	Range        string    `json:"range"`
	SystemID     int       `json:"system_id"`
	TypeID       int       `json:"type_id"`
	VolumeRemain int       `json:"volume_remain"`
	VolumeTotal  int       `json:"volume_total"`
	Price        float64   `json:"price"`
	OrderID      int       `json:"order_id"`
}

type Config struct {
	Urls struct {
		OrdersUrl string `yaml:"ordersUrl"`
		LookupUrl string `yaml:"lookupUrl"`
	} `yaml:"urls"`
}

func ReadConfig(filename string) (Config, error) {
	var c Config
	cfg, err := os.ReadFile(filename)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal(cfg, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}

// OrderList
type OrderList []Order

// UniverseNameList
type UniverseNameList []UniverseName

// FindByName
func (list UniverseNameList) FindByName(name string) (*UniverseName, error) {
	for _, universeName := range list {
		if universeName.Name == name {
			return &universeName, nil
		}
	}
	return nil, errors.New("UniverseName not found")
}

func (list UniverseNameList) LexSortByName() UniverseNameList {
	sorted := UniverseNameList{}
	names := []string{}
	for _, u := range list {
		names = append(names, u.Name)
	}
	sort.Strings(names)
	for _, name := range names {
		u, err := list.FindByName(name)
		if err != nil {
			log.Fatal(err)
		}
		sorted = append(sorted, *u)
	}
	return sorted
}

// UniqIds
func (list OrderList) UniqIds() []int {
	found := map[int]bool{}
	result := []int{}
	for _, u := range list {
		if !found[u.TypeID] {
			found[u.TypeID] = true
			result = append(result, u.TypeID)
		}
	}
	return result
}

// GetOrders
func GetOrders(cfg Config) (OrderList, error) {
	resp, err := http.Get(cfg.Urls.OrdersUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var orders OrderList
	json.NewDecoder(resp.Body).Decode(&orders)
	return orders, nil
}

// GetUniverseNames
func GetUniverseNames(cfg Config, ids []int) (UniverseNameList, error) {
	data, _ := json.Marshal(ids)
	body := bytes.NewBuffer(data)
	resp, err := http.Post(cfg.Urls.LookupUrl, "application/json", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		var result interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		fmt.Printf("%s %d\n", resp.Status, resp.StatusCode)
		errMsg := fmt.Sprintf("%s %d\n%s", resp.Status,
			resp.StatusCode, result.(map[string]interface{})["error"])
		return nil, errors.New(errMsg)
	} else {
		var universeNameList UniverseNameList
		json.NewDecoder(resp.Body).Decode(&universeNameList)
		return universeNameList, nil
	}
}
