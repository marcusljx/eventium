package main

/* AUTOGENERATED USING grpcgen
 * Only edit if you know what you are doing!
 * For template changes/edits/bugs/new features, please contact marcusljx@gmail.com
 */

import (
	"flag"

	"github.com/marcusljx/eventium/server"
	"github.com/marcusljx/eventium/server/db"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	server.StartServer(*port, make(db.LocalDB))
}

//go:generate protoc -I=./eventiumpb -I=$PROTOC/include ./eventiumpb/eventium.proto --go_out=plugins=grpc:eventiumpb
