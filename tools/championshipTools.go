package tools

import (
	"github.com/nkanish2002/token_generator"
	"strings"
)

func ChampionshipToken(size int) string {
	g := token_generator.Generator{}
	g.New()

	return strings.ToUpper(g.GetToken(size))
}