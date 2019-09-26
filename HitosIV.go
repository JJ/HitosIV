package HitosIV

// taken from https://github.com/tucnak/telebot
import (

	"time"
	"strings"
	"strconv"
	"log"
	"encoding/json"
	"io/ioutil"

)

type Hito struct {
	File string `json:"file"`
	Title string `json:"title"`
	Date string `json:"fecha"`
}

type Data struct {
	Comment string `json:"comment"`
	Hitos []Hito `json:"hitos"`
}

var hitos []Hito
var ahora = time.Now()
var fechas []time.Time
var hitos_data Data

func init() {

	// Load milestones array. Must be right here, same directory this is run from. 
	file, e := ioutil.ReadFile("./hitos.json")
	if e != nil {
		log.Fatal("No se puede leer fichero de hitos")
	}
	if err := json.Unmarshal(file,&hitos_data); err != nil {
		log.Fatal("Error en el JSON de hitos →", err)
	}

	for _,hito := range hitos_data.Hitos {

		d := strings.Split(hito.Date,"/")
		this_day, _ := strconv.Atoi(d[0])
		this_month, _ := strconv.Atoi(d[1])
		this_year, _ := strconv.Atoi(d[2])
		fechas = append( fechas,
			time.Date(this_year, time.Month(this_month), this_day,
				12,30,0,0, time.Local))

	}

}

// Returns milestone data
func Hitos() Data {
	return hitos_data
}

// Returns how many milestones are there
func CuantosHitos() uint {
	return uint(len(hitos_data.Hitos))
}

// Returns a single milestone, indentified by number
func Uno( hito_id uint) Hito {
	if hito_id > uint(len(hitos_data.Hitos)) {
		log.Fatal("Index too high")
	}
	return hitos_data.Hitos[hito_id]
}
