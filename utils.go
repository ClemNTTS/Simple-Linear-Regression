package main

import "math"

func GradientDescent(X, y []float64, w, b, learningRate float64, iterations int) (float64, float64) {
	n := float64(len(y))

	for i := 0; i < iterations; i++ {
		diff_w, diff_b := 0.0, 0.0
		yPred := Predict(X, w, b)

		for j := range y {													//using gradient to put in w and b, best values as possible
			diff_w += (yPred[j] - y[j]) * X[j]  			//diff_w take the relation with the X[j] value
			diff_b += (yPred[j] - y[j])								//diff_b take only the difference with yPred and yTrue values
		}

		w -= learningRate * diff_w / n							//change w and b with previous sum, but related with number of values and steps(learningRate)
		b -= learningRate * diff_b / n
	}
	return w, b
}

func Predict(X []float64, w, b float64) []float64 {
	yPred := make([]float64, len(X))
	for i := range X {
		yPred[i] = w*X[i] + b
	}
	return yPred
}

//using the Mean Squared Error to calcul the cost
func Cost(yPred, Ytrue []float64) float64 {
	sum := 0.0
	for i := range yPred {
		sum += math.Pow(yPred[i]-Ytrue[i], 2)
	}
	return sum/float64(len(yPred))
}