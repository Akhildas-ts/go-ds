package main

import "fmt"

func main() {

	m := make(map[string]map[string]string)

	m["akhil"] = map[string]string{
		"name": "akhil",
		"age":  "20",
	}

	// fmt.Println(m)

	//nested map with interface,
	// if if the key exist then, (upate,delete,add like question )

	nestedMap := map[string]interface{}{
		"key1": map[string]interface{}{
			"innerKey1": "value1",
			"innerKey2": 42,
			"innerKey3": true,
		},
		"key2": map[string]interface{}{
			"innerKey1": []int{1, 2, 3, 4},
			"innerKey2": map[string]string{
				"subInnerKey1": "subValue1",
				"subInnerKey2": "subValue2",
			},
		},
		"key3": "simpleValue",
	}

	if key2, ok := nestedMap["key2"].(map[string]interface{}); ok {

		if innerkey2, ok := key2["innerKey2"].(map[string]string); ok {

			if subinner, ok := innerkey2["subInnerKey1"]; ok {
				fmt.Println("subinner", subinner)
			}

		}
	}

}
