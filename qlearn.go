package main

type QLearningAgent struct {
	QTable [][]float64
	Alpha float64
	Gamma float64
	Epsilon float64
	NumActions int
	NumStates int
}

