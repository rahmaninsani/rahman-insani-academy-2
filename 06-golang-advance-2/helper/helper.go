package helper

func GenerateID(last int) int {
	if last == 0 {
		return 1
	}

	last++
	return last
}
