package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stenstromen/go-snapnote-backend/model"
)

var (
	db *sql.DB
)

func init() {
	username, password, hostname, database := os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", username, password, hostname, database)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
}

func GetFormDataByNoteID(noteID string) (model.FormData, error) {
	stmt, err := db.Prepare("SELECT json_data FROM your_table WHERE noteid = ?")
	if err != nil {
		return model.FormData{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(noteID)

	var jsonData []byte
	err = row.Scan(&jsonData)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.FormData{}, fmt.Errorf("no data found for noteid: %s", noteID)
		}
		return model.FormData{}, err
	}

	var formData model.FormData
	err = json.Unmarshal(jsonData, &formData)
	if err != nil {
		return model.FormData{}, err
	}

	return formData, nil
}

func InsertFormData(formData model.FormData) error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS your_table (
		id INT PRIMARY KEY AUTO_INCREMENT,
		noteid VARCHAR(8),
		json_data JSON
	)`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	noteID := generateRandomNoteID()

	jsonData, err := json.Marshal(formData)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO your_table (noteid, json_data) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(noteID, jsonData)
	if err != nil {
		return err
	}

	fmt.Println("Data inserted successfully", noteID)
	return nil
}

func generateRandomNoteID() string {
	charSet := "0123456789abcdefghijklmnopqrstuvwxyz"

	noteID := make([]byte, 8)
	for i := range noteID {
		randomIndex := rand.Intn(len(charSet))
		noteID[i] = charSet[randomIndex]
	}

	return string(noteID)
}
