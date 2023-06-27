package models

import "github.com/google/uuid"

type Employee struct {
	Id      uuid.UUID
	Name    string
	Address Address
}
