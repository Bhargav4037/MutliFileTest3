package maindb

import (
	"fmt"

	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
)

func Check() {
	pgxv4.RegisterDriver("ghitnj")
	fmt.Println("Hello")
}
