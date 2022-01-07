package game

import (
	"math/rand"
	"sync"
	"time"

	"github.com/AishwaryJS/snakesAndLadders/common/config"
	"github.com/AishwaryJS/snakesAndLadders/common/utils"
)

type PlayerStats struct {
	PlayerNumber int
	TotalTurns   int
	TotalRolls   int
	TotalClimb   int
	TotalSlides  int
	BiggestClimb int
	BiggestSlide int
	LongestTurn  []int
	UnluckyRolls int
	LuckyRoll    int
}

func Start(snakes []config.Snake, ladders []config.Ladder, playerNumber int, ch chan<- PlayerStats, wg *sync.WaitGroup) {
	defer wg.Done()
	//position on start of game for each player
	currentPos := 0

	//intializing stats for player
	stats := &PlayerStats{}
	stats.PlayerNumber = playerNumber

	// running game till player reaches 100
	for currentPos != 100 {
		rolls := []int{}
		movesForCurrentTurn := 0
		stats.TotalTurns = stats.TotalTurns + 1
		diceRoll := getDiceRollNumber()
		stats.TotalRolls = stats.TotalRolls + 1
		rolls = append(rolls, diceRoll)

		// player get 6 roll agian till he gets any other number
		if diceRoll == 6 {
			movesForCurrentTurn = movesForCurrentTurn + 6
			for diceRoll == 6 {
				diceRoll = getDiceRollNumber()
				stats.TotalRolls = stats.TotalRolls + 1
				rolls = append(rolls, diceRoll)
				movesForCurrentTurn = movesForCurrentTurn + diceRoll
			}
		} else {
			movesForCurrentTurn = diceRoll
		}

		// move player pieces computed no. of positions
		intialPos := currentPos + movesForCurrentTurn
		if intialPos > 100 {
			intialPos = currentPos
			continue
		}
		finalPos := getNewPositionIfLadderOrSnakePresent(snakes, ladders, intialPos, stats)
		stats.LongestTurn = updateStatsIfLongestTurn(rolls, stats.LongestTurn)
		currentPos = finalPos
	}
	ch <- *stats
}

func updateStatsIfLongestTurn(rolls []int, longestTurn []int) []int {
	if len(longestTurn) == 0 {
		return rolls
	}
	rollsTurn := utils.Sum(rolls)
	currentLongestTurn := utils.Sum(longestTurn)
	if currentLongestTurn > rollsTurn {
		return longestTurn
	}
	return rolls
}

func getNewPositionIfLadderOrSnakePresent(snakes []config.Snake, ladders []config.Ladder, pos int, stats *PlayerStats) int {
	for checkIfLaddeOrSnakePresent(snakes, ladders, pos) {
		for _, val := range ladders {
			if val.From == pos {
				pos = val.To
				stats.LuckyRoll++
				stats.TotalClimb = stats.TotalClimb + (val.To - val.From)
				if val.To-val.From > stats.BiggestClimb {
					stats.BiggestClimb = val.To - val.From
				}
			}
		}
		for _, val := range snakes {
			if val.From == pos {
				pos = val.To
				stats.UnluckyRolls++
				stats.TotalSlides = stats.TotalSlides + (val.From - val.To)
				if val.From-val.To > stats.BiggestSlide {
					stats.BiggestSlide = val.From - val.To
				}
			}
		}
	}
	return pos
}

func checkIfLaddeOrSnakePresent(snakes []config.Snake, ladders []config.Ladder, pos int) bool {
	constructPresent := false
	for _, val := range ladders {
		if val.From == pos {
			constructPresent = true
		}
	}
	for _, val := range snakes {
		if val.From == pos {
			constructPresent = true
		}
	}
	return constructPresent
}

func getDiceRollNumber() int {
	min := 1
	max := 6
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
