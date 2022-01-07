package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	viper.AddConfigPath(os.Getenv("CONFIG_PATH"))
	viper.SetConfigName(viper.GetString("ENV") + ".config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}

type Simulation struct {
	Runs    int
	Players Players
}

type Players struct {
	Count int
}

type Snake struct {
	From int
	To   int
}

type Ladder struct {
	From int
	To   int
}

// InitSnakes inits snakes info from config
func InitSnakes() ([]Snake, error) {
	snakes := []Snake{}
	construct := viper.GetStringSlice("snakes")
	for _, val := range construct {
		from, err := strconv.Atoi(strings.Split(val, ",")[0])
		if err != nil {
			return snakes, err
		}
		to, err := strconv.Atoi(strings.Split(val, ",")[1])
		if err != nil {
			return snakes, err
		}
		snakes = append(snakes, Snake{
			From: from,
			To:   to,
		})
	}

	return snakes, nil
}

// InitLadders inits ladders info from config
func InitLadders() ([]Ladder, error) {
	ladders := []Ladder{}
	construct := viper.GetStringSlice("ladders")
	for _, val := range construct {
		from, err := strconv.Atoi(strings.Split(val, ",")[0])
		if err != nil {
			return ladders, err
		}
		to, err := strconv.Atoi(strings.Split(val, ",")[1])
		if err != nil {
			return ladders, err
		}
		ladders = append(ladders, Ladder{
			From: from,
			To:   to,
		})
	}

	return ladders, nil
}

// InitPlayersCount inits players count from config
func InitSimulationConfig() Simulation {
	simulation := Simulation{}
	err := viper.UnmarshalKey("simulation", &simulation)
	if err != nil {
		log.Fatalln("Unable to unmarshal simulation config")
	}

	return simulation
}
