package main

import (
	"fmt"
	"time"
)

type MultiMapa[C comparable, V any] struct {
	mapa map[C][]V
}

// recordar: las restricciones solo se especifican en la definicion
func (m *MultiMapa[C, V]) Pon(clave C, valor V) {
	if m.mapa == nil {
		m.mapa = map[C][]V{}
	}
	m.mapa[clave] = append(m.mapa[clave], valor)
}

func (m *MultiMapa[C, V]) Lee(clave C) ([]V, bool) {
	if m.mapa == nil {
		return nil, false
	}
	valor, ok := m.mapa[clave]
	return valor, ok
}

func main() {
	pueblos := MultiMapa[string, string]{}
	pueblos.Pon("Colombia", "Cali")
	pueblos.Pon("España", "Cachorrilla")
	pueblos.Pon("Colombia", "Ibagué")
	pueblos.Pon("Argentina", "La Puerta")
	pueblos.Pon("España", "Vinuesa")

	for _, pais := range []string{
		"Colombia", "Perú", "Argentina", "España", "Malta",
	} {
		if pueblos, ok := pueblos.Lee(pais); ok {
			fmt.Printf("Pueblos de %s: %v\n", pais, pueblos)
		} else {
			fmt.Println("No se encontraron pueblos en", pais)
		}
	}

	/**
	Pueblos de Colombia: [Cali Ibagué]
	No se encontraron pueblos en Perú
	Pueblos de Argentina: [La Puerta]
	Pueblos de España: [Cachorrilla Vinuesa]
	No se encontraron pueblos en Malta
	*/

	numerosSorteo := MultiMapa[time.Time, int]{}
	ahora := time.Now()
	hoy := time.Date(ahora.Year(), ahora.Month(), ahora.Day(),
		0, 0, 0, 0, time.Local)
	numerosSorteo.Pon(hoy, 12)
	numerosSorteo.Pon(hoy, 13)
	numerosSorteo.Pon(hoy, 27)
	numerosSorteo.Pon(hoy, 32)
	numerosSorteo.Pon(hoy, 6)
	nums, _ := numerosSorteo.Lee(hoy)
	fmt.Println("La combinación premiada del sorteo de hoy es",
		nums)

	/*
		La combinación premiada del sorteo de hoy es [12 13 27 32 6]
	*/
}
