package mysql

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)
var db *sql.DB

func init(){
	var err error
	// 连接数据库
	db, err = sql.Open("mysql", "username:Password@tcp(127.0.0.1:3306)/DatabaseName")
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		errors.Wrapf(err,"连接数据库失败")
		logrus.Errorf(err.Error())
		return
	}
}


func GetInvitedNumber(userNickname string) int {
	var number int

	// 首先尝试查询用户
	err := db.QueryRow("SELECT Count FROM Invited_count WHERE Nickname = ?", userNickname).Scan(&number)
	if err != nil {
		if err == sql.ErrNoRows {
			// 如果没有找到用户，插入新用户数据，Count初始为0
			_, err := db.Exec("INSERT INTO Invited_count (Nickname, Count) VALUES (?, 0)", userNickname)
			if err != nil {
				logrus.WithError(err).Error("error inserting new user into database")
				return 0 // 或者其他错误处理方式
			}
			// 插入成功后，Count为0
			return 0
		} else {
			// 查询过程中出现其他错误
			logrus.WithError(err).Error("error querying database")
			fmt.Println("error querying database")
			return -1 // 或者其他错误处理方式
		}
	}

	// 如果用户存在，返回查询到的Count值
	return number
}


func Addinvitor(userNickname string) error {
	var number int

	// 首先检查用户是否存在
	err := db.QueryRow("SELECT Count FROM Invited_count WHERE Nickname = ?", userNickname).Scan(&number)
	if err != nil {
		if err == sql.ErrNoRows {
			// 如果没有找到用户，插入用户数据，Count初始为0
			_, err = db.Exec("INSERT INTO Invited_count (Nickname, Count) VALUES (?, 0)", userNickname)
			if err != nil {
				return errors.Wrap(err, "error inserting new user")
			}
		} else {
			// 查询过程中出现其他错误
			return errors.Wrap(err, "error querying database")
		}
	}

	// 更新用户的Count
	_, err = db.Exec("UPDATE Invited_count SET Count = Count + ? WHERE Nickname = ?", 1, userNickname)
	if err != nil {
		return errors.Wrap(err, "error updating Count")
	}

	return nil
}

func Setzero(userNickname string) error{
	_,err:=db.Exec("UPDATE Invited_count SET Count = ? WHERE Nickname = ?", 0, userNickname)
	if err != nil {
		return errors.Wrap(err, "error updating Count zero")
	}
	return nil
}