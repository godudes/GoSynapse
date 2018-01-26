package rak

type pk interface {
	encode(buf *[]byte)
	decode(buf *[]byte)
}

type pkImpl struct {
	pkId int32
}

type offlinePk struct {
	*pkImpl
	magic []byte
}

type pingTime int64

// Officially named UNCONNECTED_PING
type offlinePing struct {
	*offlinePk
	pingTime pingTime
}
