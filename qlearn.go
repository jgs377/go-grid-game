package main

import (
	"math/rand"
	"sort"
)

type GameState struct {
	grid *Grid
	Coord
}

type ObservationAction struct {
	coord Coord
	action int
}

type ActionValue struct {
	action int
	value float64
}

type QLearningAgent struct {
	QTable               map[ObservationAction]float64
	NumActions           int
	LearningRate         float64
	Discount             float64
	ExplorationRate      float64
	ExplorationDecayRate float64
	UpdatesSum           float64
	// MaxChange            float64
}

func NewQLearningAgent(numActions int, learningRate float64, discount float64, explorationRate float64, explorationDecayRate float64) (agent QLearningAgent) {
	agent = QLearningAgent{
		QTable:               make(map[ObservationAction]float64),
		NumActions:           numActions,
		LearningRate:         learningRate,
		Discount:             discount,
		ExplorationRate:      explorationRate,
		ExplorationDecayRate: explorationDecayRate,
		UpdatesSum:           0.0,
	}
	return agent
}

func (agent *QLearningAgent) act(observation GameState) (action int) {
	if rand.Float64() < agent.ExplorationRate {
		var options []int
		for i := 0; i < agent.NumActions; i++ {
			if observation.grid.IsValidTile(observation.Shift(i)) {
				options = append(options, i)
			}
		}
		action = options[rand.Intn(len(options))]
		return action
	}

	var options []ActionValue
	for i := 0; i < agent.NumActions; i++ {
		if observation.grid.IsValidTile(observation.Shift(i)) {
			options = append(options, ActionValue{i, agent.QTable[ObservationAction{Coord{observation.tileX, observation.tileY}, i}]})
		}
	}

	rand.Shuffle(len(options), func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})

	sort.Slice(options, func(i, j int) bool {
		return options[i].value > options[j].value
	})

	action = options[0].action

	return action
}

func (agent *QLearningAgent) update(observation GameState, action int, newObservation GameState, reward float64) {
	observationAction := ObservationAction{
		coord: Coord{observation.tileX, observation.tileY},
		action: action,
	}
	optimalFutureValue := -9999.9
	for i := 0; i < agent.NumActions; i++ {
		temp := agent.QTable[ObservationAction{Coord{newObservation.tileX, newObservation.tileY}, i}]
		if temp > float64(optimalFutureValue) {
			optimalFutureValue = temp
		}
	}
	currentValue := agent.QTable[observationAction]
	agent.QTable[observationAction] = (1.0 - agent.LearningRate) * currentValue + agent.LearningRate * (reward + agent.Discount * optimalFutureValue)

}