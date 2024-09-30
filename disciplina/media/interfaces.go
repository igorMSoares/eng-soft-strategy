package media

type ICalcMedia interface {
	CalculaMedia(a, b float64) float64
	Situacao(media float64) string
}
