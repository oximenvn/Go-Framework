package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"context"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type DataBase struct {
	Godb *sql.DB
	Name string
}

func (db DataBase) make() {
	fmt.Println("run make")
	query_create := "CREATE DATABASE IF NOT EXISTS " + db.Name
	query_use := "USE " + db.Name
	ctx :=context.Background()
	rows, err := db.Godb.QueryContext(ctx, query_create)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	defer rows.Close()
	fmt.Println(rows)

	rows, err = db.Godb.QueryContext(ctx, query_use)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println(rows)

	query :="SELECT *	FROM `help_topic`	LIMIT 50"
	rows, err = db.Godb.Query(query)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println(rows)


	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	query_table:="CREATE TABLE Persons (		PersonID int,		LastName varchar(255),		FirstName varchar(255),		Address varchar(255),		City varchar(255)	);"
	if ctx, err := db.Godb.QueryContext(ctx, query_table); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
		fmt.Println(ctx)
	}

}

type BaseModel struct {
	Create int
	Update string
}

type Curder interface {
	find() BaseModel
	insert() BaseModel
	update() BaseModel
	delete() BaseModel
}

func (model BaseModel) find() BaseModel {
	s := reflect.ValueOf(&model).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	return BaseModel{}
}

func main() {
	aa := BaseModel{1, "abc"}
	aa.find()

	dsn := "root:123456@tcp(godockerDB:3306)/?charset=utf8"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
    err=db.Ping();
	// defer sql.Close()

	// rows, err := db.Query("select id, first_name from user limit 10")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer rows.Close()

	godb := DataBase{db, "go_db"}
	godb.make()
}
