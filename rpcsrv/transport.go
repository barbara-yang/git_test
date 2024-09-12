// realize the transport function
package rpcsrv

import (
	"encoding/binary"
	"io"
	"net"
)

// Transport save conn
type Transport struct {
	conn net.Conn
}

// Create a Transport conn
func NewTransport(conn net.Conn) *Transport {
	return &Transport{conn}
}

// Send request to server or client
func (t *Transport) Send(req Data) error {
	b, err := encode(req)
	if err != nil {
		return err
	}
	buf := make([]byte, 4+len(b))
	binary.BigEndian.PutUint32(buf[:4], uint32(len(b)))
	copy(buf[4:], b)
	_, err = t.conn.Write(buf)
	return err
}

// Receive message from server or client
func (t *Transport) Receive() (Data, error) {
	header := make([]byte, 4)
	_, err := io.ReadFull(t.conn, header)
	if err != nil {
		return Data{}, err
	}
	datalen := binary.BigEndian.Uint32(header)
	data := make([]byte, datalen)
	_, err = io.ReadFull(t.conn, data)
	if err != nil {
		return Data{}, err
	}
	rsp, err := decode(data)
	return rsp, err
}
