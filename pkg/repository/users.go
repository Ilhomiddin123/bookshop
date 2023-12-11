package repository

import (
	"bookShop/db"
	"bookShop/models"
	"fmt"
)

/*CRUD
C - create
R - read
U - update
D - delete
*/

func CreateUser(u models.User) error {
	_, err := db.GetDBConn().Exec(
		`INSERT INTO users (full_name, age) values ($1, $2)`,
		u.FullName, u.Age)
	if err != nil {
		return err
	}

	return nil
}

func GetAllUsers() (users []models.User, err error) {
	rows, err := db.GetDBConn().Query("SELECT id, full_name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		u := models.User{}
		err = rows.Scan(&u.ID, &u.FullName, &u.Age)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}

	return users, nil
}

func GetUserByID(id int) (u models.User, err error) {
	row := db.GetDBConn().QueryRow("SELECT id, full_name, age FROM users WHERE id = $1", id)
	err = row.Scan(&u.ID, &u.FullName, &u.Age)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

func EditUserByID(u models.User) error {
	_, err := db.GetDBConn().Exec(
		"UPDATE users SET full_name = $1, age = $2 WHERE id = $3",
		u.FullName, u.Age, u.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUserByID(id int) error {
	_, err := db.GetDBConn().Exec(
		"DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
