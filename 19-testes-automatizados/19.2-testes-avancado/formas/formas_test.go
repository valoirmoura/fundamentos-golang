package formas

import (
	"math"
	"testing"
)

// t.Run Duas Funçoes que representaram dois metodos
func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Area Recebida %f é difereten da Area Esperada %f", areaRecebida, areaEsperada) //Fatalf para o Programa
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := math.Pi * 100
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Area Recebida %f é difereten da Area Esperada %f", areaRecebida, areaEsperada) //Fatalf para o Programa
		}
	})
}
