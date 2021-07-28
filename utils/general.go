package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const TypeOfMap = "map"
const TypeOfArray = "slice"

func IsKvUpdateStructure(curValue string, oldValue string) (bool, error) {
	curMap, err := JsonStrToMap([]byte(curValue))
	if err != nil {
		return false, err
	}
	oldMap, err := JsonStrToMap([]byte(oldValue))
	if err != nil {
		return false, err
	}

	curArr, err := GetKeyArray(curMap)
	if err != nil {
		return false, err
	}
	oldArr, err := GetKeyArray(oldMap)
	if err != nil {
		return false, err
	}

	fmt.Println(curArr)
	fmt.Println(oldArr)

	if len(curArr) != len(oldArr) {
		return false, nil
	}

	res := IsArraySame(curArr, oldArr)

	return res, nil
}

func JsonStrToMap(str []byte) (map[string]interface{}, error) {
	m := make(map[string]interface{}, 4)
	err := json.Unmarshal(str, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func GetKeyArray(m map[string]interface{}) ([]string, error) {
	arr := make([]string, 1, 1)
	for k, v := range m {
		if reflect.TypeOf(v).Kind().String() != TypeOfMap {
			arr = append(arr, k)
		} else {
			brr := make([]string, 1, 1)
			brr, err := GetKeyArray(v.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			for _, brrV := range brr {
				if brrV != "" {
					arr = append(arr, k+"."+brrV)
				}
			}
		}
	}
	return arr, nil
}

func IsArraySame(arr1 []string, arr2 []string) bool {
	for _, v := range arr1 {
		if InArrayAndDeleteItem(v, arr2) {
			continue
		}
		return false
	}
	return true
}

func InArrayAndDeleteItem(value string, arr []string) bool {
	for i, v := range arr {
		if v == value {
			arr = append(arr[:i], arr[i+1:]...)
			return true
		}
	}
	return false
}
