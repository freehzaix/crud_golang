package databases

import (
	"database/sql"
	"fmt"

	"freehzaix.com/crud_golang/models"
	_ "github.com/go-sql-driver/mysql"
)


func db() *sql.DB {
	// Configure the database connection (always check errors)
	myDb, err := sql.Open("mysql", "jeanluc:edyrodal@(127.0.0.1:3306)/db_test_go?parseTime=true")

	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil
	}

	// Initialize the first connection to the database to see if everything works correctly.
	// Make sure to check the error.
	errPing := myDb.Ping()
	if errPing != nil {
		fmt.Println("Error connecting to the database:", errPing)
		return nil
	}

	// Database connection successful, continue with your code...
	createTableUsers(myDb)

	return myDb
}

//Créer la table users
func createTableUsers(Mydb *sql.DB){
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT,
			nom TEXT NOT NULL,
			prenom TEXT NOT NULL,
			telephone TEXT NOT NULL,
			email TEXT NOT NULL,
			PRIMARY KEY (id)
		);`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	_, errTableUsers := Mydb.Exec(query)

	if errTableUsers != nil {
		fmt.Println("Error creating tablle users:", errTableUsers)
		return
	}
}

//Supprimer la table users
// func deleteTableUsers(db *sql.DB){
// 	query := `DROP TABLE users ;`

// 	// Executes the SQL query in our database. Check err to ensure there was no error.
// 	_, errTableUsers := db.Exec(query)

// 	if errTableUsers != nil {
// 		fmt.Println("Error creating tablle users:", errTableUsers)
// 		return
// 	}
// }

func InsertUser(user models.User){
	myDb := db()

	// Inserts our data into the users table and returns with the result and a possible error.
	// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
	result, err := myDb.Exec(`INSERT INTO users (nom, prenom, telephone, email) VALUES (?, ?, ?, ?)`, user.Nom, user.Prenom, user.Telephone, user.Email)
	if err != nil {
		fmt.Println("Error insert user table data:", err)
		return
	}

	userID, err := result.LastInsertId()
	if err == nil {
		fmt.Println("Insert user table data successful:", userID)
		return
	}

}

func ReadUser() []models.User{
	myDb := db()

	rows, err := myDb.Query(`SELECT * FROM users`)
	if err != nil {
		fmt.Println("Error user :", err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.Id, &u.Nom, &u.Prenom, &u.Telephone, &u.Email) // check err
		if err != nil {
			fmt.Println("Error user :", err)
		}
		users = append(users, u)
	}
	err = rows.Err() // check err
	if err != nil {
		fmt.Println("Error user row :", err)
	}

	return users
}

func ShowUser(id int) []models.User{
	myDb := db()

	rows, err := myDb.Query(`SELECT * FROM users WHERE id = ?`, id)
	if err != nil {
		fmt.Println("Error user :", err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.Id, &u.Nom, &u.Prenom, &u.Telephone, &u.Email) // check err
		if err != nil {
			fmt.Println("Error user :", err)
		}
		users = append(users, u)
	}
	err = rows.Err() // check err
	if err != nil {
		fmt.Println("Error user row :", err)
	}

	return users
}

func ShowUserId(id int) []models.User{
	myDb := db()

	rows, err := myDb.Query(`SELECT id FROM users WHERE id = ?`, id)
	if err != nil {
		fmt.Println("Error user :", err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.Id) // check err
		if err != nil {
			fmt.Println("Error user :", err)
		}
		users = append(users, u)
	}
	err = rows.Err() // check err
	if err != nil {
		fmt.Println("Error user row :", err)
	}

	return users
}

func UpdateUser(user models.User){
	myDb := db()

	// Inserts our data into the users table and returns with the result and a possible error.
	// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
	_, err := myDb.Exec(`UPDATE users SET nom = ?, prenom = ?, telephone = ?, email = ? WHERE id = ?`, user.Nom, user.Prenom, user.Telephone, user.Email, user.Id)
	
	if err != nil {
		fmt.Println("Error insert user table data:", err)
		return
	}

}

func DeleteUser(user models.User){
	myDb := db()

	// Exécuter la requête DELETE avec l'ID spécifié
	_, err := myDb.Exec(`DELETE FROM users WHERE id = ?`, user.Id)
	
	if err != nil {
		fmt.Println("Error delete row user:", err)
		return
	}

}