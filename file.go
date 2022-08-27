package main

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"golang.org/x/exp/slices"
	"io/ioutil"
	"os"
)

func getAnimationsDataFromFile() (gjson.Result, error) {
	// Gets Animation File
	fileAnimationsJson, err := os.Open("animations.json")
	if err != nil {
		return gjson.Result{}, err
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(fileAnimationsJson)

	// Check if the file is correct
	json := string(byteValue)

	if !gjson.Valid(json) {
		return gjson.Result{}, errors.New("invalid json")
	}

	valid := gjson.Get(json, "animations")
	if !valid.Exists() {
		return gjson.Result{}, errors.New("animations does not exist in json")
	}

	return valid, nil
}

func openFile(fileName string) {
	// open file
	fileMoveAnimationData, err := os.ReadFile(fileName + ".bin")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func removeDuplicateStr(gjsonResult gjson.Result) []string {
	var result []string

	gjsonResult.ForEach(func(key, value gjson.Result) bool {
		if slices.Contains(result, value.Str) == false {
			result = append(result, value.Str)
		}

		return true
	})

	return result
}
