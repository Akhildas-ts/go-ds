package main

import "fmt"

func main() {

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

		if innerKey2, ok := key2["innerKey2"].(map[string]string); ok {

			if _, ok := innerKey2["subInnerKey1"]; ok {

				innerKey2["subInnerKey1"] = "value marrii mwone"
			}

		}
	}

	fmt.Println(nestedMap)

}
