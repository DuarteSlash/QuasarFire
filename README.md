Título del Proyecto

Proyeecto realizado para presentación de codigo golang

Comenzando 🚀

Para instalar el software ingresar en la pagina de golang.es y 
descargar archivo.

Pre-requisitos
■ Desacargar este codigo fuente
■ Tener instalado el Golang en su máquina (solo para hacer build)

Posición de los satélites actualmente en servicio
● Kenobi: [-500, -200]
● Skywalker: [100, -100]
● Sato: [500, 100]

■ Crear un programa con las siguientes firmas:
// input: distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetLocation(distances ...float32) (x, y float32)
// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
func GetMessage(messages ...[]string) (msg string)
Consideraciones:
● La unidad de distancia en los parámetros de GetLocation es la misma que la que se
utiliza para indicar la posición de cada satélite.
● El mensaje recibido en cada satélite se recibe en forma de arreglo de strings.
● Cuando una palabra del mensaje no pueda ser determinada, se reemplaza por un string
en blanco en el array.
○ Ejemplo: [“este”, “es”, “”, “mensaje”]
● Considerar que existe un desfasaje (a determinar) en el mensaje que se recibe en cada
satélite.
○ Ejemplo:
■ Kenobi: [“”, “este”, “es”, “un”, “mensaje”]
■ Skywalker: [“este”, “”, “un”, “mensaje”]
■ Sato: [“”, ””, ”es”, ””, ”mensaje”]



Instalación
Ubicar el paquete app en el GOROOT
ejecutar desde la linea de comandos ejecutar la sentencia
go build
go install Nivel_1.go

Ejecutando las pruebas
go run Nivel_1.go o Nivel_1.exe


Respuesta del programa:

Hi, Location Found! Emisor Distance = [-500 -200]
Posición Kenobi: [-500.000000,-200.000000] 
mensaje: " ","este","es"," ","mensaje","" ""

Hi, Location Found! Emisor Distance = [100 -100]
Posición Skywalker: [100.000000,-100.000000]
mensaje: " "," ","este","es","un","mensaje"," ","" ""

Hi, Location Found! Emisor Distance = [500 100]
Posición Sato: [500.000000,100.000000]
mensaje: " "," ","este","es","un","mensaje"," ","" ""
