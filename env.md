env
===

Functional wrappers of core language features

```js
var env = require('env');
```

API
---

Types
-----

Duct tape checking

```js
env.types.isNumber({ valueOf: function () {} }); // true
env.types.isString({ toString: function () {} }); // true
env.types.isArray({ length: isNumber }); // true
```
