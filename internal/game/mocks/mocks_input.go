package mocks

import (
	"github.com/AishwaryJS/snakesAndLadders/common/config"
)

var (
	SimulationTest = config.Simulation{
		Runs: 1,
		Players: config.Players{
			Count: 1,
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
)
