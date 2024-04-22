# Saludos en Go

Este paquete proporciona una forma simple de obtener saludos personalizados en Go

## Instalacion
Ejecuta el siguiente comando para instalar el paquete:
```bash
go get -u github.com/jeissonjf/greetings
```

## Uso
Aqui tienes un ejemplo de como usar el paquete en tu codigo:

```go
package main

import(
    "fmt"
    "github.com/jeissonjf/greetings"
)

func main() {
    message,err := greetings.Hello("jeisson")

    if err != nil{
        fmt.Println("Ocurrio un error", err)
        return
    }

    fmt.Println(message)
}
```
Este ejemplo importa el paquete github.com/jeissonjf/greetings y llama la funcion Hello
imprimiendo un saludo personalizado. Si ocurre un error, se imprime un mensaje de error
