package models

type Account struct {
	id           int             `json:"id,omitempty"`
	name         string          `json:"name,omitempty"`
	currentValue int             `json:"current_value,omitempty"`
	historical   []Transaction   `json:"historical,omitempty"`
	rules        map[string]bool `json:"rules,omitempty"`
}

// AddTransaction adds a new transaction to the historic and
// updates the currentValue to reflect this change
func (a *Account) AddTransaction(transaction *Transaction) {
	a.historical = append(a.historical, *transaction)
	a.CalculateCurrentValue()
}

// CalculateCurrentValue takes all the amounts in the historic,
// adds them up and returns a value representing the current state of the account
// it also returns the current_value:
//
//	new_current_value := account.CalculateCurrentValue()
//
// (* this part is up for discussion *)
// this is to be run every time the historic changes
func (a *Account) CalculateCurrentValue() int {
	var result int = 0
	for _, transaction := range a.historical {
		result += transaction.Value
	}
	return result
}

func (a *Account) GetHistoric() []Transaction {
	return a.historical
}
