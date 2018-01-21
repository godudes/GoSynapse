package query

import (
	"net"
)

type Request struct {
	addr net.Addr
}

type Response struct {
	serverName string
	serverAddr net.Addr
	gameType string
	gameId string
	version string
	serverEngine string
	numPlayers int
	maxPlayers int
	extra map[string] string
	players map[int] string
}
