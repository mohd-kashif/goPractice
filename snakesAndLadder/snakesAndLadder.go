package snakesandladder

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

// func main() {
// 	dice := NewDice(6)
// 	board := NewBoard(100)
// 	game := NewGame([]Snake{}, []Ladder{}, []Player{}, dice, board)
// 	snakes := [][]int{{97, 52}, {37, 9}, {54, 21}, {62, 21}, {21, 4}, {81, 32}, {78, 37}, {79, 16}}
// 	game.AddSnakes(snakes)
// 	ladders := [][]int{{4, 81}, {8, 37}, {22, 60}, {16, 42}, {42, 97}, {51, 90}, {24, 62}, {27, 79}}
// 	game.AddLadders(ladders)
// 	game.AddPlayer([]string{"john", "doe", "kashif"})
// 	game.Play()
// }

type Snake struct {
	Head int
	Tail int
}

func NewSnake(head int, tail int) Snake {
	return Snake{
		Head: head,
		Tail: tail,
	}
}

type Ladder struct {
	Head int
	Tail int
}

func NewLadder(head int, tail int) Ladder {
	return Ladder{
		Head: head,
		Tail: tail,
	}
}

type Player struct {
	Name            string
	CurrentPosition int
}

func NewPlayer(name string) Player {
	return Player{
		Name:            name,
		CurrentPosition: 0,
	}
}

type Dice struct {
	DiceVale int
}

func NewDice(value int) Dice {
	return Dice{
		DiceVale: value,
	}
}

func (d *Dice) RollDice() int {
	return 1 + rand.IntN(d.DiceVale-1)
}

type Board struct {
	TotalCells int
	Cells      []int
}

func NewBoard(totalCell int) Board {
	cells := make([]int, totalCell)
	return Board{
		TotalCells: totalCell,
		Cells:      cells,
	}
}

type Game struct {
	Snakes  []Snake
	Ladders []Ladder
	Players []Player
	Dice    Dice
	Board   Board
}

func NewGame(snakes []Snake, ladder []Ladder, player []Player, dice Dice, board Board) Game {
	return Game{
		Snakes:  snakes,
		Ladders: ladder,
		Players: player,
		Dice:    dice,
		Board:   board,
	}
}

func (g *Game) AddSnakes(snakes [][]int) {
	for i := 0; i < len(snakes); i++ {
		snake := NewSnake(snakes[i][0], snakes[i][1])
		g.Snakes = append(g.Snakes, snake)
	}
}

func (g *Game) AddLadders(ladder [][]int) {
	for i := 0; i < len(ladder); i++ {
		ladder := NewLadder(ladder[i][0], ladder[i][1])
		g.Ladders = append(g.Ladders, ladder)
	}
}

func (g *Game) AddPlayer(players []string) {
	for i := 0; i < len(players); i++ {
		player := NewPlayer(players[i])
		g.Players = append(g.Players, player)
	}
}

func (g *Game) Play() {
	for i := 0; i < len(g.Players); {
		player := g.Players[i]
		fmt.Println("Current Player " + player.Name)
		fmt.Println("current position " + strconv.Itoa(player.CurrentPosition))
		diceValue := g.Dice.RollDice()
		fmt.Println("Player got " + strconv.Itoa(diceValue) + " on the dice")
		nextPosition := player.CurrentPosition + diceValue
		if nextPosition > g.Board.TotalCells {
			fmt.Println("cant finish the game")
			continue
		}
		if nextPosition == g.Board.TotalCells {
			fmt.Println("player " + player.Name + " wins")
			break
		}
		for {
			snakeOrLadderFound := false
			for _, snake := range g.Snakes {
				if nextPosition == snake.Head {
					fmt.Println("-----------------snake found " + strconv.Itoa(snake.Head) + " going to " + strconv.Itoa(snake.Tail))
					player.CurrentPosition = snake.Tail
					nextPosition = snake.Tail
					snakeOrLadderFound = true
					break
				}
			}
			for _, ladder := range g.Ladders {
				if nextPosition == ladder.Head {
					fmt.Println("+++++++++++++++ladder found at " + strconv.Itoa(ladder.Head) + " going to " + strconv.Itoa(ladder.Tail))
					player.CurrentPosition = ladder.Tail
					nextPosition = ladder.Tail
					snakeOrLadderFound = true
					break
				}
			}
			if !snakeOrLadderFound {
				break
			}
		}
		player.CurrentPosition = nextPosition
		g.Players[i] = player
		fmt.Println("final position : " + strconv.Itoa(nextPosition))
		if i == len(g.Players)-1 {
			i = 0
		} else {
			i++
		}
	}
}
