package stripansi

// String strips a string of all ansi escape sequences
func String(src string) string {
	result := ansiRegex.ReplaceAllString(src, "")
	return result
}

// Bytes strips a slice of bytes of all ansi escape sequences
func Bytes(src []byte) []byte {
	result := ansiRegex.ReplaceAll(src, []byte(""))
	return result
}
