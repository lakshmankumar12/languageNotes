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

* `=` expands the variable everytime, while `:=` expands just once.
https://www.gnu.org/software/make/manual/html_node/Flavors.html#Flavors

```
variable_name = value

target: $(var_having_prereq)
  recipe_command $(some_var)

```


## Automatic variables

`$@` - filename of the target of the rule
`$<` - first prerequistie
`$^` - all prerequisties

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

