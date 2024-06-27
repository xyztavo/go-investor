package database

import (
	"errors"
	"teste/internal/models"
)

func CreateUser(user *models.CreateUserStruct) (createdUserId string, err error) {
	// this will always create a user with a investor role
	err = db.QueryRow("INSERT INTO users (name, password, role) VALUES ($1, $2, $3) RETURNING id", user.Name, user.Password, "investor").Scan(&createdUserId)
	if err != nil {
		return "", errors.New("could not create user")
	}
	return createdUserId, nil
}

func GetUserById(userId string) (user models.User, err error) {
	err = db.QueryRow("SELECT * FROM users WHERE id = $1", userId).Scan(&user.Id, &user.Name, &user.Password, &user.Role)
	if err != nil {
		return user, errors.New("could not find user by id")
	}
	return user, nil
}

func GetUsers() (users []models.User, err error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, errors.New("could get users from db")
	}
	for rows.Next() {
		var user models.User
		rows.Scan(&user.Id, &user.Name, &user.Password, &user.Role)
		users = append(users, user)
	}
	return users, nil
}

func SetAdmin(userId string) error {
	err := db.QueryRow(`UPDATE users SET role = 'admin' WHERE id = $1`, userId).Err()
	if err != nil {
		return err
	}
	return nil
}
