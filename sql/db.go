package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"unpercent/activity"
)

//timeNow := time.Now()
//fmt.Println(timeNow.Format("2006.01.02 15:04")) //dtime
//fmt.Println(time.Hour*1 + time.Minute*30)       //duration

type Activities interface {
	All() []activity.Activity
	Get(name string) interface{}
	Create(name, dtime, duration, category string)
	Update(name, dtime, duration, category string)
	Delete(name string)
}

type SqliteDb struct {
	db *sql.DB
}

func New() SqliteDb {
	conn, err := sql.Open("sqlite3", "./identifier.sqlite")
	if err != nil {
		panic(err)
	}
	return SqliteDb{
		db: conn,
	}
}

func (s *SqliteDb) All() []activity.Activity {
	rows, err := s.db.Query("SELECT * FROM activities")

	if err != nil {
		panic(err)
	}

	as := make([]activity.Activity, 0)

	for rows.Next() {
		a := activity.Activity{}
		err = rows.Scan(&a.Id, &a.Name, &a.Dtime, &a.Duration, &a.Category)
		fmt.Print("Activity: ", a)
		as = append(as, a)
	}
	_ = rows.Close()
	return as
}

func (s *SqliteDb) Get(name string) activity.Activity {
	rows, err := s.db.Query("SELECT * FROM activities WHERE name=?", name)

	if err != nil {
		panic(err)
	}

	var id int
	var dtime string
	var duration string
	var category string

	a := activity.Activity{}

	for rows.Next() {
		err = rows.Scan(&a.Id, &a.Name, &a.Dtime, &a.Duration, &a.Category)
		fmt.Println("Result: ", id, " ", name, " ", dtime, " ", duration, " ", category)
	}
	_ = rows.Close()
	return a
}

func (s *SqliteDb) Create(name, dtime, duration, category string) {

}

func (s *SqliteDb) Update(name, dtime, duration, category string) {

}

func (s *SqliteDb) Delete(name string) {

}
