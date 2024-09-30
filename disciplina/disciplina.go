package disciplina

import "igorMSoares/eng-soft-strategy/disciplina/media"

type disciplina struct {
	calcMedia media.ICalcMedia
	media     float64
	nome      string
	p1        float64
	p2        float64
	situacao  string
}

func NewDisciplina(calcMedia media.ICalcMedia) *disciplina {
	return &disciplina{
		calcMedia: calcMedia,
	}
}

func (d *disciplina) CalcularMedia() {
	d.media = d.calcMedia.CalculaMedia(d.p1, d.p2)
	d.situacao = d.calcMedia.Situacao(d.media)
}

func (d *disciplina) GetCalcMedia() media.ICalcMedia {
	return d.calcMedia
}

func (d *disciplina) GetMedia() float64 {
	return d.media
}

func (d *disciplina) GetNome() string {
	return d.nome
}

func (d *disciplina) GetP1() float64 {
	return d.p1
}

func (d *disciplina) GetP2() float64 {
	return d.p2
}

func (d *disciplina) GetSituacao() string {
	return d.situacao
}

func (d *disciplina) SetCalcMedia(media media.ICalcMedia) {
	d.calcMedia = media
}

func (d *disciplina) SetMedia(media float64) {
	d.media = media
}

func (d *disciplina) SetNome(nome string) {
	d.nome = nome
}

func (d *disciplina) SetP1(p1 float64) {
	d.p1 = p1
}

func (d *disciplina) SetP2(p2 float64) {
	d.p2 = p2
}

func (d *disciplina) SetSituacao(situacao string) {
	d.situacao = situacao
}
