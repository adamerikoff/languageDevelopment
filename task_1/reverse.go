package stringutil

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func RunesCount(s string) int {
	var count int
	for i := 0; i < len(s); i++ {
		count++
	}
	return count
}
