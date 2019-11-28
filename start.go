package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/MyRetail/common"
	hndlr "github.com/MyRetail/handler"

	"github.com/gorilla/mux"
)

type myRetail common.Server

func main() {
	mR := newRetail()
	mR.hostProducts()
}

func newRetail() *myRetail {
	r := new(myRetail)
	r.Name = "MyRetail"
	r.Address = os.Getenv("IP_N_PORT")
	if r.Address == "" {
		r.Address = ":8088"
	}
	r.Router = mux.NewRouter().StrictSlash(true)
	r.updateRoutes()
	return r
}

func (r *myRetail) hostProducts() {
	for {
		e := http.ListenAndServe(r.Address, r.Router)
		if e != nil {
			fmt.Println("Error host retail server", e.Error())
			time.Sleep(10 * time.Second)
			continue //try to host after sleeping 10 sec
		}
		break
	}
}

func (re *myRetail) updateRoutes() {
	re.Router.StrictSlash(true)
	re.Router.HandleFunc("/", serverStarted)
	for _, r := range hndlr.MyRetailRoutes {
		re.Router.Methods(r.Method).Path(r.Path).Name(r.Name).Handler(r.Handler)
	}
}

func serverStarted(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello there mate\n")
}
