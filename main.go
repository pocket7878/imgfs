package main

import "net"
import "github.com/docker/go-p9p"
import "github.com/pocket7878/imgfs/session"
import "context"
import "log"

func main() {
	ctx := context.Background()
	listen, _ := net.Listen("tcp", ":8080")
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln("error accepting:", err)
			continue
		}

		go func(conn net.Conn) {
			session := session.NewSession(ctx)
			if err := p9p.ServeConn(ctx, conn, p9p.Dispatch(session)); err != nil {
				log.Printf("serving conn: %v", err)
			}
		}(conn)
	}
}
