# inca - Inca's Not Chromium Altered

Uzbl ideas applied on node-webkit

*   chromeless window
*   windowless daemon from which launch the windows
*   daemon with a tcp port listening
*   tcp client for the daemon

## Offering a CLI from which browse the web

### Create aliases

```sh
inca --alias google http://www.google.com
```

### Open a window

```sh
inca --open google
```

>   { "windowId": 1 }

### Open a new tab

```sh
inca --tab-open www.github.com
```

### Also from stdin

```sh
inca --open google
type weather
```

Previous example open a window with google loaded (CLI options are parsed first)
and then a google search about the *weather* is queried through the `type` command
which acts on the focused input if any.


### Input types

An `--input-type` option can be offered

```sh
inca --open google --input-type json
{
    "command": "type",
    "args": ["weather"]
}
```

Or more crazy things!

```sh
inca --alias $ http://code.jquery.com/jquery-latest.min.js
inca --alias jquery-render-plugin <a new cool project here?>
cat weather.js | inca --input-type js --require $ --require jquery-render-plugin
```

**weather.js**

```js
inca.load('google', function () {
    inca.type('wheater', function () {
        $('#wob_wc').render('png').pipe(inca.openStream);
    });
});
```
