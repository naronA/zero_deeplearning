package tensor

import (
	"fmt"
	"testing"

	"github.com/naronA/zero_deeplearning/vec"
)

func SmapleT4D() *Tensor {
	t4d := make(Tensor4D, 2)
	t3d1 := make(Tensor3D, 2)
	t3d1[0] = &Matrix{Vector: vec.Vector{4, 9, 3, 6}, Rows: 2, Columns: 2}
	t3d1[1] = &Matrix{Vector: vec.Vector{7, 9, 0, 9}, Rows: 2, Columns: 2}
	t3d2 := make(Tensor3D, 2)
	t3d2[0] = &Matrix{Vector: vec.Vector{4, 7, 3, 9}, Rows: 2, Columns: 2}
	t3d2[1] = &Matrix{Vector: vec.Vector{4, 4, 1, 9}, Rows: 2, Columns: 2}
	t4d[0] = t3d1
	t4d[1] = t3d2
	return &Tensor{T4D: t4d, Shape: []int{2, 2, 2, 2}}
}

func TestTranspose(t *testing.T) {
	sample := SmapleT4D()
	actual := sample.Transpose([]int{3, 0, 1, 2})
	t4d := make(Tensor4D, 2)
	t3d1 := make(Tensor3D, 2)
	t3d1[0] = &Matrix{Vector: vec.Vector{4, 3, 7, 0}, Rows: 2, Columns: 2}
	t3d1[1] = &Matrix{Vector: vec.Vector{4, 3, 4, 1}, Rows: 2, Columns: 2}
	t3d2 := make(Tensor3D, 2)
	t3d2[0] = &Matrix{Vector: vec.Vector{9, 6, 9, 9}, Rows: 2, Columns: 2}
	t3d2[1] = &Matrix{Vector: vec.Vector{7, 9, 4, 9}, Rows: 2, Columns: 2}
	t4d[0] = t3d1
	t4d[1] = t3d2
	exp := &Tensor{T4D: t4d, Shape: []int{2, 2, 2, 2}}
	if actual.NotEqual(exp) {
		fmt.Println(actual)
		t.Fail()
	}
}

func TestIm2Col_2_2_2_1(t *testing.T) {
	t4d := SmapleT4D()

	actual := t4d.Im2Col(2, 2, 2, 1)
	expected := NewMatrix(16, 4, vec.Vector{
		0, 0, 0, 4, 0, 0, 0, 7,
		0, 0, 9, 0, 0, 0, 9, 0,
		0, 3, 0, 0, 0, 0, 0, 0,
		6, 0, 0, 0, 9, 0, 0, 0,
		0, 0, 0, 4, 0, 0, 0, 4,
		0, 0, 7, 0, 0, 0, 4, 0,
		0, 3, 0, 0, 0, 1, 0, 0,
		9, 0, 0, 0, 9, 0, 0, 0,
	})
	if actual.NotEqual(expected) {
		fmt.Println(expected, actual)
		t.Fail()
	}
}

func TestReshape(t *testing.T) {
	t4d := SmapleT4D()
	expected := t4d.ReshapeToMat(2, -1)
	actual := NewMatrix(2, 8, vec.Vector{
		4, 9, 3, 6, 7, 9, 0, 9,
		4, 7, 3, 9, 4, 4, 1, 9,
	})
	if actual.NotEqual(expected) {
		t.Fail()
	}
}
