package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

type User [8]byte

type Message struct {
	ID       int     `json:"id"`
	Message  string  `json:"message"`
	X        float32 `json:"x"`
	Y        float32 `json:"y"`
	User     User    `json:"user"`
	Comments []Comment
}

type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	User    User
}

func RepoCreateComment(c Comment) (Comment, error) {
	var id int
	err := db.QueryRow("INSERT INTO comments(content) VALUES ($1) RETURNING id", c.Content).Scan(&id)
	if err != nil {
		return Comment{}, err
	}
	c.ID = id
	return c, nil
}

func RepoCreateMessage(m Message) (Message, error) {
	query := fmt.Sprintf("INSERT INTO messages(message, location) VALUES ($1, ST_GeographyFromText('SRID=4326;POINT(%v %v)')) RETURNING id", m.X, m.Y)
	var id int
	err := db.QueryRow(query, m.Message).Scan(&id)
	if err != nil {
		return Message{}, err
	}
	m.ID = id
	return m, nil
}

func RepoFindMessage(x float32, y float32) ([]Message, error) {
	query := fmt.Sprintf("SELECT id, message, ST_X(location::geometry) as x, ST_Y(location::geometry) as y FROM messages WHERE ST_DWithin(location, ST_GeographyFromText('SRID=4326;POINT(%v %v)'), 10000)", x, y)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	messages := []Message{}
	for rows.Next() {
		var r Message
		if err := rows.Scan(&r.ID, &r.Message, &r.X, &r.Y); err != nil {
			return nil, err
		}
		messages = append(messages, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func init() {
	var err error

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("You need to set DATABASE_URL environement variable")
	}

	db, err = sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}
}
