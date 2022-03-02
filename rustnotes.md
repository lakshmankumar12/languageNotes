# source

https://doc.rust-lang.org/book/ch01-03-hello-cargo.html

# cargo basics


```
#create a new project
cargo new project_name

# src files in project_name/src
cd project_name

# build and create executable
cargo build

# just compile
cargo check

# run
cargo run

```

# types

## number bytes

```
8-bit   | i8    | u8
16-bit  | i16   | u16
32-bit  | i32   | u32            # i32 is the integer literal default
64-bit  | i64   | u64
128-bit | i128  | u128
arch    | isize | usize          # 32/64 depending on machine-type

float   | f32
        | f64                    # defaut for float literals

bool

char   # 4 bytes in size - unicode point value.

```

## complex

```
#tuples
let tup: (i32, f64, u8) = (500, 6.4, 1);
let tup = (500, 6.4, 1);                     //  (i32, f64, i32) implicit.

## called destructing the tuple
let (x, y, z) = tup;

## directly access a member
let x: (i32, f64, u8) = (500, 6.4, 1);
let five_hundred = x.0;                    // note: 0 based index

## The tuple without any values, (), is a special type
## that has only one value, also written (). The type is called the unit type
## and the value is called the unit value. Expressions implicitly return
## the unit value if they don’t return any other value.

#Array type  - same type all elem.
let a = [1, 2, 3, 4, 5];
let a: [i32; 5] = [1, 2, 3, 4, 5];
let a = [3; 5];   // same as let a = [3, 3, 3, 3, 3];

## accessing array elments
let a = [1, 2, 3, 4, 5];
let first = a[0];
let second = a[1];

## accessing out-of-bound index throws a runtime panic.


```

# literals

## number types

```
Decimal        | 98_222
Hex            | 0xff
Octal          | 0o77
Binary         | 0b1111_0000
Byte (u8 only) | b'A'

57u32     # literal of a particular type

true/false  # bool literals

'A' , 'ℤ'   # char literals
```

## ranges

```
(1..10)   # returns a Range-type
(1..=9)   # same as (1..10)
```

# rust variables

* `let variable = value` . Rust gets the type from the value. Immutable by default.
* `let mut variable = value`. Creates a mutable type
* `let variable: u32 = 10`.  Explicity typing.
* `const LIGHT_SPEED : u64 = 3 * 1000 * 1000 * 100`

# Constructs

```

match value {
    value1 => result_value1,
    value2 => {
        break;
    }
    value3 => println('This is value3'),
}

let y = {
        let x = 3;     // is a statement - ends with ;
        x + 1          // is a expression. This is value of the {}, which is also an expression.
    };

    if number % 4 == 0 {                                  // Note no parenthesis
        println!("number is divisible by 4");
    } else if number % 3 == 0 {
        println!("number is divisible by 3");
    } else if number % 2 == 0 {
        println!("number is divisible by 2");
    } else {
        println!("number is not divisible by 4, 3, or 2");
    }

## truthy: unlike python, ints can't be directly used in if ... Only bools are allowed
## we need to explicitly check num != 0

let number = if condition { 5 } else { 6 };         // allowed
let number = if condition { 5 } else { "string" };  // compile error: all branches should return same type.


loop {

    break;

    continue;

}

'label: loop {

    break 'label;

}

    let result = loop {
        counter += 1;

        if counter == 10 {
            break counter * 2;        // result of the loop!
        }
    };


    while number != 0 {
        println!("{}!", number);

        number -= 1;
    }

    let a = [10, 20, 30, 40, 50];

    for element in a {
        println!("the value is: {}", element);
    }

    for number in (1..4).rev() {        // The (1..4) is a range!
        println!("{}!", number);
    }


fn add(x: i32, y: i32) -> i32 {
    x + y               // Note the absence of ; This is a expression & the value of the {}-as well.
                        // A ; will make it a statement and render the function value-less (compile error)
}

```

# concepts

* ownership
    * Copy-trait types (simple integers that can be copied)
    * Drop-trait types
        * string like. They point to a heap location
        * when u assign one from another, the older variable goes out of scope/life,
          as the owned heap location is moved.
        * Same goes for fun arg passing and return value. Ownership moves.
* References
    * You can have any number of immutable referenes
    * You can have only one mutuable reference.
    * References go out of scope after the last line they are used.
* slices
    * `&str` is the slice for a `String`
        * string literals are actually string slices (pointing to a location in the binary of a program)
        * they are like immutable references to the underlying String
        * eg:
            ```
            let my_string = String::from("hello world");

            // `first_word` works on slices of `String`s, whether partial or whole
            let word = first_word(&my_string[0..6]);
            let word = first_word(&my_string[..]);
            ```
    * `&[i32]` is a slice into a array of i32
        ```
        let a = [1, 2, 3, 4, 5];
        let slice = &a[1..3];
        ```

# useful function

* `string_type.parse()`. returns io::Result which is either ok(num) or Err(_)
