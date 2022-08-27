package main

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"regexp"
	"strings"
)

func userInputAnimationType(aniType *string, buf *bufio.Reader, animationTypes []string) error {
	fmt.Print("Are you wanting to import a " + strings.Join(animationTypes, ", ") + "?: ")

	input, err := buf.ReadBytes('\n')

	if err != nil {
		return err
	}

	*aniType = convertByteToStringUpper(input)

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

	*aniId = convertByteToStringLower(input)

	if *aniId == "" {
		return fmt.Errorf("input cannot be empty")
	}

	success, _ := regexp.MatchString(`(^[0-9A-F]{4}$)`, *aniId)
	if !success {
		return fmt.Errorf("input must be a valid hex input (4 chatacters long and including characters 0-1 A-F)")
	}

	return nil
}
