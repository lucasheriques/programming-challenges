package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

type Bill struct {
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, exists := units[unit]

	if !exists {
		return false
	}

	bill[item] = bill[item] + unitValue

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	qtFromBill, itemExists := bill[item]
	unitValue, unitExists := units[unit]

	if !itemExists || !unitExists || qtFromBill-unitValue < 0 {
		return false
	}

	if qtFromBill-unitValue == 0 {
		delete(bill, item)
	} else {
		bill[item] -= unitValue
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, exists := bill[item]

	return qty, exists
}
