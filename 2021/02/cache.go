package main

import (
	"fmt"
	"time"
)

type item struct {
	ID    int
	Name  string
	Value int
}

func (i item) String() string {
	return fmt.Sprintf("%v: %s %v ", i.ID, i.Name, i.Value)
}

var items = []item{

	{
		ID:    1,
		Name:  "Second",
		Value: 2,
	},
	{
		ID:    2,
		Name:  "Banana",
		Value: 54,
	},
	{
		ID:    3,
		Name:  "Basket",
		Value: 0,
	},
	{
		ID:    4,
		Name:  "Password",
		Value: 15,
	},
	{
		ID:    5,
		Name:  "Croatia",
		Value: -1,
	},
	{
		ID:    6,
		Name:  "First",
		Value: 1,
	},
	{
		ID:    7,
		Name:  "First II",
		Value: 110,
	},
	{
		ID:    8,
		Name:  "PRVZ",
		Value: 1,
	},
}

func queryCacheFast(id int) (item, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabaseSlow(id int) (item, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, item := range items {
		if item.ID == id {
			cache[id] = item
			return item, true
		}
	}
	return item{}, false
}
