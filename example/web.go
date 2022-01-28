package main

import (
	"fmt"
	"log"
	"net/http"

	flag "github.com/spf13/pflag"
	godeamon "github.com/surfaceyu/godeamon"
)

func main() {
	goDaemonF := flag.Bool("f", false, "run app with -f.")

	godeamon.App.Run()
	fmt.Println("=============>", *goDaemonF)

	mux := http.NewServeMux()
	mux.HandleFunc("/index", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("hello, golang!\n"))
	})
	log.Fatalln(http.ListenAndServe(":7070", mux))
}