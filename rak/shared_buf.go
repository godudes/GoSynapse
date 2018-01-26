package rak

import (
	"encoding/binary"
)

// Shared buffer.
// Each network card, not CPU core or so, should have a shared buffer instance.
type sharedBuf struct {
	offlinePingOut [] byte
}

func newSharedBuf() * sharedBuf {
	return &sharedBuf{
		offlinePingOut: make([]byte, binary.Size(offlinePing{})),
	}
}