package connect

import (
	"log"

	"github.com/google/uuid"
)

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

func CreateParent(id uuid.UUID, father string, mother string) {
	dbQuery := "insert into parents (uid,father_name,mother_name) values($1, $2, $3)"

	err := con.QueryRow(ctx, dbQuery, id, father, mother)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func CreateStudent(id uuid.UUID, first_name string, second_name string, pid uuid.UUID) {
	dbQuery := "insert into students (id,first_name,second_name,pid) values($1, $2, $3,$4)"

	err := con.QueryRow(ctx, dbQuery, id, first_name, second_name, pid)
	if err != nil {
		log.Fatal(err)
		return
	}

}
