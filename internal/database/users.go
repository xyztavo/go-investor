package database

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password" validate:"required"`
}
type CreateUserStruct struct {
	Id       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func CreateUser(user *CreateUserStruct) (createdUserId string, err error) {
	err = db.QueryRow("INSERT INTO users (name, password) VALUES ($1, $2) RETURNING id", user.Name, user.Password).Scan(&createdUserId)
	if err != nil {
		return "", err
	}
	return createdUserId, nil
}

func GetUserById(userId string) (user User, err error) {
	err = db.QueryRow("SELECT * FROM users WHERE id = $1", userId).Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUsers() (users []User, err error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name, &user.Password)
		users = append(users, user)
	}
	return users, nil
}
