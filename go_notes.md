:toc:

# literals

* 4 types of literals
* Integer literals
    * normal numbers are base-10, `0b...` is binary,
      `0o12321` is octal, `0x...` is hex.
    * Starting with plain `0` is octal. Avoid.
    * You can add underscores to numbers.. They have no effect `132_456`
        * not allowed at start, end and cant be contiguous.
        * makes sense to add thousand-boundaries or word/byte boundaries
* Float literals
    * with dot and/or exponent - `1.234e6`
    * can be in hex as well. here `p` indicates exponent prefix
    * underscores allowed here as well
* rune literals
    * single-quote strings represent one utf-8 character
        ```go
        'u'      // simple ascii
        '\141'   // 8-bit octal
        '\xab'   // 8-bit hex
        '\u0012' // 16-bit hex.
        '\U00001234' // 32-bit unicode (notice capital U)
        ```
    * other escape chars:
        ```
        \\     | Backslash(\)
        \000   | Unicode character with the given 3-digit 8-bit octal code point
        \'     | Single quote ('). It is only allowed inside character literals
        \"     | Double quote ("). It is only allowed inside interpreted string literals
        \a     | ASCII bell (BEL)
        \b     | ASCII backspace (BS)
        \f     | ASCII formfeed (FF)
        \n     | ASCII linefeed (LF
        \r     | ASCII carriage return (CR)
        \t     | ASCII tab (TAB)
        \uhhhh | Unicode character with the given 4-digit 16-bit hex code point. Unicode character with the given 8-digit 32-bit hex code point.
        \v     | ASCII vertical tab (VT)
        \xhh   | Unicode character with the given 2-digit 8-bit hex code point.
        ```
* String Literals
    * double-quote strings are regular strings
        * supports escape chars, including `\"`
        * not-allowed: unescaped back-slash, newline and double-quotes
    * back-tick quoted
        * called raw-string literals
        * can span multiple lines.
        * No escape chars
            * Doesn't support another backtick inside
        ```go
        a := "simple usual string"
        b := "string with an escape char\" and another\n"
        c := `raw-string literal
           that can span multiple lines`

        // get formatted strings
        string_var := fmt.Sprintf("item1: %s item2: %d", message, year)
        ```
* The last rare one is complex literal with a `i`

* Go literals are untyped.
* they can interact with any variable that's compatible with the literal
    * ints to float is okay
    * 1000 to a byte is an error
    * int/string to string/int variable is an error
* Each literal has a default type which is used when there is not explicit variable.

## more notes on string literals

* always utf-8
* strings are immutable
* to edit strings, convert to slice of runes
* string(byteslice) coverts a byte slice to a string

```go
s := "hello"               /* string */
c := []rune(s)             /* covert to slice of runes */
c[0] = 'c'                 /* modify */
s2 := string(c)            /* covert slice of runes to string */
fmt.Printf("%s\n", s2)

// checking empty string -- both are okay
if s != "" { ... }
if len(s) > 0 { ... }
```

* Watch out, string indexing doesn't give the rune, but the byte! In fact len(str) is
  also total bytes, not runes. However range on a string works over runes
* The above isn't usually a problem to iterate over a string comparing it to individual
  runes(Atleast the ascii ones), as the non-ascii ones anyway dont compare equal.
* Inside a program, you can convert a string(utf-8 encoded) to slice of runes

## composite literal

* Initializes a slice or a struct
    ```go
    var palette = []color.Color{color.White, color.Black}
    anim := gif.GIF{LoopCount: nframes}
    ```

# consts

* Constants are declared like variables, but with the const keyword. Constants
  can be character, string, boolean, or numeric values.
    ```go
    const (
        Summer int = 0
        Winter int = 1
    )
    ```
* The const value has to be compile time derivable.
    * We can't have consts whose value are found at runtime.
    * Constants cannot be declared using the := syntax.
* Type info can be absent in the const declaration. In this case, its derived
  from the literal. Or if we have a typename (typically from type declaration,
  it can be used too). When type is absent, its untyped. So its of any one
  flavor - boolean, int, rune, floating-point, complex, string.
* can be package level or function level or at any block
* iota is used for enumeration
    * but the iota is very limited.
    * Use it only if you dont care for the values
    * You can start from some number or skip a number
    * If you need explicit values dont use ioto and just assign.
    ```go
    type Season int64
    const (                   /* enum is a type, followed by const assignments to that type */
        Summer Season = iota
        Winter                /* implicit iota of the same type */
    )
    // to stringize the enum
    func (s Season) String() string {
        switch s {
        case Summer:
            return "summer"
        case Autumn:
            return "autumn"
        }
        return "Unk"
    }
    ```
* `nil` represents non-existing pointer or reference-type (for slices, interface).


# Types

## Basic Types

```go
bool  // either true or false
      // zero val: false

// integer types
int8  int16  int32  int64
uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8, byte is more common in go
int uint    // depends on cpu. may be 32 or 64
            // default type for integer literals is int
rune // alias for int32
uintptr

// represents a Unicode code point
float32
float64   // default type
complex64 complex128 // rare.

string   // zero val: empty string
```

* Variables w/o Initialization is set to 0/false/empty-string
* No automatic type conversion between types.
    ```go
    var x int = 10
    var y float64 = 30.2
    var z float64 = float64(x) + y
    var d int = x + int(y)
    ```
* No implicitly truthy to any variables. i.e no automation boolean conversion
    * Use comparison operator
* Note, there is only one basic type which is a pointer.
  This is big enuf to hold any poitner.

## operations on basic types

### integers

* Standard arithmetic
    ```
    + - * / % .. and all support +=, -=, .. %= versions
    ```
* integers can be compared - `== != < > <= >=`
* Integers have bit operators - `<< >> & | ^ &^`.
    * these support assignment variants `&= |= <<= >>= ^= &^=`
    * `&^`  -- (golang special) `x &^ y`  is same as C's   `x & (~y)`

### floats

* no modulus ofcourse
* Dont compare floats with `== !=`.. although language allows it.

## strings

* Comparisions are okay -  `== != < > <= >=`
* Can be concatenated by `+`
* Are immutable
* Note that a string is an array of bytes (not runes). So, if you slice it you get bytes.
    * unless its ascii, the results will be unpredictable.
* bytes and runes can be type-casted to string.
* int gets type-casted .. but is best avoided.


## arrays

* [n]T is an array of n values of type T
    ```go
    func main() {
        var a [2]string
        a[0] = "Hello"
        a[1] = "World"
        fmt.Println(a[0], a[1])
        fmt.Println(a)

        primes := [6]int{2, 3, 5, 7, 11, 13}
        auto_size_detected_array := [...]int{2, 3, 5, 7, 11, 13}
        fmt.Println(primes)
    }
    ```
* Go's array are value-types. Think of it as struct with indexed members. Passing
  arrays to function will pass entire copies. (No decaying of name to pointer)
* array literals are like [n]type{val1,val2,..}. The n can be (...) in which
  case its auto derived.
  ```go
     a := [...]int{28,13,14} // compiler determines a is of type [3]int
     // initialize specific indices with values
     a := [5]int{1:10, 3:30}
     b := [4]*int{ 2:new(int) }
  ```
* every size is different type. `[5]int` is different from `[3]int`
    * This makes it very rigid and usable directly in most places
        * we cant type-cast from one size to another
        * we cant assign 2 diff size arrays to same variable.
* len(array) gives its length
* Array copy works as long as both are of same types
  ```go
    // Declare a string array of five elements.
    var array1 [5]string
    // Declare a second string array of five elements.
    // Initialize the array with colors.
    array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
    // Copy the values from array2 into array1.
    array1 = array2
  ```
* (Yet to grasp this fully: Be wary of saying/mentioning arrays in go. May be
  the slice is more appropriate). Note that []T is a slice of T, not array of T,
  but [n]T is an array.
* multi-dimensional arrays are supported
  ```go
     a := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, //this comma is necessary.
                               //The compiler will complain if you omit this comma
        }
  ```
* Array supports comparison operator `==` and `!=`
* out of bound access with compile-time decided index value will give compilation error
* out of bound access with run time index value will cause panic


## slices

* An array has a fixed size. A slice, on the other hand, is a
  dynamically-sized, flexible view into the elements of an array. In practice,
  slices are much more common than arrays.
* slices are reference-types
* The type `[]T` is a slice with elements of type T.
* Slice - ptr, len and cap and has the underlying array. len is the number of
  slice elements, cap is the number of elems in underlying array from the loc
  where ptr is pointing. Always len `<=` cap
  * there are function `len()`, `cap()` that give the len and cap of slices
    * Also for arrays and channels
* sequence is a term that can indexible. (its either a slice, array or ptr to
  array). slice-operator on a sequence produces a slice. This expression creates
  a slice of the first five elements of the sequence a.
    ```go
    a[i:j]  // 0 <= i <= j <= cap(a). resulting slice has j-1 elements
    a[0:5]
    c := []int{5,6,7}  // creates an array, and returns a slice to that array
    ```
* A slice does not store any data, it just describes a section of an underlying
  array.  Changing the elements of a slice modifies the corresponding elements
  of its underlying array.  Other slices that share the same underlying array
  will see those changes.
* slice with no underlying array is nil. This is the zero value for a slice
    * there is also a slice with empty capacity
    ```go
    // nil slice
    var slice []int

    // Use make to create an empty slice of integers.
    slice := make([]int, 0)
    // Use a slice literal to create an empty slice of integers.
    slice := []int{}
    ```
* slices cannot be compared with one another, but can be compared against `nil`
* Can be created with a built-in function - make. Note the odditity. The first
  args is a type-name (and not a var-name). This creates a unnamed array and
  then returns a slice to that array. The returned slice is the only way
  to access that array
    ```go
    func make([]T, len, cap) []T
    //Eg:
    i := make([]int, 5, 5)
    ```
* new slice from existing slice
    ```go
    // Create a slice of integers.
    // Contains a length and capacity of 5 elements.
    slice := []int{10, 20, 30, 40, 50}
    // Create a new slice.
    // Contains a length of 2 and capacity of 4 elements.
    newSlice := slice[1:3]
    ```
* append adds to a slice and returns a new slice. Note: It may or may not modify the
  slice (depends on if there is a capacity). So copy it back as return value to
  update the variable. Its a compile error if append's return value isn't captured
  ```go
  slice = append(slice, 10)
  // add more than one value
  slice = append(slice, 20, 30)
  // if s1 and s2 are themselves slices
  slice = append(slice, s1, s2...)
  ```
  * append() may add to the underlying array if there is capacity or it might
    create a new array and copy all elements
  * at any rate, the capacity of the slice is increased by 1.
* Use the capability when creating a new slice from existing slice/array.
  Making this same as length, forces the next append to allocate a new array
* `range` on a slice gives a copy of the element. So dont take pointers while
  iterating over slices

## map

* Basically an unordered key-value hash-map
* Keys are any type on which == works. Value can be anything. == is good to
  for for integer, boolean, string, rune. Not == is bad for float(Nan). (Complex?)
  If a struct is absolutely just made of the above (to any depth) that is good
  for equivality too.
* Map created with make(map[K] V) or using map literal.
* We can't get address to a map. However the map is itself a reference type.

```go
var a map[string]int   // creates a nil map
var a map[string]int{} // creates a empty map. This is a map literal.
teams := map[string][]string {        // directly creates a map.
        "Orcas": []string{"Fred", "Ralph", "Bijou"},
        "Lions": []string{"Sarah", "Peter", "Billie"},
        "Kittens": []string{"Waldo", "Raul", "Ze"},
    }
ages := make(map[int][]string, 10)  // create a map with size 10


// check if map is empty
if len(m) == 0 {

}

// lookup:
// value will be the 0-type if the keyval does't exist
// ok(exists) is a boolean.
if val, ok := map_var[key] ; ok {

}

// non existing keys
totalWins := map[string]int{}
totalWins["Orcas"] = 1
totalWins["Lions"] = 2
fmt.Println(totalWins["Orcas"])
fmt.Println(totalWins["Kittens"])    // non-existing key.. prints 0-value for the int-type
totalWins["Kittens"]++               // non-existing key.. automatically assign 0s and ++. So finally 1
fmt.Println(totalWins["Kittens"])
totalWins["Lions"] = 3
fmt.Println(totalWins["Lions"])

//iterate.. see for section for more
for key, value := range m {

}

// deleting..
// no problem if key doesn't exist or even if m is nil
// no return value
delete(m, key)

```

## structs

* A struct is a collection of fields.
    ```go
    type Vertex struct {
        X int
        Y int
    }
    ```
* struct literals:
    ```go
    type Point struct { X, Y  int }
    var p Point{}  // members get their nil values
    var p Point    // same as above. (unlike maps .. there is no nil struct value)

    // type-1 .. comma separated list. Depends on order of defn.
    Point{1,2}
    // type-2 .. map-style. explicitly called out name of fields to value
    Point{X:1, Y:2}
    ```
* Fields are accessed using dot. (dot is called selector in go. It selects
  which field or method to use)
* To access the field X of a struct when we have the struct pointer p we could
  write `(*p).X`. However, that notation is cumbersome, so the language permits
  us instead to write just p.X, without the explicit dereference.
  * whether u have direct struct var or pointer to struct, u can still use dot
* Capital letter rules follow for struct too. If the struct type name is
  capitalized, the type is exported. If the individual members are caps, they
  are exported. A struct can have mix of exported and non-exported members.
* Functions that return struct, can better return struct-pointer. This will
  make function call be a L-value
* Struct can't have the same struct inside, but have a pointer of itself.
  (like c)
* Structs can embed other structures inside (anonyomous members). This has
  benefit of accessing members directly + invoking methods of the embedded
  type directly. However, this implicit behavior as receiver is only
  limited to receiver. We cant pass as args the outer Type, where embedded
  type is expected.
* Structs can have field-tags. These are any literal string. One eg is the
    ```go
        type Identity struct {
            Value isIdentity_Value `protobuf_oneof:"Value"`
        }
        type Feed struct {
            Name string `json:"site"`
            URI  string `json:"link"`
            Type string `json:"type"`
        }
    ```
* empty structs are allowed. They take 0 bytes. Typically used to define types
  with methods on them (to fit into a interface) and offer default functionality
* structs can be defined globally(pkg scope) or defined within function.
    * the in-function defined structs are only visible inside the function
    * well.. structs can be scoped to any block level.

### anonymous structs

* used is marshalling/unmarshalling external data like json/protobuf
* writing tests is also when this use pops up

```go
// using var varname
var person struct {
    name string
    age int
    pet string
}
person.name = "bob"
person.age = 50
person.pet = "dog"

// using :=
pet := struct {
    name string
    kind string
}{
    name: "Fido",
    kind: "dog",
}

```

### methods

* The object on which method is called is referred as receiver
* We can define methods for any type (basic-types, named-types, slices, maps)
  We can't however define methods for pointers and interfaces. (Pointer methods
  are treated as methods of the pointed-type itself)
  * if u have a pointer/actual variable, the compiler will automatically choose
    the right way to pass to receiver
  * Methods declared with pointer receivers can only be called by
    interface type values that contain pointers.
  * Methods declared with value receivers can be called by interface type values
    that contain both values and pointers.
* methods can be defined only in the same package where the type is defined.
* by convention either declare all methods on type or on pointer. Only pointer
  methods can change the receiver though.
    * Compiler will implicitly make a T as *T if only *T is a receiver. However,
       when T is passed as-is as interface{} (Like to fmt.Printf), then there is
       no implicit conversion! So `(*T).String()` may not work!
* its okay for a receiver to be nil as long as its pointer type reciever, and
  the method guards itself against that.


## channel

* communication mechanism
* Is always of a given type

```go
//create a unbuffered channel
ch := make(chan int)

//create a buffered channel
ch := make(chan int, 10)

// read from a channel
// ok if false if the channel is closed, or true
v, ok := <-ch

// write to a channel
ch<- v

// for-range loop
// runs till channel is closed
func runThingConcurrently(in <-chan int, out chan<- int) {
    go func() {
        for val := range in {
            result := process(val)
            out <- result
        }
    }()
}

// close a channel
close(ch)
```
* select .. do any one of the cases, that is doalbe. Usually wrapped in a `for()`
```go
select {
case v := <-ch:
    fmt.Println(v)
case v := <-ch2:
    fmt.Println(v)
case ch3 <- x:
    fmt.Println("wrote", x)
case <-ch4:
    fmt.Println("got value on ch4, but ignored it")
}
```
* If all go-coroutines are stopped , then go-runtime panics


## goroutine

* concurrent function execution
* go statement creates it

## interface

* Collection of methods
* composed of a type/value. These are dynamic and at runtime point to the concrete type and the value
* nil interface has both the type & value to nil. But beware, there can be cases where just the value
  is nil. Such interfaces dont compare to nil.
* interfaces are comparable (==) if the underlying type is comparable or if both are nil.
  Otherwise, comparing uncomparable types causes runtime panic. (So this is not caught at compile time)
* convention - if there is only method, the interface name typicall ends in er. Eg: Matcher
* idomatic go is to name the interface with a `-er` suffix.

c-equivalent of `void*` is `interface{}`
There is now a
```go
type any = interface{}
```

### wrapping interfaces - embedded type pattern

```go
type nopCloser struct {
    io.Reader
}
func (nopCloser) Close() error { return nil }
func NopCloser(r io.Reader) io.ReadCloser {
    return nopCloser{r}
}

```


## generics

```go
type Processor[T any] interface {
    ProcessGrant(frequency int64, bandwidth int64) T
}
```

## Pointers

* Like c, * is used for type. `*T` is a pointer of type T. & is for getting a
  variable's pointer, and `*var` is for deferencing or indirecting. However,
  there is no pointer arithmetic in go.
* Its okay to take pointers to struct members
* Pointers are useful to pass by reference (like a slice that might be
  modified in a function)
* Note that idiomatic go is
    * pass by pointer if you need to mute the value
    * pass by value if you dont need to mute
    * return by value

# More on types

## Comparability/compare/equality/==

* basic types are comparable.
* struct made of comparable basic types or structs are comparable.
* interfaces are comparable if its current-type is comparable.
* slices/maps/functions are NOT comparable.
* array compares true if size is same and each of underlying value is
  comparable and equal.


## reference types

* maps, channels, slices, pointers, functions are reference-types. When you pass
  these in functions, you pass a reference to them. So, there are multiple references
  of them pointing to the same underlying type.
* structures, arrays, interfaces that contain reference-type also kind of become
  referenced.

## type declaration

Used to create new types from existing types - although they share same representation
they are different types

```go
type newTypeName underlyingType
type Celcius float64
// this creates a function-type
type HandlerFunc func(http.ResponseWriter, *http.Request)
```

* Explicity type conversion is then used to covert one to another.
    * one wont implicitly cast to the other
* Initialization however can be direct from a literal

```go
const boilingPoint Celcius = 100.0
var freezingPoint Celcius
freezingPoint = Celcius(someFloatVar)
```

Type names from basic-types are referred as named basic types. Eg. time.Duration

### aliases

```go
// For all intents and purposes, Bar is same as Foo
// compiler allows to use them interchangeably
type Bar = Foo

```


## type-convertion

* T(value) converts value to the type T
    ```go
    []rune("Hello World")

    //native type are cast like in c
    aUint64Int = uint64(aInt64Var)
    ```
* Usually a new type is same as the other type, but defines extra methods so
  that it can be passed as interfaces. In such cases, you will see these
  type conversions done from type X to Y (although both types are internally
  same)

## type-assertion

```go
// converts interface-i to a concrete type-T.
// will panic is i is not of type T
v = i.(T)

// for safe cast
// ok is true or false.
v, ok = i.(T)

//switching on type
func doThings(i interface{}) {
    switch j := i.(type) {
    case nil:
        // i is nil, type of j is interface{}
    case int:
        // j is of type int
    case MyInt:
        // j is of type MyInt
    case io.Reader:
        // j is of type io.Reader
    case string:
        // j is a string
    case bool, rune:
        // i is either a bool or rune, so j is of type interface{}
    default:
        // no idea what i is, so j is of type interface{}
    }
}
```

# Variables

* Initialization using the var statement. This is possible both inside
  functions and in global scope
    ```go
    var name type = expression        // Everything present
    var name = expression             // type is inferred from expression
    var name type                     // zero-initialized name for that type
    name := expression                // var keyword ommited because of := short-hand.
                                      // type also ommited
                                      // cannot be used at global(pkg) scope

    var i, j int = 1, 2

    var (                             // called a declaration list
        x int
        y = 20
        z int = 30
        d, e = 40, "hello"
        f, g string
    )

    ```
* Using the `:=` construct, var is skipped and type is assumed. This also
  help in initializing variables of different types in same statement. So
  `:=` is for declaration and `=` is assignment.
    ```go
    k := 3
    c, python, java := true, false, "no!"
    ```
* var statements can also be factored like import statements
* `_` can be used in place where a variable name isnt required.
* for all names, case matters. HeapSort and heapSort are different.
* Go typically uses camel case. Abbreviations may be all-caps.
* Multiple assignmenents are done in one go.
    ```go
    i , j = j , i  // swap i and j
    ```
* Always have default values
    * 0 for numeric types, False for bool, nil for pointers.
    * Composite type anyway build on the simple types above.

## scope

* very different from that of c.
* global scope is package scope in go.

* pointers to local variables can be passed back. (very different from
  c/cpp)
* each function invocation will result in a different local-variable
  pointer.

# Flow control statements

## semi-colons

* dont bother putting them.
* The compiler will auto-fill them for compilation intelligently.

## comments

* comments are like cpp.
* `//` for one line and `/* .. */` for multi line

## godoc

* `//` comments just above package or a function or a struct

## for

* complete-for statement.
    * has init/ condition/ post parts
    ```go
    {
        sum := 0
        for i := 0; i < 10; i++ {
            sum += i
        }
        fmt.Println(sum)
    }
    ```
* init and post are optional. At that point you can drop the semicolons: C's
  while is spelled for in Go. Omitting condition makes it a infinite loop
    ```go
    {
        sum := 1
        for sum < 1000 {
            sum += sum
        }
        fmt.Println(sum)
        for {
            fmt.Println("This will be infinitely printed")
        }
    }
    ```
* Variables declared in for's initialization part have loop's scope
* there is `break` and `continue`
* iterating over a slice/array (note string iteration will give runes)
  ```go
    var a := { 1,2,3}
    for i,v := range a {  // use _, v if u need only value.
      fmt.Println("%d %d",i,v)
    }
  ```
* iterating a string gives the runes.
    * Note that slice a string gives bytes.
    ```go
    sample := "apple_π!"
    for i, r := range sample {
        fmt.Println(i, r, string(r)) // you get to see the unicode points
    }
    ```
    * if a string contains junk utf-8 char, you get 0xfffd as the rune
* iterating over a map
  ```go
  for k, v := range m {
    fmt.Printf("key[%s] value[%s]\n", k, v)
  }
  for k := range m {
    fmt.Printf("only key[%s]\n", k)
  }
  for _,v := range m {
    fmt.Printf("only value[%s]\n", v)
  }
  ```
* iterate over a channel
  ```go
  for val_copy := range channel_var {

  }
  ```

* Note that for-range, gets a copy of the value. You dont modify the underlying type

* Loop labelling
    ```go
    func main() {
        samples := []string{"hello", "apple_π!"}
        outer:
        for _, sample := range samples {
            for i, r := range sample {
                fmt.Println(i, r, string(r))
                if r == 'l' {
                    continue outer // note the label
                }
            }
            fmt.Println()
        }
    }
    ```


## if

* if statements are like its for loops;
    * the expression is NOT surrounded by parentheses ( )
    * the braces { } are required
* the if statement can start with a short statement to execute before the
  condition. A var initailized here is availabe in if, else if and else.
    ```go
    func pow(x, n, lim float64) float64 {
        if v := math.Pow(x, n); v < lim {
            return v
        }
        return lim
    }
    ```
* combine a stmt and err check like this, limiting the err's scope
    ```go
    if err := r.ParseForm(); err != nil {
       log.Print(err)
    }
    ```
* if .. else if
    ```go
    if n := rand.Intn(10) ; n == 0 {   // n is available for the while of if/else/if
        fmt.Println("That's too low")
    } else if n > 5 {
        fmt.Println("That's too big:", n)
    } else {
        fmt.Println("That's a good number:", n)
    }
    ```

## switch

* Switch cases evaluate cases from top to bottom, stopping when a case succeeds
* A case body breaks automatically, unless it ends with a fallthrough statement
    ```go
    func main() {
        fmt.Print("Go runs on ")
        switch os := runtime.GOOS; os {
        case "darwin":
            fmt.Println("OS X.")
        case "linux":
            fmt.Println("Linux.")
        default:
            // freebsd, openbsd,
            // plan9, windows...
            fmt.Printf("%s.", os)
        }
    }
    ```
* break is implied -- need not be coded. f isn't called if i == 0 or 1
    ```go
    switch i {
      case 0, 1:
      case f():
    }
    ```
    * use `fallthrough` if you realy want it. but should be rare.
    * `break` can break before the end of case.
* we can have multiple values in a single case.
* Switch without a condition is the same as switch true. This construct can be
  a clean way to write long if-then-else chains.
    ```go
    // a better if else
    switch { // note not condition
      case a > b:
        //...
      case c < d:

    }
    ```

## defer

* A defer statement defers the execution of a function until the surrounding
  function returns. The args to any function called, are however, evaulated
  immediately
* Deferred function calls are pushed onto a stack. When a function returns, its
  deferred calls are executed in last-in-first-out order.
  ```go
    func main() {
        if len(os.Args) < 2 {
            log.Fatal("no file specified")
        }
        f, err := os.Open(os.Args[1])
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()    // mind the () ... you pass args to the function via that.
        data := make([]byte, 2048)
        for {
            count, err := f.Read(data)
            os.Stdout.Write(data[:count])
            if err != nil {
                if err != io.EOF {
                    log.Fatal(err)
                }
                break
            }
        }
    }
  ```

# functions

```go
func name(parameter-list) (result-list) {
    body
}
```

* 0 or more args. 0 or more return values
    ```go
    func add(x int, y int) int {
        return x + y
    }
    ```
* When two or more consecutive named function parameters share a type,
  you can omit the type from all but the last.
    ```go
    x int, y int
    //to
    x, y int
    ```
* Naked returns (to be avoided)
    * This simply returns the variable with same name as parameters
    ```go
    func split(sum int) (x, y int) {
        x = sum * 4 / 9
        y = sum - x
        return
    }
    ```
* go doesn't have named/optional params. If you need them use a struct as arg
* functions associated with types are called methods. These sport the anchoring
  type in their declaration like this:
    ```go
    func (m *MyType) methodFunc(args int) (result int) {
    }
    ```
* Function calls can precede function declaration within the package. Unlike c,
  there is no declaration/definition distinction. Its just one place and go
  calls it declaration.
* variadic function have ellipsis at the last arg's type.
  This makes the funciton take any number of args of that type.
  Internally its accessible as a slice of that type.
  ```go
    func addTo(base int, vals ...int) []int {
        out := make([]int, 0, len(vals))
        for _, v := range vals {
            out = append(out, base+v)
        }
        return out
    }
    // call
    func main() {
        fmt.Println(addTo(3))
        fmt.Println(addTo(3, 2))
        fmt.Println(addTo(3, 2, 4, 6, 8))
        a := []int{4, 3}
        fmt.Println(addTo(3, a...))  // expands the slice as individual int args
        fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
    }
  ```
* `...interface{}` makes the funciton take any type of arg any
  number of times.

## anonymous functions

* Useful as other function args, closures or in defer statements

```go
func main() {
    for i := 0; i < 5; i++ {
        func(j int) { // anonymous function
                fmt.Println("printing", j, "from inside of an anonymous function")
            }(i)
        }
    }
}

// sort by last name
sort.Slice(people, func(i int, j int) bool {     // anonymous func and also a closure. Captures the people
    return people[i].LastName < people[j].LastName
    })

```


## closures

* go supports closures
* Dont pass loop-vars to a function over a closure.
* basically the closure gets a reference to the variables
* see the slice sort example above

## go builtin functions

* make
    * used to create a slice, map, chan
    ```go
    // args to make in T is the type name itself
    make ([]T, len, cap) T[]
    make (chan T)
    make (map[T1] T2)
    ```
* len
* cap
* new
    * `new(T)` creates a unnamed variable of type T, initializes it to
      zero-value, and returns `T*`, a pointer to the type.
    * new is only a syntactic convenience, and avoids having to create a
      name. (Different from c/cpp in this regard. In go, every variable
      is like from heap (compiler chooses stack/heap depending on how
      its used)
    * To confirm:
        ```go
        a := new(MyType)
        //is same as
        a := &MyType{}
        ```
* append
    * adds an element to a slice. If slice has capacity, its very fast.
      If slice capacity doesn't fit, it creates a new array, copies
      existing elems and then appends.
    * Never forget to assign the result of append back to the original
      varaible, as it could have changed.
      ```go
      runes = append(runes, r)
      ```
* copy
    * copies from one slice to another. It stops as the shortest of the two.
        ```go
        num_copied = copy(dst, src)
        ```
    * cap of the two doesnt matter. its the len that matters
    * Safe against overlapping slices.
    * Returns number of elements actually copied - the smaller of the 2 slices.
      So its safe againt unavailable sizes too
* close
    * close a channel
* delete
    * used to delete a key in a map
* complex
* real
* imag
* panic
    * Takes any arg
* recover


# error

* definition
    ```go
    type error interface {
        Error() string
    }
    ```
* To generate a error:
  ```go
  err := errors.New("this is a new error")
  err2 := fmt.Errorf("whatever went wrong:%d", int_arg)
  ```
* Idiomatic go
    * Dont start error strings with caps
    * dont end them with punctuations
* wrapping and unwrapping
  ```go
  // error struct should support this
  func (eimpl ErrorImpl) Unwrap() error

  //use %w to wrap another error
  return fmt.Errorf("in my file, faced err:%w from lib", err)

  // or use this
  errors.Wrap(myerr, "read failed")

  // to get the wrapped error.. nil if there is none or the err doesn't implement Unwrap
  wrappederr := errors.Unwrap(err)
  ```
  * wrapped errors are called a error chain
* Use `%+v` verb to print a error with its stack trace

## sentinel errors

* Errors which are expected to be checked against directly.

```go
retVal, err := function_with_sentinel_err(args)
if err == ASentinelErr {

}
```
* Sample
```go
func fileChecker(name string) error {
    f, err := os.Open(name)
    if err != nil {
        return fmt.Errorf("in fileChecker: %w", err)
    }
    f.Close()
    return nil
}
func main() {
    err := fileChecker("not_here.txt")
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {  // checks in the error Chain for this sentinel error
            fmt.Println("That file doesn't exist")
        }
    }
}

```

## panics

* equivalent of assert
```go
panic("your message")

```

# Managing Go code

* Roughly
```
repository > module > packages > files
```
## repository

* Is where the code lives - typically something that is source controlled
* While you can keep multiple modules in one repo, practise is to have
  one repository per module so that a module can be version controlled
  on its own properly

## modules

* Globally identified by its repository's path and its name. Eg:
    ```
    github.com/some-company/some-project/module
    ```
* Module management commands

```sh
## create a new module giving its MODULE_PATH
##   this creates a go.mod file
go init mod public-domain.com/path/to/module/
```

* roughly how a go.com looks like
    * the optional `require` captures the dependent pkgs
    * a optional `exclude` section prevents a specific version
      of a  module from begin used
```
module github.com/learning-go-book/money

go 1.15

require (
    github.com/learning-go-book/formatter v0.0.0-20200921021027-5abc380940ae
    github.com/shopspring/decimal v1.2.0
)
```
* you might want to let go know that for now, modules are available locally
    ```sh
    go mod edit -replace public-domain.com/module/i/need/modA=../local/path/to/modA
    ```
* this (re)generates the mod file
    ```sh
    go mod tidy
    ```
* GOROOT/GOPATH will always be searched.
    * GOPATH can have multiple values like PATH
    * GOPATH/src is where source code lives.

* Other module mgmt commands
```sh
## list all versions in a module
go list -m -versions github.com/learning-go-book/simpletax

## replace to a paraticular version -- note the @
go get github.com/learning-go-book/simpletax@v1.0.0
```


## Packages

* Every Go program is made up of packages. Programs start running in package main.
* By convention, the package name is the same as the last element of the import path
    * This means, typically you put all files of a package in a folder
    * All files in a folder must have the same package name
        * Excpetion is main
        * Or when dir name cant be a valid pkg name. (avoid this)
* package can be spread across files (like namespaces in cpp)
* Each sourcefile calls out its package as its first line.
    * (unlike cpp) one source file is fully one package only.
* import is how one file uses another package's identifiers. imports can be
  grouped into a parenthesized, "factored" import statement.
  This is preferred over individual imports.
  ```go
  import (
       //  importing from std library
      "fmt"
      "math"

      // alternate name for a module in this file
      // used to avoid conflicts if 2 modules have same name
      crand "crypt/rand"

      // import form other modules
      // later used with just formatter.FunctionThere()
      "github.com/whatever-repository/formatter"
  )
  import myhttp "mypath/http"
  ```
  * imports have file scope only. If you want to import on 2 files of same package
    you have import in both files
* In Go, a name is exported if it begins with a capital letter. Otherwise its private
  to the package its in.
* Avoid relative paths.

## init

* Avoid this if possible
* `init()` is a special no-arg, no-return-type function that can appear any number of times
  in every file, in any package. These are called in declaraion order and are invoked
  before the package is considered initialized.
  init()s of all packages are called before main()
* calling out
  ```go
  import _ "mypath/http"     // blank identifier
  ```
  triggers a intialization of the package even if no references to the package is done.
  Otherwise compiler will complain of redundant import


# Packages in standard library

## os

* os.Args[] - slice of cmd line args. os.Args[0] is the command itself.
* os.Stdin  - a io.Reader for stdin
* os.Exit(1) - exit with a error code.

```go
// open a file. file of type *os.File
// that *File is a io.Reader, io.Writer
file, err := os.Open(dataFile)
if err != nil {
    return nil, err
}
```


## fmt

* fmt.Println
* fmt.Printf
* fmt.Fprintf
    ```go
    fmt.Println(split(17))
    var i int
    fmt.Println(i, c, python, java)

    fmt.Printf("Regular c style printing with formats:%d", i)
    ```

```go
type Stringer interface {
    String() string
}

```


### Verbs

Format specifier in go is called a verb. THe one between % and verb is an adverb

* %v is verb(name in go for format-specifier) to choose the default format for the
   passed type
* %T is for type of the value
   * %x for strings prints 2-hex-digits for each byte in string. An option space(adverb
     in go) adds a space between each byte.
   * %x for []runes prints the runes in utf encoded hex values


## log

```go
// Change the device for logging to stdout.
log.SetOutput(os.Stdout)

// regular log
log.Printf("fmt", values)
log.Println - 

// printf and then exit
log.Fatalf("fmt", exit)
log.Fatal(stringvarorliteral)


```


## strings

* strings.Join(a []string, sep string)
    * concatenates elements of a to make a bing string using sep
* strings.LastIndex

## bytes

* bytes.Buffer - efficient type for manipulation of []byte
    * Implements the io.Writer and fmt.Stringer interface
    * bytes.Buffer.WriteByte()

## strconv

* `intVar, err := strconv.Atoi(stringVar)`

### unicode/utf8

* utf8.DecodeRuneInString - gets rune at a index i
* utf8.RuneCountInString

## bufio

* Scanner
    * Reads a input and breaks it into lines
    * Scanner.Scan() - reads one line, strips the newline. Returns True/false on whether a line was read or not.
    * Scanner.Text() - gets the line previous read by Scan()
* bufio.NewScanner
    * returns a *Scanner from a io.Reader

## io/ioutil

* ReadFile
    * Given a filename returns byte slice/err of file contents
* io.EOF  // an error type
* Discard - sth like /dev/null sink

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
type Closer interface {
    Close() error
}
type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
    // whence is one of io.SeekStart, io.SeekCurrent, and io.SeekEnd.
}
io.ReadCloser
io.ReadSeeker
io.ReadWriteCloser
io.ReadWriteSeeker
io.ReadWriter
io.WriteCloser
io.WriteSeeker

// write/read in one slurp
ioutil.ReadAll
ioutil.ReadFile
ioutil.WriteFile

```

* string to io.Reader:
    ```go
    s := "The quick brown fox jumped over the lazy dog"
    sr := strings.NewReader(s)
    ```
* io.MultiReader
    * gives a io.Reader from many io.Reader one after another
* io.LimitReader
    * gives a io.Reader that will stop reading after a limit num bytes
* io.MultiWriter
    * gives a io.Writer from many that will write to all at the same time





## net/http

* http.Get(url) resp,err

## time

```go
// datattype
time.Time      // hold a wall-clock time akin to datetime.datetime of python
time.Duration  // delta between 2 Time s. akine to datetime.timedelta of python
               // tracks in ns. Can diff upto 290 years.

// Duration constants
time.Second
time.Minute
time.Hour
a := 5 * time.Second // a is of type time.Duration

var t time.Time = time.Now()      // gives the current time. search: now current epoch
d := time.Until(t)                // gives the time.Duration from now till t
n := now.Add(timeout)             // add (timeout time.Duration) in sections to now

// time pkg constants for formats
time.Day
time.Month, Year, Hour, Minute, Second, Weekday,
Clock (just hh:mm:ss)
Date  (just yyyy-mm-dd)

//compare time.Time
After, Before, Equal


// returns a channel that outputs once after said time
time.After
// returns a channel that outputs every perdiodic time
// avoid-- as the ticker cannot be stopped/gc'ed.
time.Tick
// prefer this.. returns *time.Ticker
// has the channel to read do and stop function
time.NewTicker()
// invoke a function after some time in its own go-routine
time.Afterfunc()

var h int = math.Floor(d.Hours()) // converts the duration to hours




// time from epoch
// the 64 is to tell give a int64
i, err := strconv.ParseInt("1405544146", 10, 64)
if err != nil {
    panic(err)
}
tm := time.Unix(i, 0) // sec and nsec
fmt.Println(tm)

// Run a forever loop in go
go PeriodicallyReportGatewayStatus(time.Second*60)

func PeriodicallyReportGatewayStatus(dur time.Duration) {
    for range time.Tick(dur) {           // «««--- is the forever syntax
        err := reportGatewayStatus()
        if err != nil {
            glog.Errorf("err in reportGatewayStatus: %v\n", err)
        }
    }
}

```


* time.Sleep(d Duration)

## sync

```go
// Setup a wait group so we sync all go-routings
var waitGroup sync.WaitGroup

// Set the number of goroutines we need to wait
waitGroup.Add(len(go_routines_created))

// mutex type
sync.Mutex
```


## sort

* sort.Interface
    * Needs Len, Less, Swap
* sort.Reverse

## regex

## json

```go
import 'encoding/json'
```

* json.Marshall
* json.MarshallIndent
* json.Unmarshall  -- ignores json fields which aren't in the struct declaration

## bits

```go
bits.TrailingZeros32(auint32arg)  // counts the number of trails 0z in the unit32 arg
bits.Len32(x)                     // gives the number of bits required to represent this number
                                  // In other words, gives the 1-idx'ed MSB-1 position

```

## context

```go

// basic context
ctx := context.Background()

// for http req
ctx := req.Context()
req = req.WithContext(ctx)

// ctx with cancel
ctx, cancel := context.WithCancel(ctx)
defer cancel()

// ctx with tiemout
parent, cancel := context.WithTimeout(ctx, 2*time.Second)
defer cancel()

// ctx with Value
func ContextWithUser(ctx context.Context, user string) context.Context {
    return context.WithValue(ctx, key, user)
}
func UserFromContext(ctx context.Context) (string, bool) {
    user, ok := ctx.Value(key).(string)
    return user, ok
}

```

* supprot context in your own code

```go
func longRunningThingManager(ctx context.Context, data string) (string, error) {
    type wrapper struct {
        result string
        err error
    }
    ch := make(chan wrapper, 1)
    go func() {
        // do the long running thing
        result, err := longRunningThing(ctx, data)
        ch <- wrapper{result, err}
    }()
    select {
    case data := <-ch:
        return data.result, data.err
    case <-ctx.Done():
        return "", ctx.Err()
    }
}


```




# Go tools

```sh
go build         # creates a exe in same dir with same name as package main file name
go build -o /path/to/exec file.go
## other args
## -trimpath   .. dont include fullpath in stack-traces
go run file.go   # Just run as a script
go install       # build, but put exe in $GOPATH/bin

go test path1/path2/a.go   # Not sure. check

go get abc.com/repo_name/path/file.go  # pulls that file (repo) in $GOPATH/src

# the @is the version to pull. latest is well, latest
# same command for both first time and for install.
# it fetch and also compiles
go install github.com/rakyll/hey@latest

# cleans executables
go clean file.go

# will create a go.mod that list all deps of this package
go mod init path/from/gopath/src/to/this/package

go mod edit -replace public-domain.com/module/i/need/modA=../local/path/to/modA

# verifies unnecessary imports and adds imports
# but can be wrong.
goimports -l -w .
## args
##  -l      .. print lines with incorrect formatting
##  -w      .. write files in place

# format a file
go fmt file.go

## line the entire project
golint ./...

## detect printf args correct etc..
go vet file.go

```

# documentation

* comments above any of packages, function, types and global vars

```go
// Retrieve connects to the configuration repository and gathers
// various connection settings, usernames, passwords. It returns a
// config struct on success, or an error.
func Retrieve() (config, error) {
// ... omitted
}
```


# vim

```vim
Plugin 'fatih/vim-go'
"should do most of the stuff. Just add this and plugin install it.
:Godoc
```

# Reading Help

* use godoc
    ```sh
    godoc <pkg-name>
    godoc image/gif
    godoc time.Now  # doesn't work in my m/c though. But good
    ```
* Has CONSTANTS, FUNCTIONS, TYPES

# Stuff not available in go

* no implicit numer conversions
* no constructors or destructors
* no operator overloading
* no default parameter values
* no inheritance
* no generics
* no exceptions
* no macros
* no function annotations
* no thread-local storage

# Read later

* https://blog.golang.org/defer-panic-and-recover[Defer-panic-and-recover]

* go in action  ch: 4.2


