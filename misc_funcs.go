//adds spaces to front of string based on length to create proper fixed width string
func FixedLengthString(length int, str string) string {
	return fmt.Sprintf(fmt.Sprintf("%%%d.%ds", length, length), str)
}

//clamps integer number to provided range
func clamp(n, min, max int) int {
	if n < min {
		return min
	}
	if n > max {
		return max
	}
	return n
}

