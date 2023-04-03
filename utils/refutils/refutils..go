package refutils

import (
	"github.com/Ericwyn/EzeFormat/log"
	"reflect"
)

func ReflectCall(i any, methodName string, params []any) {
	valueS := reflect.ValueOf(i)
	if !valueS.IsValid() {
		log.E("reflect call error, can't get value of ", i)
		return
	}

	method := valueS.MethodByName(methodName)
	if !method.IsValid() {
		log.E("reflect call error, can't get method of " + valueS.String())
		return
	}

	paramList := make([]reflect.Value, 0)
	for p := range params {
		paramList = append(paramList, reflect.ValueOf(p))
	}
	log.D("reflect call, methodName:", methodName, ", paramList:", params)
	method.Call(paramList)
}
