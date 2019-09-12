package prime

func checkPrimeNumber(in int) bool {
	if in <= 2 {
		return true
	}
	for i := 2; i <= in/2; i++ {
		if in%i == 0 {
			return false
		}
	}
	return true
}
