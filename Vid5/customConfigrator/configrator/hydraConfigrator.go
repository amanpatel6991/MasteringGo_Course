package configrator

import (
	"github.com/pkg/errors"
	"reflect"
)

const (
	CUSTOM = iota
)

var typeError error = errors.New("Type must be a pointer or struct")

func GetConf(confType int , obj interface{} , filename string) error{
       //fmt.Println("in getConf")
	mysValue := reflect.ValueOf(obj)
	//checking for pointer
	if mysValue.Kind() != reflect.Ptr || mysValue.IsNil() {
		return typeError
	}

	//get and cofirm struct value
	mysValue = mysValue.Elem()
	if mysValue.Kind() != reflect.Struct {
		return typeError
	}

	var err error
	switch confType {
	case CUSTOM:
		err = MarshalCustomConfig(mysValue , filename)
	}

	return err

}