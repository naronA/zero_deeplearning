package num

import (
	"errors"
	"log"
)

type Mat64 struct {
	Array   []float64
	Rows    int
	Columns int
}

func NewMat64(row int, column int, array []float64) (*Mat64, error) {
	if row == 0 || column == 0 {
		return nil, errors.New("row/columns is zero.")
	}
	return &Mat64{
		Array:   array,
		Rows:    row,
		Columns: column,
	}, nil
}

func (m *Mat64) Element(r int, c int) float64 {
	index := r*m.Columns + c
	log.Printf("Row: %d / Column: %d / Index: %d\n", r, c, index)
	return m.Array[index]
}

func Mul(m1 *Mat64, m2 *Mat64) (*Mat64, error) {
	// 左辺の行数と、右辺の列数があっていないの掛け算できない
	if m1.Columns != m2.Rows {
		return nil, errors.New("cant multiply. lengths are not matched ")
	}
	mat := make([]float64, m1.Rows*m2.Columns)
	for i := 0; i < m1.Columns; i++ {
		for r := 0; r < m1.Rows; r++ {
			for c := 0; c < m2.Columns; c++ {
				index := r*m1.Columns + c
				mat[index] += m1.Element(r, i) * m2.Element(i, c)
			}
		}
	}
	return &Mat64{
		Array:   mat,
		Rows:    m1.Rows,
		Columns: m2.Columns,
	}, nil
}
