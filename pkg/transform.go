package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func Data2Bytes(object interface{}, field string) ([]byte, error) {
	metadata := make([]byte, 0)
	metadata, err := json.Marshal(object)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s json序列化失败", field))
	}
	return metadata, nil
}

func Bytes2Data(rb []byte, object interface{}, field string) error {
	err := json.Unmarshal(rb, object)
	if err != nil {
		return errors.New(fmt.Sprintf("%s json解析失败", field))
	}
	return nil
}

func StringToUint(s string) uint {
	id, err := strconv.Atoi(s)
	if err != nil {
		log.Error(err)
		return 0
	}
	return uint(id)
}
