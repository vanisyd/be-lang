package authservice

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"time"
	"web/data"
)

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AuthToken struct {
	UID   uint64 `json:"uid"`
	Token string `json:"token"`
}

type AuthUser struct {
	User      User      `json:"user"`
	AuthToken AuthToken `json:"token"`
}

func Auth(userObj User) (AuthUser, error) {
	dbConnection := data.DBConnection()

	user, err := findUser(dbConnection, userObj.ID)
	if err != nil {
		return AuthUser{}, err
	}
	if user.Password != userObj.Password {
		return AuthUser{}, WrongPassword(userObj)
	}

	authUser := AuthUser{
		User:      user,
		AuthToken: getToken(user),
	}

	return authUser, nil
}

func Register(userObj User) (AuthUser, error) {
	var user AuthUser
	dbConnection := data.DBConnection()

	rows, err := dbConnection.Query("SELECT * FROM `users` WHERE `name` = ?", userObj.Name)
	defer rows.Close()
	if err != nil || rows.Next() == true {
		return user, UserAlreadyExists(userObj)
	}

	newUser, err := dbConnection.Exec("INSERT INTO `users` (`name`, `password`) VALUES (?, ?)", userObj.Name, userObj.Password)
	if err != nil {
		return user, err
	}

	userId, err := newUser.LastInsertId()
	if err != nil {
		return user, err
	}
	userObj.ID = uint64(userId)
	user = AuthUser{
		User:      userObj,
		AuthToken: getToken(userObj),
	}

	return user, nil
}

func getToken(user User) AuthToken {
	currentTime := time.Now()
	token := sha256.Sum256([]byte(fmt.Sprint(user.ID, user.Name, currentTime)))
	userToken := AuthToken{
		UID:   user.ID,
		Token: hex.EncodeToString(token[:]),
	}

	return userToken
}

func findUser(db *sql.DB, id uint64) (user User, err error) {
	row := db.QueryRow("SELECT * FROM `users` WHERE `id` = ?", id)
	err = row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		return user, UserNotFound(User{
			ID: id,
		})
	}

	return
}
