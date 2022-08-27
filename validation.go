package main

import "github.com/tidwall/gjson"

func validAnimationTypes(jsonData gjson.Result) []string {
	return removeDuplicateStr(gjson.Get(jsonData.String(), "#.animation_type"))
}
