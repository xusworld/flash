package util

import "reflect"

func Contains(slice interface{}, target interface{}) bool {
	sliceVal := reflect.ValueOf(slice)
	kind := sliceVal.Kind()

	if kind == reflect.Slice || kind == reflect.Array {

		for idx := 0; idx < sliceVal.Len(); idx++ {
			if sliceVal.Index(idx).Interface() == target {
				return true
			}
		}
	}

	return false
}
