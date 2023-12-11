package repository

import (
	"bookShop/db"
	"bookShop/models"
	"fmt"
)

func CreateBook(b models.Book) error {
	_, err := db.GetDBConn().Exec(`
				INSERT INTO books 
				    (title, description, author)
				    VALUES ($1, $2, $3)`, b.Title, b.Description, b.Author)
	if err != nil {
		return err
	}

	return nil
}

func GetAllBooks() (books []models.Book, err error) {
	/////
	rows, err := db.GetDBConn().Query("SELECT id, title, description, author FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		u := models.Book{}
		err = rows.Scan(&u.ID, &u.Title, &u.Description, &u.Author)
		if err != nil {
			fmt.Println(err)
			continue
		}
		books = append(books, u)
	}

	return books, nil
}
func GetBookByID(id int) (u models.Book, err error) {
	row := db.GetDBConn().QueryRow("SELECT id, title, description, author FROM books WHERE id = $1", id)
	err = row.Scan(&u.ID, &u.Title, &u.Description, &u.Author)
	if err != nil {
		return models.Book{}, err
	}

	return u, nil
}
func EditBookByID(u models.Book) error {
	_, err := db.GetDBConn().Exec(
		"UPDATE books SET title = $1, description = $2, author = $3 WHERE id = $4",
		u.Title, u.Description, u.Author, u.ID)
	if err != nil {
		return err
	}

	return nil
}
func DeleteBookByID(id int) error {
	_, err := db.GetDBConn().Exec(
		"DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
