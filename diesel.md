diesel.js
=========

HTML-based DSL compiler

Angular approach
----------------

Given a DOM tree, angular traverses the tree matching each node with defined directives

Even if the DOM node is a valid node, angular looks for a directive with the same name and if it is found,
the directive is called with the node tree and a model

Instead of looking each directive into the DOM through getElementByTagName, the DOM tree is traversed and
each node is matched in the directive list

This way, the directives can be resolved in the DOM hierarchy order (either upwards or downwards)

My first approach was using getElementByTagName, but then the nodes does not have a notion of parents or children
(they must be retrieved through the DOM API)

I think the angular's approach is cleaner (and maybe faster?) so I will use it

Resolution order
----------------

The next question is to resolve the directives upwards or downwards in the DOM tree

Let's assume that a directive can define inner directives:

```html
<datatable>
  <header>...</header>
  <body>...</body>
  <footer>...</footer>
  <pagination>...</pagination>
</datatable>
```

Here `header`, `body`, `footer` and `pagination` will be treated as potential sub directives from the
`datatable` directive, meaning that if they are defined as directive components in `datatable` they
will be parsed as `datatable` components instead of isolated directives

So, if a `pagination` directive is declared, it will be ignored in the context of `datatable`

To accomplish this feature, **the resolution order must be downwards** (from root to leaves)
