package main

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/google/uuid"
	database "github.com/tankip/go-social/db/postgres"
)

// function to load 100 users in the database - (id: uuid string, name: string, year: int) and 100 friends for each user - (userid: uuid string, friendid: uuid string)
func main() {
	database.InitDB()
	stmt, err := database.Db.Prepare("INSERT INTO users (id, name, year) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var addedusers []string
	for i := 0; i < 100; i++ {
		id := uuid.New()
		addedusers = append(addedusers, id.String())
		_, err := stmt.Exec(uuid.New(), "user"+strconv.Itoa(i), rand.Intn(100))
		if err != nil {
			log.Fatal(err)
		}
	}
	stmt, err = database.Db.Prepare("INSERT INTO friends (id, userid, friendid) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rand.Shuffle(len(addedusers), func(i, j int) { addedusers[i], addedusers[j] = addedusers[j], addedusers[i] })
	for i := 0; i < 100; i++ {
		userId := addedusers[i]
		friendid := addedusers[rand.Intn(len(addedusers))]
		_, err := stmt.Exec(uuid.New(), userId, friendid)
		if err != nil {
			log.Fatal(err)
		}
	}
}
