package database

import "errors"

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}
type CreateUserStruct struct {
	Id       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}

func CreateUser(user *CreateUserStruct) (createdUserId string, err error) {
	// this will always create a user with a investor role
	err = db.QueryRow("INSERT INTO users (name, password, role) VALUES ($1, $2, $3) RETURNING id", user.Name, user.Password, "investor").Scan(&createdUserId)
	if err != nil {
		return "", errors.New("could not create user")
	}
	return createdUserId, nil
}

func GetUserById(userId string) (user User, err error) {
	err = db.QueryRow("SELECT * FROM users WHERE id = $1", userId).Scan(&user.Id, &user.Name, &user.Password, &user.Role)
	if err != nil {
		return user, errors.New("could not find user by id")
	}
	return user, nil
}

func GetUsers() (users []User, err error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, errors.New("could get users from db")
	}
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name, &user.Password, &user.Role)
		users = append(users, user)
	}
	return users, nil
}
