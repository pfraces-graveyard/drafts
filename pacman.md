pacman
======

JavaScript package manager

Features
--------

  * Modules are annonymous
  * No registry: modules can be fetched from any url

Module API
----------

**log.js**

```js
module(function (require) {
  var js = require('js');
  
  var log = function (message) {
    js.console.log(message);
  };

  return log;
});
```

**defer.js**

```js
module(function (require) {
  var js = require('js');

  var defer = function (deferred) {
    js.setTimeout(deferred, 0);
  };

  return defer;
});
```

**sum.js**

```js
module(function () {
  var sum = function (x, y) {
    return x + y;
  };

  return sum;  
});
```

**main.js**

```js
module(function (require) {
  var defer = require('./defer');
  var log = require('./log');
  var sum = require('./sum');

  var main = function () {
    log(sum(2, 3));
  };
  
  defer(main);
});
```

Configuration API
-----------------

Configuration is defined in `pacman.json`

**pacman.json**

```json
{
}
```
