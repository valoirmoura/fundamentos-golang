package main

type instrumento struct {
	nome  string
	naipe string
}

func main() {

	int1 := instrumento{"Clarinet", "Wood"}
	int2 := instrumento{"Trompa", "Metais"}
	int3 := instrumento{"Violoncelo", "String"}
	int4 := instrumento{"Timpano", "Percussao"}

	var instrumentos []instrumento
	instrumentos = append(instrumentos, int1)
	instrumentos = append(instrumentos, int2)
	instrumentos = append(instrumentos, int3)
	instrumentos = append(instrumentos, int4)

}
