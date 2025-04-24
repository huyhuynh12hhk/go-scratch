package fwpropagate

import (
	"math"
)

type Neuron struct {
	Weights []float64
	Output  float64
}

type Network [][]Neuron


// Neuron Activation
// 
// activation = sum(weight_i * input_i) + bias
func activate(weights, inputs []float64) float64 {
	activation := weights[len(weights)-1]
	// Add weighted inputs
	for i := 0; i < len(weights)-1; i++ {
		activation += weights[i] * inputs[i]
	}

	return activation
}


// Neuron Transfer
// 
// Transfer an activation function using the sigmoid function:
// 
// output = 1 / (1 + e^(-activation))
func transfer(activation float64) float64 {

	return 1.0 / (1.0 + math.Exp(-activation))
}

// Forward Propagation
func ForwardPropagate(network Network, row []float64) []float64 {
	inputs := row

	for _, layer := range network {
		newInputs := make([]float64, 0, len(layer))
		for _, neuron := range layer {
			activation := activate(neuron.Weights, inputs)
			neuron.Output = transfer(activation)
			newInputs = append(newInputs, neuron.Output)
		}
		inputs = newInputs
	}
	return inputs
}
