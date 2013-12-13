# idea de juego

*   aprender mecanografia
*   multijugador, online
*   emulando magos lanzando hechizos
*   para lanzar un hechizo lo has de escribir
*   como guitar hero pero con teclado
*   un hechizo es un conjunto de keystrokes hasta que se pulsa enter
*   se da valor al hechizo según la dificultad que tiene escribirlo
    *   número de pulsaciones
    *   distancia de letras
    *   intercalación de manos

*   requiere conocer el sistema de teclado (default querty)
*   requiere que el usuario haya aprendido antes el hechizo
    *   sino uno puede empezar a pulsar teclas a lo loco
    *   no habría reto
    *   el reto es escribir algo correctamente, no sólo escribir algo
    *   por eso se necesita conocer la forma del hechizo
    *   aprender un hechizo puede ser escribirlo en un fichero
    *   el libro de hechizos: un fichero con hechizo por linea
    *   cuanto más larga es la linea mayor es la bonificación (exponencial)

*   los hechiceros tienen vida
*   los hechizos hacen daño

*   o pueden ser más versátiles y hacer una acción?
    *   vida: foo man choo
    *   daño: foz baz qux

*   curación propia o ajena?
    *   ajena en area circular? 
        *   cuantos más hay juntos curando al resto mas cantidad de hechizo
            de curación se aprovecha

En su forma más simple

*   Los hechizos hacen x daño donde x es el resultado de aplicar un
    algoritmo (siempre el mismo para todos los hechizos) al hechizo
    que no es más que una cadena de texto

*   La vida es la cantidad de caracteres en el libro de hechizos?
    *   Los hechizos menos usados se van eliminando
    *   Hay que penalizar para evitar copy paste o scripts
    *   No parece factible fácilmente
    *   para evitar scripts se pueden tener los hechizos en una base de
        datos

*   Los hechiceros tienen vida
*   La vida máxima puede ser la experiencia
*   La experiencia puede ser tiempo vivo
*   Al morir: nuevo personaje sin experiencia
*   Recuperación de daño: factor relativo a tiempo y nivel
*   Nivel: pulsaciones y complejidad medias
*   Nivel relativo: Si juegas mucho subes, sino bajas
*   Si juegas mucho recuperas muy rápido
*   otro dato a tener en cuenta (puede definir el nivel) es el número de
    palabras conocidas...

