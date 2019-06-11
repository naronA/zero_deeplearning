package tensor

import "github.com/naronA/zero_deeplearning/vec"

type Tensor5D []Tensor4D

func (t Tensor5D) Shape() (int, int, int, int, int) {
	a := len(t)
	b := len(t[0])
	c := len(t[0][0])
	d, e := t[0][0][0].Shape()
	return a, b, c, d, e

}

func (t Tensor5D) Element(b, n, c, h, w int) float64 {
	return t[b].Element(n, c, h, w)
}

func (t Tensor5D) Assign(value float64, b, n, c, h, w int) {
	t[b].Assign(value, n, c, h, w)
}

func (t Tensor5D) ReshapeTo2D(row, col int) *Matrix {
	a, b, c, d, e := t.Shape()
	size := a * b * c * d * e
	if row == -1 {
		row = size / col
	} else if col == -1 {
		col = size / row
	}

	return &Matrix{
		Vector:  t.Flatten(),
		Rows:    row,
		Columns: col,
	}
}

func (t Tensor5D) Flatten() vec.Vector {
	v := vec.Vector{}
	for _, e := range t {
		v = append(v, e.Flatten()...)
	}
	return v
}

func ZerosT5D(a, b, c, h, w int) Tensor5D {
	t5d := make(Tensor5D, a)
	for i := range t5d {
		t5d[i] = ZerosT4D(b, c, h, w)
	}
	return t5d
}

func ZerosLikeT5D(x Tensor5D) Tensor5D {
	t5d := make(Tensor5D, len(x))
	for i, v := range x {
		t5d[i] = ZerosLikeT4D(v)
	}
	return t5d
}

func (t Tensor5D) window(x, y, h, w int) Tensor5D {
	newT5D := make(Tensor5D, len(t))
	for i, mat := range t {
		newT5D[i] = windowT4D(mat, x, y, h, w)
	}
	return newT5D
}

func (t Tensor5D) transpose(a, b, c, d, e int) Tensor5D {
	u, v, w, x, y := t.Shape()
	shape := []int{u, v, w, x, y}
	t5d := zerosT5D([]int{shape[a], shape[b], shape[c], shape[d], shape[e]})
	for i, et4d := range t {
		for j, et3d := range et4d {
			for k, emat := range et3d {
				for l := 0; l < emat.Rows; l++ {
					for n := 0; n < emat.Columns; n++ {
						oldIdx := []int{i, j, k, l, n}
						idx := make([]int, 5)
						idx[0] = oldIdx[a]
						idx[1] = oldIdx[b]
						idx[2] = oldIdx[c]
						idx[3] = oldIdx[d]
						idx[4] = oldIdx[e]
						// fmt.Println(i, j, k, l)
						// fmt.Println(" ", idx[0], idx[1], idx[2], idx[3])
						v := t.element([]int{i, j, k, l, n})
						t5d.assign(v, []int{idx[0], idx[1], idx[2], idx[3], idx[4]})
					}
				}
			}
		}
	}
	return t5d
}

func (t Tensor5D) element(point []int) float64 {
	a := point[0]
	return t[a].element(point[1:])
}

func (t Tensor5D) assign(value float64, point []int) {
	a := point[0]
	t[a].assign(value, point[1:])
}

func zerosT5D(shape []int) (t5d Tensor5D) {
	t5d = make(Tensor5D, shape[0])
	for i := range t5d {
		t5d[i] = zerosT4D(shape[1:])
	}
	return
}

func (t Tensor5D) pad(size int) Tensor5D {
	newT5D := make(Tensor5D, len(t))
	for i, t4d := range t {
		padded := t4d.pad(size)
		newT5D[i] = padded
	}
	return newT5D
}

func (t Tensor5D) equal(x Tensor5D) bool {
	for i := range t {
		if !t[i].equal(x[i]) {
			return false
		}
	}
	return true
}