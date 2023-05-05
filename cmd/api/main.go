package main

import (
	"flag"
	"log"

	"git_p/test/insert/web"
	"git_p/test/pkg/bd"
)

var (
	addr    = flag.String("addr", "0.0.0.0:80", "addres server")
	connStr = flag.String("postgres", "postgres://pet:1234@127.0.0.1:5432/postgres?sslmode=disable", "connection parameters")
)

func main() {
	flag.Parse()

	db, err := bd.ConectPostgras(connStr)
	if err != nil {
		log.Fatal(err)
	}

	var cesh = &web.Cesh{
		PostgreasQL: db,
	}

	log.Fatal(cesh.StartServer(addr))
}
