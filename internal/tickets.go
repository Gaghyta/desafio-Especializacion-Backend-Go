// package tickets
package internal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errNotFoundTicket = errors.New("Not found ticket for this destination")
	errNotFoundTime   = errors.New("Not found time")
)

type Ticket struct {
	ID          string
	Nombre      string
	Email       string
	PaisDestino string
	HoraVuelo   string
	Precio      string
}

type Storage struct {
	Tickets []Ticket
}

// GetTotalTickets devuelve, para un destino, la cantidad total de tickets
func (s *Storage) GetTotalTickets(destination string) (int, error) {
	totalTickets := 0

	for _, ticket := range s.Tickets {
		if ticket.PaisDestino == destination {
			totalTickets++
		}
	}

	if totalTickets > 0 {
		return totalTickets, nil
	} else {
		return 0, ErrNotFoundTicket
	}
}

// Calcula cuantas personas viajan en los turnos pedidos
func (s *Storage) GetCountByPeriod(time string) (int, error) {

	totalPersonas := 0

	for i := 1; i < len(s.Tickets); i++ {
		ticket := s.Tickets[i]
		horaString := strings.Split(string(ticket.HoraVuelo), ":")[0]
		hora, err := strconv.Atoi(horaString)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}

		switch {
		case time == "madrugada":
			if hora >= 0 && hora < 7 {
				totalPersonas++
			}

		case time == "mañana":
			if hora >= 7 && hora < 13 {
				totalPersonas++
			}

		case time == "tarde":
			if hora >= 13 && hora < 20 {
				totalPersonas++
			}

		case time == "noche":
			if hora >= 20 && hora <= 24 {
				totalPersonas++
			}

		default:
			totalPersonas = 0
		}
	}

	if totalPersonas > 0 {
		return totalPersonas, nil
	} else {
		return 0, ErrNotFoundTime
	}

}

// Calcula el porcentaje por destino
func (s *Storage) AverageDestination(destination string, Tickets *[]Ticket) (float64, error) {

	var totalTickets float64
	var targetCountryTickets float64

	for _, ticket := range *Tickets {
		totalTickets++
		if ticket.PaisDestino == destination {
			targetCountryTickets++
		}
	}

	if targetCountryTickets == 0 {
		return 0.0, ErrNotFoundTicket
	}

	avg := (targetCountryTickets / totalTickets) * 100
	return avg, nil
}

// ejemplo 1
//func GetTotalTickets(destination string) (int, error) {}

// ejemplo 2
//func GetMornings(time string) (int, error) {}

// ejemplo 3
//func AverageDestination(destination string, total int) (int, error) {}
