package users

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	database "github.com/tankip/go-social/db/postgres"
	"github.com/tankip/go-social/graph/model"
)

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Year    int64  `json:"year"`
	Friends []User `json:"friends"`
}

func GetUsers(filter *model.UserFilter) []User {
	var query = QueryBuilder(filter)
	mys := "SELECT u.id, u.name, u.year, f.name, f.year FROM users AS u LEFT JOIN ( SELECT * from friends AS fr LEFT JOIN users as us ON fr.friendid = us.id ) AS f ON u.id = f.userid  " + query + " LIMIT 100"
	stmt, err := database.Db.Prepare(mys)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	cols, _ := rows.Columns()
	row := make([]interface{}, len(cols))
	rowPtr := make([]interface{}, len(cols))
	for i := range row {
		rowPtr[i] = &row[i]
	}

	var users []User
	for rows.Next() {
		err := rows.Scan(rowPtr...)
		if err != nil {
			log.Fatal(err)
		}
		user := User{
			ID:   row[0].(string),
			Name: row[1].(string),
			Year: row[2].(int64),
		}

		users = append(users, user)
	}
	return users
}

func AddUser(name string, year int) {
	stmt, err := database.Db.Prepare("INSERT INTO users (id, name, year) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(uuid.New(), name, year)
	if err != nil {
		log.Fatal(err)
	}
}

func QueryBuilder(filter *model.UserFilter) string {
	query := ""
	if filter.And != nil {
		for _, filter := range filter.And {

			if filter.Expression.Field != nil {
				if query == "" {
					query += "WHERE "
				} else {
					query += "AND "
				}
				if filter.Expression.Like != nil {
					query += fmt.Sprintf("u.%s LIKE '%%%s%%' ", *filter.Expression.Field, *filter.Expression.Like)
				}
				if filter.Expression.Gte != nil {
					query += fmt.Sprintf("u.%s >= %s ", *filter.Expression.Field, *filter.Expression.Gte)
				}
				if filter.Expression.Lte != nil {
					query += fmt.Sprintf("u.%s <= %s ", *filter.Expression.Field, *filter.Expression.Lte)
				}
			}

		}
	}
	return query
}
