package HitosIV

import (
	"reflect"
	"testing"
	"os"
)

func TestMain(m *testing.M) {
	_ = ReadsFromFile("./hitos_test.json") // Alternative test file
	os.Exit(m.Run())
}

func TestBadFile(t *testing.T) {
	error := ReadsFromFile("/doesnotexist.json");
	if error == nil  {
		t.Error( "That file should not exist")
	}
}

func TestHitos(t *testing.T) {
	t.Log("Test Id")
	if len(Fechas) <= 0   {
		t.Error("No milestones")
	}
	
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
	var x uint = uint(CuantosHitos())
	if x == 3 {
		t.Log("El número de hitos es correcto")
	} else {
		t.Errorf("El número de hitos es incorrecto; esperábamos %d", 3)
	}
	var too_big uint = x + 3
	_, e := Uno( too_big )
	if e != nil {
		t.Log("Devuelve error si es demasiado grande")
	} else {
		t.Error("No devuelve error y debería")
	}

}

// Comprueba que el tipo de los hitos sea correcto (lo que devuelve la función Uno(...))
func TestUno(t *testing.T) {
	t.Log("Test Uno")
	var x uint = CuantosHitos()
	for i := uint(0); i < x; i++ {
		hito, err := Uno(i)
		if err != nil {
			t.Error( "El hito ", i , " es erróneo" )
		}
		if reflect.TypeOf(hito).String() == "HitosIV.Hito" {
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
