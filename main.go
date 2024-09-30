package main

import (
	"fmt"
	"igorMSoares/eng-soft-strategy/disciplina"
	"igorMSoares/eng-soft-strategy/disciplina/media"
	"os"
)

func exitWithError() {
	fmt.Println(">> Argumento inválido.\n>> Informe -a para média aritmética ou -g para média geométrica.")
	os.Exit(1)
}

func main() {
	args := os.Args

	if len(args) < 2 {
		exitWithError()
	}

	var calculo media.ICalcMedia

	switch args[1] {
	case "-a":
		calculo = media.NewAritmetica()
	case "-g":
		calculo = media.NewGeometrica()
	default:
		exitWithError()
	}

	d := disciplina.NewDisciplina(calculo)

	d.SetNome("Padrões de Desenvolvimento")
	d.SetP1(10)
	d.SetP2(5)
	d.CalcularMedia()

	fmt.Printf("P1: %.2f P2: %.2f Media: %.2f Situação: %s\n",
		d.GetP1(), d.GetP2(), d.GetMedia(), d.GetSituacao())
}
