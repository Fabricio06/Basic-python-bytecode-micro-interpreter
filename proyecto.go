package main

// Fabricio Porras Morera 2021144223

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var storage = make(map[string]interface{}) //Almacén donde se guardan los datos
type PILA []interface{}                    //Crea el tipo del slice que va a almacenar cualquier valor
var pila PILA                              //La pila que vamos a utilizar para las operaciones

func main() {
	var instruccionesArchivo = leerTexto()  //Obtiene la lista con las instrucciones ordenadas del texto
	leerInstrucciones(instruccionesArchivo) //Lee y ejecuta las instrucciones en BYTECODE
}

//*********************************Funciones extras***********************************************

// Función que se encarga de ingresar un elemento en el tope de la pila
func (p *PILA) push(n interface{}) {
	*p = append(*p, n)
}

// Función que se encarga de sacar el elemento del tope de la pila
func (p *PILA) pop() interface{} {
	valor := (*p)[len(*p)-1]
	*p = append((*p)[:len(*p)-1])
	return valor
}

// Función que se encarga de imprimir la cantidad de argumentos enviados
func printCall(numArgs int) {
	for i := numArgs; i > 0; i-- {
		fmt.Println(pila[i])
	}
	pila = append(pila[0:0]) //La deja vacía
}

// Función que se encarga de revisar si dos variables son del mismo tipo
func mismoTipo(var1 interface{}, var2 interface{}) bool {
	if reflect.TypeOf(var1).Kind() == reflect.TypeOf(var2).Kind() {
		return true
	}
	return false
}

//********************La siguiente es la lista de instrucciones que se deberán programar:************************************

// LOAD_CONST Función que carga un valor de cualquier tipo a la pila
func LOAD_CONST(cons interface{}) {
	pila.push(cons)
}

// LOAD_FAST Función que coloca el valor de una variable en la pila
func LOAD_FAST(vr string) {
	pila.push(storage[vr])
}

// STORE_FAST Función que guarda el valor de una variable
func STORE_FAST(vr string) {
	storage[vr] = pila.pop()
}

// LOAD_GLOBAL Función con el mismo funcionamiento de LOAD_FAST con la diferencia de que también puede recibir ref a funciones
func LOAD_GLOBAL(vr string) {
	if vr == "print" {
		var miFunc func(int) = printCall
		pila.push(miFunc)
	} else {
		pila.push(storage[vr])
	}
}

// CALL_FUNCTION Función que solo va a llamar a la función printCall para efectos de este programa
func CALL_FUNCTION(numArgumentos int) {
	fu := pila[0]
	pr, err := fu.(func(int)) //Revisa si es una funcion
	if err != false {
		pr(numArgumentos)
	} else {
		fmt.Println("El valor no es una función")
	}
}

// COMPARE_OP Funcion que se encarga de comparar dos valores en pila mediante un operando recibido por el usuario
func COMPARE_OP(op string) {
	valor1 := pila.pop()
	valor2 := pila.pop()

	if mismoTipo(valor1, valor2) {

		switch reflect.ValueOf(valor1).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			switch op {
			case ">":
				pila.push(valor2.(int) > valor1.(int))
			case ">=":
				pila.push(valor2.(int) >= valor1.(int))
			case "<":
				pila.push(valor2.(int) < valor1.(int))
			case "<=":
				pila.push(valor2.(int) <= valor1.(int))
			case "==":
				pila.push(valor2.(int) == valor1.(int))
			case "!=":
				pila.push(valor2.(int) != valor1.(int))
			default:
				fmt.Println("Operador desconocido")
				return
			}
		case reflect.Float32, reflect.Float64:
			switch op {
			case ">":
				pila.push(valor2.(float64) > valor1.(float64))
			case ">=":
				pila.push(valor2.(float64) >= valor1.(float64))
			case "<":
				pila.push(valor2.(float64) < valor1.(float64))
			case "<=":
				pila.push(valor2.(float64) <= valor1.(float64))
			case "==":
				pila.push(valor2.(float64) == valor1.(float64))
			case "!=":
				pila.push(valor2.(float64) != valor1.(float64))
			default:
				fmt.Println("Operador desconocido")
				return
			}
		default:
			switch op {
			case ">":
				pila.push(valor2.(string) > valor1.(string))
			case ">=":
				pila.push(valor2.(string) >= valor1.(string))
			case "<":
				pila.push(valor2.(string) < valor1.(string))
			case "<=":
				pila.push(valor2.(string) <= valor1.(string))
			case "==":
				pila.push(valor2.(string) == valor1.(string))
			case "!=":
				pila.push(valor2.(string) != valor1.(string))
			default:
				fmt.Println("Operador desconocido")
				return
			}

		}

	} else {
		fmt.Println("Las variables tienen que ser del mismo tipos")
	}
}

// BINARY_SUBSTRACT Función que se encarga de realizar una resta de dos valores y subirlo a la pila
func BINARY_SUBSTRACT() {
	var result interface{}
	valor1 := pila.pop()
	valor2 := pila.pop()
	if mismoTipo(valor1, valor2) {
		switch reflect.ValueOf(valor1).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result = valor2.(int) - valor1.(int)
		case reflect.Float32, reflect.Float64:
			result = valor2.(float64) - valor1.(float64)
		default:
			fmt.Println("No se puede realizar una resta a un tipo string")
		}
		pila.push(result)
	} else {
		fmt.Println("Los valores tienen que ser del mismo tipo")
	}
}

// BINARY_ADD Función que se encarga de realizar una suma de dos valores y subirlo a la pila
func BINARY_ADD() {
	var result interface{}
	valor1 := pila.pop()
	valor2 := pila.pop()
	if mismoTipo(valor1, valor2) {
		switch reflect.ValueOf(valor1).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result = valor2.(int) + valor1.(int)
		case reflect.Float32, reflect.Float64:
			result = valor2.(float64) + valor1.(float64)
		default:
			result = valor2.(string) + valor1.(string)
		}
		pila.push(result)
	} else {
		fmt.Println("Los valores tienen que ser del mismo tipo")
	}
}

// BINARY_MULTIPLY Función que se encarga de realizar una multiplicación de dos valores y subirlo a la pila
func BINARY_MULTIPLY() {
	var result interface{}
	valor1 := pila.pop()
	valor2 := pila.pop()
	if mismoTipo(valor1, valor2) {
		switch reflect.ValueOf(valor1).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result = valor2.(int) * valor1.(int)
		case reflect.Float32, reflect.Float64:
			result = valor2.(float64) * valor1.(float64)
		default:
			fmt.Println("No se pueden multiplicar strings ")
		}
		pila.push(result)
	} else {
		fmt.Println("Los valores tienen que ser del mismo tipo")
	}
}

// BINARY_DIVIDE Función que se encarga de realizar una division de dos valores y subirlo a la pila
func BINARY_DIVIDE() {
	var result interface{}
	valor1 := pila.pop()
	valor2 := pila.pop()
	if mismoTipo(valor1, valor2) {
		switch reflect.ValueOf(valor1).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result = valor2.(int) / valor1.(int)
		case reflect.Float32, reflect.Float64:
			result = valor2.(float64) / valor1.(float64)
		default:
			fmt.Println("No se pueden dividir strings")
		}
		pila.push(result)
	} else {
		fmt.Println("Los valores tienen que ser del mismo tipo")
	}

}

// BINARY_AND Función que se encarga de realizar un AND lógico de dos valores y subirlo a la pila
func BINARY_AND() {
	var result interface{}
	valor1 := pila.pop()
	valor2 := pila.pop()
	if mismoTipo(valor1, valor2) {
		switch reflect.ValueOf(valor1).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result = valor2.(int) & valor1.(int)
		case reflect.Float32, reflect.Float64:
			fmt.Println("No se pueden operar con un AND a números decimales")
		default:
			fmt.Println("No se pueden operar con un AND a string")
		}
		pila.push(result)
	} else {
		fmt.Println("Los valores tienen que ser del mismo tipo")
	}
}

// BINARY_OR Función que se encarga de realizar un OR lógico de dos valores y subirlo a la pila
func BINARY_OR() {
	var result interface{}
	valor1 := pila.pop()
	valor2 := pila.pop()
	if mismoTipo(valor1, valor2) {
		switch reflect.ValueOf(valor1).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result = valor2.(int) | valor1.(int)
		case reflect.Float32, reflect.Float64:
			fmt.Println("No se pueden operar con un OR lógico a un tipo decimal")
		default:
			fmt.Println("No se pueden operar con un OR lógico a un tipo string")
		}
		pila.push(result)
	} else {
		fmt.Println("Los valores tienen que ser del mismo tipo")
	}
}

// BINARY_MODULO Función que se encarga de realizar un módulo de dos valores y subirlo a la pila
func BINARY_MODULO() {
	var result interface{}
	valor1 := pila.pop()
	valor2 := pila.pop()

	if mismoTipo(valor1, valor2) {
		switch reflect.ValueOf(valor1).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result = valor2.(int) % valor1.(int)
		case reflect.Float32, reflect.Float64:
			fmt.Println("No se pueden operar con modulo a un tipo decimal")
		default:
			fmt.Println("No se pueden operar con modulo a un tipo string")
		}
		pila.push(result)
	} else {
		fmt.Println("Los valores tienen que ser del mismo tipo")
	}
}

// STORE_SUBSCR Función que se encarga de cargar a la pila un valor en una posición en específico
func STORE_SUBSCR() {
	value := pila.pop()
	array := pila.pop().(PILA)
	index := pila.pop().(int)
	array[index] = value
}

// BINARY_SUBSCR Función que se encarga de subir un valor de un arreglo en el índice mandado en el tope de la pila
func BINARY_SUBSCR() {
	indice := pila.pop().(int)
	array := pila.pop().(PILA) // []interface{}
	pila.push(array[indice])
}

// JUMP_ABSOLUTE Funcion que se encarga de saltar sin importar que a un indice en especifico
func JUMP_ABSOLUTE(target int) int {
	return target
}

// JUMP_IF_TRUE Funcion que se encargar de saltar si en la pila se encuentra un "true" en un indice en especifico
func JUMP_IF_TRUE() bool {
	if pila.pop().(bool) {
		return true
	} else {
		return false
	}
}

// JUMP_IF_FALSE Funcion que se encargar de saltar si en la pila se encuentra un "false" en un indice en especifico
func JUMP_IF_FALSE() bool {
	if !pila.pop().(bool) {
		return true
	} else {
		return false
	}
}

// BUILD_LIST Función que se encarga de agregar la cantidad de elementos que hay en la pila a un slice
func BUILD_LIST(elements int) {
	var nuevaLista = make(PILA, elements)
	for x := 0; x < elements; x++ {
		nuevaLista[(elements-1)-x] = pila.pop()
	}
	pila.push(nuevaLista)
}

// END Función que se encarga de terminar el programa si ya no hay valores en la pila ni más instrucciones por seguir
func END() bool {
	switch len(pila) {
	case 0:
		return true
	default:
		fmt.Println("No se puede terminar el programa con valores aun en la pila")
		time.Sleep(3 * time.Second)
		return false
	}
}

//*********************************************Funciones en la conversion y lectura de textos y funciones*******************************

// convertirPara Funcion que se encargar de obtener todos los parametros tipo []interface y convertirlo a su respectivo tipo de variable
func convertirPara(listaAConvertir []string) []interface{} {
	var instruccionesPerfectas []interface{}
	for _, elemento := range listaAConvertir {
		valorNum, errN := strconv.Atoi(elemento)
		valorFlotante, errF := strconv.ParseFloat(elemento, 64)
		valorBooleano, errB := strconv.ParseBool(elemento)
		if errN == nil {
			instruccionesPerfectas = append(instruccionesPerfectas, valorNum)
		} else if errF == nil {
			instruccionesPerfectas = append(instruccionesPerfectas, valorFlotante)
		} else if errB == nil {
			instruccionesPerfectas = append(instruccionesPerfectas, valorBooleano)
		} else {
			instruccionesPerfectas = append(instruccionesPerfectas, elemento)
		}
	}
	return instruccionesPerfectas
}

// leerTexto Funcion que se encarga de procesar el archivo y separara las lineas/parametros que ocuparemos para el contexto del proyecto
func leerTexto() []interface{} {
	fileData, err := os.ReadFile("archivoInstrucciones3.txt") //Obtiene el archivo

	if err != nil { //Se confirma si se encuentra el archivo.txt
		fmt.Println("Hubo un error encontrando el archivo")
		time.Sleep(2 * time.Second) //Espera para que pueda leer el mensaje
		return nil                  //Se sale de la ejecución
	}
	texto := string(fileData)                    //Obtiene todo el texto del archivo como un string
	var lineas = strings.SplitAfter(texto, "\n") //Primero lo separa en líneas
	instruccionesArchivo := make([]string, 0)

	//Los corta por espacios entre instrucciones y parámetros, teniendo en cuenta los strings que pueden venir con espacios también "hello world"
	for x := 0; x < len(lineas); x++ {
		texto = lineas[x]
		chop := strings.SplitN(texto, " ", 3)

		recortarSaltos := []byte(chop[1])
		if recortarSaltos[len(recortarSaltos)-1] == 10 { //Corta bits residuales en el texto de las intrucciones
			recortarSaltos = append(recortarSaltos[:len(recortarSaltos)-2])
		}
		instruccionesArchivo = append(instruccionesArchivo, string(recortarSaltos))
		if len(chop) > 2 { //Significa que tiene parámetro
			recortarSaltos = []byte(chop[2]) //Corta bits residuales en los parametros de las instrucciones
			recortarSaltos = append(recortarSaltos[:len(recortarSaltos)-2])
			instruccionesArchivo = append(instruccionesArchivo, string(recortarSaltos))
		}
	}
	instruccionesPerfectas := convertirPara(instruccionesArchivo)
	return instruccionesPerfectas
}

// Funcion que se encarga de leer la pila de instrucciones y ejecutar cada una correspondientemente
func leerInstrucciones(instruccionesArchivo []interface{}) {
	for i := 0; i < len(instruccionesArchivo); i++ {
		switch instruccionesArchivo[i] {
		case "LOAD_CONST":
			LOAD_CONST(instruccionesArchivo[i+1])
		case "LOAD_FAST":
			LOAD_FAST(instruccionesArchivo[i+1].(string))
		case "STORE_FAST":
			STORE_FAST(instruccionesArchivo[i+1].(string))
		case "LOAD_GLOBAL":
			LOAD_GLOBAL(instruccionesArchivo[i+1].(string))
		case "CALL_FUNCTION":
			CALL_FUNCTION(instruccionesArchivo[i+1].(int))
		case "COMPARE_OP":
			COMPARE_OP(instruccionesArchivo[i+1].(string))
		case "BINARY_SUBSTRACT":
			BINARY_SUBSTRACT()
		case "BINARY_ADD":
			BINARY_ADD()
		case "BINARY_MULTIPLY":
			BINARY_MULTIPLY()
		case "BINARY_DIVIDE":
			BINARY_DIVIDE()
		case "BINARY_AND":
			BINARY_AND()
		case "BINARY_OR":
			BINARY_OR()
		case "BINARY_MODULO":
			BINARY_MODULO()
		case "STORE_SUBSCR":
			STORE_SUBSCR()
		case "BINARY_SUBSCR":
			BINARY_SUBSCR()
		case "JUMP_ABSOLUTE":
			i = JUMP_ABSOLUTE(instruccionesArchivo[i+1].(int)) - 1
		case "JUMP_IF_TRUE":
			if JUMP_IF_TRUE() {
				i = instruccionesArchivo[i+1].(int) - 1
			}
		case "JUMP_IF_FALSE":
			if JUMP_IF_FALSE() {
				i = instruccionesArchivo[i+1].(int) - 1
			}
		case "BUILD_LIST":
			BUILD_LIST(instruccionesArchivo[i+1].(int))
		case "END":
			if END() {
				return
			}
		}
	}
}
