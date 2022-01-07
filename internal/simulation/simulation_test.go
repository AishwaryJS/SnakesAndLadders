package simulation

import (
	"testing"

	"github.com/AishwaryJS/snakesAndLadders/common/utils"
	"github.com/AishwaryJS/snakesAndLadders/internal/simulation/mocks"
)

func TestSimulationStart(t *testing.T) {
	stats, results := Start(mocks.SnakesArrInput, mocks.LaddersArrInput, mocks.SimulationTest)
	if stats.TotalRolls.Minimum == 0 || stats.TotalRolls.Average == float64(0) || stats.TotalRolls.Maximum == 0 {
		t.Errorf("Simulation Test Failed")
	}
	if stats.TotalClimb.Average == float64(0) || stats.TotalClimb.Maximum == 0 {
		t.Errorf("Simulation Test Failed")
	}
	if stats.TotalSlides.Average == float64(0) || stats.TotalSlides.Maximum == 0 {
		t.Errorf("Simulation Test Failed")
	}
	if len(stats.LongestTurn) == 0 {
		t.Errorf("Simulation Test Failed")
	}
	for _, result := range results {
		if result.PlayerWon == checkWinningPlayer(result) {
			t.Errorf("Simulation Test Failed.. Wrong winning player computation")
		}
	}
}

func TestSimulationStatsCalculation(t *testing.T) {
	stats := calculateSimulationStats(mocks.PlayerStatsInput)
	if stats.TotalRolls.Maximum != 70 || stats.TotalRolls.Minimum != 51 || stats.TotalRolls.Average != 58.2 {
		t.Errorf("Calcualting simulation test failed >> incorrect total rolls calculation")
	}
	if stats.TotalClimb.Maximum != 130 || stats.TotalClimb.Minimum != 100 || stats.TotalClimb.Average != 114 {
		t.Errorf("Calcualting simulation test failed >> incorrect total climb calculation")
	}
	if stats.TotalSlides.Maximum != 140 || stats.TotalSlides.Minimum != 70 || stats.TotalSlides.Average != 88 {
		t.Errorf("Calcualting simulation test failed >> incorrect total slide calculation")
	}
	if stats.BiggestClimb != 33 || stats.BiggestSlide != 37 {
		t.Errorf("Calcualting simulation test failed >> incorrect biggest climb or slide calculation")
	}
	if !utils.IntArrayEquals(stats.LongestTurn, []int{6, 6, 6, 2}) {
		t.Errorf("Calcualting simulation test failed >> incorrect longest turn calculation")
	}
	if stats.UnluckyRolls.Maximum != 30 || stats.UnluckyRolls.Minimum != 10 || stats.UnluckyRolls.Average != 18 {
		t.Errorf("Calcualting simulation test failed >> incorrect unlucky rolls calculation")
	}
	if stats.LuckyRolls.Maximum != 20 || stats.LuckyRolls.Minimum != 10 || stats.LuckyRolls.Average != 16 {
		t.Errorf("Calcualting simulation test failed >> incorrect lucky rolls calculation")
	}
}

func checkWinningPlayer(result SimulationResult) int {
	playerWon := -1
	minTurns := 0
	for index, stats := range result.PlayerStats {
		if index == 0 {
			playerWon = stats.PlayerNumber
			minTurns = stats.TotalTurns
		} else {
			if stats.TotalTurns < minTurns {
				playerWon = stats.PlayerNumber
				minTurns = stats.TotalTurns
			}
		}
	}
	return playerWon
}
