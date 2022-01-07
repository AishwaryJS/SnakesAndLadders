package input

import (
	"testing"

	"github.com/AishwaryJS/snakesAndLadders/internal/input/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSnakeAndLadderOnSamePositionCheck(t *testing.T) {
	assert := assert.New(t)
	err := Check(mocks.SnakesArrForDuplicateCheck, mocks.LaddersArrForDuplicateCheck, mocks.SimulationTest)
	assert.Equal(err.Error(), mocks.SnakesAndLaddersOnSamePositionError)
	if err == nil {
		t.Errorf("Test failed")
	}
	if err.Error() != mocks.SnakesAndLaddersOnSamePositionError {
		t.Errorf("Test failed")
	}
}

func TestSnakeInputCheck(t *testing.T) {
	err := Check(mocks.SnakesArrForInvalidCheck, mocks.LaddersArrForDuplicateCheck, mocks.SimulationTest)
	if err == nil {
		t.Errorf("Test failed")
	}

	if err.Error() != mocks.SnakeInvalidInput {
		t.Errorf("Test failed")
	}
}

func TestLadderInputCheck(t *testing.T) {
	err := Check(mocks.SnakesArrForDuplicateCheck, mocks.LaddersArrForInvalidCheck, mocks.SimulationTest)
	if err == nil {
		t.Errorf("Test failed")
	}
	if err != nil && err.Error() != mocks.LadderInvalidInput {
		t.Errorf("Test failed")
	}
}

func TestSnakeAndLadderCircularInputCheck(t *testing.T) {
	err := Check(mocks.SnakesArrForDuplicateCheck, mocks.LaddersArrForCircularCheck, mocks.SimulationTest)
	if err == nil {
		t.Errorf("Test failed")
	}
	if err.Error() != mocks.SnakeAndLadderCircularErrorCheck {
		t.Errorf("Test failed")
	}
}
