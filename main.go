package main

import "fmt"
import "net"
import "github.com/docker/go-p9p"
import "context"
import "log"

type imgFS struct {
}

func (fs *imgFS) Handle(ctx context.Context, msg p9p.Message) (p9p.Message, error) {
	log.Printf("Incoming message type: %v\n", msg.Type())
	switch msg.Type() {
	case p9p.Tauth:
		return fs.HandleAuth(msg.(p9p.MessageTauth))
	default:
		return fs.HandleUnsupported(msg)
	}
}

func (fs *imgFS) HandleAuth(msg p9p.MessageTauth) (p9p.Message, error) {
	log.Printf("Afid: %v\n", msg.Afid)
	return nil, fmt.Errorf("Yappy")
}

func (fs *imgFS) HandleUnsupported(msg p9p.Message) (p9p.Message, error) {
	return nil, fmt.Errorf("Unsupported message type: %v", msg.Type())
}

func main() {
	listen, _ := net.Listen("tcp", ":8080")
	for {
		conn, _ := listen.Accept()
		go func() {
			newFS := new(imgFS)
			p9p.ServeConn(context.Background(), conn, newFS)
		}()
	}
}
