package http

import (
	"fmt"
	"net/http"
	"unpercent/sql"
)

//https://hackandsla.sh/posts/2021-06-18-embed-vuejs-in-go/#building-the-backend
// https://dev.to/ksckaan1/creating-a-website-using-golang-and-sveltejs-together-55g8
func Start(db *sql.SqliteDb) error {
	router := http.NewServeMux()
	router.HandleFunc("/activities/", activityHandler(db))
	router.Handle("/", http.FileServer(http.Dir("./public")))

	return http.ListenAndServe(":5000", router)
}

/*
	GET
		All()
		Get(name)
	POST
		Create(name, dtime, duration, category)
	PUT
		Update(name, dtime, duration, category)
	DELETE
		Delete(name)
*/

func activityHandler(db *sql.SqliteDb) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/activities/" {
			if r.Method == http.MethodGet {
				activities := db.All()
				w.Write([]byte(fmt.Sprint(activities)))
				return
			}
			if r.Method == http.MethodDelete {
				w.Write([]byte("Can not call method 'DELETE' on all activities"))
				return
			}
			return
		}

		switch r.Method {
		case http.MethodGet:
			activity := db.Get("Run")
			w.Write([]byte(fmt.Sprint(activity)))
			return
		}

		w.Write([]byte("hello, world!"))
	}

}
