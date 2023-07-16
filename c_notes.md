# C Notes

:toc:

# Declarations

* What is really guaranteed on sizes? http://c-faq.com/decl/inttypes.html[Faq 1.1]
    * type char can hold values up to 127(1 byte);
    * types short int and int can hold values up to 32,767(2 bytes); and
    * type long int can hold values up to 2,147,483,647(4 bytes).
    * something like the relation
    ```
    sizeof(char) <= sizeof(short) <= sizeof(int) <= sizeof(long) <= sizeof(long long)
    ```
* Use size_t when dealing with sizes.
* Use int, double when possible - fastest data-types.
    * Using things like char, short, float may result in more code.
* typedefs have scope rules. So they should be preferred over #defines.

## String Initialization
A string literal (the formal term for a double-quoted string in C source) can
be used in two slightly different ways:

```
char a[] = "string literal";
char *p  = "string literal";
```

* As the initializer for an array of char, as in the declaration of char a[] ,
  it specifies the initial values of the characters in that array (and, if
  necessary, its size).
* Anywhere else, it turns into an unnamed, static array of characters, and this
  unnamed array may be stored in read-only memory, and which therefore cannot
  necessarily be modified. In an expression context, the array is converted at
  once to a pointer, as usual (see section 6), so the second declaration
  initializes p to point to the unnamed array's first element.


## Some complex typedef styles

* Typedefing function pointers
    ```
    typedef int (*funcptr)();
    ```
    Now both the following are same.
    ```
    funcptr pf1, pf2;
    int (*pf1)(), (*pf2)();
    ```
* Typedefing pointers
    ```
    typedef char *charp;
    const charp p;
    ```
    Will make p a `char *const p` and not `const char *p+`. This is desirable.

## Array

* when you feel like indicating the index as part of initialization, you can use array-designators.
  Eg:
  ```c
    enum { RED, GREEN, BLUE };
    const char *nm[] = {
        [RED]   = "red",
        [GREEN] = "green",
        [BLUE]  = "blue",
                  NULL
    };
  ```

## Namespaces

* Detailed in this http://c-faq.com/decl/namespace.html[Faq 1.29]

There are 4 major namespaces

* labels (i.e. goto targets);
* tags (names of structures, unions, and enumerations; these three aren't separate even though they theoretically could be);
* structure/union members (one namespace per structure or union); and
* everything else (functions, variables, typedef names, enumeration constants), termed 'ordinary identifiers' by the Standard.
    * Note that the struct-names comes under 2nd category, while typedef-names sit along-side variable names

# Some facts to remember

* size_t is unsigned. There is a ssize_t which is signed (used as ret-val for send())
* struct compare isn't directly possible, as there are unused/padding holes in the struct.
* ANSI C has a offsetof(type,field) macro in <stddef.h>
* bit-fields are possible only within a struct/union definition
* `void*` is a generic type only for data-objects. There is no guarantee for
  function-pointers to be stored to `void*`. However all function pointers are
  inter-castable explicitly. If you need to store both types in one place, have
  a union of `void*` and some function-ptr type.
  http://c-faq.com/ptrs/int2ptr.html[Faq 4.14]

## volatile

* change done to a variable outside of compiler knowledge (by hardware, kernel, another thread).
  So reload this value from memory on every access.

# Sequence points

Best explained here in http://c-faq.com/expr/seqpoints.html[faq 3.8]

A sequence point is a point in time at which the dust has settled and all side
effects which have been seen so far are guaranteed to be complete. The sequence
points listed in the C standard are:

* at the end of the evaluation of a full expression (a full expression is an
  expression statement, or any other expression which is not a
  subexpression within any larger expression);
* at the ||, &&, ?:, and comma operators; and
* at a function call (after the evaluation of all the arguments, and just before the actual call).

The Standard states that

```
Between the previous and next sequence point an object shall have its stored
value modified at most once by the evaluation of an expression. Furthermore,
the prior value shall be accessed only to determine the value to be
stored.
```

* The following are all undefined (not even unspecified)
    ```
    a[i] = i++;
    i = i++;
    a ^= b ^= a ^= b
    ```
* Operator precedance doesn't gurantee function evaulation sequence
    ```
    f() + g() * h()  /* although * happens before +, f,g,h can be called in any order */
    ```
* Comma opeartor guarantees left to right evaulation (each comma is a sequence point). But function arguments
  are not comma operators and doesn't guranatee order of evaluation
    ```
    printf("%d %d", f1(), f2());   /* no guarantee on order of f1 or f2 */
    ```

## value preserving vs unsigned preserving

Talked about in http://c-faq.com/expr/unswarn.html[faq 3.19]

Also note that signed int overflow is UNDEFINED. Unsigned overflow follows modulo-arithmentic.

# Signatures of common routines

* main
    ```
    int main(int argc, char **argv)
    ```
* strcpy
    ```
    char *strcpy(char *dest, const char *src);
    char *strncpy(char *dest, const char *src, size_t n);
    void *memcpy(void *dest, const void *src, size_t n);
    void *memmove(void *dest, const void *src, size_t n);
    ```
    * strcpy/memcpy return the char*/void* of dst.
* strtol, atoi
    ```
    long int strtol(const char *nptr, char **end_ptr, int base);
    long long int strtoll(...);
    ```
    * long atol(const char*) (or atoi) is same as strtol(ptr, NULL, 10);
    * base is any of 0 to 36 (inclusive). 0 means 0x-begin, 0-begin or 10-base.
    * end_ptr if non-NULL is set to the location of ptr which has first non-0 bad value. if it points to '\0', the entire string was good.
    * errno is ERANGE if return is LONG_MAX/MIN and long doesn't fit.
* memset
    ```
    void *memset(void *s, int c, size_t n);
    ```
    * memset returns the same s that is passed.


# Compilation Tips

* Argument -E in gcc stops at preprocessing stage
* Argument -dM in processing stage dumps all macro definitions used.

## preprocessor output study

https://gcc.gnu.org/onlinedocs/cpp/Preprocessor-Output.html

* linemarkers
    ```
    # linenum filename flags
    ```
* flags:
    * 1: start of new file
    * 2: returning to file (after end of a included file)
    * 3: contents from sys header file
    * 4: implicit extern C block

# Some other library functions

* Double handling
    ```
    double modf(double input, double *integral_op);
    ```
    * The above function splits the double into integral part and fractional part.
    * The integral part is also returned as double, as a double can represent far more integer numbers than INT_MAX.

```
Last read FAQ: 3.7
```

# Malloc

Good links:

* http://g.oswego.edu/dl/html/malloc.html
* https://sourceware.org/glibc/wiki/MallocInternals
* https://danluu.com/malloc-tutorial/
* https://manybutfinite.com/post/anatomy-of-a-program-in-memory/

* http://www.cs.cmu.edu/afs/cs/academic/class/15213-s03/src/interposition/mymalloc.c

# bit operations


```c
// remote the Least sig 1-bit off
lastBit = x & -x

```

