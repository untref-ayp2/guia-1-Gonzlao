package ordenamientos

func OrdenamientoBorbujeo(f []int) []int {

	for i := 0; i < len(f)-1; i++ {

		for j := 0; j < len(f)-i-1; j++ {

			if f[j] > f[j+1] {

				temp := f[j]
				f[j] = f[j+1]
				f[j+1] = temp
			}
		}
	}
	return f
}
