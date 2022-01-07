package input

import (
	"errors"
	"fmt"

	"github.com/AishwaryJS/snakesAndLadders/common/config"
)

//Check Verifies the input and return error if there in any invalid game data detected in config
func Check(snakes []config.Snake, ladders []config.Ladder, simulation config.Simulation) error {
	laddersMap := make(map[int]int)
	snakesMap := make(map[int]int)
	//checking if player count is valid
	if simulation.Players.Count < 1 {
		return errors.New("player count must be greater than 1")
	}
	//checking if simulation is set to run atleast once
	if simulation.Runs < 1 {
		return errors.New("please provide valid simulation runs number to start the game")
	}
	//checking if snakes input is valid
	for _, val := range snakes {
		if val.From-val.To < 1 {
			return fmt.Errorf("invalid snake input >> From : %d, To : %d", val.From, val.To)
		}
		if val.From <= 0 || val.From > 100 || val.To <= 0 || val.To > 100 {
			return fmt.Errorf("invalid snake input with value >> From : %d, To : %d", val.From, val.To)
		}
		_, ok := snakesMap[val.From]
		if ok {
			return fmt.Errorf("invalid snake input, There cannot be two snakes on same position >> From : %d, To : %d", val.From, val.To)
		} else {
			snakesMap[val.From] = val.To
		}
	}
	//checking if ladders input is valid\
	for _, val := range ladders {
		if val.To-val.From < 1 {
			return fmt.Errorf("invalid ladder input >> From : %d, To : %d", val.From, val.To)
		}
		if val.From <= 0 || val.From > 100 || val.To <= 0 || val.To > 100 {
			return fmt.Errorf("invalid snake input with value >> From : %d, To : %d", val.From, val.To)
		}
		_, ok := laddersMap[val.From]
		if ok {
			return fmt.Errorf("invalid snake input, There cannot be two ladders on same position >> From : %d, To : %d", val.From, val.To)
		} else {
			laddersMap[val.From] = val.To
		}
	}

	//check for circular dependency between snakes and ladders
	for snakeFrom, snakeTo := range snakesMap {
		if ladderTo, ok := laddersMap[snakeTo]; ok {
			if ladderTo == snakeFrom {
				return fmt.Errorf("invalid input, There is ciruclar dependency of snake and ladder on position >> From : %d, To : %d", snakeFrom, snakeTo)
			}
		}
		//check that snake and ladder are not present on same position
		if _, ok := laddersMap[snakeFrom]; ok {
			return fmt.Errorf("invalid input, There cannot be snake and ladder on same position >> From : %d", snakeFrom)
		}
	}

	return nil
}
