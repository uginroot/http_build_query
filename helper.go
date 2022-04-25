package http_build_query

import (
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

func Encode(data map[string]interface{}) string {
	return encode(data)
}

func encode(data interface{}) string {
	result := encodeValue(data, "")
	var params []string
	for key, value := range result {
		key = url.QueryEscape(key)
		switch vv := value.(type) {
		case string:
			params = append(params, key+"="+url.QueryEscape(vv))
		case uint, int, uint8, int8, uint16, int16, uint32, int32, uint64, int64:
			params = append(params, key+"="+url.QueryEscape(fmt.Sprintf("%d", vv)))
		case float32, float64:
			params = append(params, key+"="+url.QueryEscape(fmt.Sprintf("%f", vv)))
		case bool:
			if vv {
				params = append(params, key+"=1")
			} else {
				params = append(params, key+"=0")
			}
		default:
			panic(fmt.Sprintf("Failed convert %T to string", value))
		}
	}

	sort.Strings(params)

	return strings.Join(params, "&")
}

func encodeValue(value interface{}, key string) map[string]interface{} {
	switch typedValue := value.(type) {
	case string, uint, int, uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64, bool:
		return map[string]interface{}{key: typedValue}
	default:
		reflectValue := reflect.ValueOf(typedValue)
		switch reflectValue.Kind() {
		case reflect.Slice:
			result := map[string]interface{}{}
			for i := 0; i < reflectValue.Len(); i++ {
				oneDimensionalSlice := encodeValue(reflectValue.Index(i).Interface(), fmt.Sprintf("%s[%d]", key, i))
				for resultKey, resultValue := range oneDimensionalSlice {
					result[resultKey] = resultValue
				}
			}
			return result
		case reflect.Map:
			result := map[string]interface{}{}
			for _, reflectMapKey := range reflectValue.MapKeys() {
				var prefixKey string
				if key == "" {
					prefixKey = fmt.Sprintf("%v", reflectMapKey.Interface())
				} else {
					prefixKey = fmt.Sprintf("%s[%v]", key, reflectMapKey.Interface())
				}
				oneDimensionalMap := encodeValue(reflectValue.MapIndex(reflectMapKey).Interface(), prefixKey)
				for resultKey, resultValue := range oneDimensionalMap {
					result[resultKey] = resultValue
				}
			}
			return result
		case reflect.Struct:
			result := map[string]interface{}{}
			for i := 0; i < reflectValue.NumField(); i++ {
				keyPrefix := fmt.Sprintf("%s[%v]", key, reflectValue.Type().Field(i).Name)
				oneDimensionalStruct := encodeValue(reflectValue.Field(i).Interface(), keyPrefix)
				for resultKey, resultValue := range oneDimensionalStruct {
					result[resultKey] = resultValue
				}
			}
			return result
		default:
			panic(fmt.Sprintf("Failed convert %T to string", typedValue))
		}
	}
}
