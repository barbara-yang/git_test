package rpcsrv

import (
	"bytes"
	"encoding/gob"
)

//Data will be convert to binary
type Data struct {
	Name string
	Args []interface{}
	Err  string
}

func encode(data Data) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode(b []byte) (Data, error) {
	buf := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buf)
	var data Data
	err := decoder.Decode(&data)
	if err != nil {
		return Data{}, err
	}
	return data, nil
}
