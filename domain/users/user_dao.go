package users

import (
	"fmt"
	"github.com/sranjan001/bookstore_users-api/datasources/mysql/usres_db"
	"github.com/sranjan001/bookstore_users-api/util/errors"
	"github.com/sranjan001/bookstore_users-api/util/mysql_utils"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?)"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created, status FROM users where  id= ?"
	queryUpdateUser  = "UPDATE users SET first_name=?, last_name=?,email=? WHERE id=?"
	queryDeleteUser  = "DELETE from users where id=?"

	queryFindUserByStatus = "SELECT id, first_name, last_name, date_created, status FROM users WHERE status=?;"
)

//Get user
func (user *User) Get() *errors.RestError {
	stmt, err := usres_db.Client.Prepare(queryGetUser)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := usres_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
		//return errors.NewInternalServerError(fmt.Sprintf("Error while saving user: %s", err.Error()))
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, err := usres_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := usres_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := usres_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user matching status %s", status))
	}
	return results, nil
}
