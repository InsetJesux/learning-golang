package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Ejemplo testeable, permite añadir ejemplos a la documentación generada por godoc.
// Estos ejemplos se compilan por lo que te asegura que al actualizar el código fuente los ejemplos siguen funcionando
// Añadir un comentario con 'Output' en el ejemplo hará que se ejecute como un test
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
