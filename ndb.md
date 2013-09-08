# ndb

    getlist lista
    
    0 esto
    1 es
    2 una
    3 lista

Una tabla es una join de listas

La gestion del sistema de tablas ha de hacerse mediante el lenguaje interfaz de
ndb, el cual ha de hacer uso de herramientas gnu

Las tablas son listas

Cada item de una lista es una lista enlazada

Se ha de atravesar una lista enlazada para acceder a su contenido por lo que 

Una lista enlazada nos que nos da?

*   cada item es indivisible ( todo proceso de actuslizacion sera externo )
*   un item no tiene asociado un identificador unico
*   se el unico modo de pedir un item es por el lugar que ocupa en la lista en
    ese momento

*   se puede pedir el numero de items
*   se puede añadir un item en una posicion concreta al principio, mitad o fin
    de la lista

*   se puede eliminar un item

La base de datos ha de protegerse de que scripts externos puedan corromper los
datos

Eso impide que se pueda dar acceso al archivo de datos mediante llamadas
similares a las del sistema de archivos

Asi que ha de controlar nativamente cosas como

*   huecos de datos sin usar 

De 2 listas enlazadaas sincronizadas podemos obtener una lista con id de fila:

    fruta.data
    - banana
    - manzana
    - pera
    - melocoton

    fruta.id
    - 0
    - 1
    - 2
    - 3

Se elimina la pera:

    removeitem fruta.data 3
    removeitem fruta.id 3

    fruta.data
    - banana
    - manzana
    - melocoton

    fruta.id
    - 0
    - 1
    - 3

Las listas "columna" son necesarias para hacer tablas por ejemplo

Ya hablaremos de los indices los cuales necesitan las tablas...

Esta funcionalidad se puede hacer desde scripts externos mientras el motor
maneja el espacio que va quedando inutilizado para poder reaprovecharlo,
asi como la paginacion, para ello el motor ha de exigir y mantener el
etiquetado de las listas enlazadas, siendo la parte externa la que
elige la nomenclatura ( como aqui lista.data, lista.id u otra cualquiera )

Se permite asi hacer tambien scripts de pilas y demas usos que se le puedan
dar a las listas enlazadas

El motor entonces ha de manejar:

*   listas enlazadas
*   lista de las listas que hay y que nombre tienen
    aunque esto sera reprogramado en scripts, estas listas manejan direcciones
    de datos de nivel no accesible para los scripts asi que han de ser lo
    suficientemente sencillas como para que no valga la pena ofrecerlas como
    sistema de tablas

*   paginacion
*   reciclador

Pero los indices tambien requieren el uso de direcciones de memoria y de tablas
asi que hay que replantear lo arriba expuesto
