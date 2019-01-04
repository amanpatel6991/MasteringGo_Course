package hydraChat

import (
	"bufio"
	"net"
)

type Client struct {
	*bufio.Reader                  //a buffer where client is sending the data on chat server
	*bufio.Writer
	wc chan string
}

func StartClient(msgCh chan string , cn net.Conn , quit chan struct{}) (chan<- string , chan struct{}) {
	client := new(Client)
	client.Reader = bufio.NewReader(cn)
	client.Writer = bufio.NewWriter(cn)

	client.wc = make(chan string)                     //create client write channel
	done := make(chan struct{})

	//setup Reader (WRITING MSGS)
	//(reading what msg client wrote and SENDING it to chatroom msgCh to broadcast it for all other clients in the chat Room)
	go func() {
		scanner := bufio.NewScanner(client.Reader)          //Allows to read buffer data
		for scanner.Scan() {
			logger.Println(scanner.Text())      //to convert received data to string
			//fan in
			msgCh <- scanner.Text()
		}
		done <- struct {}{}                           //when data buffer does not give any data , signal done
	}()

	//setup Writer (READING MSGS)
	client.WriteMonitor()

	go func() {
		select {
		case <-quit:
			cn.Close()               //close underlying connection of the client
		//case <-done:
		}
	}()

	return client.wc , done

}

func(c *Client) WriteMonitor()  {
	go func() {
		for msg:=range c.wc {          //When the msg is received in the clients write channel
			c.WriteString(msg + "\n")        //write it to the underlying client network buffer
		        c.Flush()
		}
	}()
}