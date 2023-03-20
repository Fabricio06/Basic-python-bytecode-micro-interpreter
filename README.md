Enlace del video explicativo: https://www.youtube.com/watch?v=imxzkmd7JWA

En el siguiente proyecto se prueba la implementacion de un micro interprete de BYTECODE de .python en codigo .go

Dicho documento obtiene un archivo con esta estructura:
Indice (no tiene relevancia en el codigo) | Intruccion | Parametro (puede tener o no, depende de la instruccion)

0 LOAD_CONST 1
2 STORE_FAST c
4 LOAD_CONST 1
6 STORE_FAST x
8 LOAD_FAST x
10 LOAD_CONST 10
12 COMPARE_OP >=
14 JUMP_IF_TRUE 50
16 LOAD_FAST c
18 LOAD_FAST x
20 BINARY_MULTIPLY
21 STORE_FAST c
23 LOAD_GLOBAL print
25 LOAD_CONST el valor actual de c es igual a:
27 CALL_FUNCTION 1
29 LOAD_GLOBAL print
31 LOAD_FAST c
33 CALL_FUNCTION 1
35 LOAD_FAST x
37 LOAD_CONST 1
39 BINARY_ADD
40 STORE_FAST x
42 LOAD_FAST x
44 LOAD_CONST 10
46 COMPARE_OP >
48 JUMP_IF_FALSE 16
50 LOAD_GLOBAL print
52 LOAD_CONST El resultado final de c es:
54 CALL_FUNCTION 1
56 LOAD_GLOBAL print
58 LOAD_FAST c
60 CALL_FUNCTION 1
62 END

El codigo se encarga de procesar todo el archivo y separa instrucciones y parametros en una lista en comun, cada parametro se guarda del tipo que corresponga (string, int, float, bool)

Luego se procesan todas las instrucciones y llaman a su funcion correspondiente, por ejemplo si la instrucciones es "LOAD_CONST" llama a la funcion:

func LOAD_CONST(cons interface{}) {
	pila.push(cons)
}

Que se encarga de subir la constante de cualquier tipo a la pila
