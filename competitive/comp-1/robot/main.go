package main

import "fmt"

type robot struct {
	x float64
	y float64
}

type hole struct {
	x float64
	y float64
}

type scenario struct {
	robots []robot
	holes []hole
}

func (s scenario) compute(order int)  {

}

func parseBots() []robot  {
	var numBots int
	fmt.Scanln(&numBots)

	var bots []robot

	for i := 0; i < numBots; i++ {
		bot := robot{}
		fmt.Scanf("%f %f", bot.x, bot.y)
		bots = append(bots, bot)
	}

	return bots
}

func parseHoles() []hole  {
	var numBots int
	fmt.Scanln(&numBots)

	var holes []hole

	for i := 0; i < numBots; i++ {
		hole := hole{}
		fmt.Scanf("%f %f", hole.x, hole.y)
		holes = append(holes, hole)
	}

	return holes
}

func parseScenario() scenario {
	return scenario{
		robots: parseBots(),
		holes: parseHoles(),
	}
}



func main() {
	var numScenarios int
	fmt.Scanln(&numScenarios)

	for i := 0; i < numScenarios; i++ {

	}
}


