# spider web framework

Conjunto de pequeñas librerias propias y de terceros para desarrollo de
aplicaciones web

# librerias

## ecstatic 

_(jesusabdullah/node-ecstatic)_

Servidor de ficheros estaticos

### Instalacion

    $ npm install ecstatic

## dnode

_(substack/dnode)_

Servidor rpc

### Alternativas

*   zeromq

### Instalacion

    # npm install dnode

### Instalacion con dependencias

    $ npm install dnode shoe domready ecstatic
    $ npm install -g browserify

### Compilar cliente con _browserify_

    $ browserify client.js -o static/bundle.js

### Incluir el bundle en fichero estatico html

    <script src="/bundle.js"></script>
    <div id="result"></div>

## shoe

_(substack/shoe)_

Sockjs straming

### Instalacion

    $ npm install shoe

## domready

_(ded/domready)_

Retiene la ejecucion de javascript hasta que el DOM esta completamente cargado

### Instalacion

    $ npm install domready

## browserify

_(substack/browserify)_

Share js between node and the browser

### Alternativas

*   ender

### Instalacion

    # npm install -g browserify

## uglify

_(mishoo/UglifyJS)_

Compress js for sending to the browser

### Instalacion

    # npm install -g uglify-js

## mustache

_(janl/mustache.js)_

Plantillas

### Instalacion

  $ npm install mustache

# workflow

    $ browserify client.js | uglifyjs > static/bundle.min.js

# todo

## Cómo enviar gzipped data?

Ejemplo que no funciona:

    $ browserify client.js | uglifyjs | gzip > static/bundle.min.js
    
Probar con el middleware para connect
