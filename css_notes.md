# Study links

W2school: https://www.w3schools.com/css/default.asp


# selectors

* p                 -> w/o any prefix : actual html element
* .class_name       -> dot-prefix for classnames
* #id_name          -> hash-prefix for id
* p.class_name      -> p elements with that class name
* h1, h2, p {  }    -> you can group different elements to the same description
* .class_name img   -> img elements within the class_name
* a:state           -> element,colon,state. For eg, a-tag have 4 states: link, visited, hover, active

## Precedence

* class override in last-written gets weight.
* Precedence
    ```
     plain-outer <  plain-inner  < class-name-first-defined < class-name-last defined < id < inline
    ```
* `!important` in class names overrides everthing else.

# various css properties

* color                     -> plain text color

* background-color          -> background color
* background-image: url('../wherever/image.jpg')
* background-repeat: repeat-x/repeat-y/no-repeat
* background-position: right top
* background-attachment: fixed          (means the image wont scroll along with page)
* background: <color> <image> <repeat> <attachment> <position>         (all in one)

* border-style: dotted/dashed/solid/double/groove/ridge/inset/outset/none/hidden
* border-size: px/pt/cm/em/../  or thin/medium/thick
* border-color
* border:<size> <style> <color>       -> border specification, size, style and color
* border-{top,right,bottom,left}-style
* border-radius: px                   -> for rounded border, the radius
* box-sizing: content-box|border-box|initial|inherit;
* outline has similar properties

* margin                              -> auto, length, percent(of containing element), inherit
* margin-{top,right,bottom,left}

* width                               -> width of element alone. (total = left+padding + width + right-padding)
* max-width
* min-width
* padding                             -> length
* padding-{top,right,bottom,left}
* box-sizing: border-box              -> ensure the broder + padding + content-width == total width of element
              content-box             -> only content

* text
* color:
* text-align: center/left/right/justify
* text-decoration: none/overline/underline/line-through
* text-transform: uppercase/lowercase/capitalize
* text-indent: 50px                    -> for first line
* letter-spacing: 3px                  -> between chars in a text
* line-height: 0.8                     -> between lines (note, not in pixel but as mutiple of a line)
* direction: rtl
* word-spacing: 10px
* text-shadow: <horizontal> <vertical> <color>
* font-family: Helvetica Neue, Arial

* *display*  -  block/inline/hidden     -> block is a for starting on a new line, inline is to continue

* visibility - hidden (also hides, but the layout is affected)

## properties

clear: both
clear: left
clear: right

No floating elements are allowed on mentioned side

display: None

  Hides the element.


margin : Outline
padding : distance from margin



# Including css

## External file

```
<head>
<link rel="stylesheet" type="text/css" href="mystyle.css">
</head>
```

## Internal style sheet

```
<head>
<style>
body {
    background-color: linen;
}

h1 {
    color: maroon;
    margin-left: 40px;
}
</style>
</head>
```

## Plain inlining

```html
<h1 style="color:blue;margin-left:30px;">This is a heading</h1>
```
# CSS grid

Sample html

```html
<html>
    <head>
        <link rel="stylesheet" href="index.css">
        <link rel="stylesheet" href="basic.css">
    </head>
    <body>
        <div class="container">
            <div class="header">HEADER</div>
            <div class="menu">MENU</div>
            <div class="content">CONTENT</div>
            <div class="footer">FOOTER</div>
        </div>
    </body>
</html>
```

Css for that:

```css
.container > div {
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 2em;
    color: #ffeead;
}

html, body {
  background-color: #ffeead;
  margin: 10px;
}

.container > div:nth-child(1n) {
  background-color: #96ceb4;	
}

.container > div:nth-child(3n) {
  background-color: #88d8b0;
}

.container > div:nth-child(2n) {
      background-color: #ff6f69;
}

.container > div:nth-child(4n) {
      background-color: #ffcc5c;
}

.container {
    display: grid;
    grid-gap: 3px;
    grid-template-columns: repeat(2, 1fr);
    grid-template-rows: 40px 200px 40px;
}

.header {
    grid-column-start: 1;
    grid-column-end: 3;
}

.menu {}

.content {}

.footer {}
```



# Bootstrap:

## Including

add-this:
```
<head>
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
</head>
```

* To get started, we should nest all of our HTML in a div element with the class container-fluid.

## For images:

* Add class "img-responsive" to auto-adjust images! This goes into img tag.

## For text:

* class "text-center" to center-align. These directly go to the h*,p elements
* class "text-primary"
* class "text-danger"

## For buttons:

* Add class "btn" to button. Button is only as big as text. Always required for buttons
* Add class "btn-block" to make button strech to entire hori-width
* "btn-primary" is the main color for a btton. Use this for your main button
* "btn-info" for info buttons, "btn-danger" for delete'ish buttons
* "btn-default"
    **  btn-default
    **  btn-primary
    **  btn-success
    **  btn-warning
    **  btn-info
    **  btn-danger

    **  btn-block
    **  btn-large

## Grid Systems

Bootstrap has a responsive grid system (of size 12).

These classes should be in div elements of their own.

.row class specifies one-row

Medium
.col-md-1 (12-such-blocks)
.col-md-2 (6 -such-blocks)
.col-md-4 (3 -such-blocks)
.col-md-6 (2 -such-blocks)
.col-md-8 (1 -such-block, can be stacked with another say 4)

xtra-small
.col-xs-1

<div id="gridContainer">
    <div class="row grid-row">
        <div class="col-xs-1">1</div>
        <div class="col-xs-1">1</div>
        <div class="col-xs-2">2</div>
        <div class="col-xs-2">2</div>
        <div class="col-xs-6">6</div>
    </div>

## Grid in basic css

https://css-tricks.com/snippets/css/complete-guide-grid/

## Other notes

<span>...</span> is used to create a division in text.

jumbotron - grey box

panel  ->  A panel in bootstrap is a bordered box with some padding around its content
  .panel-default
  .panel-heading
  .panel-body
well

navbar
  navbar-inverse
  navbar-brand

page-header
input-group
form-control
input-group-btn
table
table-striped
label
label-default

# Fonts or Icons

## font-awesome

<link rel="stylesheet" href="//maxcdn.bootstrapcdn.com/font-awesome/4.5.0/css/font-awesome.min.css"/>
Icons: https://github.com/driftyco/ionicons
Google Material Icons

* <i class="fa"></i>
* fa-thumbs-up
* fa-info-circle
* fa-trash
* fa-paper-plane

# Animate.css

* class="animated" precedes
* "bounce"
* "shake"
* "fadeOut"
* "hinge"

# Colors

Color picker: https://www.w3schools.com/colors/colors_picker.asp
