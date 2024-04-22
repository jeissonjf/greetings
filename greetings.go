package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello devuelve un saludo para la persona especifica
func Hello(name string) (string, error) {

	// Devuelve un mensaje si el nombre es un string vacio
	if name == "" {
		return "", errors.New("nombre vacio")
	}

	//Devuelve un saludo que incluye el nombre en un mensaje.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// hellos devuelve un mapa que asocia cada nombre con un mensaje
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

// randomFormat devuelve uno o varios formatos de mensajes de saludo
// que se seleccionan de forma aleatoria
func randomFormat() string {
	formats := []string{
		"!Hola, %v¡ !Bienvenido Parcero¡",
		"¡Que bueno verte %v!",
		"¡Como son vueltas %v!",
		"¡Saludo, %v! ¡Encantado de conocerte!",
	}

	return formats[rand.Intn(len(formats))]

}
