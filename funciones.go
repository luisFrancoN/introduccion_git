package funciones

func suma(a int, b int) int {
	return a + b
}

func resta(a int, b int) int {
	if a < 0 {
		return 0;
	}
	return a - b
}