package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	width  = 20
	height = 10
)

type Point struct {
	x, y int
}

var (
	snake     []Point
	direction Point
	food      Point
	gameOver  bool
	score     int
)

func main() {
	initGame()

	for !gameOver {
		draw()
		input()
		update()
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Printf("Game Over! Your score: %d\n", score)
}

func initGame() {
	snake = []Point{{width / 2, height / 2}}
	direction = Point{1, 0}
	placeFood()
	gameOver = false
	score = 0
}

func placeFood() {
	food = Point{rand.Intn(width), rand.Intn(height)}
}

func update() {
	head := snake[0]
	newHead := Point{head.x + direction.x, head.y + direction.y}

	if newHead.x < 0 || newHead.x >= width || newHead.y < 0 || newHead.y >= height {
		gameOver = true
		return
	}

	for _, p := range snake {
		if p == newHead {
			gameOver = true
			return
		}
	}

	snake = append([]Point{newHead}, snake...)

	if newHead == food {
		score++
		placeFood()
	} else {
		snake = snake[:len(snake)-1]
	}
}

func draw() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if contains(snake, Point{x, y}) {
				fmt.Print("O")
			} else if food.x == x && food.y == y {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func contains(snake []Point, point Point) bool {
	for _, p := range snake {
		if p == point {
			return true
		}
	}
	return false
}

func input() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "w":
		if direction != (Point{0, 1}) {
			direction = Point{0, -1}
		}
	case "s":
		if direction != (Point{0, -1}) {
			direction = Point{0, 1}
		}
	case "a":
		if direction != (Point{1, 0}) {
			direction = Point{-1, 0}
		}
	case "d":
		if direction != (Point{-1, 0}) {
			direction = Point{1, 0}
		}
	}
}
