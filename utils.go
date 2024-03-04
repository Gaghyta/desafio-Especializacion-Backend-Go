package main

import (
	"os"
	"strings"

	"github.com/Gaghyta/desafio-Especializacion-Backend-Go/internal"
)

func ReadFile1(filename string) []internal.Ticket {
	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	data := strings.Split(string(file), "\n")

	var resultado []internal.Ticket
	for i := 0; i < len(data); i++ {
		if len(data[i]) > 0 {
			file := strings.Split(string(data[i]), ",")
			//horaVuelo, _ := time.Parse("2006-01-02 15:04:05", file[4]) // Ajustar el formato según cómo se almacene la hora en el archivo
			ticket := internal.Ticket{
				ID:          file[0],
				Nombre:      file[1],
				Email:       file[2],
				PaisDestino: file[3],
				HoraVuelo:   file[4],
				Precio:      file[5],
			}
			resultado = append(resultado, ticket)
		}
	}

	return resultado
}
