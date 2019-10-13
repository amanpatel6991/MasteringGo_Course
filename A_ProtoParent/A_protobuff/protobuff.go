package A_protobuff

import (
	"github.com/golang/protobuf/proto"
	"errors"
	"net"
	"log"
	"io/ioutil"
)

func ProtoEncodeAndSend(obj interface{}, destination string) error {
	//from message to send format
	val, ok := obj.(*Ship)
	if !ok {
		return errors.New("Proto: Unknown message type")
	}
	data, err := proto.Marshal(val)
	if err != nil {
		return err
	}

	return sendMessage(data, destination)
}

func sendMessage(data []byte, destination string) error {
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		return err
	}
	defer conn.Close()

	log.Println("sending on :", destination)
	_, err = conn.Write(data)
	return err
}

func ProtoReceiveAndDecode(address string) (chan interface{}, error) {
	out := make(chan interface{})

	l, err := net.Listen("tcp", address)
	if err != nil {
		return out, nil
	}

	log.Println("listening on :", address)
	go func() {
		defer l.Close()
		for {
			conn, err := l.Accept()
			if err != nil {
				break
			}
			log.Println("accepted connection from :", conn.RemoteAddr())

			go func(conn net.Conn) {
				defer conn.Close()
				for {
					buffer, err := ioutil.ReadAll(conn)
					if err != nil {
						break
					}
					if len(buffer) == 0 {
						continue
					}
					decodedData, err := DecodeProto(buffer)
					if err != nil {
						continue
					}
					select {
					case out <- decodedData:
					default:
					}
				}
			}(conn)

		}
	}()
	return out , nil

}

func DecodeProto(buffer []byte) (*Ship, error) {
	//receive format to message
	pb := new(Ship)
	return pb, proto.Unmarshal(buffer, pb)
}
