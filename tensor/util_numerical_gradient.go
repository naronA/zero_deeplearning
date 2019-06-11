package tensor

import (
	"github.com/naronA/zero_deeplearning/vec"
)

func NumericalGradient(f func(vec.Vector) float64, x *Tensor) *Tensor {
	if len(x.Shape) == 2 {
		mat := x.Mat
		return &Tensor{
			Mat:   NumericalGradientMat(f, mat),
			Shape: x.Shape,
		}
	}
	if len(x.Shape) == 3 {
		mat := x.T3D
		return &Tensor{
			T3D:   NumericalGradientT3D(f, mat),
			Shape: x.Shape,
		}
	}
	if len(x.Shape) == 4 {
		mat := x.T4D
		return &Tensor{
			T4D:   NumericalGradientT4D(f, mat),
			Shape: x.Shape,
		}
	}
	panic(x)
}

func NumericalGradientMat(f func(vec.Vector) float64, x *Matrix) *Matrix {
	grad := vec.NumericalGradient(f, x.Vector)
	mat := &Matrix{Rows: x.Rows, Columns: x.Columns, Vector: grad}
	return mat
}

func NumericalGradientT3D(f func(vec.Vector) float64, x Tensor3D) Tensor3D {
	result := make(Tensor3D, len(x))
	for i, v := range x {
		result[i] = NumericalGradientMat(f, v)
	}
	return result
}

func NumericalGradientT4D(f func(vec.Vector) float64, x Tensor4D) Tensor4D {
	result := make(Tensor4D, len(x))
	for i, v := range x {
		result[i] = NumericalGradientT3D(f, v)
	}
	return result
}