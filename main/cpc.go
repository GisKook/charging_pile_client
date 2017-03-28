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

	//conn := conn.NewConn(1000000000000151, configuration)
	//conn := conn.NewConn(1000000000000152, configuration)
	//conn := conn.NewConn(1000000000000153, configuration)
	//conn := conn.NewConn(1000000000000154, configuration)
	//conn := conn.NewConn(1000000000000155, configuration)
	//conn := conn.NewConn(1000000000000156, configuration)
	//conn := conn.NewConn(1000000000000157, configuration)
	//conn := conn.NewConn(1000000000000158, configuration)
	//conn := conn.NewConn(1000000000000159, configuration)
	conn := conn.NewConn(1000000000000144, configuration)
	//conn := conn.NewConn(1000000000000161, configuration)
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
