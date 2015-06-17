package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
)

func main() {

	if len(os.Args) == 2 {
		json := os.Args[1]
		body := ParseJson([]byte(json))
		Output(body)
	} else {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		if len(bytes) == 0 {
			fmt.Println("please provide json on stdin using pipes")
			os.Exit(1)
		} else {
			body := ParseJson(bytes)
			Output(body)
		}
	}
	
}

func ParseJson(b []byte) interface{} {
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	return f
}

// Consider the following JSON document:
// {
//   "foo": "bar",
//   "baz": "boa"
// }
// An equivalent SHON structure:
// nodes=foo,baz
// foo_type=string
// foo_value=bar
// baz_type=string
// baz_value=boa

// ## Complex Example ##
// Consider the following JSON document:
// {
//   "foo": [
//   "a",
//   "b",
//   {
//     "subguy": "object"
//   }
//   ],
//   "baz": "boa"
// }
// An equivalent SHON structure:
// nodes=foo,baz
// foo_type=array
// foo_length=3
// foo_0_type=string
// foo_0_value=a
// foo_1_type=string
// foo_1_value=b
// foo_2_type=node
// foo_2_nodes=subguy
// foo_2_subguy_type=string
// foo_2_subguy_value=object
// baz_type=string
// baz_value=boa

func Output(v interface{}) {
	switch v.(type) {
	case map[string]interface{}:
		OutputMap("", v)
	case []interface{}:
		OutputArray("", v)
	default:
		fmt.Println("ONOSE")
		os.Exit(1)
	}
	// fmt.Println(v)
}

func OutputMap(preface string, v interface{}) {
	nodes := preface + "nodes="
	m := v.(map[string]interface{})
	for key, value := range m {
		if nodes == preface+"nodes=" {
			nodes = nodes + key
		} else {
			nodes = nodes + "," + key
		}
		OutputSwitch(preface+key, value)
	}
	fmt.Println(nodes)

}

func OutputArray(preface string, v interface{}) {
	a := v.([]interface{})
	for index, value := range a {
		OutputSwitch(preface+strconv.Itoa(index), value)
	}
}

func OutputSwitch(key string, v interface{}) {
	switch v.(type) {
	case map[string]interface{}:
		fmt.Println(key + "_type=map")
		OutputMap(key+"_", v)
	case []interface{}:
		fmt.Println(key + "_type=array")
		fmt.Println(key + "_length="+strconv.Itoa(len(v.([]interface{}))))
		OutputArray(key+"_", v)
	case string:
		fmt.Println(key + "_type=string")
		fmt.Print(key + "_value=")
		fmt.Printf("%q\n", v.(string))
	case int:
		fmt.Println(key + "_type=int")
		fmt.Println(key + "_value=" + strconv.Itoa(v.(int)))
	case float64:
		if v.(float64) == float64(int(v.(float64))) {
			fmt.Println(key + "_type=int")
			fmt.Println(key + "_value=" + strconv.Itoa(int(v.(float64))))
		} else {
			fmt.Println(key + "_type=float")
			fmt.Println(key + "_value=" + strconv.FormatFloat(v.(float64), 'f', 4, 64))
		}
	case bool:
		fmt.Println(key + "_type=bool")
		fmt.Println(key + "_value=" + strconv.FormatBool(v.(bool)))
	case nil:
		fmt.Println(key + "_type=nil")
		fmt.Println(key + "_value=")
	default:
		fmt.Println("I dont know how to deal with ", reflect.TypeOf(v))
		os.Exit(1)
	}
}
