package client_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/marcusljx/eventium/eventiumpb"
	"github.com/marcusljx/eventium/server"
	"github.com/marcusljx/eventium/server/db"
	"google.golang.org/grpc"
)

const (
	testServerPort = 10000
)

var (
	dbHandler     = &db.LocalDB{}
	et            eventiumpb.EventiumClient
	serverAddress = fmt.Sprintf("localhost:%d", testServerPort)
)

func establishClientConnection(address string) *grpc.ClientConn {
	log.Print("Initialising client...")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while initlising client: %v", err)
	}
	et = eventiumpb.NewEventiumClient(conn)
	return conn
}

func TestMain(m *testing.M) {
	// Start server
	//comms := make(chan trigger.Comm)
	log.Print("Starting server...")
	// Start and wait for server
	go server.StartServer(testServerPort, dbHandler)

	//time.Sleep(5 * time.Second)
	conn := establishClientConnection(serverAddress)
	defer conn.Close()

	log.Print("Running Tests...")
	retcode := m.Run()

	os.Exit(retcode)
}
