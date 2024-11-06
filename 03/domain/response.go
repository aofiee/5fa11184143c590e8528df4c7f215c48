package domain

var (
	Target = []string{
		"t-bone",
		"fat-back",
		"pastrami",
		"pork",
		"meatloaf",
		"jowl",
		"enim",
		"bresaola",
	}
)

type (
	Response struct {
		Beef map[string]int `json:"beef"`
	}
)
