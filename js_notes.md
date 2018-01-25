# General rules

* All variables must be declared before use with var
* javascript is case-sensitive. lastname and lastName are 2 diff vars.
* Does type coersion when say `+` is operated on different data-types.
* scope is either global or function - not per block like in c. All blocks withing function
  are in same scope. If you need scope private - use IIFE. Also read notes on
  execution blocks


# Literals

## bool literal

* true
* false

## string literal

* Strings can be single quotes or double quotes
* \ escapes. \' and \" respectively escape single and double quote in a double and single quote string.
* double/single is automatically escaped in single/double quote string.

### stringize

```
var num = 15;
var s = num.toString();

var s = '120';
var n = Number(s);
```


## object literal

* In object literal, the property name is always a string. It needs to be enclosed in quotes if it has spaces/hyphen. Otherwise string can be omitted.
* Empty string is a property too
* Properties are separated by comma. Colon and the value are separated by a space. Note the finishing }; like in c-struct definition.
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

cpp like. Single line with `//` and multiple line with `/* .. */`

# Operators

* `+` adds numbers or does string concat.. It coverts number to string

## logical operators

`===` does equality checking without type coercion
`==`  does equality checking after type coercion

# Constructs

/* pretty much c-like looking */

[//]: # Search Terms: statements

```
if (expression) {

} else if (expression) {

} else {

}

switch (variable) {
    case 'value1':
         break;
    ...
    default:

}

for ( initial_expr ; condition ; incr_expr ) {

}

for ( var i = 0 ; i < array_var.length ; ++i ) {

}

break;
continue; // as in c.

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

## closures

```javascript
function outerFunction(arg1, arg2) {
    var localvar = something;
    return function (closureArg1, closureArg2) {
        return arg1 + localvar + closureArg1;
    }
}

var func1 = outerFunction(1,2);
var finalResult = func1(3,4);
```

* Pretty much standard like in python
* closure captures the variables of its creator.

## IIFE

Immediately invoked function expressions.

* Javascript's ways of providing scope-private variables.

```
(function() {
   var priv_var1 = 5;
   ...
})()
```

* you can pass args and get things returned as well if you want.

# Basic Types

## Primitive Types

* Number
  * Always floating point.
* strings
* booleans
* null
* Undefined

Primities directly store their values. Objects store references.
Every var is either a primitive or an object

# Native Objects

* provided by the EMEA implementation. Always available

## Arrays

```
  var names = ['John', 'Jane', 'Mark'];
  var years = new Array(1990, 1968, 1948); /* same as above */
```


* 0 indexed.
* can be heteregenous
* array.length     // Its a property, so no paranthesis
* array.push(obj)  // Adds a element at end of array
* array.pop()      // Removes a element at beg of array
* array.join(" ")  // Uses arg to build a string
* array.split( )
* array.concat( )
* array.slice( )
* array.push()      // Adds to end of array
* array.pop()
* array.unshift()   // Adds to beginning of array
* array.shift()
* array.indexOf(element)   // returns index if present, else -1

* setting to an index that doesn't exist -> what happens?
    * If its just length+1, it gets set and array grows.
    * Not sure when its some arbitrary index.

## Date


# Objects

* See object literal above for defining new objects statically

```
#accessing can either be directly or using []
john.lastName === john['lastName']
# The latter form is needed when property name has spaces.
# Further latter form allows property name to be held in a variable
```

Creating using new:
```
var jane = new Object();
jane.name = 'Jane';  // Note the arbitrary addition of property name
jane['lastName'] = 'Smith';
```
As mentioned, objects can have properties arbitrarily added anywhere.

Functions inside object are referred as method
The keyword `this` refers to the object on which the method is called.

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

There are 3 properties to know

* .`__proto__`
    * This is how the property searching happens. This builds up the prototype chain
    * For a normal object, this points to its constructor's prototype
    * For a constructor(or any function object), this points to the
      Function.prototype, whose `__proto__` points to Object.prototype,
      whose `__proto__` points to null
* .prototype
    * This is not used in itself at deciding property. But this is
      present in a constructor object, pointing to the prototype object,
      ConstructorFunction.Prototype, so that new objects can be setup.
    * You can have a prototype propery for any object (not just for constructor
      functions). Whatever property and method are in here, is used
      later in inheritance chain
* .constructor
    * This is just a read-only property in a object to know its
      constructor



## constructor

* Imagine constructor as object-template(class definition)
* Dont return anything explicitly in constructor (if it does, this is not normal and confuses ppl)
* Constructor functions are noramlly started with a capital letter.
* `this` in constructor refers to newly created object.
* Note the parenthesis in the new call.
    ```
    var a = new Car();
    ```
What happens when u call new

* A new object is created
* constructor property of this new object is set to the constructor Function Object.
* the new object's .__proto__ is set to the constructor's.prototype.
* The constructor is called with this as the new object.

self observation: Like in python:
* The constructor.prototype has a property and a value, which is used by everybody
  when they themselves dont have that set.
* But moment you set in an individual object, it gets that for itself.

## Object.create

* The second way of creating objects
```
var personProto = {
  method1: function() { console.log(this.prop1); }
};

var aPerson = Object.create(personProto);
aPerson.prop1 = "value1";
```
* You have a prototype built, and have ur new object setup its `__proto__` directly
  as argument to constructor


# Null/Undefined

* null
    * no value
* undefined
    * deeper absense ..
    * value of a variable that has no initialization.
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

# execution context

* The global (outside of fucntion) is referred as global conext.
* Each function called has its down execution context.

A execution conext is created when a function is called.

Each execution context kind of conceptually has:
* Variable Object(VO)
    * argument object is created, which contains all arguments passed
    * code is scanned for function declarations. And each function is setup as a property poiting to the function
    * code is scanned for variable declarations. And each variable is setup as a property. The value is set to `undefined`
        * The above 2 actions are called hoisting. The functions and variables are availabe even before declaraion
* Scope chain
* `this` variable
    * points to the window object for a regular function call. (even for functions inside methods, but regular fns)
    * points to the object upon which a method is called
    * Note that this is assigned value only at time of execution just before the fn-call


* In a browser, that's the window object
* document object refers to the full html document we are rendering

## scope chain

* each function is a scope.
* each if{}, for{}, while{} is *NOT* a scope!
* lexical scoping: A function can be written inside another funciton.
                   The inner function gets accesss to the scope of the outer function.


# Js functions in a browser

* console.log(string) prints the string in browser console
* str_var = prompt(string) prints the string and gets a string and assigns it to the str_var
* alert(string) will just print the string in a pop-up-ish box
* console.info(x) to get information on an object
    * For array, etc.. plain x will display the arraized version.
      So, to inspect the other properties of x, u can use console.info(x)

document.querySelector('')  // either #id, .class - gets the first element with that class
document.getElementbyId('id_without_hash')   // faster than querySelector()

someDomElement.addEventListener('event-name', function)
someDomElement.classList.add('className')        // Dont call if you aren't sure. Remove and add
someDomElement.classList.remove('className')     // Okay to remove even if its not there
someDomElement.classList.toggle('className')

getElementsByTagName

## DOM element properties

someElement.textContent
someElement.innerHTML
someElement.style.display = "none"
someElement.

## add and remove elements

add

```
var newElem = document.createElement('DIV');
var existingElem = document.getElementbyId('id1');
existingElem.appendChild(newElem);
```


remove

```
var myNode = document.getElementById("foo");
myNode.innerHTML = '';

//or
var myNode = document.getElementById("foo");
while (myNode.firstChild) {
    myNode.removeChild(myNode.firstChild);
}
```


## things that can be done in console

Just type in a var, and the browser will let you know its value.

# HTML Div events

```
<div onclick="message('hi')">
onmouseover
onresize
onload
```

# JS libraries

Math.random()   /* gives a random number between 0 and 1 */
```javascript
//Inclusive of 1 and 6
Math.floor(Math.random() * 6) + 1;
```


# Jquery

## Including

```html
<head>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
</head>
```

## Basic syntax

```
$(selector).action()
```

eg:

```
$(this).hide()
$("p").hide()
$(".someclass").hide()
$("#specficElemByID").hide()
```

## Jquery functions

```
$(document).ready(function() {
    your code to run on document ready
});
```

```
$("#someButton").click(function() {
  // event handler for a button click!
});
```

Using the on-method to add event handler
```
$("#someButton").on('click', function() {

});

$("#someButton").on({
    mouseenter: function() {
        $(this).css("background-color","light-gray");
    },
    mouseleave: funciton() {
        $(this).css("background-color","light-blue");
    },
    'click': function() {
    });
```

* text, html and value

```
var textAt = $("#someElem").text();
var htmlAt = $("#someElem").html();
var valAt = $("#someElem").val();

$("#someElem").text("new repalce text")
```

# Google charts

[Link][https://developers.google.com/chart/interactive/docs/quick_start]


# Kitchen Sink
