package game

import (
	"log"
	"sync"
	"testing"

	"github.com/AishwaryJS/snakesAndLadders/common/utils"
	"github.com/AishwaryJS/snakesAndLadders/internal/game/mocks"
)

func TestRollDice(t *testing.T) {
	for i := 0; i < 10; i++ {
		roll := getDiceRollNumber()
		if roll < 1 || roll > 6 {
			t.Errorf("Roll dice test failed")
		}
	}
}

func TestCheckIfLadderAndSnakesPresent(t *testing.T) {
	res := checkIfLaddeOrSnakePresent(mocks.SnakesArrInput, mocks.LaddersArrInput, 54)
	if !res {
		t.Errorf("TestCheckIfLadderAndSnakesPresent Failed")
	}
	res = checkIfLaddeOrSnakePresent(mocks.SnakesArrInput, mocks.LaddersArrInput, 44)
	if res {
		t.Errorf("TestCheckIfLadderAndSnakesPresent Failed")
	}
}

func TestStart(t *testing.T) {
	var wg sync.WaitGroup
	statsChannel := make(chan PlayerStats)
	var gameStats []PlayerStats
	for j := 0; j < mocks.SimulationTest.Players.Count; j++ {
		wg.Add(1)
		// Running simulation for each player concurently using goroutines
		go Start(mocks.SnakesArrInput, mocks.LaddersArrInput, j+1, statsChannel, &wg)
	}
	_ = func() error {
		for {

			// Get stats for each finished game
			stat := <-statsChannel
			gameStats = append(gameStats, stat)
			if len(gameStats) == mocks.SimulationTest.Players.Count {
				return nil
			}
		}
	}()

	//Wait for all player to finish the game
	wg.Wait()
	close(statsChannel)

	if len(gameStats) != mocks.SimulationTest.Players.Count {
		t.Errorf("start game test failed")
	}

	log.Println(utils.AsJSON(gameStats))
}
