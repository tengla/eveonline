package eveapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const ordersUrl = "https://esi.evetech.net/latest/markets/10000002/orders/?datasource=tranquility&order_type=all&page=1"
const lookupUrl = "https://esi.evetech.net/latest/universe/names/?datasource=tranquility"

// UniverseNames
type UniverseName struct {
	Category string `json:"category"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
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

// OrderList
type OrderList []Order

// UniverseNameList
type UniverseNameList []UniverseName

func (list UniverseNameList) FindByName(name string) (UniverseName, error) {
	for _, universeName := range list {
		if universeName.Name == name {
			return universeName, nil
		}
	}
	var null UniverseName
	return null, errors.New("UniverseName not found")
}

// PrettyPrintOrder
func (o Order) PrettyPrintOrder() string {
	props := []string{
		fmt.Sprintf("Duration=%d", o.Duration),
		fmt.Sprintf("IsBuyOrder=%t", o.IsBuyOrder),
		fmt.Sprintf("Issued=%s", o.Issued),
		fmt.Sprintf("LocationID=%d", o.LocationID),
		fmt.Sprintf("MinVolume=%d", o.MinVolume),
		fmt.Sprintf("Range=%s", o.Range),
		fmt.Sprintf("SystemId=%d", o.SystemID),
		fmt.Sprintf("TypeID=%d", o.TypeID),
		fmt.Sprintf("VolumeRemain=%d", o.VolumeRemain),
		fmt.Sprintf("VolumeTotal=%d", o.VolumeTotal),
		fmt.Sprintf("Price=%.2f", o.Price),
		fmt.Sprintf("OrderID=%d", o.OrderID),
	}
	return fmt.Sprintf("Order(%s)", strings.Join(props, ",\n    "))
}

// GetOrders
func GetOrders() (OrderList, error) {
	resp, err := http.Get(ordersUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var orders OrderList
	json.NewDecoder(resp.Body).Decode(&orders)
	return orders, nil
}

// GetUniverseNames
func GetUniverseNames(ids []int) (UniverseNameList, error) {
	data, _ := json.Marshal(ids)
	body := bytes.NewBuffer(data)
	resp, err := http.Post(lookupUrl, "application/json", body)
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
