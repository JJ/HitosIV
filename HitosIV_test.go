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

func TestUno(t *testing.T){
	t.Log("Test Uno");
	var x int = int(CuantosHitos());
	for i:=0; i < x; i++{
		var aux uint = uint(i);
		if reflect.TypeOf(Uno(aux)).String() == "HitosIV.Hito"{
			t.Log("El hito ", i, "ha sido comprobado.");
		}else{
			t.Error("El hito", i, " no es del tipo correcto");
		}
	}
}


/*
*	Comprueba el tipo de cada uno de los hitos.

func TestUno(t *testing.T){
	t.Log("Test Uno");

		var prueba1 = Uno(0);
		var prueba2 = Uno(1);
		if reflect.TypeOf(prueba1).String() != "HitosIV.Hito"{
			t.Error("El hito 1 no es del tipo correcto");
		}

		if reflect.TypeOf(prueba2).String() != "HitosIV.Hito"{
			t.Error("El hito 2 no es del tipo correcto");
		}

	t.Log(reflect.TypeOf(prueba1));

	if reflect.TypeOf(prueba1).String() == "Data"{
		t.Log("Ha salido por aquí");
	}else{
		t.Log("Ha salido por allá");
	}
}
*/
