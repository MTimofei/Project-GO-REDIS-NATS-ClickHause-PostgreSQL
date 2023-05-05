package main

import (
	"flag"
	"log"

	"git_p/test/insert/web"
	"git_p/test/pkg/connect/bd"
)

var (
	addr      = flag.String("addr", "0.0.0.0:80", "addres server")
	connStr   = flag.String("postgres", "postgres://pet:1234@127.0.0.1:5432/postgres?sslmode=disable", "connection parameters")
	redisAddr = flag.String("redis_addr", "127.0.0.1:6379", "addtes redis")
)

func main() {
	flag.Parse()

	db, err := bd.ConectPostgras(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var cesh = &web.Cesh{
		PostgreasQL: db,
		RediaAddr:   redisAddr,
	}

	log.Fatal(cesh.StartServer(addr))
}
