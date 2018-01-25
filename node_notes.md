# node interpreter stuff

* Like python use _ to capture last value.
* .help list of all commands.
* .break exit from multiline expression.
* .clear exit from multiline expression.
* .save filename save the current Node REPL session to a file.
* .load filename load file content in current Node REPL session.

# General npm understanding

* npm install sth will install in the node_modules in pwd.
* use -g flag to install globally!

## package.json

Attributes of Package.json

* name
    * name of the package
* version
    * version of the package
* description 
    * description of the package
* homepage 
    * homepage of the package
* author 
    * author of the package
* contributors 
    * name of the contributors to the package
* dependencies 
    * list of dependencies. NPM automatically installs all the dependencies mentioned here in the node_module folder of the package.
* repository 
    * repository type and URL of the package
* main 
    * entry point of the package
* keywords
    * keywords

# Including

```javascript
var a = require('express');
```

# common libraries

## fs

```javascript
var fs = require("fs");

var data = fs.readFileSync('input.txt');

fs.readFile('input.txt', function (err, data) {
    if (err) {
        console.log(err.stack);
        return console.error(err);
    }
    console.log(data.toString());
});
```

## events

```javascript
// Import events module
var events = require('events');

// Create an eventEmitter object
var eventEmitter = new events.EventEmitter();

// Bind event and event  handler as follows
eventEmitter.on('eventName', eventHandler);

// Fire an event 
eventEmitter.emit('eventName');
```

## serve-index

See [https://github.com/expressjs/serve-index]
