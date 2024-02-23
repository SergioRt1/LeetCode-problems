package main

import (
	"context"
	"fmt"
	"sync"
)

const (
	limit            = 1 << 32
	concurrencyLimit = 20
)

type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Team struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}

type PlayerResponse struct {
	Player
	Teams []Team
}

type nothing struct{}

func mockRequest(ctx context.Context, id int) (Team, error) {
	// ctx may be canceled and thus stop http requests in progress
	return Team{
		ID:   id,
		Name: "some dummy",
		Players: []Player{
			{ID: 1, Name: "some"},
		},
	}, nil
}

func readInput() map[string]bool {
	return map[string]bool{
		"dummy name": true,
	}
}

func producer(ctx context.Context, responseChan chan<- Team) {
	concurrencyLimiter := make(chan nothing, concurrencyLimit) // ensures only concurrencyLimit requests are done at a given time
	var wg sync.WaitGroup                                      // wg ensures all ongoing requests are done before closing responseChan

	for i := 0; i < limit; i++ {
		select {
		case <-ctx.Done():
			wg.Wait()
			close(responseChan)
			return
		default:
			concurrencyLimiter <- nothing{}
			wg.Add(1)
			go func(ID int) {
				defer func() {
					<-concurrencyLimiter
					wg.Done()
				}()

				if team, err := mockRequest(ctx, ID); err == nil {
					responseChan <- team
				} else {
					// TODO log error, validate type of error (5XX, 4XX...), retry chain...
					fmt.Println(err)
				}
			}(i)
		}
	}
}

func consumer(responseChan <-chan Team, names map[string]bool, cancelFunc context.CancelFunc) map[string]*PlayerResponse {
	response := make(map[string]*PlayerResponse)
	nPlayers := len(names)

	for team := range responseChan {
		for _, player := range team.Players {
			if names[player.Name] {
				if playerResponse, ok := response[player.Name]; ok {
					playerResponse.Teams = append(playerResponse.Teams, team)
				} else {
					response[player.Name] = &PlayerResponse{
						Player: player,
						Teams:  []Team{team},
					}
				}
				if len(response) >= nPlayers {
					cancelFunc() // this will stop the producer and close responseChan finishing the outer loop
				}
			}
		}
	}

	return response
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	playerNames := readInput()

	teams := make(chan Team, concurrencyLimit)

	go producer(ctx, teams)
	response := consumer(teams, playerNames, cancel)

	fmt.Println(response)
}
