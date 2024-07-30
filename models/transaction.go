package models

import "time"

type Transaction struct {
	id          int
	Value       int
	date        time.Time
	Description string
}
