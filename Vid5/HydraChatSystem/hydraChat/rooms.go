package hydraChat

import (
	"sync"
	"fmt"
	"net"
)

//var logger = hydraLogger.GetInstance()

type Room struct {
	Name    string
	Msgch   chan string                 //receives msg from chat clients (main msg channel of the room)
	Clients map[chan <- string]struct{} //people currently talking on the chatroom (empty struct here represents set data structure)
	Quit    chan struct{}               // signals a room exit
	*sync.RWMutex                       //To protect Clients from concurrent access
}

func CreateRoom(roomName string) *Room {

	room := &Room{
		Name: roomName,
		Msgch: make(chan string),
		Clients: make(map[chan <- string]struct{}),
		Quit: make(chan struct{}),
		RWMutex: new(sync.RWMutex),
	}
	room.Run()
	return room
}

func (r *Room) Run() {
	logger.Println("Starting Chat Room :", r.Name)
	go func() {
		for msg := range r.Msgch {
			//check for msg available from any client
			r.broadCastMsg(msg)       //broadcast that msg to all clients on the room
		}
	}()

}

func (r *Room) AddClient(c net.Conn) {
	//c is client connection
	logger.Println("Adding client", c.RemoteAddr())
	r.Lock()                    //Acquire write lock as we are to write to Clients map now

	wc, done := StartClient(r.Msgch, c, r.Quit)                   //wc is write channel this client will be listening on
	r.Clients[wc] = struct{}{}
	r.Unlock()

	//remove client if done is signalled (NICE CLEANUP APPROACH IN GO)
	go func() {
		<- done
		r.RemoveClient(wc)
	}()

}

func(r *Room) RemoveClient(wc chan<- string) {                 //wc is write channel
	logger.Println("Removing Client :", r.Name)

	r.Lock()
	close(wc)
	delete(r.Clients , wc)
	r.Unlock()

	go func() {
		select {
		case <-r.Quit:
			if len(r.Clients) == 0 {
				close(r.Msgch)            //closing main msg channel of the chat room and loop of Run() method exits automatically
			}
		default:                                  //is used here just to avoid a code blockage if above IF condition is not ready yet
		}
	}()

}

func (r *Room) broadCastMsg(msg string) {
	r.RLock()                    //Acquire read lock as we are to which clients are present
	defer r.RUnlock()

	fmt.Println("Received msg :", msg)

	//check for which clients are active on chat room (from the clients set) and send the msg to wc (write channel) of all those clients
	for wc := range r.Clients {
		//Used goroutine here so that even if some client has error in reading other are able to read as running in seperate funs
		go func(wc chan<- string) {
			wc <- msg
		}(wc)
	}
}
