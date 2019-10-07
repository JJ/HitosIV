package HitosIV

import (
	"reflect"
	"testing"
)

func TestHitos(t *testing.T) {
	t.Log("Test Id")
	if CuantosHitos() <= 0 {
		t.Error("No milestones")
	}
}

func TestTodosHitos(t *testing.T) {
	t.Log("Test Todos")
	these_milestones := Hitos()
	if reflect.TypeOf(these_milestones).String() == "Data" {
		t.Error("No milestones here")
	}
}

// Comprueba si el número de hitos es correcto
func TestNumHitos(t *testing.T) {
	t.Log("Test Número Hitos")
	var x int = int(CuantosHitos())
	if x == 2 {
		t.Log("El número de hitos es correcto")
	} else {
		t.Error("El número de hitos es incorrecto")
	}

}

// Comprueba que el tipo de los hitos sea correcto (lo que devuelve la función Uno(...))
func TestUno(t *testing.T) {
	t.Log("Test Uno")
	var x uint = CuantosHitos()
	for i := uint(0); i < x; i++ {
		if reflect.TypeOf(Uno(i)).String() == "HitosIV.Hito" {
			t.Log("El hito", i, "ha sido comprobado.")
		} else {
			t.Error("El hito", i, " no es del tipo correcto")
		}
	}
}

// Comprueba que el número de fechas sea igual al número de hitos
func TestHitosOrdenados(t *testing.T) {
	var hitos = HitosOrdenados()
	var x uint = CuantosHitos()
	for i := uint(0); i < x-1; i++ {
		fecha_i := dateFromStr(hitos[i].Date)
		fecha_j := dateFromStr(hitos[i+1].Date)
		if fecha_i.Before(fecha_j) {
			t.Log("Las fechas", fecha_i, "y", fecha_j, "han sido comprobadas")
		} else {
			t.Error("La ordenación de fechas no es correcta")
		}
	}
}
