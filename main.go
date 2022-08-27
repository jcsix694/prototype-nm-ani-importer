package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

const ANIMATION_SUB_BY_HEX = "3322"
const TOKI1OFFSET_SUB_BY_HEX = "2dd9"
const TOKI1OFFSET_MULTIPLY_BY_HEX = "24"

func main() {
	// Loads the animation json data from file
	animationJson, err := getAnimationsDataFromFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	aniType := ""
	aniId := ""
	aniValue := ""
	toki1Offset := ""

	// Gets a list of valid Animation Types
	animationTypes := validAnimationTypes(animationJson)

	// Creates new reader
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("-----------")

	// user input animation type
	for {
		err := userInputAnimationType(&aniType, buf, animationTypes)
		if err == nil {
			break
		}
		fmt.Println(err)
	}
	fmt.Println("-----------")

	// user input animation id
	for {
		err := userInputAnimationId(&aniId, buf)
		if err == nil {
			break
		}
		fmt.Println(err)
	}

	// Calculates the animation value for 01F0 from animation ID
	aniValue, err = calculateAnimationValue(aniId)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Calculates the animation value for 01F0 from animation ID
	toki1Offset, err = calculateToki1Offset(aniId)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Load up the move properties data file
	fileMoveProperties, err := os.ReadFile("01F0 - Move Properties Data.bin")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search for Animation Value in the file Move Properties
	movePropertyOffset, err := getOffset(fileMoveProperties, aniValue)

	if err != nil {
		// step 3
		os.Exit(1)
	}

	fmt.Println(toki1Offset)

	// step 5
	editDamageMod(movePropertyOffset)
}

func getOffset(file []byte, search string) (string, error) {
	offset := strings.Index(hex.EncodeToString(file), search)

	if offset == -1 {
		return "", fmt.Errorf("error finding offset for search value: %s", search)
	}

	offset = offset / 2

	return intToHex(offset, "8"), nil
}

func editDamageMod(movePropertyOffset string) string {
	// open file
	fileMoveAnimationData, err := os.ReadFile("01F1 - Move Animation Data.bin")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(fileMoveAnimationData)
	fmt.Println(movePropertyOffset)

	return "test"
}
