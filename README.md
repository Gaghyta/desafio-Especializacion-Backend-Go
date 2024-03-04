# Desafío

Realizar un programa que sirva como herramienta para
calcular diferentes datos estadísticos. Para lograrlo,
deberás descargar este repositorio que contiene un
archivo CSV con datos generados y un esqueleto del
proyecto.


¡Atención! Los ejemplos a continuación son solo a modo de guía. El
desafío se puede resolver de múltiples maneras.

# 1

## Requerimiento 1
Una función que calcule cuántas personas viajan a un país determinado.
func GetTotalTickets(destination string) (int, error) {}
(ejemplo 1)

Tip: VS Code nos permite buscar una palabra en un archivo con Ctrl + F o ⌘ + F.


## 2

Requerimiento 2:

Una o varias funciones que calculen cuántas personas viajan en madrugada (0 → 6),
mañana (7 → 12), tarde (13 → 19), y noche (20 → 23).
func GetCountByPeriod(time string) (int, error) {}
(ejemplo 2)

Tip: En Go, para manipular caracteres, tenemos el paquete strings.


##  3
Requerimiento 3:
Calcular el porcentaje de personas que viajan a un país determinado en un día.
func PercentageDestination(destination string, total int) (float64, error) {}
(ejemplo 3)

Tip: El porcentaje de x se calcula como: x̄ =
∑ x
n
. 100

2

##  4
Requerimiento 4:
Ejecutar al menos una vez cada requerimiento en la función main. Las ejecuciones deben
realizarse de manera concurrente (utilizando diferentes goroutines).
go func(){
totalTickets, err := GetTotalTickets(“Brazil”)
// process all the data
}()
(ejemplo 4)

Tip: Se pueden enviar parámetros a la función anónima de la siguiente manera: go func(param string)
{ // function detail } (“param string”)