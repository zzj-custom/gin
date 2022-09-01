package config

import (
	"github.com/gookit/goutil/strutil"
	"reflect"
)

func (h Handler) Handle(dataId string, content string) {
	rv := reflect.ValueOf(h)
	methodName := strutil.UpperFirst(dataId)
	methodByName := rv.MethodByName(methodName)
	methodByName.Call([]reflect.Value{reflect.ValueOf(content)})
}
