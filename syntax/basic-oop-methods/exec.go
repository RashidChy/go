package methods

func Exec() {
	mobydick := book{
		title: "moby dick",
		price: 10,
	}

	minecraft := game{
		title: "minecraft",
		price: 20,
	}

	tetris := game{
		title: "tetris",
		price: 5,
	}

	// #3
	mobydick.print()  // sends `mobydick` value to `book.print`
	minecraft.print() // sends `minecraft` value to `game.print`
	tetris.print()    // sends `tetris` value to `game.print`

	//or we can print in this way as well
	game.print(minecraft)

}
