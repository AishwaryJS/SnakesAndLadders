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
	SnakesArrForDuplicateCheck = []config.Snake{
		{
			From: 20,
			To:   10,
		},
		{
			From: 40,
			To:   30,
		},
		{
			From: 80,
			To:   40,
		},
	}

	SnakesArrForInvalidCheck = []config.Snake{
		{
			From: 10,
			To:   20,
		},
		{
			From: 40,
			To:   30,
		},
		{
			From: 80,
			To:   40,
		},
	}

	LaddersArrForDuplicateCheck = []config.Ladder{
		{
			From: 15,
			To:   20,
		},
		{
			From: 40,
			To:   47,
		},
		{
			From: 50,
			To:   80,
		},
	}

	LaddersArrForInvalidCheck = []config.Ladder{
		{
			From: 15,
			To:   20,
		},
		{
			From: 40,
			To:   37,
		},
		{
			From: 50,
			To:   80,
		},
	}
	LaddersArrForCircularCheck = []config.Ladder{
		{
			From: 15,
			To:   20,
		},
		{
			From: 30,
			To:   40,
		},
		{
			From: 50,
			To:   80,
		},
	}
)
