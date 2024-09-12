package rpcsrv

import (
	"errors"
	"net"
	"reflect"
)

// Client save the RPC conn
type Client struct {
	conn net.Conn
}

// NewClient return a Client to user
func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

// Conn is Client method to return rpc conn
func (c *Client) Conn() net.Conn {
	return c.conn
}

// Call init function for remote call
func (c *Client) Call(name string, fptr interface{}) {
	container := reflect.ValueOf(fptr).Elem()
	f := func(req []reflect.Value) []reflect.Value {
		transporter := NewTransport(c.conn)

		errorHandler := func(err error) []reflect.Value {
			resultArgs := make([]reflect.Value, container.Type().NumOut())
			for i := 0; i < len(resultArgs)-1; i++ {
				resultArgs[i] = reflect.Zero(container.Type().Out(i))
			}
			resultArgs[len(resultArgs)-1] = reflect.ValueOf(&err).Elem()
			return resultArgs
		}
		// package request arguments
		inArgs := make([]interface{}, 0, len(req))
		for i := range req {
			inArgs = append(inArgs, req[i].Interface())
		}
		// send request
		err := transporter.Send(Data{Name: name, Args: inArgs})
		if err != nil {
			return errorHandler(err)
		}
		// receive response
		rsp, err := transporter.Receive()
		if err != nil {
			return errorHandler(err)
		}
		if rsp.Err != "" {
			return errorHandler(errors.New(rsp.Err))
		}

		if len(rsp.Args) == 0 {
			rsp.Args = make([]interface{}, container.Type().NumOut())
		}
		// unpackage response argument
		resultNum := container.Type().NumOut()
		result := make([]reflect.Value, resultNum)
		for i := 0; i < resultNum; i++ {
			if i != resultNum-1 {
				if rsp.Args[i] == nil {
					result[i] = reflect.Zero(container.Type().Out(i))
				} else {
					result[i] = reflect.ValueOf(rsp.Args[i])
				}
			} else {
				result[i] = reflect.Zero(container.Type().Out(i))
			}
		}
		return result
	}
	container.Set(reflect.MakeFunc(container.Type(), f))
}
