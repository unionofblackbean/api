package common

import "reflect"

func Nil(ptrs ...interface{}) bool {
	for _, ptr := range ptrs {
		if reflect.ValueOf(ptr).Kind() == reflect.Ptr && ptr != nil {
			return false
		}
	}

	return true
}
