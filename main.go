package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"os"

	"github.com/Khigashiguchi/go-ecs-example/config"
	_ "github.com/go-sql-driver/mysql" // mysql driverを使うため
)

// NewDB return database global connection handle.
func NewDB(conf config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			conf.User,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Name))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// Handler represents each routing handler.
type Handler struct {
	DB *sql.DB
}

// GetPostsHandler handle GET request to /posts.
func (h *Handler) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	// use database here.
}

func main() {
	var err error

	// Get configuration
	conf, err := config.NewConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get configuration: %s", err)
		panic(err.Error())
	}

	// Get database Handle
	db, err := NewDB(conf.DB)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get connection with database: %s", err)
		panic(err.Error())
	}

	// Router
	r := mux.NewRouter()
	h := Handler{DB: db}
	r.Methods("GET").Path("/posts").HandlerFunc(h.GetPostsHandler)
	r.Methods("GET").Path("/.healthcheck").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Serve HTTP service
	fmt.Fprint(os.Stdout, ">> Start to listen http server post :80\n")
	if err = http.ListenAndServe(":80", r); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start http server: %s", err)
		panic(err.Error())
	}
}
