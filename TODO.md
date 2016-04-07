TODO
====

Higlights
---------

  * ngl
      * publish minified sources
      * ngl.is + ngl.fp / ngl.lodash ?
  * umdify (wip: mu.is)
  * centralize starters in base.prj recipes
  * module2
      * load sources from URLs
      * submodules + nodejs support
  * module.lodash: wrap lodash in mujs modules
      * Lodash npm package builder
      * Lodash fp builder
  * freddie refactor
  * git-install refactor
  * Sanstorm + Gitlab + WeKan
  * Ember
  * TypeScript
  * Atom

Centralize tasks
----------------

### Repositories

  * [gists](https://gist.github.com/pfraces)
  * [plunks](https://gist.github.com/pfraces/5259db062652d00c9238)
  * [github](https://github.com/pfraces?tab=repositories)
  * [wip](https://github.com/pfraces-wip)
  * [drafts](https://github.com/pfraces-wip/drafts)
  * [labs](https://github.com/pfraces-labs)
  * [deprecated](https://github.com/pfraces-deprecated)

### Services

  * [SandStorm](https://github.com/sandstorm-io/sandstorm)
  * [Gitlab](https://gitlab.com/gitlab-org/gitlab-ce/)
  * [WeKan](https://github.com/wekan/wekan)

### Hosting

  * [Linode](https://www.linode.com/pricing)
  * [Heroku](https://www.heroku.com/pricing)

Projects
--------

  * git-install
  * scaffold
  * freddie
  * reset.css
  * browserify.prj: project starter

### mujs

  * project starter
  * core: `noop`, `identity`, `extend`
  * `thenable`
  * `iterator`
  * `dirty`
  * event delegation
  * types
  * signatures
  * function overloading
 
`module.umdify`

```js
if (typeof module === 'object') { /* node */ }
else if (typeof module === 'function') { /* browser mujs/module */ }
else if (typeof self === 'object') { /* browser worker global */ }
else if (typeof window === 'object') { /* browser global */ }
else { throw Error('unable to register module: unknown environment'); }
```

Finish project

  * lock API (zepto, lodash)
  * build system: eslint, mocha, umdify, uglify
  * unit tests
  * documentation
  * [Github Pages](https://pages.github.com/)
  * publish packages: npm, bower

### fp

  * iteratee first
  * inmutable data mutation
  
3-dimension helpers

  * Binary: `function (a, b) { ... }`
  * List: `function (list) { ... }` (reduce using binary methods)
  * Tree: `function (tree) { ... }` (recursive traversal using list methods)

### LightScript

  * [spec](https://github.com/pfraces-wip/LightScript/blob/master/README.md)
  * standard library

### ngl

  * style guide
  * project starter
  * request
  * Angular Mock

### Misc

  * dotfiles
  * Theater Mode
  * Gotick

Drafts
------

  * Listme (task manager)
  * Moghub (documentation server)
  * Inca (scriptable browser)
  * Wed (code editor)
  * Ytv (Youtube remote commander)
  * Mtgo (Magic Online)
  * Tufo
  * Grid (communication aid boards)
  * DOM-based Window Manager
  * Spaghetti: Workers + Virtual DOM
  * module-ng: submodules, umd
  * script to fix package.json dependencies versions

JavaScript
----------

### Workers

  * [Context detection](http://stackoverflow.com/questions/7507638)
  * [Cross platform parallel JavaScript](http://stackoverflow.com/questions/10773564/29088224#29088224)
  * Shared Workers
  * Load balancer
  * Promise based data flow

IPC

  * [Transferable objects](https://developers.google.com/web/updates/2011/12/Transferable-Objects-Lightning-Fast)
  * [Typed Arrays](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Typed_arrays)
  * [PSON](https://github.com/dcodeIO/PSON)
  * <https://developers.google.com/web/updates/2012/06/How-to-convert-ArrayBuffer-to-and-from-String>

### Virtual DOM

  * Nesting components
  * `shouldComponentUpdate`

HTML to Hyperscript

  * [dom2hscript](https://github.com/AkeemMcLennon/dom2hscript)
  * [html2hscript](https://github.com/twilson63/html2hscript)
  * [React Templates](http://wix.github.io/react-templates/)

### Frameworks

  * Backbone + Marionette
  * Angular
  * Angular2 / Typescript
  * Ember / ES6
  * React
  * Polymer

Ember

  * [Full Stack Fest 2015: Inside Glimmer](https://www.youtube.com/watch?v=VY-r7Ac06ho)

Example apps

  * todo app
  * dbmon

Markup
------

  * [SMACSS](https://smacss.com/)
  * [A Book Apart](https://abookapart.com/products)
  * Responsive Design
  * HTML/CSS Patterns

### CSS Modules

  * <https://medium.com/seek-ui-engineering/the-end-of-global-css-90d2a4a06284#.1ls34xsfu>
  * <https://github.com/css-modules/css-modules>

### CSS Frameworks

  * 960 grid system
  * unsemantic
  * Bootstrap
  * Foundation
  * Lost

XP
--

  * TDD
  * Agile Estimations
  * Extreme Programming Explained

Code Evenings
-------------

  * Functional JavaScript
  * Modular JavaScript
  * Angular 1.x introduction

Project Management
------------------

  * Kanban
  * GTD
  * Pomodoro
