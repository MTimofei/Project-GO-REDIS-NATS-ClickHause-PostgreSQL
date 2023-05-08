package main

import (
	"bytes"
	"flag"
	"time"

	"git_p/test/insert/migration"
	"git_p/test/insert/natspkg"
	"git_p/test/insert/web"
	"git_p/test/pkg/connect/bd"
	"git_p/test/pkg/errmy"
)

var (
	addr          = flag.String("addr", "0.0.0.0:80", "addres server")
	connStr       = flag.String("postgres", "postgres://pet:1234@127.0.0.1:5432/postgres?sslmode=disable", "connection parameters")
	redisAddr     = flag.String("redis_addr", "127.0.0.1:6379", "addtes redis")
	clichauseAddr = flag.String("clichause_addr", "127.0.0.1:9000", "addtes clichause")

	idlog  = 1
	bufLog = bytes.Buffer{}
)

func main() {
	flag.Parse()

	log := errmy.Log(&bufLog)

	db, err := bd.ConectPostgras(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	nc, err := natspkg.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	var cesh = &web.Cesh{
		PostgreasQL: db,
		RediaAddr:   redisAddr,
		Log:         log,
	}

	go func() {
		for {
 		time.Sleep(10 * time.Minute)
	 		migration.Migration(db, clichauseAddr)
	 	}
	 }()


	go func() {
		for {
			natspkg.SetLog(&idlog, &bufLog, nc)
			time.Sleep(10 * time.Second)
		}
	}()

	log.Fatal(cesh.StartServer(addr))
}
