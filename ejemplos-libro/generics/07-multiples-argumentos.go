package main

type MapaDoble[C1, C2 comparable, V any] struct {
	elems map[C1]map[C2]V
}

// recordar: las restricciones solo se especifican en la definicion
func (m *MapaDoble[C1, C2, V]) Pon(clave1 C1, clave2 C2, valor V) {
	if m.elems == nil {
		m.elems = map[C1]map[C2]V{}
	}

	// seguir
}
