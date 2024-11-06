package greetings

var (
	Characters                      = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/!@#$%^&*()-_=<>[]{}|;:',.?~\"\\ "
	CharToBitMapping                = make(map[int32]byte)
	BitToCharMapping                = make(map[byte]int32)
	TAG_2_LEN, TAG_3_LEN, TAG_5_LEN = 8, 8, 4
)

func init() {
	for i, char := range Characters {
		CharToBitMapping[char] = byte(i + 1)
		BitToCharMapping[byte(i+1)] = char
	}
}
