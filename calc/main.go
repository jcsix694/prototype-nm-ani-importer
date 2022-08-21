package main

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"golang.org/x/exp/slices"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const ANIMATION_SUB_BY_HEX = "3322"
const TOKI1OFFSET_SUB_BY_HEX = "2dd9"
const TOKI1OFFSET_MULTIPLY_BY_HEX = "24"

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

func hexToInt(input string) (int, error) {
	number, err := strconv.ParseUint(input, 16, 16)

	if err != nil {
		return 0, fmt.Errorf("error converting hex to int: %w", err)
	}

	return int(number), nil
}

func intToHex(input int) string {
	hex := strings.ToUpper(fmt.Sprintf("%04x", input))

	return hex
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

func convertByteToString(input []byte) string {
	return strings.ToUpper(strings.TrimSpace(string(input)))
}

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
	testfile, err := os.ReadFile("01F0 - Move Properties Data.bin")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(hex.EncodeToString(testfile))

	// 	fmt.Println(hex.Dump(testfile))
	// fmt.Println(bytes.Index([]byte(testfile), []byte("00 88")))

	tt, err := os.Open("01F0 - Move Properties Data.bin")
	// f, err := io.ReadSeeker("01F0 - Move Properties Data.bin")
	for {
		o2, err := tt.Seek(0, 1)
		if err != nil {
			log.Fatal(err)
		}
		b1 := make([]byte, 20)
		n1, err := tt.Read(b1)
		if err != nil {
			log.Fatal("Failed here ", err)
		}
		fmt.Printf("start:%d %d bytes: %s\n", o2, n1, string(b1))
	}

	fmt.Println("Animation Value (01F0) is " + aniValue)
	fmt.Println("Toki 1 Offset (02BF) is " + toki1Offset)
	fmt.Println("program end")
}

func userInputAnimationType(aniType *string, buf *bufio.Reader, animationTypes []string) error {
	fmt.Print("Are you wanting to import a " + strings.Join(animationTypes, ", ") + "?: ")

	input, err := buf.ReadBytes('\n')

	if err != nil {
		return err
	}

	*aniType = convertByteToString(input)

	if *aniType == "" {
		return fmt.Errorf("input cannot be empty")
	}

	if slices.Contains(animationTypes, *aniType) == false {
		return fmt.Errorf("invalid input")
	}

	return nil
}

func userInputAnimationId(aniId *string, buf *bufio.Reader) error {
	fmt.Print("Please enter an Animation ID to import into: ")

	input, err := buf.ReadBytes('\n')

	if err != nil {
		return err
	}

	*aniId = convertByteToString(input)

	if *aniId == "" {
		return fmt.Errorf("input cannot be empty")
	}

	success, _ := regexp.MatchString(`(^[0-9A-F]{4}$)`, *aniId)
	if !success {
		return fmt.Errorf("input must be a valid hex input (4 chatacters long and including characters 0-1 A-F)")
	}

	return nil
}

func calculateAnimationValue(aniId string) (string, error) {
	aniIdInt, err := hexToInt(aniId)

	if err != nil {
		return "", fmt.Errorf("error converting hex to int: %w", err)
	}

	aniValSubBy, err := hexToInt(ANIMATION_SUB_BY_HEX)

	if err != nil {
		return "", fmt.Errorf("error converting hex to int: %w", err)
	}

	return intToHex(aniIdInt - aniValSubBy), nil
}

func calculateToki1Offset(aniId string) (string, error) {
	aniIdInt, err := hexToInt(aniId)

	if err != nil {
		return "", fmt.Errorf("error converting hex to int: %w", err)
	}

	toki1OffsetSubBy, err := hexToInt(TOKI1OFFSET_SUB_BY_HEX)

	if err != nil {
		return "", fmt.Errorf("error converting hex to int: %w", err)
	}

	toki1OffsetMultiplyBy, err := hexToInt(TOKI1OFFSET_MULTIPLY_BY_HEX)

	if err != nil {
		return "", fmt.Errorf("error converting hex to int: %w", err)
	}

	var calc1 = aniIdInt - toki1OffsetSubBy

	return intToHex(calc1 * toki1OffsetMultiplyBy), nil
}

func validAnimationTypes(jsonData gjson.Result) []string {
	return removeDuplicateStr(gjson.Get(jsonData.String(), "#.animation_type"))
}
