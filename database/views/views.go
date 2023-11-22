package views

import "github.com/google/uuid"

type Student struct {
	id         uuid.UUID `json:"id"`
	first_name string    `json:"first_name"`
	last_name  string    `json:"last_name"`
	pid        uuid.UUID `json:"pid"`
}

type Parents struct {
	id          uuid.UUID `json:"id"`
	father_name string    `json:"father_name"`
	mother_name string    `json:"mother_name"`
}
