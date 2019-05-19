# variables in make

* Space after var at declaration.
* Use with dollar-parentheses
* `=` expands the variable everytime, while `:=` expands just once.

```
variable_name = value

target: $(var_having_prereq)
  recipe_command $(some_var)

```

# Targets

----
.PHONY : target_name
----

* dot-targets will not be used for default-goals


The % matches any target! Its an always on target. Typically '%.c' is used.
To prevent an infinite loop of pre_req also matching %, we have an explicit
rule for pre_req with a empty recipe.

----
%: pre_req
    recipie

pre_req: ;
----


# CMake

* to run cmake

