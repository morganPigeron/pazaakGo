package pazaak

type card struct {
	value int8
	name  string
}

var mainCards = [...]card{
	{value: 1, name: "normal"},
	{value: 2, name: "normal"},
	{value: 3, name: "normal"},
	{value: 4, name: "normal"},
	{value: 5, name: "normal"},
	{value: 6, name: "normal"},
	{value: 7, name: "normal"},
	{value: 8, name: "normal"},
	{value: 9, name: "normal"},
	{value: 10, name: "normal"},
}

var sideCards = [...]card{
	{value: -6, name: "side"},
	{value: -5, name: "side"},
	{value: -4, name: "side"},
	{value: -3, name: "side"},
	{value: -2, name: "side"},
	{value: -1, name: "side"},
	{value: 6, name: "side"},
	{value: 5, name: "side"},
	{value: 4, name: "side"},
	{value: 3, name: "side"},
	{value: 2, name: "side"},
	{value: 1, name: "side"},
	{value: 6, name: "switchable"},
	{value: 5, name: "switchable"},
	{value: 4, name: "switchable"},
	{value: 3, name: "switchable"},
	{value: 2, name: "switchable"},
	{value: 1, name: "switchable"},
}

var testSideDeck = []card{
	{value: -6, name: "side"},
	{value: -5, name: "side"},
	{value: -4, name: "side"},
	{value: -3, name: "side"},
	{value: -2, name: "side"},
	{value: -1, name: "side"},
	{value: 6, name: "side"},
	{value: 5, name: "side"},
	{value: 4, name: "side"},
	{value: 3, name: "side"},
}
