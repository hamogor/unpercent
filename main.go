package main

import (
	"unpercent/http"
	"unpercent/sql"
)

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.3.html

func main() {

	db := sql.New()

	db.All()
	db.Get("Run")

	err := http.Start(&db)
	if err != nil {
		panic(err)
	}

}
