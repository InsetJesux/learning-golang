package main

import "fmt"

// Se pueden agrupar las constantes en un bloque
const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

// Se crea automáticamente la variable prefix con un valor predeterminado, en el caso de string es ""
// Si el método empieza con una letra minúscula es privado, si empieza con una letra mayúscula es público
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	// No es obligatorio especificar la variable a retornar, se usa por defecto la variable especificada en la firma del método, en este caso prefix
	return
}

func main() {
	fmt.Println(Hello("Go", ""))
}
