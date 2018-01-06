# General rules

* All variables must be declared before use with var
* javascript is case-sensitive. lastname and lastName are 2 diff vars.

# Literals

## bool literal

* true
* false

## string literal

* Strings can be single quotes or double quotes
* \ escapes. \' and \" respectively escape single and double quote in a double and single quote string.
* double/single is automatically escaped in single/double quote string.

## object literal

* In object literal, the property name is always a string. It needs to be enclosed in quotes if it has spaces/hyphen. Otherwise string can be omitted.
* Empty string is a property too
* Properties are separated by spaces
* Objects can nest

```
var flight= {
  airline: "Oceanic",
  number: 815,
  departure: {
    IATA: "SYD",
    time: "2004-09-22 14:55",
    city: "Sydney"
  },
  arrival: {
    IATA: "LAX",
    time: "2004-09-23 10:42",
    city: "Los Angeles"
  }
};

```

## regex literal

Inside slashes


```
/ ... /
```

## comments

cpp like. Single line with // and multiple line with `/* .. */`

# Constructs

/* pretty much c-like looking */

[//]: # Search Terms: statements

```
if (expression) {

} else if (expression) {

} else {

}

for ( initial_expr ; condition ; incr_expr ) {

}

while (expression) {

}
```

# Functions

```
function FunctionName () {

    return variable;
} <- No semi-colon
```

* No return statement, then return undefined
* No overloading. Last definition overrides other definitions
* arguments is like a array that can access function args w/o naming them


# Basic Types

## Primitive Types

* Number
  * Always floating point.
* strings
* booleans
* null
* Undefined

# Native Objects

* provided by the EMEA implementation. Always available

## Arrays

* 0 indexed.
* array.length     // Its a property, so no paranthesis
* array.push(obj)  // Adds a element at end of array
* array.pop()      // Removes a element at beg of array
* array.join(" ")  // Uses arg to build a string
* array.split( )
* array.concat( )
* array.slice( )
* array.push()
* array.pop()

## Date


# Objects

* dictionary of name/value pairs called properties
* Properties
  * Attributes
    * writable
    * enumerable
    * configurable
  * prototype, class, extensible-flag
    * own(directly-defined)/inherited(defined in prototype)
* types
  * native objects
    * arrays, functions, dates, regex
  * host objects
    * HTMLElement objects
  * user-defined
    * Creation
      * object literal. type will be Object
      * new Constructor(). type will be Constructor.
        Inside you can set o.prototype = Something. If unset, it will be Constructor.prototype

Attempting to retrieve values from undefined will throw a TypeError exception. This can be guarded against with the && operator

```
    flight.equipment                               // undefined
    flight.equipment.model                         // throw "TypeError"
    flight.equipment && flight.equipment.model     // undefined
```


# Prototype/Contructor/__proto__ relations:

See picture @ http://stackoverflow.com/questions/9959727/proto-vs-prototype-in-javascript
Read at https://blog.pivotal.io/labs/labs/javascript-constructors-prototypes-and-the-new-keyword

## constructor

var a = new Car();


# Null/Undefined

* null
    * no value
    * value of a variable that has no initialization.
* undefined
    * deeper absense ..
    * (non existing variable, return value of fn that didn't return anything)

# Useful functions

* anyobject.hasOwnProperty("property")
* anyobject.isPrototypeOf(object)
* anyobject.properyIsEnumerable("property")
* anyobject.toString()
* anyobject.valueOf()
* parseInt(), parseFloat()
* Object.getOwnPropertyDescriptor(object, "property")
* Object.defineProperty(object, { value: ".." , writable: true/false, enumerable: true/false, configurable: true/false } )
* object.isFunction


# string functions

string manipulations

str.length
string[index]  .. gives single char at index
string.slice(beg_index[,end_index])  .. gives from beg_index, but exclusing end-index
str.substring(beg_index,end_index)

new_string = string.replace("old_literal_text","new_literal_text");

# Jquery

## Objects

```
$(document)      <-- Entire doc
$("button")      <-- All elemetns of that type
$(".class-name") <-- All elements of this class
$("#id")         <-- Exact Element with this id
```

## Functions

```
$("#elementSelector").addClas("animated-bounce");
$("#elementSelector").css("property","value");
$("#elementSelector").prop("property",value);

$("#elementSelector").html(htmlContent);   // Replaces html content
$("#elementSelector").text(textContent);   // Replaces html content w/o tags
$("#elementSelector").remove();            // Removes an entire element
$("#elementSelector").appendTo("#targetElement"); // Moves this to target
$("#elementSelector").clone();             // Clones and returns the new element, so that u can chain the netxt function
.parent()
.children()
$("...:nth-child(1-based-index)")
$("...:odd")

$("#elementSelector").append(htmlContent);
```

# Global object


* document object refers to the full html document we are rendering
* window is a much higher global object.


# Js functions in a browser

getElementsByTagName

# HTML Div events

```
<div onclick="message('hi')">
onmouseover
onresize
onload
```


# Kitchen Sink
