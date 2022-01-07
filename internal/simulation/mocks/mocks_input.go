package mocks

import (
	"github.com/AishwaryJS/snakesAndLadders/common/config"
	"github.com/AishwaryJS/snakesAndLadders/internal/game"
)

var (
	SimulationTest = config.Simulation{
		Runs: 2,
		Players: config.Players{
			Count: 2,
		},
	}
	SnakesArrInput = []config.Snake{
		{
			From: 99,
			To:   41,
		},
		{
			From: 89,
			To:   53,
		},
		{
			From: 76,
			To:   58,
		},
		{
			From: 66,
			To:   45,
		},
		{
			From: 54,
			To:   31,
		},
		{
			From: 43,
			To:   18,
		},
		{
			From: 40,
			To:   3,
		},
		{
			From: 27,
			To:   5,
		},
	}

	LaddersArrInput = []config.Ladder{
		{
			From: 4,
			To:   25,
		},
		{
			From: 13,
			To:   46,
		},
		{
			From: 33,
			To:   49,
		},
		{
			From: 42,
			To:   63,
		},
		{
			From: 50,
			To:   69,
		},
		{
			From: 62,
			To:   81,
		},
		{
			From: 74,
			To:   92,
		},
	}

	PlayerStatsInput = []game.PlayerStats{
		{
			PlayerNumber: 1,
			TotalTurns:   30,
			TotalRolls:   50,
			TotalClimb:   100,
			TotalSlides:  80,
			BiggestClimb: 21,
			BiggestSlide: 25,
			LongestTurn:  []int{6, 6, 2},
			UnluckyRolls: 20,
			LuckyRoll:    10,
		},
		{
			PlayerNumber: 2,
			TotalTurns:   25,
			TotalRolls:   60,
			TotalClimb:   120,
			TotalSlides:  70,
			BiggestClimb: 19,
			BiggestSlide: 37,
			LongestTurn:  []int{6, 6, 6, 2},
			UnluckyRolls: 10,
			LuckyRoll:    20,
		},
		{
			PlayerNumber: 1,
			TotalTurns:   41,
			TotalRolls:   51,
			TotalClimb:   130,
			TotalSlides:  140,
			BiggestClimb: 16,
			BiggestSlide: 25,
			LongestTurn:  []int{6, 6, 1},
			UnluckyRolls: 20,
			LuckyRoll:    10,
		},
		{
			PlayerNumber: 2,
			TotalTurns:   38,
			TotalRolls:   70,
			TotalClimb:   100,
			TotalSlides:  80,
			BiggestClimb: 33,
			BiggestSlide: 25,
			LongestTurn:  []int{6, 6, 6, 1},
			UnluckyRolls: 30,
			LuckyRoll:    20,
		},
	}
)
