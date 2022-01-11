package models

import (
	"testing"
)

func TestAdd(t *testing.T) {

	db := NewOrderDB()

	db.Add("1", Location{Lat: 33.4, Lng: 11.22}, -1)
	db.Add("1", Location{Lat: 99.9, Lng: 11.55}, -1)
	db.Add("2", Location{Lat: 33.4, Lng: 11.22}, -1)

	if len(db.Orders) != 2 {
		t.Fatalf("expected 2 order, got %d", len(db.Orders))
	}

	order := db.List("1", 0)

	if len(order.History) != 2 {
		t.Fatal("failed to list correctly")
	}

}

func TestAddAllUnique(t *testing.T) {

	db := NewOrderDB()

	db.Add("1", Location{Lat: 33.4, Lng: 11.22}, -1)
	db.Add("2", Location{Lat: 99.9, Lng: 11.55}, -1)
	db.Add("3", Location{Lat: 33.4, Lng: 11.22}, -1)

	if len(db.Orders) != 3 {
		t.Fatalf("expected 3 order, got %d", len(db.Orders))
	}

	order := db.List("1", 0)

	if len(order.History) != 1 {
		t.Fatal("failed to list correctly")
	}

	db.Delete("2")

	if len(db.Orders) != 2 {
		t.Fatal("failed to delete order")
	}
}
