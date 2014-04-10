# drafts

To do projects

# Libs

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

## json-reduce

JSON parser. Reduce JSON object into simpler one based on custom filters

1.  json-reduce sync
2.  shell tool sync
3.  json-stream-reduce async
4.  shell tool sync

### Acknowledgements

*   [dominictarr's JSONStream][1]

## need

system-agnostic dependency manager

# Services

## knowledge dependency manager

Bits of human worth reading knowledge ordered by dependency

# Tools

## Owl compiler

Own web language compiler

# Specs

## Mog

Markdown over git

A simple spec for wiki systems

## Owl

Own web language

What if the coffescript headlines are really implemented in a language spec?

*   Line to line JavaScript
*   Hide JavaScript *bad* parts

Hints

*   `fn` is the new `function`
*   no automatic semicolon
*   no function named declaration
*   `self` inside functions refer to the function itself (recursion)
*   no `this`, no `new`, no classes, no inheritance, no prototype
*   no globals: no Object, no Date, no Array
*   no prototyped methods
*   block scope, no hoisting

Use `define` as the unique global:

Revealed module pattern with caching, uses the same amount of memory
as prototyped methods and is more versatile and less verbose thanks to
the lexical scope

## Own programming languages

A collection of *phases* of precompilators

# Games

## magic online

Web platform to build magic-like online games

## dungeon master assistant

Web UI for `AD&D` masters

## Fairy Tales

Browser puzzle game based on sokoban

## Helmet

Browser game with Thor's son as main character and a magic helmet as main
(maybe unique) weapon

[1]: https://github.com/dominictarr/JSONStream
