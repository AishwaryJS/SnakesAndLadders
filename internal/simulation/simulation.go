package simulation

import (
	"sync"

	"github.com/AishwaryJS/snakesAndLadders/common/config"
	"github.com/AishwaryJS/snakesAndLadders/common/utils"
	"github.com/AishwaryJS/snakesAndLadders/internal/game"
)

type AggregatedValues struct {
	Minimum int
	Maximum int
	Average float64
}

type SimulationStats struct {
	TotalRolls   AggregatedValues
	TotalClimb   AggregatedValues
	TotalSlides  AggregatedValues
	BiggestClimb int
	BiggestSlide int
	LongestTurn  []int
	UnluckyRolls AggregatedValues
	LuckyRolls   AggregatedValues
}

type SimulationResult struct {
	SimulationNumber int
	PlayerWon        int
	PlayerStats      []game.PlayerStats
}

func Start(snakes []config.Snake, ladders []config.Ladder, simulation config.Simulation) (SimulationStats, []SimulationResult) {
	var simulationStats []game.PlayerStats
	var simulationResults []SimulationResult
	// Run each simulation one by one
	for i := 0; i < simulation.Runs; i++ {
		var wg sync.WaitGroup
		statsChannel := make(chan game.PlayerStats)
		var gameStats []game.PlayerStats
		for j := 0; j < simulation.Players.Count; j++ {
			wg.Add(1)
			// Running simulation for each player concurently using goroutines
			go game.Start(snakes, ladders, j+1, statsChannel, &wg)
		}
		_ = func() error {
			for {

				// Get stats for each finished game
				stat := <-statsChannel
				gameStats = append(gameStats, stat)
				if len(gameStats) == simulation.Players.Count {
					return nil
				}
			}
		}()
		//Wait for all player to finish the game
		wg.Wait()
		close(statsChannel)

		simulationResults = append(simulationResults, GetSimulationResult(gameStats, i+1))
		simulationStats = append(simulationStats, gameStats...)
	}

	return calculateSimulationStats(simulationStats), simulationResults
}

func GetSimulationResult(gameStats []game.PlayerStats, simulationNumber int) SimulationResult {
	minTurn := 0
	playerWon := -1
	for index, val := range gameStats {
		if index == 0 {
			minTurn = val.TotalTurns
			playerWon = val.PlayerNumber
		} else {
			if val.TotalTurns < minTurn {
				minTurn = val.TotalTurns
				playerWon = val.PlayerNumber
			}
		}
	}

	return SimulationResult{
		PlayerWon:        playerWon,
		PlayerStats:      gameStats,
		SimulationNumber: simulationNumber,
	}

}

func calculateSimulationStats(stats []game.PlayerStats) SimulationStats {
	simulationStats := SimulationStats{}
	var totalRolls, totalClimb, totalSlides, unluckyRolls, luckyRolls []int
	for index, gameStats := range stats {
		if index == 1 {
			// Assigning first encountered value as basis of comparision
			simulationStats.TotalRolls.Minimum = gameStats.TotalRolls
			simulationStats.TotalRolls.Maximum = gameStats.TotalRolls
			totalRolls = append(totalRolls, gameStats.TotalRolls)

			simulationStats.TotalClimb.Minimum = gameStats.TotalClimb
			simulationStats.TotalClimb.Maximum = gameStats.TotalClimb
			totalClimb = append(totalClimb, gameStats.TotalClimb)

			simulationStats.TotalSlides.Minimum = gameStats.TotalSlides
			simulationStats.TotalSlides.Maximum = gameStats.TotalSlides
			totalSlides = append(totalSlides, gameStats.TotalSlides)

			simulationStats.BiggestClimb = gameStats.BiggestClimb
			simulationStats.BiggestSlide = gameStats.BiggestSlide
			simulationStats.LongestTurn = gameStats.LongestTurn

			simulationStats.UnluckyRolls.Minimum = gameStats.UnluckyRolls
			simulationStats.UnluckyRolls.Maximum = gameStats.UnluckyRolls
			unluckyRolls = append(unluckyRolls, gameStats.UnluckyRolls)

			simulationStats.LuckyRolls.Minimum = gameStats.LuckyRoll
			simulationStats.LuckyRolls.Maximum = gameStats.LuckyRoll
			luckyRolls = append(luckyRolls, gameStats.LuckyRoll)
		}
		// Calculating Min and Max
		if gameStats.TotalRolls > simulationStats.TotalRolls.Maximum {
			simulationStats.TotalRolls.Maximum = gameStats.TotalRolls
		}
		if gameStats.TotalRolls < simulationStats.TotalRolls.Minimum {
			simulationStats.TotalRolls.Minimum = gameStats.TotalRolls
		}
		totalRolls = append(totalRolls, gameStats.TotalRolls)

		if gameStats.TotalClimb > simulationStats.TotalClimb.Maximum {
			simulationStats.TotalClimb.Maximum = gameStats.TotalClimb
		}
		if gameStats.TotalClimb < simulationStats.TotalClimb.Minimum {
			simulationStats.TotalClimb.Minimum = gameStats.TotalClimb
		}
		totalClimb = append(totalClimb, gameStats.TotalClimb)

		if gameStats.TotalSlides > simulationStats.TotalSlides.Maximum {
			simulationStats.TotalSlides.Maximum = gameStats.TotalSlides
		}
		if gameStats.TotalSlides < simulationStats.TotalSlides.Minimum {
			simulationStats.TotalSlides.Minimum = gameStats.TotalSlides
		}
		totalSlides = append(totalSlides, gameStats.TotalSlides)

		if gameStats.BiggestClimb > simulationStats.BiggestClimb {
			simulationStats.BiggestClimb = gameStats.BiggestClimb
		}
		if gameStats.BiggestSlide > simulationStats.BiggestSlide {
			simulationStats.BiggestSlide = gameStats.BiggestSlide
		}
		if utils.Sum(gameStats.LongestTurn) > utils.Sum(simulationStats.LongestTurn) {
			simulationStats.LongestTurn = gameStats.LongestTurn
		}

		if gameStats.UnluckyRolls > simulationStats.UnluckyRolls.Maximum {
			simulationStats.UnluckyRolls.Maximum = gameStats.UnluckyRolls
		}
		if gameStats.UnluckyRolls < simulationStats.UnluckyRolls.Minimum {
			simulationStats.UnluckyRolls.Minimum = gameStats.UnluckyRolls
		}
		unluckyRolls = append(unluckyRolls, gameStats.UnluckyRolls)

		if gameStats.LuckyRoll > simulationStats.LuckyRolls.Maximum {
			simulationStats.LuckyRolls.Maximum = gameStats.LuckyRoll
		}
		if gameStats.LuckyRoll < simulationStats.LuckyRolls.Minimum {
			simulationStats.LuckyRolls.Minimum = gameStats.LuckyRoll
		}
		luckyRolls = append(luckyRolls, gameStats.LuckyRoll)

	}
	// Calculating Avg
	simulationStats.TotalRolls.Average = utils.Average(totalRolls)
	simulationStats.TotalClimb.Average = utils.Average(totalClimb)
	simulationStats.TotalSlides.Average = utils.Average(totalSlides)
	simulationStats.LuckyRolls.Average = utils.Average(luckyRolls)
	simulationStats.UnluckyRolls.Average = utils.Average(unluckyRolls)

	return simulationStats
}
