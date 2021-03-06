package tensor

import (
	"github.com/naronA/zero_deeplearning/vec"
)

type Tensor4D []Tensor3D

func (t Tensor4D) Size() int {
	n, c, h, w := t.Shape()
	return n * c * h * w
}

func (t Tensor4D) element(n, c, h, w int) float64 {
	return t[n].element(c, h, w)
}

func (t Tensor4D) assign(value float64, n, c, h, w int) {
	t[n].assign(value, c, h, w)
}

func (t Tensor4D) flatten() vec.Vector {
	v := make(vec.Vector, 0, len(t)*len(t[0].flatten()))
	for _, e := range t {
		v = append(v, e.flatten()...)
	}
	return v
}

func (t Tensor4D) Shape() (int, int, int, int) {
	N := len(t)
	C, H, W := t[0].Shape()
	return N, C, H, W
}

func zerosT4D(n, c, h, w int) Tensor4D {
	t4d := make(Tensor4D, n)
	for i := range t4d {
		t4d[i] = zerosT3D(c, h, w)
	}
	return t4d
}

func zerosLikeT4D(x Tensor4D) Tensor4D {
	t4d := make(Tensor4D, len(x))
	for i, v := range x {
		t4d[i] = zerosLikeT3D(v)
	}
	return t4d
}

type Tensor4DIndex struct {
	N int
	C int
	H int
	W int
}

type Tensor4DSlice struct {
	Actual   Tensor4D
	Indices  []*Tensor4DIndex
	NewShape []int
}

func (t4s *Tensor4DSlice) ToTensor4D() Tensor4D {
	newT4D := zerosT4D(t4s.NewShape[0], t4s.NewShape[1], t4s.NewShape[2], t4s.NewShape[3])
	for i, idx := range t4s.Indices {
		val := t4s.Actual[idx.N][idx.C].element(idx.H, idx.W)
		matrixLength := t4s.NewShape[2] * t4s.NewShape[3]
		newMatIdx := i - idx.C*matrixLength - idx.N*(matrixLength*t4s.NewShape[1])
		newT4D[idx.N][idx.C].Vector[newMatIdx] = val
	}
	return newT4D
}

func (t Tensor4D) window(x, y, h, w int) Tensor4D {
	newT4D := make(Tensor4D, len(t))
	for i, t3d := range t {
		newT4D[i] = t3d.window(x, y, h, w)
	}
	return newT4D
}

func (t Tensor4D) transpose(a, b, c, d int) Tensor4D {
	w, x, y, z := t.Shape()
	shape := []int{w, x, y, z}
	t4d := zerosT4D(shape[a], shape[b], shape[c], shape[d])
	for i, e1t3d := range t {
		for j, e2mat := range e1t3d {
			for k := 0; k < e2mat.Rows; k++ {
				for l := 0; l < e2mat.Columns; l++ {
					oldIdx := []int{i, j, k, l}
					idx := make([]int, 4)
					idx[0] = oldIdx[a]
					idx[1] = oldIdx[b]
					idx[2] = oldIdx[c]
					idx[3] = oldIdx[d]
					v := t.element(i, j, k, l)
					t4d.assign(v, idx[0], idx[1], idx[2], idx[3])
				}
			}
		}
	}
	return t4d
}

func addAssignT4D(t1 *Tensor4DSlice, t2 Tensor4D) {
	t2flat := t2.flatten()
	for i, idx := range t1.Indices {
		add := t1.Actual[idx.N][idx.C].element(idx.H, idx.W) + t2flat[i]
		t1.Actual[idx.N][idx.C].assign(add, idx.H, idx.W)
	}
}

func (t Tensor4D) pad(size int) Tensor4D {
	newT4D := make(Tensor4D, len(t))
	for i, t3d := range t {
		padded := t3d.pad(size)
		newT4D[i] = padded
	}
	return newT4D
}

func (t Tensor4D) abs() Tensor4D {
	t4d := make([]Tensor3D, len(t))
	for i, t3d := range t4d {
		t4d[i] = t3d.abs()
	}
	return t4d
}

func (t Tensor4D) argMaxAll() int {
	max := 0
	for _, t3d := range t {
		max += t3d.argMaxAll()
	}
	return max
}

func (t Tensor4D) crossEntropyError(x Tensor4D) float64 {
	r := vec.Zeros(len(t))
	for i := range t {
		r[i] = t[i].crossEntropyError(x[i])
	}
	return vec.Sum(r) / float64(len(t))
}

func (t Tensor4D) equal(x Tensor4D) bool {
	for i := range t {
		if !t[i].equal(x[i]) {
			return false
		}
	}
	return true
}

func t4DT4D(a Arithmetic, x1 Tensor4D, x2 Tensor4D) Tensor4D {
	t4d := make(Tensor4D, len(x1))
	for i := range x1 {
		t4d[i] = t3DT3D(a, x1[i], x2[i])
	}
	return t4d
}

func t4DT3D(a Arithmetic, x1 Tensor4D, x2 Tensor3D) Tensor4D {
	t4d := make(Tensor4D, len(x1))
	for i := range x1 {
		t4d[i] = t3DT3D(a, x1[i], x2)
	}
	return t4d
}

func t4DMat(a Arithmetic, x1 Tensor4D, x2 *Matrix) Tensor4D {
	t4d := make(Tensor4D, len(x1))
	for i := range x1 {
		t4d[i] = t3DMat(a, x1[i], x2)
	}
	return t4d
}

func t4DVec(a Arithmetic, x1 Tensor4D, x2 vec.Vector) Tensor4D {
	t4d := make(Tensor4D, len(x1))
	for i := range x1 {
		t4d[i] = t3DVec(a, x1[i], x2)
	}
	return t4d
}

func t4DFloat(a Arithmetic, x1 Tensor4D, x2 float64) Tensor4D {
	t4d := make(Tensor4D, len(x1))
	for i := range x1 {
		t4d[i] = t3DFloat(a, x1[i], x2)
	}
	return t4d
}

func t3DT4D(a Arithmetic, x1 Tensor3D, x2 Tensor4D) Tensor4D {
	return t4DT3D(a, x2, x1)
}

func matT4D(a Arithmetic, x1 *Matrix, x2 Tensor4D) Tensor4D {
	return t4DMat(a, x2, x1)
}

func vecT4D(a Arithmetic, x1 vec.Vector, x2 Tensor4D) Tensor4D {
	return t4DVec(a, x2, x1)
}

func floatT4D(a Arithmetic, x1 float64, x2 Tensor4D) Tensor4D {
	return t4DFloat(a, x2, x1)
}
func (t Tensor4D) exp() Tensor4D {
	t4d := make([]Tensor3D, len(t))
	for i, t3d := range t4d {
		t4d[i] = t3d.exp()
	}
	return t4d
}

func (t Tensor4D) log() Tensor4D {
	t4d := make([]Tensor3D, len(t))
	for i, t3d := range t4d {
		t4d[i] = t3d.log()
	}
	return t4d
}

func (t Tensor4D) maxAll() float64 {
	max := 0.0
	for _, t3d := range t {
		max += t3d.maxAll()
	}
	return max
}

func (t Tensor4D) meanAll() float64 {
	return t.sumAll() / float64(len(t))
}

func (t Tensor4D) pow(p float64) Tensor4D {
	t4d := make([]Tensor3D, len(t))
	for i, t3d := range t {
		t4d[i] = t3d.pow(p)
	}
	return t4d
}

func (t Tensor4D) sumAll() float64 {
	sum := 0.0
	for _, t3d := range t {
		sum += t3d.sumAll()
	}
	return sum
}

func (t Tensor4D) sqrt() Tensor4D {
	t4d := make([]Tensor3D, len(t))
	for i, t3d := range t {
		t4d[i] = t3d.sqrt()
	}
	return t4d
}

func (t Tensor4D) softmax() Tensor4D {
	t4d := make([]Tensor3D, len(t))
	for i, t3d := range t {
		t4d[i] = t3d.softmax()
	}
	return t4d
}

func (t Tensor4D) sigmoid() Tensor4D {
	t4d := make([]Tensor3D, len(t))
	for i, t3d := range t {
		t4d[i] = t3d.sigmoid()
	}
	return t4d
}

func (t Tensor4D) relu() Tensor4D {
	t4d := make([]Tensor3D, len(t))
	for i, t3d := range t {
		t4d[i] = t3d.relu()
	}
	return t4d
}

func (t Tensor4D) numericalGradient(f func(vec.Vector) float64) Tensor4D {
	result := make(Tensor4D, len(t))
	for i, v := range t {
		result[i] = v.numericalGradient(f)
	}
	return result
}

func (t Tensor4D) slice(y, yMax, x, xMax int) Tensor4D {
	t4dslice := t.strideSlice(y, yMax, x, xMax, 1)
	sliced := t4dslice.ToTensor4D()
	return sliced
}

func (t Tensor4D) strideSlice(y, yMax, x, xMax, stride int) *Tensor4DSlice {
	indLen := 0
	for _, imgT3D := range t {
		for k := 0; k < len(imgT3D); k++ {
			for i := y; i < yMax; i += stride {
				for j := x; j < xMax; j += stride {
					indLen++
				}
			}
		}
	}

	indices := make([]*Tensor4DIndex, 0, indLen)
	totalRows := (yMax - y) / stride
	totalColumns := (xMax - x) / stride
	for n, imgT3D := range t {
		for c := range imgT3D {
			for i := y; i < yMax; i += stride {
				for j := x; j < xMax; j += stride {
					index := &Tensor4DIndex{
						N: n,
						C: c,
						H: i,
						W: j,
					}
					indices = append(indices, index)
				}
			}
		}
	}
	n, c, _, _ := t.Shape()
	return &Tensor4DSlice{
		Actual:   t,
		Indices:  indices,
		NewShape: []int{n, c, totalRows, totalColumns},
	}
}

func (t Tensor4D) im2Col(fw, fh, stride, pad int) *Matrix {
	t4d := t
	nVLen := 0
	for _, t3d := range t4d {
		for x := 0; x <= t3d[0].Columns-fw+2*pad; x += stride {
			for y := 0; y <= t3d[0].Rows-fh+2*pad; y += stride {
				for i := 0; i < len(t3d); i++ {
					nVLen++
				}
			}
		}
	}

	colVec := make(vec.Vector, 0, len(t4d))
	for _, t3d := range t4d {
		nV := make(vec.Vector, 0, nVLen)
		for x := 0; x <= t3d[0].Columns-fw+2*pad; x += stride {
			for y := 0; y <= t3d[0].Rows-fh+2*pad; y += stride {
				for _, ma := range t3d {
					padE := ma.pad(pad)
					nV = append(nV, padE.window(x, y, fw, fh).Vector...)
				}
			}
		}
		colVec = append(colVec, nV...)
	}

	col := fw * fh
	row := len(colVec) / col

	return &Matrix{
		Vector:  colVec,
		Rows:    row,
		Columns: col,
	}
}

func (t Tensor4D) reshapeToMat(row, col int) *Matrix {
	t4d := t
	size := t4d.Size()
	if col == -1 {
		col = size / row
	} else if row == -1 {
		row = size / col
	}
	flat := t4d.flatten()
	return &Matrix{
		Vector:  flat,
		Rows:    row,
		Columns: col,
	}
}
