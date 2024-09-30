package media

import "math"

type geometrica struct{}

func NewGeometrica() *geometrica {
	return &geometrica{}
}

func (g *geometrica) CalculaMedia(a, b float64) float64 {
	return math.Sqrt(a * b)
}

func (g *geometrica) Situacao(media float64) string {
	if media > 7.0 {
		return "APROVADO"
	}

	return "REPROVADO"
}
