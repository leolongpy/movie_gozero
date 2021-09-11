package db

import (
	"database/sql"
	"fmt"
	"movie_gozero/rpc/user/internal/entity"
	"time"
)

func InsertUser(userName string, password string, email string) error {
	today := time.Now().Format("2006-01-02")
	_, err := db.Exec("INSERT INTO `user` (`user_name`,`password`,`create_at`,email) VALUES (?,?,?,?)", userName, password, today, email)
	return err
}

func SelectUserByEmail(email string) (*entity.User, error) {
	//var user  *entity.User
	user := entity.User{}
	fmt.Println(user)
	err := db.Get(&user, "SELECT * FROM user WHERE `email` = ? ", email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func SelectUserByPasswordName(email string, password string) (*entity.User, error) {
	user := entity.User{}
	err := db.Get(&user, "SELECT * FROM user where `email` = ? AND `password` = ? LIMIT 1", email, password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func UpdateUserNameProfile(userName string, userId int64) error {
	_, err := db.Exec("UPDATE user SET `user_name` = ? WHERE user_id = ?", userName, userId)
	return err
}
func UpdateUserEmailProfile(email string, userId int64) error {
	_, err := db.Exec("UPDATE `user` SET `email` = ? WHERE user_id = ?", email, userId)
	return err
}

func UpdateUserPhoneProfile(phone string, userId int64) error {
	_, err := db.Exec("UPDATE `user` SET `phone` = ? WHERE user_id = ?", phone, userId)
	return err
}
