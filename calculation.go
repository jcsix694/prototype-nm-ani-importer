package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hexToInt(input string) (int, error) {
	number, err := strconv.ParseUint(input, 16, 16)

	if err != nil {
		return 0, fmt.Errorf("error converting hex to int: %w", err)
	}

	return int(number), nil
}

func intToHex(input int, length string) string {
	hex := strings.ToLower(fmt.Sprintf("%0"+length+"x", input))

	return hex
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

	hex := intToHex(aniIdInt-aniValSubBy, "4")

	if hex[0:1] == "-" {
		return "", fmt.Errorf("error converting hex to int: %s", "Animation Value is a negative number")
	}

	return hex, nil
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

	return intToHex(calc1*toki1OffsetMultiplyBy, "4"), nil
}
