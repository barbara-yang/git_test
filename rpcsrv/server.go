package rpcsrv

import (
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
)

// Server singleton object
type Server struct {
	addr  string
	funcs map[string]reflect.Value
}

// NewServer creat a server
func NewServer(addr string) *Server {
	return &Server{addr: addr, funcs: make(map[string]reflect.Value)}
}

// Run start server 开始
func (s *Server) Run() {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Printf("listen on %s err: %v\n", s.addr, err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept err: %v\n", err)
			continue
		}
		go s.Handle(conn)
	}
}

// Handle handle connect request
func (s *Server) Handle(conn net.Conn) {
	defer conn.Close() // Ensure the connection is closed when the function returns
	transporter := NewTransport(conn)

	for {
		data, err := s.recvReq(transporter)
		if err != nil {
			if err == net.ErrClosed {
				log.Printf("connection closed by client: %v\n", err)
				return
			}
			log.Printf("recv req err: %v\n", err)
			return
		}

		resultData := s.callFunc(data.Name, data.Args)
		s.resSend(transporter, resultData)
	}
}

// recvReq receive rpc request
func (s *Server) recvReq(transporter *Transport) (Data, error) {
	data, err := transporter.Receive()
	if err != nil {
		if err != io.EOF {
			log.Printf("read err: %v\n", err)
		}
	}
	return data, err
}

// call function to solve request
func (s *Server) callFunc(funName string, args []interface{}) Data {
	f, ok := s.funcs[funName]
	if !ok {
		e := fmt.Sprintf("func %s dosen't exist", funName)
		log.Println(e)
		//err = transporter.Send(Data{Name: funName, Err: e})
		//if err != nil {
		//	log.Printf("transport write err:%v\n", err)
		//}
		return Data{Name: funName, Err: e}
	}
	// parse request argument
	inArgs := make([]reflect.Value, len(args))
	for i := range args {
		inArgs[i] = reflect.ValueOf(args[i])
	}
	// invoke requested method
	result := f.Call(inArgs)
	// package response arguments (except error)
	resultArgs := make([]interface{}, len(result)-1)
	for i := 0; i < len(result)-1; i++ {
		resultArgs[i] = result[i].Interface()
	}
	// package error argument
	e := ""
	_, ok = result[len(result)-1].Interface().(error)
	if ok {
		e = result[len(result)-1].Interface().(error).Error()
	}
	return Data{Name: funName, Args: resultArgs, Err: e}

}

// resSend send response to user
func (s *Server) resSend(transporter *Transport, data Data) {
	err := transporter.Send(data)
	if err != nil {
		log.Printf("transporter write err: %v\n", err)
	}

}

// Register register function to map
func (s *Server) Register(name string, f interface{}) {
	_, ok := s.funcs[name]
	if ok {
		return
	}
	s.funcs[name] = reflect.ValueOf(f)
}
