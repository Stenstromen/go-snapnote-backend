package service

import (
	"bytes"
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
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

	var compressedData []byte
	err = row.Scan(&compressedData)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.FormData{}, fmt.Errorf("no data found for noteid: %s", noteID)
		}
		return model.FormData{}, err
	}

	gz, err := gzip.NewReader(bytes.NewReader(compressedData))
	if err != nil {
		return model.FormData{}, err
	}
	defer gz.Close()

	decompressedData, err := io.ReadAll(gz)
	if err != nil {
		return model.FormData{}, err
	}

	var formData model.FormData
	err = json.Unmarshal(decompressedData, &formData)
	if err != nil {
		return model.FormData{}, err
	}

	return formData, nil
}

func InsertFormData(formData model.FormData) (string, error) {
	createTableQuery := `CREATE TABLE IF NOT EXISTS snapnote (
        id INT PRIMARY KEY AUTO_INCREMENT,
        noteid VARCHAR(8),
        json_data LONGBLOB
    )`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		return "", err
	}

	noteID := generateRandomNoteID()

	var buf bytes.Buffer
	gz, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	if err != nil {
		return "", err
	}

	if err := json.NewEncoder(gz).Encode(formData); err != nil {
		return "", err
	}
	if err := gz.Close(); err != nil {
		return "", err
	}

	stmt, err := db.Prepare("INSERT INTO snapnote (noteid, json_data) VALUES (?, ?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(noteID, buf.Bytes())
	if err != nil {
		return "", err
	}

	fmt.Println("Data inserted successfully", noteID)
	return noteID, nil
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
