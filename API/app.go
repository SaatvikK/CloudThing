package main
import (
	"database/sql
	"github.com/lib/pq"
)
type App struct {
	Router *mux.Router
	DB     *sqp.DB
}

func (a *App) initialize(user, pwd, dbname string) {

}

func (a *App) run(addr string) {

}
