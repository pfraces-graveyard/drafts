MuJS API
========

Modules
-------

  * module
  * udmify

Core
----

Low level utilities. Should provide enough to avoid using native methods of
`Array`, `Object`, `Function`, `Number`, `String` and `Boolean`

Other host objects will be covered as standalone modules, one per host object

`core` will be the foundation for the `fp` higher abstraction

Use `fp` utilities for as long as you can and go back to `core` utilities for low level
routines hopefully wrapped with a nice api in general purpose module

### types

  * isBoolean
  * isNumber
  * isString
  * isObject
  * isArray
  * isFunction

Extensions

  * isNumeric
  * isPlainObject
  * isArrayLike
  * isScalar

### object

  * extend
  * forEachAttr

### array

  * append
  * forEach

### tree

  * forEachNode

### function

  * apply
  * defer
  * debounce
  * throttle

FP
--

### types

  * isList
  * isIterator

### object

  * merge

### array

  * concat

### list

  * indexOf
  * keys
  * values

### iterators

  * map
  * filter
  * reduce
  * every
  * some

Generators

  * iterateArray
  * iterateObject
  * iterateTree
  * iteratorList
  
### function

  * partial
  * noop
  * identity
