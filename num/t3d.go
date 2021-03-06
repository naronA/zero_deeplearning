package num

import (
	"errors"

	"github.com/naronA/zero_deeplearning/vec"
)

func (t Tensor3D) Flatten() vec.Vector {
	v := make(vec.Vector, 0, len(t)*len(t[0].Vector))
	for _, e := range t {
		v = append(v, e.Vector...)
	}
	return v
}

func (t Tensor3D) Channels() int {
	return len(t)
}

func (t Tensor3D) Element(c, h, w int) float64 {
	return t[c].Element(h, w)
}

func (t Tensor3D) Assign(value float64, c, h, w int) {
	t[c].Assign(value, h, w)
}

func (t Tensor3D) Shape() (int, int, int) {
	C := t.Channels()
	H, W := t[0].Shape()
	return C, H, W
}

func (t Tensor3D) Window(x, y, h, w int) Tensor3D {
	newT3D := make(Tensor3D, len(t))
	for i, mat := range t {
		newT3D[i] = mat.Window(x, y, h, w)
	}
	return newT3D
}

func (t Tensor3D) Pad(size int) Tensor3D {
	newT3d := make(Tensor3D, len(t))
	for i, m := range t {
		newT3d[i] = m.Pad(size)
	}
	return newT3d
}

func NewRandnT3D(c, h, w int) (Tensor3D, error) {
	if c == 0 || h == 0 || w == 0 {
		return nil, errors.New("row/columns is zero")
	}
	t3d := make(Tensor3D, c)
	for i := 0; i < c; i++ {
		mat, _ := NewRandnMatrix(h, w)
		t3d[i] = mat
	}
	return t3d, nil
}

func ZerosT3D(c, h, w int) Tensor3D {
	t3d := make(Tensor3D, c)
	for i := range t3d {
		t3d[i] = Zeros(h, w)
	}
	return t3d
}

func ZerosLikeT3D(x Tensor3D) Tensor3D {
	matrixes := make(Tensor3D, len(x))
	for i, v := range x {
		matrixes[i] = ZerosLike(v)
	}
	return matrixes
}

// func (t *Tensor3D) Reshape(h, w int) Tensor3D {
// 	matrixes := Tensor3D{}
// 	for _, v := range t {
// 		matrixes = append(matrixes, v.Reshape(h, w))
// 	}
// 	return matrixes
//
// }

func calcT3d(a ArithmeticT3D, x1 interface{}, x2 interface{}) *Matrix {
	switch a {
	case ADDT3D:
		return Add(x1, x2)
	case SUBT3D:
		return Sub(x1, x2)
	case MULT3D:
		return Mul(x1, x2)
	case DIVT3D:
		return Div(x1, x2)
	}
	return nil
}

func t3dT3d(a ArithmeticT3D, x1 Tensor3D, x2 Tensor3D) Tensor3D {
	mats := make(Tensor3D, len(x1))
	x1mat := x1
	x2mat := x2
	for i := range x1 {
		mats[i] = calcT3d(a, x1mat[i], x2mat[i])
	}
	return mats
}

func t3dAny(a ArithmeticT3D, x1 Tensor3D, x2 interface{}) Tensor3D {
	mats := make(Tensor3D, len(x1))
	for i, x1mat := range x1 {
		mats[i] = calcT3d(a, x1mat, x2)
	}
	return mats
}

func anyT3d(a ArithmeticT3D, x1 interface{}, x2 Tensor3D) Tensor3D {
	tensor := ZerosLikeT3D(x2)
	mats := tensor
	for i, x2mat := range x2 {
		mats[i] = calcT3d(a, x1, x2mat)
	}
	return tensor
}

type ArithmeticT3D int

const (
	ADDT3D ArithmeticT3D = iota
	SUBT3D
	MULT3D
	DIVT3D
)

func calcArithmetic(a ArithmeticT3D, x1 interface{}, x2 interface{}) Tensor3D {
	if x1v, ok := x1.(Tensor3D); ok {
		switch x2v := x2.(type) {
		case Tensor3D:
			return t3dT3d(a, x1v, x2v)
		case *Matrix:
		case vec.Vector:
		case float64:
			return t3dAny(a, x1v, x2v)
		case int:
			return t3dAny(a, x1v, float64(x2v))
		}
	} else if x2v, ok := x2.(Tensor3D); ok {
		switch x1v := x1.(type) {
		case *Matrix:
		case vec.Vector:
		case float64:
			return anyT3d(a, x1v, x2v)
		case int:
			return anyT3d(a, float64(x1v), x2v)
		}
	}
	return nil
}

func PowT3D(x Tensor3D, p float64) Tensor3D {
	result := ZerosLikeT3D(x)
	for i, v := range x {
		result[i] = Pow(v, p)
	}
	return result
}

func SqrtT3D(x Tensor3D) Tensor3D {
	result := ZerosLikeT3D(x)
	for i, v := range x {
		result[i] = Sqrt(v)
	}
	return result
}

func AddT3D(x1 interface{}, x2 interface{}) Tensor3D {
	return calcArithmetic(ADDT3D, x1, x2)
}

func SubT3D(x1 interface{}, x2 interface{}) Tensor3D {
	return calcArithmetic(SUBT3D, x1, x2)
}

func MulT3D(x1 interface{}, x2 interface{}) Tensor3D {
	return calcArithmetic(MULT3D, x1, x2)
}

func DivT3D(x1 interface{}, x2 interface{}) Tensor3D {
	return calcArithmetic(DIVT3D, x1, x2)
}

func EqualT3D(t1, t2 Tensor3D) bool {
	for i := range t1 {
		if NotEqual(t1[i], t2[i]) {
			return false
		}
	}
	return true
}

func SoftmaxT3D(x Tensor3D) Tensor3D {
	t3d := ZerosLikeT3D(x)
	for i, v := range x {
		t3d[i] = Softmax(v)
	}
	return t3d
}

func DotT3D(x Tensor3D, y *Matrix) Tensor3D {
	result := ZerosLikeT3D(x)
	for i, v := range x {
		result[i] = Dot(v, y)
	}
	return result
}

func CrossEntropyErrorT3D(y, t Tensor3D) float64 {
	r := vec.Zeros(len(y))
	for i := range y {
		r[i] = CrossEntropyError(y[i], t[i])
	}
	return vec.Sum(r) / float64(len(y))
}

func NumericalGradientT3D(f func(vec.Vector) float64, x Tensor3D) Tensor3D {
	result := ZerosLikeT3D(x)
	for i, v := range x {
		result[i] = NumericalGradient(f, v)
	}
	return result
}

func (t Tensor3D) ToMatrix(axis int) *Matrix {
	c, h, w := t.Shape()
	if axis == 0 {
		var zeroVec vec.Vector
		// zeroVec := vec.Zeros(len(t) * len(t[0].Vector))
		for _, mat := range t {
			zeroVec = append(zeroVec, mat.Vector...)
		}
		return &Matrix{
			Vector:  zeroVec,
			Rows:    c,
			Columns: h * w,
		}
	} else if axis == 1 {
		var zeroVec vec.Vector
		for _, mat := range t {
			for i := 0; i < h; i++ {
				zeroVec = append(zeroVec, mat.SliceRow(i)...)
			}
		}
		return &Matrix{
			Vector:  zeroVec,
			Rows:    h,
			Columns: c * w,
		}
	} else if axis == 2 {
		var zeroVec vec.Vector
		for _, mat := range t {
			for i := 0; i < w; i++ {
				zeroVec = append(zeroVec, mat.SliceColumn(i)...)
			}
		}
		return &Matrix{
			Vector:  zeroVec,
			Rows:    w,
			Columns: c * h,
		}

	}
	return nil
}
