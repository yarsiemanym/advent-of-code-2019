package main

func IntToBytes(input int) []byte {
	bytes := make([]byte, 4)
	bytes[0] = byte(input)
	bytes[1] = byte(input >> 8)
	bytes[2] = byte(input >> 16)
	bytes[3] = byte(input >> 24)
	return bytes
}

func BytesToInt(input []byte) int {
	value := int(input[0])
	value |= int(input[1]) << 8
	value |= int(input[2]) << 16
	value |= int(input[3]) << 24
	return value
}
