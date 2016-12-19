package main

import (
	"fmt"
	"github.com/giskook/charging_pile_client/conf"
	"github.com/giskook/charging_pile_client/conn"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// read configuration
	configuration, err := conf.ReadConfig("./conf.json")

	checkError(err)

	conn := conn.NewConn(1000000000000050, configuration)
	go conn.Start()
	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)

	// stops service
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
