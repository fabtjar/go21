package main

import (
	"fmt"
)

type Command struct {
	Name     string
	Help     string
	Callback func()
}

func getCommands(g *Game) []Command {
	return []Command{
		{
			Name:     "hit",
			Help:     "Draw a card from the deck",
			Callback: g.Hit,
		},
		{
			Name:     "stand",
			Help:     "End your turn",
			Callback: g.Stand,
		},
		{
			Name:     "cards",
			Help:     "Show your hand",
			Callback: g.Hand,
		},
		{
			Name:     "score",
			Help:     "Show your score",
			Callback: g.Score,
		},
		{
			Name:     "exit",
			Help:     "Exit the game",
			Callback: g.Exit,
		},
	}
}

func getCommand(commands []Command, name string) (Command, bool) {
	for _, command := range commands {
		if command.Name == name {
			return command, true
		}
	}
	return Command{}, false
}

func showHelp(commands []Command) {
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

		command, ok := getCommand(commands, commandName)

		if !ok {
			fmt.Println("Invalid command!")
			continue
		}

		command.Callback()
	}
}
