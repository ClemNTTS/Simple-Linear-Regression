<image src="https://github.com/ClemNTTS/Simple-Linear-Regression/blob/SimpleLinearRegression/img.png" width=140/>

# Simple Linear Regression in Go

Welcome to the **Simple Linear Regression in Go** project! This project implements a basic linear regression model from scratch, using Go, to predict values based on a single feature. It's a great starting point for learning about machine learning concepts like gradient descent and model evaluation.

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Project](#running-the-project)
- [Code Overview](#code-overview)
  - [Data](#data)
  - [Model](#model)
  - [Training](#training)
- [Further Reading and Resources](#further-reading-and-resources)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This project demonstrates a simple implementation of linear regression using the Go programming language. Linear regression is a fundamental machine learning algorithm used to model the relationship between a dependent variable and one or more independent variables. Here, we focus on a single-variable case, often called simple linear regression.

## Getting Started

### Prerequisites

To run this project, you'll need:

- [Go](https://golang.org/doc/install) installed on your system (version 1.14 or higher is recommended).

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/ClemNTTS/Simple-Linear-Regression
   ```
2. Navigate to the project directory:
   ```sh
   cd Simple-Linear-Regression
   ```

### Running the Project

To run the linear regression model and see the output:

```sh
go run main.go
```

You should see output indicating the final weight, bias, predicted values, and the cost.

## Code Overview

### Data

The dataset used is a small, manually defined set of points:

```go
var (
    X = []float64{1, 2, 3, 4}
    y = []float64{2, 2.5, 3.5, 5}
)
```

### Model

The model implements the linear regression equation:
\[ y\_{\text{pred}} = w \cdot X + b \]

### Training

We train the model using gradient descent, an optimization algorithm that iteratively adjusts the weights \( w \) and bias \( b \) to minimize the mean squared error.

```go
w, b = gradientDescent(X, y, w, b, learningRate, iterations)
```

## Further Reading and Resources

To learn more about linear regression and machine learning concepts, consider exploring the following resources:

- **Books**:

  - [An Introduction to Statistical Learning](https://www.statlearning.com/)
  - [Pattern Recognition and Machine Learning](https://www.springer.com/gp/book/9780387310732) by Christopher Bishop

- **Online Courses**:

  - [Machine Learning by Andrew Ng (Coursera)](https://www.coursera.org/learn/machine-learning)
  - [Deep Learning Specialization (Coursera)](https://www.coursera.org/specializations/deep-learning)

- **Documentation and Tutorials**:
  - [Go Documentation](https://golang.org/doc/)
  - [A Visual Introduction to Machine Learning](http://www.r2d3.us/visual-intro-to-machine-learning-part-1/)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request or open an Issue if you have any suggestions or improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
