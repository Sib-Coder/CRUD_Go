package db

import (
	"CRUD_Go/internal/app/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sqlx.DB
}

func New(dsn string) (*Database, error) {
	//dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Db.Host, cfg.Db.Port, cfg.Db.Username)
	//"host='10.10.0.136' port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		fmt.Println("DataBase NOT WORK")
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

// получение пользователя из бд по имени
func (db *Database) ExtractUser(name string) (model.User, error) {
	var user model.User
	res, err := db.db.Query("SELECT name,lastname, surname, status, gender  FROM employees WHERE name=$1;", name)
	if err != nil {
		return model.User{}, err
	}
	for res.Next() {
		err = res.Scan(&user.Name, &user.LastName, &user.SurName, &user.Status, &user.Gender)
		if err != nil {
			return model.User{}, err
		}
	}
	return user, nil
}

// добавление данных пользователя - мне не нравится надо переделать идею добавления
func (db *Database) AddUser(user model.User) (string, error) {
	var count_users int

	res, err := db.db.Query("SELECT COUNT(id) FROM employees WHERE name = $1 AND lastname = $2 AND surname = $3;", user.Name, user.LastName, user.SurName)
	if err != nil {
		return "", err
	}
	for res.Next() {
		err = res.Scan(&count_users)
		//fmt.Println(count_users)
		if err != nil {
			return "", err
		}
	}

	if count_users == 0 {
		_, err := db.db.Exec("insert into employees (name, lastname, surname, gender, status ) values ($1, $2, $3, $4, $5);", user.Name, user.LastName, user.SurName, user.Gender, user.Status)
		if err != nil {
			return "", err
		}
		return user.Name, err
	} else {
		return "", err
	}

}

// удаление пользователя на основе его статуса и имени и фамилии- готова
// TODO продумать больше нормальных параметров для удаления
func (db *Database) DeleteUser(user model.User) (string, error) {
	_, err := db.db.Exec("DELETE FROM employees WHERE name = $1 and status = $2 and lastname =$3", user.Name, user.Status, user.LastName)
	if err != nil {
		return "", err
	} else {
		return user.Name, err
	}
}

// получение всех пользователей из бд -готова
func (db *Database) ExtractUsers() ([]model.User, error) {
	var user model.User
	var users []model.User
	res, err := db.db.Query("SELECT name,lastname,surname,gender,status FROM employees;")
	if err != nil {
		return nil, err
	}
	for res.Next() {
		err = res.Scan(&user.Name, &user.LastName, &user.SurName, &user.Gender, &user.Status)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil

}

// TODO подумать над обновлением и его параметрами - НЕ ГОТОВА
func (db *Database) UpdateUser(user model.User) (string, error) {
	var count_users int

	res, err := db.db.Query("SELECT COUNT(id) FROM employees WHERE name = $1 AND lastname = $2 AND surname = $3;", user.Name, user.LastName, user.SurName)
	//fmt.Println(res)
	if err != nil {
		return "", err
	}
	for res.Next() {
		err = res.Scan(&count_users)
		fmt.Println(count_users)
		if err != nil {
			return "", err
		}
	}
	if count_users != 0 {
		_, err := db.db.Exec("UPDATE employees set  surname =$3, gender = $4 ,status =$5 WHERE name =$1 AND lastname =$2 ;", user.Name, user.LastName, user.SurName, user.Gender, user.Status)
		if err != nil {
			return "", err
		} else {
			return user.Name, nil
		}
		return "", err
	}
	return user.Name, err

}
