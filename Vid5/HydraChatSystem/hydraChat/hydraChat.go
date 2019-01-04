package hydraChat

import (
	"MasteringGo_Course/Vid4/SingletonPattern/hydraLogger"
	"net"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

var logger = hydraLogger.GetInstance()

//Start Chat
func Run(port string) error {
	l, err := net.Listen("tcp", port)
	if err != nil {
		logger.Println("Error establishing tc connection for chatroom : ", err)
		return err
	}

	room := CreateRoom("Hydra Chat Room !")

	//Cleanup and close chat server
	go func() {
		ch := make(chan os.Signal)
		//Notify causes package signal to relay incoming signals to c.
		// If no signals are provided, all incoming signals will be relayed to c.
		// Otherwise, just the provided signals will (here : syscall.SIGINT, syscall.SIGTERM)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)          //when interrupt or terminate signal occurs

		<-ch

		l.Close()
		logger.Println("Closing chatroom")
		fmt.Println("Closing chatroom")
		close(room.Quit)

		<-room.Msgch

		os.Exit(0)
	}()

	go func(l net.Listener) {
		for {
			conn, err := l.Accept()
			if err != nil {
				logger.Println("Error accepting chat connection from client : ", err)
				break
			}
                        //fan out
			go handleConnection(room, conn)

		}
	}(l)

	return nil
}

func handleConnection(room *Room, conn net.Conn) {
	logger.Println("received client request from : ", conn.RemoteAddr())
	room.AddClient(conn)
}
