package tool

import jsoniter "github.com/json-iterator/go"

// jsonTool jsoniter專用的比原生快的套件
var jsonTool = jsoniter.ConfigCompatibleWithStandardLibrary

// JSONDeSerialize Json反序列化
func JSONDeSerialize(msg []byte, dataSt interface{}) error {
	err := jsonTool.Unmarshal(msg, &dataSt)
	return err
}

// JSONSerialize Json序列化
func JSONSerialize(dataSt interface{}) ([]byte, error) {
	data, err := jsonTool.Marshal(dataSt)
	return data, err
}
