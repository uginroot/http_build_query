package http_build_query

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/url"
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	data := map[string]interface{}{
		"int":        13,
		"str":        "str",
		"bool_true":  true,
		"bool_false": false,
		"arr_int":    []int16{1, -2, 4},
		"m_arr": map[string][]int16{
			"test": []int16{1, -2, 4},
		},
		"m_m": []interface{}{
			map[string]string{"mo1": "v", "mo2": "v2"},
			map[string]string{"mo2": "v"},
		},
		"m_m_m": map[string]interface{}{
			"mm": struct{ Name string }{"name"},
		},
	}

	stringParams := Encode(data)
	paramsParts := strings.Split(stringParams, "&")

	params := []string{
		"int=13",
		"str=str",
		"bool_true=1",
		"bool_false=0",
		"arr_int[0]=1",
		"arr_int[1]=-2",
		"arr_int[2]=4",
		"m_arr[test][0]=1",
		"m_arr[test][1]=-2",
		"m_arr[test][2]=4",
		"m_m[0][mo1]=v",
		"m_m[0][mo2]=v2",
		"m_m[1][mo2]=v",
		"m_m_m[mm][Name]=name",
	}

	assertions := assert.New(t)
	for _, v := range params {
		val := strings.Split(v, "=")
		kk := url.QueryEscape(val[0])
		vv := url.QueryEscape(val[1])
		assertions.Contains(paramsParts, fmt.Sprintf("%s=%s", kk, vv))
	}

	fmt.Println(stringParams)
}
