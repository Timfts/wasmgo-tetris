package main

// abstract later

type Game struct {
	GameOver bool
}

func (s *Game) New() Game {
	// setup initial state
	return Game{
		GameOver: false,
	}
}

func (s *Game) update() {

}
