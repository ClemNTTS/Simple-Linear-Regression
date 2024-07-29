package main

import "fmt"

var (
	X = []float64{1, 2, 3, 4}
	y = []float64{2, 2.5, 3.5,5}
)

func main() {
	//Initial vlaues
	w, b := 0.0,0.0
	learningRate := 0.01
	iterations := 1000

	//Training
	w, b = GradientDescent(X, y, w, b, learningRate, iterations)

	//Display Result
	fmt.Println("Bias : ",b,"  Weight : ",w)
	fmt.Println("Predict : ",Predict(X,w,b))
	fmt.Println("Cost : ", Cost(Predict(X,w,b),y))
}