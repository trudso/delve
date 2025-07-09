package engine

import (
	"net"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type NetworkServer struct {
	Port         int
	clientsMutex sync.Mutex
	clients      []clientConnection
}

type clientConnection struct {
	Id   string
	Conn net.Conn
}

func (s *NetworkServer) Run() {
	listener, err := net.Listen("tcp4", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			rl.TraceLog(rl.LogWarning, "unable to accept client connection: %+v", err.Error())
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *NetworkServer) UpdateClients() bool {
	lockAcquired := s.clientsMutex.TryLock()
	if !lockAcquired {
		return false
	}

	// TODO[mt]: do your stuff here

	// unlock and return success
	s.clientsMutex.Unlock()
	return true
}

func (s *NetworkServer) closeClient(c net.Conn) {
	defer c.Close()
	s.clientsMutex.Lock()
	newClients := make([]clientConnection, 0, len(s.clients)-1)
	for _, client := range s.clients {
		if client.Conn != c {
			newClients = append(newClients, client)
		}
	}
	s.clients = newClients
	s.clientsMutex.Unlock()
}

func (s *NetworkServer) addClient(c net.Conn) clientConnection {
	client := clientConnection{
		Conn: c,
	}
	s.clientsMutex.Lock()
	s.clients = append(s.clients, client)
	s.clientsMutex.Unlock()
	return client
}

func (s *NetworkServer) handleConnection(c net.Conn) {
	defer s.closeClient(c)
	client := s.addClient(c)
	data := make([]byte, 8192)

	for {
		_, err := client.Conn.Read(data)
		if err != nil {
			rl.TraceLog(rl.LogWarning, "Error reading from client connection: %+v", err.Error())
			return
		}

		// TODO[mt]: handle data
	}
}
