# csr
La primera capa debia ser un set de comandos escritos en bash que permitiesen
la creación y administración de diferentes hojas de cálculo, sin embargo, a la
hora de diseñar el render de celda se ha visto que bash se vuelve complejo con
el tratamiento de arrays de 2 dimensiones y lento ante la recursividad.

Es por ello que se considera oportuno la creación de una capa inferior que se
dedique exclusivamente a renderizar archivos de hoja de cálculo

Se elija c ya que su rendimiento es excelente y su código es simple y elegante
ante el diseño del algoritmo de renderizado

Esta capa sólo contempla renderizar ficheros únicos creados de forma externa

Estos ficheros no tienen identificador de fila como se pensó para las otras
capas, sólo un conjunto de filas y columnas formando celdas que pueden
referenciarse unas a otras dentro de un mismo archivo

Las siguientes capas se encargaran del resto, se podrán escribir en bash o rc
y usarán este render para la resolución de las hojas de cálculo


## formato de fichero
Siguiendo con el diseño inicial, el fichero se compone de filas separadas por
el caracter de salto de linea '\n' cada una de ellas compesta por columnas
separadas por pipes '|'

La primera y última fila, así como la primera y última columna no tienen
el caracter separador al principio o al final respectivamente

Siguiendo el modelo matemático, cada celda se identifica con la coordenada del
eje x (horizontal/columna) al principio seguida de la coordenada del eje y
(vertical/fila)

### Ejemplo
    celda 1, 1 | celda 2, 1 | celda 3, 1
    celda 1, 2 | celda 2, 2 | celda 3, 2
    celda 1, 3 | celda 2, 3 | celda 3, 3


## formato de celda
Cada celda puede contener únicamente uno de los dos tipos de datos disponibles
o una combinación de ellos:
- Literal
- Fórmula

Un literal es un valor final, mientras que una fórmula deberá ser renderizada
hasta obtener su valor final

Sólo cuando todas las celdas tengan su valor final asignado se entregará la
tabla a la salida estándar

Lo único que diferencia un literal de una fórmula es que la fórmula está
contenida entre llaves '{}'

