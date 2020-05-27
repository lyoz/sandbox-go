package model

type User struct {
	Id       int
	Username string
}

func AddUser(username string) error {
	_, err := DB.Exec("INSERT INTO users (username) VALUES (?)", username)
	return err
}

func FindUser(username string) (User, error) {
	user := User{}
	err := DB.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(
		&user.Id, &user.Username,
	)
	return user, err
}
