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
	for k, v := range result {
		k = url.QueryEscape(k)
		switch vv := v.(type) {
		case string:
			params = append(params, k+"="+url.QueryEscape(vv))
		case uint, int, uint8, int8, uint16, int16, uint32, int32, uint64, int64:
			params = append(params, k+"="+url.QueryEscape(fmt.Sprintf("%d", vv)))
		case float32, float64:
			params = append(params, k+"="+url.QueryEscape(fmt.Sprintf("%f", vv)))
		case bool:
			if vv {
				params = append(params, k+"=1")
			} else {
				params = append(params, k+"=0")
			}
		default:
			panic(fmt.Sprintf("Failed convert %T to string", v))
		}
	}

	sort.Strings(params)

	return strings.Join(params, "&")
}

func encodeValue(value interface{}, key string) map[string]interface{} {
	switch v := value.(type) {
	case string, uint, int, uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64, bool:
		return map[string]interface{}{key: v}
	default:
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Slice:
			result := map[string]interface{}{}
			for i := 0; i < rv.Len(); i++ {
				for vkk, vvv := range encodeValue(rv.Index(i).Interface(), fmt.Sprintf("%s[%d]", key, i)) {
					result[vkk] = vvv
				}
			}
			return result
		case reflect.Map:
			result := map[string]interface{}{}
			for _, ri := range rv.MapKeys() {
				var kkk string
				if key == "" {
					kkk = fmt.Sprintf("%v", ri.Interface())
				} else {
					kkk = fmt.Sprintf("%s[%v]", key, ri.Interface())
				}

				for vkk, vvv := range encodeValue(rv.MapIndex(ri).Interface(), kkk) {
					result[vkk] = vvv
				}
			}
			return result
		case reflect.Struct:
			result := map[string]interface{}{}
			for i := 0; i < rv.NumField(); i++ {
				kkk := fmt.Sprintf("%s[%v]", key, rv.Type().Field(i).Name)
				for vkk, vvv := range encodeValue(rv.Field(i).Interface(), kkk) {
					result[vkk] = vvv
				}
			}
			return result
		default:
			panic(fmt.Sprintf("Failed convert %T to string", v))
		}
	}
}
