# Basic python bytecode micro interpreter


Enlace del video explicativo: https://www.youtube.com/watch?v=imxzkmd7JWA

En el siguiente proyecto se prueba la implementacion de un micro interprete de BYTECODE de .python en codigo .go

Dicho documento obtiene un archivo con esta estructura:
Indice (no tiene relevancia en el codigo) | Intruccion | Parametro (puede tener o no, depende de la instruccion)

![image](https://github.com/Fabricio06/Basic-python-bytecode-micro-interpreter/assets/82431338/5e99899f-514f-41be-befb-eb58fd6688dc)
### Ejemplo de las instrucciones enviadas por archivo .txt imagen 1

El codigo se encarga de procesar todo el archivo y separa instrucciones y parametros en una lista en comun, cada parametro se guarda del tipo que corresponga (string, int, float, bool)

Luego se procesan todas las instrucciones y llaman a su funcion correspondiente, por ejemplo si la instrucciones es "LOAD_CONST" llama a la funcion:

func LOAD_CONST(cons interface{}) {
	pila.push(cons)
}

Que se encarga de subir la constante de cualquier tipo a la pila


El repositorio cuenta con 3 archivos .txt para probar, un ejemplo de la salida de la consola con las llamadas bytecode

![image](https://github.com/Fabricio06/Basic-python-bytecode-micro-interpreter/assets/82431338/67d77091-ee52-4797-bb94-65e01863bfdc)
### Ejemplo de un archivo de texto factorial imagen 2
