package web

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
)

type Cesh struct {
	PostgreasQL *sql.DB
	RediaAddr   *string
	Log         *log.Logger
}

func (cesh *Cesh) StartServer(addr *string) (err error) {
	n := net.ListenConfig{}
	lis, err := n.Listen(context.Background(), "tcp", *addr)
	if err != nil {
		return err
	}
	fmt.Println("Server Start")
	err = http.Serve(lis, cesh.router())
	return err
}

func (cesh *Cesh) router() (mux *http.ServeMux) {
	mux = http.NewServeMux()

	mux.HandleFunc("/item/create", cesh.handlerPost)
	mux.HandleFunc("/item/update", cesh.handlerPatch)
	mux.HandleFunc("/item/remove", cesh.handlerDelete)
	mux.HandleFunc("/items/list", cesh.handlerGet)

	return mux
}
