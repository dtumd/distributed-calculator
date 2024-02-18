package model

import "time"

type ComputingResource struct {
	Uuid        string
	Status      string
	Name        string
	LastPing    time.Time
	Description string
}
