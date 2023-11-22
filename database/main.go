package database

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type Parent struct {
	// ID          uuid.UUID `json:"id"`
	father_name string `json:"father_name"`
	mother_name string `json:"mother_name"`
}

var conn *pgx.Conn

func main() {
	var err error
	ctx := context.Background()
	connString := "postgres://yash:yash1234@localhost:5432/student_info?sslmode=disable"

	conn, err = pgx.Connect(ctx, connString)
	fmt.Println(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	fmt.Println("connected")
	r := gin.Default()
	r.GET("/parents", GetParent)
	r.Run(":3000")

}

// func connect() {

// }
func GetParent(c *gin.Context) {
	rows, err := conn.Query(context.Background(), "SELECT father_name,mother_name FROM parents")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		// log.Fatal(err)
		return
	}
	defer rows.Close()
	// if rows == nil {

	// }
	var parents []Parent
	for rows.Next() {
		var parent Parent

		if err := rows.Scan(&parent.father_name, &parent.mother_name); err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		parents = append(parents, parent)
	}
	fmt.Println("Parents:")
	for _, parent := range parents {
		fmt.Printf("father name: %d, mother name: %s\n", parent.father_name, parent.mother_name)
	}
	c.JSON(http.StatusOK, parents)
}
