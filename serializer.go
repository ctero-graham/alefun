package alefun

import "encoding/json"

type Serializer interface {
    Serialize(interface{}) (string, error)
    Deserialize(string, interface{}) error
}

type JSONSerializer struct {}

func (JSONSerializer) Serialize(v interface{}) (string, error) {
    bts, err := json.Marshal(v)
    return string(bts), err
}

func (JSONSerializer) Deserialize(s string, v interface{}) error {
    return json.Unmarshal([]byte(s), v)
}