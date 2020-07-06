package repository

import (
	"time"
	"fmt"
	"context"
	"simple_api_go/model"
	"simple_api_go/utils"
	"log"
)

const (
	table = "user"
	layoutDateTime = "2006-01-02 15:04:05"
)

//for get
func GetAll(ctx context.Context) ([]model.User, error){

	var users []model.User

	db, err := utils.MySQL()

	if err != nil{
		log.Fatal("Can't 	connect to Mysql")
	}

	queryText := fmt.Sprintf("SELECT * FROM %v ORDER BY id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil{
		log.Fatal(err)
	}

	for rowQuery.Next(){
		var user model.User
		var createAt string

		if err = rowQuery.Scan(
			&user.ID,
			&user.Nama,
			&user.Gender,
			&createAt); err != nil{
				return nil, err
			}
		
		user.CreateAt, err = time.Parse(layoutDateTime, createAt)
		
		if err !=  nil{
			log.Fatal(err)
		}

		users = append(users, user)

	}

	return users, nil


}

// for insert 
func Insert(ctx context.Context, usr model.User) error{
	db, err := utils.MySQL()
	if err != nil{
		log.Fatal("cant connect msql")
	}
	queryText := fmt.Sprintf("INSERT INTO %v (id, nama, gender, create_at) values (%v, '%v','%v','%v')", table,
		usr.ID,
		usr.Nama,
		usr.Gender,
		time.Now().Format(layoutDateTime),
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil{
		return err
	}

	return nil
}

// for update
func Update(ctx context.Context, usr model.User) error{
	db, err := utils.MySQL()
	if err != nil{
		log.Fatal("cant connect msql")
	}
	queryText := fmt.Sprintf("UPDATE %v set nama = '%v', gender = '%v' WHERE id = '%v'", table,
		
		usr.Nama,
		usr.Gender,
		usr.ID,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil{
		return err
	}

	return nil
}

//delete
func Delete(ctx context.Context, usr model.User) error{
	db, err := utils.MySQL()
	if err != nil{
		log.Fatal("cant connect msql")
	}
	queryText := fmt.Sprintf("DELETE FROM %v  WHERE id = '%v'", table,
		usr.ID,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil{
		return err
	}

	return nil
}