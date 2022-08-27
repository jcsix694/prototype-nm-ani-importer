package main

import "strings"

func convertByteToStringLower(input []byte) string {
	return strings.ToLower(convertByteToString(input))
}

func convertByteToStringUpper(input []byte) string {
	return strings.ToUpper(convertByteToString(input))
}

func convertByteToString(input []byte) string {
	return strings.TrimSpace(string(input))
}
