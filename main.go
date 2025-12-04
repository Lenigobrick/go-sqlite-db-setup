package main

import (
    "database/sql"
    "fmt"
    _ "modernc.org/sqlite"
)

func main() {
    db, err := sql.Open("sqlite", "test.db")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

	// DROP table users
    _, err = db.Exec(`DROP TABLE IF EXISTS users;`)
    if err != nil {
        fmt.Println(err)
    }

	// DROP table post
	_, err = db.Exec(`DROP TABLE IF EXISTS post;`)
    if err != nil {
        fmt.Println(err)
    }


    // CREATE a first table
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT, 
		username VARCHAR(50) NOT NULL,
		age INTEGER NOT NULL,
		bio TEXT

	);`)
    if err != nil {
        fmt.Println(err)
    }

	// CREATE a second table with FOREIGN KEY
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS post (
		post_id INTEGER PRIMARY KEY AUTOINCREMENT, 
		title VARCHAR(250) NOT NULL,

		user_id INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(user_id)

	);`)
    if err != nil {
        fmt.Println(err)
    }

    // INSERT users
	user1 := "Bob"
	user2 := "John"
    _, err = db.Exec(`INSERT INTO users (username, age) VALUES 
	(?, 25),
	(?, 52);
	`, user1, user2)
    if err != nil {
        fmt.Println(err)
    }

	// INSERT an other user
	user3 := "Lenicode"
	user3Age := 20
	user3Bio := "Some text"
	_, err = db.Exec(`INSERT INTO users (username, age, bio) VALUES 
	(?, ?, ?);
	`, user3, user3Age, user3Bio)
    if err != nil {
        fmt.Println(err)
    }

	// INSERT a post from user3
	postTitle := "My First Post"
	_, err = db.Exec(`INSERT INTO post (title, user_id) VALUES 
	(?, (SELECT user_id FROM users WHERE username = ?));
	`, postTitle, user3)
    if err != nil {
        fmt.Println(err)
    }

	// SELECT number of users
	var count int
	err = db.QueryRow(`
	SELECT count(user_id) FROM users
	`).Scan(&count)
    if err != nil {
        fmt.Println(err)
    }
	fmt.Println("Number of users :", count)

	// SELECT the OLDER user
	var name string
    err = db.QueryRow(`
	SELECT u.username FROM users u
	ORDER BY u.age DESC 
	-- use ASC for the youngest
	LIMIT 1
	`).Scan(&name)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Older user :", name)

    // SELECT username who post with post_id = 1
    var name2 string
    err = db.QueryRow(`
	SELECT u.username FROM users u
	JOIN post p ON p.user_id = u.user_id
	WHERE p.post_id = 1
	`).Scan(&name2)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Name :", name2)
}
