# General Notes

## conditionals

```vim
:if 0
:    echom "if"
:elseif "nope!"
:    echom "elseif"
:else
:    echom "finally!"
:endif
```

### Operators

```
==  Ideally never used for strings (as it depneds on ignorecase for behavior)
==? case-insensitive
==# case-sensitive, both options work fine for integers. So just use this mostly!

:help expr4 for all operators
```

## normal

normal! will not turn normal keys into maps.


## Simple print

```vim
#echo "message and show them"
echom "message"
messages
messages clear
```

## Query a setting

```vim
set number
set number?
set nonumber
set number?
echo &number
```



## Get current line number

```vim
line(".") "Get the current line number
```

## External command output

```vim
#get output of a command into a varilable
let varName = system("command with args in a single string")
```


# notes on variables

```vim
:help internal-variables
```

* buffer-variable    b:     Local to the current buffer.
* window-variable    w:     Local to the current window.
* tabpage-variable   t:     Local to the current tab page.
* global-variable    g:     Global.
* local-variable     l:     Local to a function.
* script-variable    s:     Local to a :source'ed Vim script.
* function-argument  a:     Function argument (only inside a function).
* vim-variable       v:     Global, predefined by Vim.

