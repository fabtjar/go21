package main

import (
	"fmt"
)

type Command struct {
	Name     string
	Help     string
	Callback func()
}

func getCommands(g *Game) map[string]Command {
	return map[string]Command{
		"hit": {
			Name:     "hit",
			Help:     "Draw a card from the deck",
			Callback: g.Hit,
		},
		"stand": {
			Name:     "stand",
			Help:     "End your turn",
			Callback: g.Stand,
		},
		"hand": {
			Name:     "cards",
			Help:     "Show your hand",
			Callback: g.Hand,
		},
		"score": {
			Name:     "score",
			Help:     "Show your score",
			Callback: g.Score,
		},
		"exit": {
			Name:     "exit",
			Help:     "Exit the game",
			Callback: g.Exit,
		},
	}
}

func showHelp(commands map[string]Command) {
	fmt.Println("help:")
	for _, command := range commands {
		fmt.Printf("- %v: %v\n", command.Name, command.Help)
	}
}

func startRepl(g *Game) {
	g.Start()
	commands := getCommands(g)
	var commandName string
	for g.Playing {
		fmt.Print("> ")
		fmt.Scanf("%s", &commandName)

		if commandName == "help" {
			showHelp(commands)
			continue
		}

		command, ok := commands[commandName]

		if !ok {
			fmt.Println("Invalid command!")
			continue
		}

		command.Callback()
	}
}
