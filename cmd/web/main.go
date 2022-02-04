package main

import (
	"flag"
	"net/http"

	"github.com/rhodeon/sniphub/pkg/prettylog"
)

const PORT = "4000"

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	app := application{}

	prettylog.InfoF("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, app.routesHandler())
	prettylog.Error(err.Error())
}
