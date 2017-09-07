package matrix

// SquareMultiply is the naive implementation resulting form the definition of
// matrix multiplication. If the input matricies are not square an error is returned
// It runs in O(n^3) time, where n := len(input)
func SquareMultiply(a, b [][]int) ([][]int, error) {
	if err := validateMatricies(a, b); err != nil {
		return nil, err
	}
	n := len(a)
	c := newSquareMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c, nil
}

// SquareMultiplyStrassen is a naive implementation of the Strassen algorithm for maxtrix
// multiplication. It reduces the requried Multipications from 8 to 7, therefore reducing
// the asymptotic complexity at the expense of more additions.
// It runs in O(n^(lg 7)) time, where n := len(input)
func SquareMultiplyStrassen(a, b [][]int) ([][]int, error) {
	if err := validateMatriciesStrassen(a, b); err != nil {
		return nil, err
	}
	if len(a) == 1 {
		res := newSquareMatrix(1)
		res[0][0] = a[0][0] * b[0][0]
		return res, nil
	}
	a11 := submatix(a, 1, 1)
	a12 := submatix(a, 1, 2)
	a21 := submatix(a, 2, 1)
	a22 := submatix(a, 2, 2)
	b11 := submatix(b, 1, 1)
	b12 := submatix(b, 1, 2)
	b21 := submatix(b, 2, 1)
	b22 := submatix(b, 2, 2)
	s1 := sub(b12, b22)
	s2 := add(a11, a12)
	s3 := add(a21, a22)
	s4 := sub(b21, b11)
	s5 := add(a11, a22)
	s6 := add(b11, b22)
	s7 := sub(a12, a22)
	s8 := add(b21, b22)
	s9 := sub(a11, a21)
	s10 := add(b11, b12)
	p1, _ := SquareMultiplyStrassen(a11, s1)
	p2, _ := SquareMultiplyStrassen(s2, b22)
	p3, _ := SquareMultiplyStrassen(s3, b11)
	p4, _ := SquareMultiplyStrassen(a22, s4)
	p5, _ := SquareMultiplyStrassen(s5, s6)
	p6, _ := SquareMultiplyStrassen(s7, s8)
	p7, _ := SquareMultiplyStrassen(s9, s10)
	c11 := sub(add(add(p5, p4), p6), p2)
	c12 := add(p1, p2)
	c21 := add(p3, p4)
	c22 := sub(sub(add(p5, p1), p3), p7)
	return merge(c11, c12, c21, c22), nil
}

func merge(c11, c12, c21, c22 [][]int) [][]int {
	n := len(c11)
	result := newSquareMatrix(n * 2)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = c11[i][j]
			result[i][j+n] = c12[i][j]
			result[i+n][j] = c21[i][j]
			result[i+n][j+n] = c22[i][j]
		}
	}
	return result
}

func add(a, b [][]int) [][]int {
	c := newSquareMatrix(len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

func sub(a, b [][]int) [][]int {
	c := newSquareMatrix(len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			c[i][j] = a[i][j] - b[i][j]
		}
	}
	return c
}

func submatix(input [][]int, row, col int) [][]int {
	half := len(input) / 2
	sub := newSquareMatrix(half)
	inputRow := (row - 1) * half
	inputCol := (col - 1) * half
	tempCol := 0
	for i := 0; i < half; i++ {
		tempCol = inputCol
		for j := 0; j < half; j++ {
			sub[i][j] = input[inputRow][tempCol]
			tempCol++
		}
		inputRow++
	}
	return sub
}
