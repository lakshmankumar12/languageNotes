# Makefile

Link : https://www.gnu.org/software/make/manual/make.html
Currently reading: 3.5

## Rule

A rule looks like this:

```
target … : prerequisites …
    recipe
    …
    …
```

* The first non-dot target is called the default goal.
    * Goals are the targets that make strives ultimately to update

* Make has some implicit rules. You can omit spelling out rules, if the receipe's of these implicit rules are good enough.

* A rule can have multiple targets. The same file can be a/one-of-the target(s) in many rules.

## Targets

* Phony tells a target isnt an actual file
```
.PHONY : target_name
```

* dot-targets will not be used for default-goals


The % matches any target! Its an always on target. Typically '%.c' is used.
To prevent an infinite loop of pre_req also matching %, we have an explicit
rule for pre_req with a empty recipe.

```
%: pre_req
    recipie

pre_req: ;
```

### Grouped targets

* Useful when the receipe builds all targets in one-go.

```
foo bar biz &: baz boz
    echo $^ > foo
    echo $^ > bar
    echo $^ > biz

```


## Recipie

* Use `@` in the beginning dif you want suppress echo'ing of the commands


# Variables

* Simple declaration and usage
    ```
    objects = main.o kdb.o \
                another.o yetanother.o

    app: $(objects)
        receipe_to_link
    ```

* Space after var at declaration.
* Use with dollar-parentheses

## Assignment with equal operator

Search: = :=

* There has to be a space on either side of `=` or `:=` symbol
* `=` expands the variable everytime, while `:=` expands just once.
https://www.gnu.org/software/make/manual/html_node/Flavors.html#Flavors

```
variable_name = value

target: $(var_having_prereq)
  recipe_command $(some_var)

```


## Automatic variables

Full list: https://www.gnu.org/software/make/manual/make.html#Automatic-Variables

`$@` - filename of the target of the rule
`$<` - first prerequistie
`$^` - all prerequisties
`$*` - stem which a implicit rule matched. Basically the `%` part that matched in the rule

## Special Make variables

`.DEFAULT_GOAL`
`.RECIPEPREFIX`

# Makefile contents

* explicit rules
    * The std rule, has target(s), pre-requisite(s) and receipe(s)
* implicit rules
    * It describes how a target may depend on a file with a name similar
      to the target and gives a recipe to create or update such a target.
* variable definitions
* directives
    * an instruction for make to do something special while reading the makefile
        * include another makefile
        * conditional
        * multi-line variables
* comments
    * line starting with `#` is a comment.
    * a `\` at end continues a comment to nextline.
    * Escape `#` with a backslash.
    * commands can't be used within variable references or function calls
    * comments within recipes is passed as-is to shell.
    * define directive keeps the `#` intact (part of the var definition)

## splitting lines

* make uses linebase syntax. A line can be arbitrarliy long
* split long lines using `\` and immediate newline after it.
* make will condense white-space before this and the white-space in next-lnie all into one single white-space.
* Use the `$\ ` trick to avoid this single white space, as `$ ` evaulates to empty value.
* splittling lines on receipe to read later.

## default order of makefiles

* GNUmakefile
* makefile
* Makefile (Recommended)

## Include

* `include file` will include another file.
* there should be space between include, and each filename. Extra space at end if okay.
* file can be abs path. Otherwise in this order
    * current directory
    * -I/--include-dir
    * prefix/include i.e {/usr/local,/usr/gnu,/usr}/include
* If file isn't found, warning is printed, but not fatal.
    * Make will see if it can remake those makefiles. Only at end its prints error
* `-include` makes the inclusion not an error when not found
* MAKEFILES is an env variable, that contains list of makefiles to include before other makefiles
    * better to use explicit include directive


# CMake

https://cmake.org/cmake/help/latest/guide/tutorial/index.html#guide:CMake%20Tutorial

* We have a file called `CMakeLists.txt` along with src code. Here is a simple such file
    ```
    cmake_minimum_required(VERSION 3.10)

    # set the project name
    project(Tutorial)

    # add the executable
    add_executable(Tutorial tutorial.cxx)
    ```
* to run cmake do
    ```
    #pwd has source files and CMakeLists.txt
    mkdir builddir

    #this will generate the actual Makefiles
    cmake ..

    # fire the make
    cmake --build .

    # your executable should be available now
    ```

## Version numbering

```
# This will make @Tutorial_VERSION_MAJOR@ and @Tutorial_VERSION_MINOR@ available
# in the in-file for configure_file() command
project(Tutorial VERSION 1.0)

# The in-file has the @..@ vars which will be translated to the out-file
# you will most likely include the out-file in your source-code.
# you will src-control the .in file.
configure_file(TutorialConfig.h.in TutorialConfig.h)


# Will add a -I to all files in your project
#  ${PROJECT_BINARY_DIR} is where the generated files are present.
target_include_directories(Tutorial PUBLIC
                           "${PROJECT_BINARY_DIR}"
                           )
```

* Sample TutorialConfig.in file
```
#define Tutorial_VERSION_MAJOR @Tutorial_VERSION_MAJOR@
#define Tutorial_VERSION_MINOR @Tutorial_VERSION_MINOR@
```


# Sample makefile tricks

## Progress bar

* https://stackoverflow.com/a/455390/2587153 [Progress printing]

```
# PLACE AT THE TOP OF YOUR MAKEFILE
#---------------------------------
# Progress bar defs
#--------------------------------
#  words = count the number of words
ifneq ($(words $(MAKECMDGOALS)),1) # if no argument was given to make...
.DEFAULT_GOAL = all # set the default goal to all
#  http://www.gnu.org/software/make/manual/make.html
#  $@ = target name
#  %: = last resort recipe
#  --no-print-directory = don't print enter/leave messages for each output grouping
#  MAKEFILE_LIST = has a list of all the parsed Makefiles that can be found *.mk, Makefile, etc
#  -n = dry run, just print the recipes
#  -r = no builtin rules, disables implicit rules
#  -R = no builtin variables, disables implicit variables
#  -f = specify the name of the Makefile
%:                   # define a last resort default rule
      @$(MAKE) $@ --no-print-directory -rRf $(firstword $(MAKEFILE_LIST)) # recursive make call,
else
ifndef ECHO
#  execute a dry run of make, defining echo beforehand, and count all the instances of "COUNTTHIS"
T := $(shell $(MAKE) $(MAKECMDGOALS) --no-print-directory \
      -nrRf $(firstword $(MAKEFILE_LIST)) \
      ECHO="COUNTTHIS" | grep -c "COUNTTHIS")
#  eval = evaluate the text and read the results as makefile commands
N := x
#  Recursively expand C for each instance of ECHO to count more x's
C = $(words $N)$(eval N := x $N)
#  Multipy the count of x's by 100, and divide by the count of "COUNTTHIS"
#  Followed by a percent sign
#  And wrap it all in square brackets
ECHO = echo -ne "\r [`expr $C '*' 100 / $T`%]"
endif
#------------------
# end progress bar
#------------------

# REST OF YOUR MAKEFILE HERE
any_target: whatever_pre_req
    @$(ECHO) Doing $@

#----- Progressbar endif at end Makefile
endif
----
```

without comment

```
----
ifneq ($(words $(MAKECMDGOALS)),1) # if no argument was given to make...
.DEFAULT_GOAL = all # set the default goal to all
%:                   # define a last resort default rule
      @$(MAKE) $@ --no-print-directory -rRf $(firstword $(MAKEFILE_LIST)) # recursive make call, 
else
ifndef ECHO
T := $(shell $(MAKE) $(MAKECMDGOALS) --no-print-directory \
      -nrRf $(firstword $(MAKEFILE_LIST)) \
      ECHO="COUNTTHIS" | grep -c "COUNTTHIS")
N := x
C = $(words $N)$(eval N := x $N)
ECHO = echo -ne "\r [`expr $C '*' 100 / $T`%]"
endif

# ...

endif
```

