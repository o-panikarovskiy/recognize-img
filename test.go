package main

// import (
// 	"flag"
// 	"fmt"

// 	"./nn"
// 	"gonum.org/v1/gonum/mat"
// )

// func main() {
// 	// 784 inputs - 28 x 28 pixels, each pixel is an input
// 	// 200 hidden neurons - an arbitrary number
// 	// 10 outputs - digits 0 to 9
// 	// 0.1 is the learning rate
// 	net := nn.CreateNetwork(784, 200, 10, 0.1)

// 	mnist := flag.String("mnist", "", "Either train or predict to evaluate neural network")
// 	file := flag.String("file", "", "File name of 28 x 28 PNG file to evaluate")
// 	flag.Parse()

// 	// train or mass predict to determine the effectiveness of the trained network
// 	switch *mnist {
// 	case "train":
// 		nn.MnistTrain(net)
// 	case "predict":
// 		nn.MnistPredict(net)
// 	default:
// 		// don't do anything
// 	}

// 	// predict individual digit images
// 	if *file != "" {
// 		answer, outputs := nn.ImagePredict(net, *file)
// 		fmt.Println("prediction:", answer)
// 		fa := mat.Formatted(outputs, mat.Squeeze())
// 		fmt.Printf("%v\n\n", fa)
// 	}
// }