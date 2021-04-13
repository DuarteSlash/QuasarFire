package app

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type app interface {
	DescubrirMensaje()
}

//Se crea estructura para ubicar las coordenadas de los satellites
type satellite struct {
	name        string
	coordinateX float32
	coordinateY float32
	message     []string
	desfasaje   int
}

//Crea Satellite con distancia al emisor
func (s *satellite) Init(coordx float32, coordy float32) {
	s.coordinateX = coordx
	s.coordinateY = coordy
}

//Constructor Satellite con distancia al emisor
func (s *satellite) newSatelliteMessage(name string, desf int, mensajeSat ...string) {
	s.name = name
	s.desfasaje = desf
	s.message = mensajeSat
}

// input: distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetLocation(distances ...float32) (x, y float32) {
	var sat = new(satellite)
	sat.Init(distances[0], distances[1])
	//var sat = satellite(distances[0], distances[1])
	fmt.Printf("Hi, Location Found! ")
	fmt.Print("Emisor Distance = ", distances)
	//fmt.Printf("\nEmisor Distance = %f %f\n", sat.coordinateX, sat.coordinateY)
	return sat.coordinateX, sat.coordinateY
}

// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
//No podemos pasar un segmento de cadenas como argumento. ...[]string
//Esto se debe a que el compilador espera argumentos discretos de cadenas
func GetMessage(desfasaje int, messages ...string) (msg string) {
	var mensaje = ""
	var textomensaje = []string{}
	var sat = new(satellite)
	sat.newSatelliteMessage(sat.name, desfasaje, messages...)

	textomensaje = mensajeEmisorSat(desfasaje, len(messages))

	rand.Seed(time.Now().UnixNano())
	var desfasajeInt = rand.Intn(desfasaje)

	messages = append(messages, textomensaje...)
	for i := 0; i < desfasaje; i++ {
		if i == desfasajeInt-1 {
			textomensaje = append(textomensaje, "\" \"")
		} else {
			textomensaje = append(textomensaje, "\""+messages[i]+"\"")
		}

	}
	mensaje = strings.Join(textomensaje, ",")
	//var fields = strings.Fields(mensaje)
	//fmt.Printf("%q", fields)

	return mensaje
}

//Como se tiene un desfase en los mensajes se debe agregar espacios en blanco
func mensajeEmisorSat(desfasaje int, tamanioMessage int) []string {
	var textomensajeSat = []string{}
	for i := 0; i < desfasaje-tamanioMessage; i++ {
		textomensajeSat = append(textomensajeSat, "\" \"")
	}
	return textomensajeSat
}

func DescubrirMensaje() {
	var mensajeEmisor = []string{"este", "es", "un", "mensaje"}
	var Kenobi = new(satellite)
	Kenobi.Init(-500, -200)
	Kenobi.newSatelliteMessage("Kenobi", 5, mensajeEmisor...)
	coordxKe, coordyKe := GetLocation(Kenobi.coordinateX, Kenobi.coordinateY)
	fmt.Printf("\nPosición Kenobi: [%f,%f] \n", coordxKe, coordyKe)
	fmt.Printf("mensaje: %s\n\n", GetMessage(Kenobi.desfasaje, Kenobi.message...))

	var Skywalker = new(satellite)
	Skywalker.Init(100, -100)
	Skywalker.newSatelliteMessage("Skywalker", 6, mensajeEmisor...)
	coordxsk, coordysk := GetLocation(Skywalker.coordinateX, Skywalker.coordinateY)
	fmt.Printf("\nPosición Skywalker: [%f,%f] \n", coordxsk, coordysk)
	fmt.Printf("mensaje: %s\n\n", GetMessage(Skywalker.desfasaje, Skywalker.message...))

	var Sato = new(satellite)
	Sato.Init(500, 100)
	Sato.newSatelliteMessage("Sato", 6, mensajeEmisor...)
	coordxSa, coordySa := GetLocation(Sato.coordinateX, Sato.coordinateY)
	fmt.Printf("\nPosición Sato: [%f,%f] \n", coordxSa, coordySa)
	fmt.Printf("mensaje: %s\n\n", GetMessage(Sato.desfasaje, Sato.message...))
}
