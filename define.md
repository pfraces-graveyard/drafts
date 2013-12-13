# define

```js
define('main', function (require) {
  var foo = require('foo');
  ...
});
```

# Motivation

Define modules for node and the browser without clobbering the global
object

This is not a script loader, is just the module pattern wrapped in a lib

This library exposes `define` in the global object and is expected that
using it you don't use any other globals but require what you need with
the `require` function offered as the unique argument in the `define`
callback
