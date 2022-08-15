package main

import (
	"fmt"
	"github.com/satishbabariya/sql/query"
)

type User struct {
	UserId   int    `db:"user_id,primary",json:"user_id"`
	Email    string `db:"email",json:"email"`
	Password string `db:"password"`
}

func main() {

	builder := query.NewBuilder[User]("user")

	q := builder.Find(User{
		Email: "satish.babariya@gmail.com",
	})

	fmt.Println(q)
}
