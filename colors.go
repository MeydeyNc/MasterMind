package main

import (

"math/rand"
"time"
)

var colors = []string{"bleu", "vert", "orange", "rouge", "blanc", "rose", "gris"}

func pickColor() string {
    rand.Seed(time.Now().UnixNano())
    return colors[rand.Intn(len(colors))]
}