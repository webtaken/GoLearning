// El siguiente paquete tiene casos donde falla, es preferible usarlo en sistemas
// de ecuaciones lineales donde se sepa que hay una o infinitas soluciones
package main

import (
	"fmt"
	"math"
)

type Mat struct {
	m   int
	n   int
	arr []float64
}

func (mat Mat) get(i int, j int) float64 {
	return mat.arr[mat.n*i+j]
}

func (mat Mat) printMat() {
	for i := 0; i < mat.m*mat.n; i++ {
		if i != 0 && i%mat.n == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%v ", mat.arr[i])
	}
	fmt.Printf("\n")
}

// Row echelon form of the matrix
func (A Mat) rref() Mat {
	EPS := 1e-9
	R := Mat{
		m:   A.m,
		n:   A.n,
		arr: make([]float64, A.m*A.n),
	}
	copy(R.arr, A.arr)
	// descendiendo por las filas
	for r, c := 0, 0; r < R.m && c < R.n; r, c = r+1, c+1 {
		piv := r
		// seleccionamos el pivot
		for y := r + 1; y < R.m; y++ {
			if math.Abs(R.arr[y*R.n+c]) > math.Abs(R.arr[piv*R.n+c]) {
				piv = y
			}
		}

		if R.arr[piv*R.n+c] < EPS {
			continue
		}

		// una vez seleccionado el índice de la fila pivote
		// hacemos un swap de todos los valores de la fila r
		for i := 0; i < R.n; i++ {
			R.arr[r*R.n+i], R.arr[piv*R.n+i] =
				R.arr[piv*R.n+i], R.arr[r*R.n+i] // swap en go
		}

		// the diagonal element
		a_ii := R.arr[r*R.n+c]
		for i := 0; i < R.n; i++ {
			// procedemos también a dividir por el número en la diagonal actual
			R.arr[r*R.n+i] /= a_ii
		}

		R.printMat()

		// por último sumamos los elementos de la fila r-ésima a cada fila restante
		for i := r + 1; i < R.m; i++ {
			temp := R.arr[i*R.n+c]
			for j := 0; j < R.n; j++ {
				R.arr[i*R.n+j] -= temp * R.arr[r*R.n+j]
			}
		}
		R.printMat()
		fmt.Printf("\n")
	}

	// ascendiendo por la filas
	for r := R.m - 1; r >= 0; r-- {
		// primero buscaremos el elemento líder de la fila (el que no es cero en la columna)
		for c := 0; c < R.n; c++ {
			if R.arr[r*R.n+c] == 1.0 {
				// una vez encontrado el elemento líder de la fila, pasamos a restar a todas las filas
				// superiores este elemento ya que este es igual a uno para resolver el sistema de ecuaciones
				for y := r - 1; y >= 0; y-- {
					temp := R.arr[y*R.n+c]
					for x := c; x < R.n; x++ {
						R.arr[y*R.n+x] -= temp * R.arr[r*R.n+x]
					}
				}
			}
		}
	}

	return R
}

func main() {
	A := Mat{
		m:   3,
		n:   4,
		arr: make([]float64, 0),
	}
	A.arr = []float64{1, 2, 3, 5, 4, 5, 6, 5, 7, 8, 9, 5}
	A.printMat()
	fmt.Printf("%v\n", A.get(1, 0))

	R := A.rref()
	R.printMat()
}
