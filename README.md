# Go 21

Go 21 is a simple Blackjack card game made in the Go.

## Installation

Once you have Go installed, you can clone the Go 21 repository to your local machine:

```shell
git clone https://github.com/fabtjar/go21.git
```

## Running the Game

To start playing Go 21, navigate to the game's directory in your terminal and run the following command:

```shell
go run .
```

This command will start the game, and you'll be greeted with a welcome message.

## Game Commands

Go 21 supports several commands that you can use to interact with the game:

- hit: Draw a card from the deck.
- stand: End your turn.
- cards: Show your current hand.
- score: Show your current score.
- exit: Exit the game.

To use a command, simply type its name into the terminal and press Enter.

## How to Play

1. Start the game by typing start.
2. You and the dealer will be dealt two cards each. Your cards will be displayed.
3. You can use the hit command to draw additional cards from the deck and try to get as close to 21 as possible without busting.
4. Use the stand command to end your turn and let the dealer play.
5. The dealer will draw cards until their score is at least as high as yours or until they bust.
6. The game will determine the winner based on the final scores.

## Example

```
$ go run .
Welcome to Go 21!
Dealer draws a King of Spades
You draw a 4 of Clubs and a 9 of Hearts
> hit
You draw a 8 of Spades
> score
Your score: 21
> stand
You stand!
Dealer's turn!
Dealer hits and gets a 2 of Clubs
Dealer hits and gets a Ace of Clubs
Dealer hits and gets a 2 of Diamonds
Dealer hits and gets a King of Spades
Dealer BUST! Score: 25
You win!
```

## Rules

- Aces can be worth 1 or 11 points, whichever benefits you more and keeps your score below 22.
- Face cards (Jack, Queen, King) are worth 10 points.
- Number cards are worth their face value (e.g., 2 is worth 2 points).
