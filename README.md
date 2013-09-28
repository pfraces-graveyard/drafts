# drafts

To do projects

# Libs

## Clio

`prompt`, `optimist` and `rc` integration

## Typed

Structure definions

Checks based on *go interfaces*

## Argue

Define function interfaces based on `typed`

Create an API for function overloading based on one main function and multiple
argument preprocessor functions

## Pathy

Object path traversal

Catch undefined property exceptions by returning undefined

    var obj = pathy({}, 'foo.bar.qux', 123) // { foo: { bar: { qux: 123 } } }
    console.log(pathy(obj, 'foo.bar')) // { qux: 123 }

Create `pathy-lib` with a simple function definition, and `pathy` as different
package with improved JavaScript API and a `clio`-based CLI API

If this pattern succeeds document it at `module-design-patterns.md`

# Services

## knowledge dependency manager

Bits of human worth reading knowledge ordered by dependency

# Tools

## Owl compiler

Own web language compiler

# Specs

## Owl

Own web language

What if the coffescript headlines are really implemented in a language spec?

*   Line to line JavaScript
*   Hide JavaScript *bad* parts

Hints

*   `fun`
*   no automatic semicolon
*   no function named declaration

## Own programming languages

A collection of *phases* of precompilators

# Games

## dungeon master assistant

Web UI for `AD&D` masters

## Fairy Tales

Browser puzzle game based on sokoban

## Helmet

Browser platform game with Thor's son as main character and a magic helmet as
main (maybe unique) weapon
