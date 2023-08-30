package dataproviders

import (
	"math/rand"
)

type GoRandom32CharsGenerator struct{}

func (goRandom32CharsGenerator *GoRandom32CharsGenerator) Execute() string {
	output := make([]rune, 32)

	min := 0
	max := 127

	for i := 0; i < 32; i++ {
		randomNumber := rand.Intn(max-min+1) + min
		output[i] = rune(randomNumber)
	}
	return string(output)
}
