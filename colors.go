package MasterMind

import (
	"math/rand"
	"time"
)

var colors = []string{"bleu", "vert", "orange", "rouge", "blanc", "rose", "gris"}

func pickColor() string { // pick a random color in the list "colors"
	rand.Seed(time.Now().UnixNano())
	return colors[rand.Intn(len(colors))]
}

func comb() []string {
	tab := []string{}
	for i := 0; i < 4; i++ {
		tab = append(tab,pickColor())
	}
	return tab
}
