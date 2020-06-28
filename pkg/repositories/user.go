package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"simple/pkg/driver"
	model "simple/pkg/models"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository() *UsersRepository {
	connection, err := driver.Connect()
	if err != nil{
		panic(err)
	}
	return &UsersRepository{db: connection}
}

func (u UsersRepository) FetchAll() []model.User {
	rows, err := u.db.Query("select * from users")
	if err != nil{
		panic(err)
	}
	defer rows.Close()
	var users []model.User
	for rows.Next(){
		p := model.User{}
		err := rows.Scan(&p.Id, &p.Email, &p.Password, &p.CreatedAt)
		if err != nil{
			fmt.Println(err)
			continue
		}
		users = append(users, p)

	}
	return users

}