package media

type aritmetica struct{}

func NewAritmetica() *aritmetica {
	return &aritmetica{}
}

func (media *aritmetica) CalculaMedia(a, b float64) float64 {
	return (a + b) / 2
}

func (a *aritmetica) Situacao(media float64) string {
	if media > 5.0 {
		return "APROVADO"
	}

	return "REPROVADO"
}
