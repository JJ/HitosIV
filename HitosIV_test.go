package HitosIV

import (
	"testing"
	"reflect"
)

func TestHitos (t *testing.T){
	t.Log("Test Id");
	if CuantosHitos() <= 0 {
		t.Error("No milestones")
	}
}

func TestTodosHitos (t *testing.T){
	t.Log("Test Todos");
	these_milestones := Hitos()
	if reflect.TypeOf(these_milestones).String() == "Data" {
		t.Error("No milestones here")
	}
}

// Comprueba si el número de hitos es correcto
func TestNumHitos (t *testing.T){
	t.Log("Test Número Hitos");
	var x int = int(CuantosHitos());
	if x == 2{
		t.Log("El número de hitos es correcto")
	} else{
		t.Error("El número de hitos es incorrecto")
	}

}

// Comprueba que el tipo de los hitos sea correcto (lo que devuelve la función Uno(...))
func TestUno(t *testing.T){
	t.Log("Test Uno");
	var x int = int(CuantosHitos());
	for i:=0; i < x; i++{
		var aux uint = uint(i);
		if reflect.TypeOf(Uno(aux)).String() == "HitosIV.Hito"{
			t.Log("El hito", i, "ha sido comprobado.");
		}else{
			t.Error("El hito", i, " no es del tipo correcto"	);
		}
	}
}
