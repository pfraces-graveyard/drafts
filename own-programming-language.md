# own programming language

# targets

*   minimum language specification
*   simple parsing
*   avoid magic numbers and string literals
*   single coding convention
*   multiple decoupled preprocessing steps
*   use common sintax when possible

## simple parsing

*   polish notation
*   single argument to/from functions

# multiple decoupled phases

## 00. JavaScript generation 

*   Use techniques for fast JIT parsing
*   `"use strict";`
*   `asm.js`

## 01. fun

JavaScript good parts

*   what coffeescript must be but don't
*   really one-to-one js syntax
*   avoid js pitfalls
*   just a js subset

## 02. owl - own web language

*   let's start playing with semantics!
*   object and behavior with same syntax

# notes

*   from tcl: "everything is a string"
*   go: interfaces
*   json: entrada y retorno de variables
*   from python:
    *   newline as instruction separator
    *   indentation as block delimiters

*   reference passsing by default
*   nested comments

# brainstorm

*   se elimina if ya que se puede hacer lo mismo con switch
*   se ha de discernir entre switch con o sin break
*   se eliminan for y do ya que se puede hacer lo mismo con while
*   una variable puede contener un valor o una funcion
*   la asignacion se hace por referencia
*   si se asigna una funcion, se asigna su proceso, no su resultado
*   si se asigna una variable, se asigna su direccion de memoria
*   se debe elegir un operador para obtener valor en lugar de referencia
*   la indentacion indica que  se siguen pasando argumentos a la funcion
*   es voluntaria. el espacio separa los distintos argumentos
*   los literales van aparte
*   como se haran las llamadas al sistema?
*   las definiciones reciben y devuelven un unico argumento, que puede ser
    cualquier tipo de dato como una estructura de varios miembros
*   se puede suprimir los bucles en favor de la recursividad

# blocks

    ()
    []
    {}

*   can be nested

# member access

    .

# operators

*   grouping not required
*   fixed number of arguments

# unary operators

    !

# binary operators

## math

    + decimal augmentation
    - decimal substraction
    * decimal multiplication
    / decimal division

## boolean

    = strict equal
    > greater than
    < less than
    | or
    & and

## flow control

    ? switch
    @ loop

# snippets

    // c
    2 * 3

    // opl 1
    * 2 3

    // opl 2
    *
        2
        3

    // opl 3
    * {2 3)

    // opl 4
    * {2, 3)

    // opl 5
    * {
        2
        3
    )

    // c
    a==b

    // opl
    = a b

    // c
    a<b

    // opl
    < a b

    // c
    a>b

    // opl
    > a b

    // c
    !a

    // opl
    ! a

    // c
    a>=b

    // opl
    ! < a b

    // c
    a<=b

    // opl
    ! > a b

    // c
    if (a<=b)
        c=a;
    else
        c=b;

    // opl 1
    ? ! > a b
        true
            : c a
        false
            : c b

    // opl 1b
    ? true
        ! > a b
            : c a
        otherwise
            : c b

    // opl 2
    ?!> a b
        : c a
        : c b

    // c
    mult (int a, int b) {
        return a * b;
    }

    main() {
        printf ("%d\n", mult (2, 3));
    }

    // opl 1
    : * sys . expr . @ mult ; ;
    :
        sys . @ echo @ * 2 3

    // opl 2
    :* sys .expr .!mult ; ;
    :
        sys .!echo !* 2 3

    // opl 3
    :* sys.expr.mult ; ;
    :
        sys.echo * 2 3

    // opl 4
    :* sys .expr .mult ; ;
    :
        sys .echo * 2 3

    // opl 5
    :* sys .expr .!mult ; ;
    : sys .!echo !* {2 3}

    // opl 6
    :* sys .expr .!mult ; ;
    : sys .!echo !*
        2
        3

    // opl 7
    : mult * a b

    // opl 8
    : mult * . @ a
             . @ b

    // opl 9
    : mult * @.a @.b

    // c
    foo.bar.quux = 2

    // opl
    . foo
      bar
      quux
