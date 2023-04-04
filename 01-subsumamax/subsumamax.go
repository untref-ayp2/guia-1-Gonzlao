package subsumamax

// Dado un arreglo devuelve la posicion inicial, la posición final y el valor
// de la subsecuencia cuya suma es máxima dentro del arreglo original
// el orden de la función SubsecuenciaSumaMaxima es de O(n^2)
func SubsecuenciaSumaMaxima(arreglo []int) (int, int, int) {
	sumaMaxima := 0
	posInicial, posFinal := -1, -1
	for i := 0; i < len(arreglo); i++ { // 1+n*(n+1)  es equivalente a n^2
		sumaLocal := 0                      // es 1
		for j := i; j < len(arreglo); j++ { // 1+n*(1+1) es equivalente a n
			sumaLocal += arreglo[j]     // es 1
			if sumaLocal > sumaMaxima { // 1+1+1 es equivalente a 1
				sumaMaxima = sumaLocal // es 1
				posInicial = i         // es 1
				posFinal = j           // es 1
			}
		}
	}
	return sumaMaxima, posInicial, posFinal // es 1
}
