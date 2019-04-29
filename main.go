package main

import (
	"fmt"

	"github.com/naronA/zero_deeplearning/array"
	"github.com/naronA/zero_deeplearning/mat"
)

func initNetwork() map[string]*mat.Mat64 {
	network := map[string]*mat.Mat64{}

	network["W1"], _ = mat.NewMat64(2, 3, array.Array{
		0.1, 0.3, 0.5,
		0.2, 0.4, 0.6,
	})
	network["b1"], _ = mat.NewMat64(1, 3, array.Array{
		0.1, 0.2, 0.3,
	})
	network["W2"], _ = mat.NewMat64(3, 2, array.Array{
		0.1, 0.4,
		0.2, 0.5,
		0.3, 0.6,
	})
	network["b2"], _ = mat.NewMat64(1, 2, array.Array{
		0.1, 0.2,
	})
	network["W3"], _ = mat.NewMat64(2, 2, array.Array{
		0.1, 0.3,
		0.2, 0.4,
	})
	network["b3"], _ = mat.NewMat64(1, 2, array.Array{
		0.1, 0.2,
	})

	return network

}

func forward(network map[string]*mat.Mat64, x *mat.Mat64) array.Array {
	W1 := network["W1"]
	W2 := network["W2"]
	W3 := network["W3"]
	b1 := network["b1"]
	b2 := network["b2"]
	b3 := network["b3"]

	mul1 := x.Mul(W1)
	a1 := mul1.Add(b1)
	z1 := mat.Sigmoid(a1)

	mul2 := z1.Mul(W2)
	a2 := mul2.Add(b2)
	z2 := mat.Sigmoid(a2)

	mul3 := z2.Mul(W3)
	a3 := mul3.Add(b3)
	y := array.IdentityFunction(a3.Array)
	return y
}

func main() {
	network := initNetwork()
	x, _ := mat.NewMat64(1, 2, array.Array{
		1.0, 0.5,
	})
	y := forward(network, x)
	fmt.Println(y)
}
