package models

import "time"

type OrderDB struct {
	Orders map[string]Order // key is order_id
}

// NewOrderDB return an initialized order db.
func NewOrderDB() *OrderDB {
	return &OrderDB{
		Orders: map[string]Order{},
	}
}

// Delete an order.
func (d *OrderDB) Delete(id string) {
	delete(d.Orders, id)
}

// Add an order.
func (d *OrderDB) Add(id string, loc Location, expiry int) {
	history := Location{
		Lat: loc.Lat,
		Lng: loc.Lng,
	}

	if expiry > 0 {
		t := time.Now().Add(time.Second * time.Duration(expiry))
		history.Expiry = &t
	}

	newOrder := Order{
		ID: id,
		History: []Location{
			history,
		},
	}

	// if there is an existing order, add the lat long.
	// else, add the new order.
	if o, found := d.Orders[id]; found {
		o.History = append(o.History, history)
		d.Orders[id] = o
		return
	}

	d.Orders[id] = newOrder
}

// id: order id
// max : max number of history to return
//
// TODO:
//	handle negative max number
func (d *OrderDB) List(id string, max int) Order {

	if o, found := d.Orders[id]; found {
		if max == 0 {
			return o
		}

		prunedOrder := pruneExpiredLocations(o)
		d.Orders[o.ID] = prunedOrder

		// get the last N items indicated by max
		newList := truncateToMax(o.History, max)
		newOrder := Order{
			ID:      o.ID,
			History: newList,
		}

		return newOrder
	}

	return Order{}
}

// truncate to determined max size
func truncateToMax(list []Location, max int) []Location {

	// Make sure we are working with non-expired items.
	list = pruneLocations(list)

	l := len(list)
	if l <= max {
		return list
	}

	return list[l-max:]
}

// Removes the expired locations from an order

func pruneExpiredLocations(o Order) Order {
	pruned := pruneLocations(o.History)
	o.History = pruned

	return o
}

func pruneLocations(list []Location) []Location {
	truncated := []Location{}

	for _, h := range list {
		if h.Expiry != nil && h.Expiry.Before(time.Now()) {
			continue
		}

		truncated = append(truncated, h)
	}

	return truncated
}
