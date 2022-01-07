package main

import (
	"fmt"
	"log"

	"github.com/AishwaryJS/snakesAndLadders/common/config"
	"github.com/AishwaryJS/snakesAndLadders/common/utils"
	"github.com/AishwaryJS/snakesAndLadders/internal/input"
	"github.com/AishwaryJS/snakesAndLadders/internal/simulation"
)

func main() {
	config.InitConfig()
	snakes, err := config.InitSnakes()
	if err != nil {
		fmt.Println(err)
		return
	}
	ladders, err := config.InitLadders()
	if err != nil {
		fmt.Println(err)
		return
	}
	simulationEnv := config.InitSimulationConfig()
	err = input.Check(snakes, ladders, simulationEnv)
	if err != nil {
		fmt.Println(err)
		return
	}

	stats, simStats := simulation.Start(snakes, ladders, simulationEnv)
	log.Println(utils.AsJSON(stats))
	log.Println(utils.AsJSON(simStats))
}
