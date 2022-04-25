# http_build_query
go 版本的http_build_query 实现


install
~~~~shell
go get github.com/uginroot/http_build_query
~~~~
code
~~~~go
data := map[string]interface{}{
    "int":     1,
    "str":     "str",
	"bool_true":  true,
    "bool_false": false,
    "arr_int": []int16{1, -2, 4},
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
str := Encode(data)
// Original
// arr_int%5B0%5D=1&
// arr_int%5B1%5D=-2&
// arr_int%5B2%5D=4&
// bool_false=0&
// bool_true=1&
// int=13&
// m_arr%5Btest%5D%5B0%5D=1&
// m_arr%5Btest%5D%5B1%5D=-2&
// m_arr%5Btest%5D%5B2%5D=4&
// m_m%5B0%5D%5Bmo1%5D=v&
// m_m%5B0%5D%5Bmo2%5D=v2&
// m_m%5B1%5D%5Bmo2%5D=v&
// m_m_m%5Bmm%5D%5BName%5D=name&
// str=str
// 
// Preview
// int=1&
// str=str&
// bool_true=1&
// bool_false=0&
// arr_int[0]=1&
// arr_int[1]=-2&
// arr_int[2]=4&
// m_arr[test][0]=1&
// m_arr[test][1]=-2&
// m_arr[test][2]=4&
// m_m[0][mo1]=v&
// m_m[0][mo2]=v2&
// m_m[1][mo2]=v&
// m_m_m[mm][Name]=name
~~~~
