CPP Notes
=========

## Miscellaneous notes

### New

* allocate memory  .. operator new (size_t)
* call constructor

### Delete

* call destructor
* operator delete(void*)

## $$C++14$$ Features

### Enum class

```
enum class Traffic_light { Red, green, yellow };
```

* strict type-checking
* we can define operators for this.
  `Traffic_light& operator++(Traffic_light& t)`

```
static_assert(A,S) prints S as a compiler error message if A is not true.
```

### Const Expressions

```
const int a=5; <-- a is a named-const
constexpr double max1 = 1.4∗square(dmv); // constexpr: meaning roughly ‘‘to be evaluated at compile time’’
                                         // In above example, the function square should also be a constexpr like below
constexpr double square(double x) { return x∗x; }   // a constexpr function can be called from anywhere. Just that the result wont be constexpr.
                                                    // This allows same function for both variables and constexpr
```

### Initializer list

* Use the {} when in doubt. It can initialize simple things like int, or vectors!
* It enforces type check. `int i=7.2;` will be accepted, but `int i{7.2}` will not.

```
std::initializer_list<T> lst;
 Has the following attributes
   lst.size()

eg: std::initializer_list<double> lst{2.2,2.3,2.4};
```

https://stackoverflow.com/questions/629017/how-does-array100-0-set-the-entire-array-to-0

### Unique Ptrs

```
unique_ptr<T>
```

* Better than auto_ptr
* when it goes out of scope, it deletes the object underneath

```
shared_ptr<ptr,deleter_func>
```

A shared_ptr is twice the size of a raw pointer, uses dynamically allocated
memory for bookkeeping and deleter-specific data, uses a virtual function
call when invoking its deleter, and incurs thread synchronization overhead
when modifying the reference count in an application it believes is
multithreaded. (You can disable multithreading support by defining a
preprocessor symbol.)

### Moves

```
std::move(x);
  -> get an rvalue for x. This helps in invoking a move-assignment
      z=x;            // copy-assignment
      z=std::move(x); // move assignment
```

```
=delete; can be used to suppress any operation, that the compiler does for free.
```

### static asserts

```
assert(4<=sizeof(int), "integers are too small");
```



## Signatures for common function

search: operator overloading

### binary operators

```
const Object operator+(const Object &lhs, const Object &rhs);
```

* note the return is const-value, while args are const-references

### assignment

```
Object& Object::operator=(const Object &rhs);
Object& Object::operator=(const DiffTypeObject &rhs);
```
* note non-const reference return and const reference arg.
    * Only then constructs like  (a=b).nonConstMember() will work.
    * std::string for eg allows this. (a=b).clear()
* same for all assignments like +=,-=,*=,/=

### prefix / postfix

```
Object& operator++()         // prefix
const Object operator++(int) // postfix

Object& operator--()         // prefix
const Object operator--(int) // postfix

non-member:
friend Object& operator++( Object& )      // Prefix increment
friend Object& operator++( Object&, int ) // Postfix increment
```

### print operator

```cpp
//add this friend to your class
friend ostream& operator<<(ostream& os, const MyClass& cl);

```


## Miscellaneous notes

Adding explicit to a constructor prevents implicit invocation to convert types

### What are aggregate/non-aggregate and PODs

https://stackoverflow.com/questions/4178175/what-are-aggregates-and-pods-and-how-why-are-they-special

In short, Aggregates:
* Shouldn't have explicit copy/default constructor
* No private/protectd non-static members
* Can have user-defined assignment-operator/destructor
* Array is an aggregate, even if its array of non-aggregate type

Basically, Aggregates can be in C++03, brace-intialized.

POD is a stricter aggregrate
* It should not have destructor/assign-oper
* It should be built by PODs only. No non-POD non-static members

Basically, a POD can be memcpy'ed back and forth, reinterpret-casted and should still work

## Links

* Memory model , single-producer, single-consumer no lock
   https://www.codeproject.com/Articles/43510/Lock-Free-Single-Producer-Single-Consumer-Circular
* Lockless, multi-producer, multi-consumer
   http://www.linuxjournal.com/content/lock-free-multi-producer-multi-consumer-queue-ring-buffer?page=0,2
* On rvalue references
   http://thbecker.net/articles/rvalue_references/section_01.html
* On auto/decltype
   http://thbecker.net/articles/auto_and_decltype/section_01.html
* New features
   https://medium.com/free-code-camp/some-awesome-modern-c-features-that-every-developer-should-know-5e3bf6f79a3c

# Exception

code look:
```c
    try {
        whatever();
    } catch (derived_exception &e) {
        cout << e.what() << endl;
    } catch (base_exception &e) {
        cout << e.what() << endl;
    } catch (std::exception &grand_base) {
        cout << e.what() << endl;
    }
```

throw(a,b,c) is called the exception specifier
    * violation is caught at run-time, not at compile time
        * I saw that a throw in the same function-scope raised a compile error.
        * But this is just delusion, as called functions can still throw and that is caught at runtime
noexcept is the C++11 keyword
    * it is a compile time thing? [to confirm]

exception specifier is a bad idea - https://stackoverflow.com/a/88905/2587153
    * Template class is impossible to write
    * they prohibit extensiblity
    * if some function called inside throws sth else, it calls terminate()
        * this dies silently/violently w/o stack-trace

noexcept - offers a binary choice, functions that throw and that that dont

* Note that when catching exceptions, u should catch by reference.
* When catching exceptions, you should list derived first, followed by base.
    * Putting base first, will match all derived too.

## On the roll

### Bad stuff list

* Defining variables too soon can cause a drag on performance.
* Overuse of casts can lead to code that’s slow, hard to maintain, and infected
  with subtle bugs.
* Returning handles to an object’s internals can defeat encapsulation and leave
  clients with dangling handles. Failure to consider the impact of exceptions can
  lead to leaked resources and corrupted data structures. Overzealous inlining
  can cause code bloat.
* Excessive coupling can result in unacceptably long build times.


## Effective C++

### Accustoming to C++

#### Item 1: parts of c++ -> c, oo-part, template-part and STL

#### Item 2: const, enum and inlines to #define
* static const inside class needn't have definition (integral
  types - as long as address is not taken)
* `#define` dont have scope
* inline functions have type-checking, and can accept expressions
  as arguments, which wil be evaluated once.

#### Item 3: use  const wherever possible
* const Rational operator*(const Rational &lhs, const Rational &rhs);
    * const helps to avoid unnecessary assignment to a return value
* member functions can be overloaded with const
* bitwise constness, logical constness


### Constructors, Destructors and Assignment Operators

#### Item 4: Objects should be initialized before their use
* use initialization list
* order of initailzation is base-class, member definition order.
* avoid initialization order problem, by having data-members
  local to compilation unit and static

#### Item 5: Functions that are auto-generated by compiler
* Default constructor (But this wont be provided if some
  other constructor is explicity declared)
* copy constructor, assignment operator , destructor
* copy/assignment will be rejected if the class has any const member
  or a reference (references are kind of `*const`)

#### Item 6: Explicitly disallow functions if u dont want
* Make copy constructor and assignment private if the class shouldn't copy.
* One can use a Uncopyable base class as show in the item.
    * Note the private inheritance

#### Item 7: Polymorphic class should have virtual destructor
* Polymorphic classes should have virtual destructor
* non polymorphic classes shoudn't

#### Item 8: Destructors should not throw exceptions
* Destructor may already be called in a exception unwinding stack. So
  it may leaad to double-throw
* Swallow it in destructor (use try {} catch(...) {} blocks in
  destructors - Note the ellipsis)
* If class clients need to be able to react to exceptions thrown during
  an operation, the class should provide a regular (i.e., non-destructor)
  function that performs the operation.

#### Item 9: Never call virtual functions from constructors or destructors
* In a constructor, constructor will only statically call the base-class
  function from within constructor scope!
  Even if u call out to another method and call another virtual function, it
  will disastrous, as the derived object hasn't been built yet.

#### Item 10: Follow assignment operator signature convention (see above for the signature)

#### Item 11: Handle self-assignment in assignment operator.
* Ensure order of operatorions is right. Dont delete the rhs's data before lhs is ready

#### Item 12: Initialize all parts of a class
* Ensure to initialize/copy/assign all members of the class
* Ensure to call the right copy/assingment operators of all base classes
* Dont implement one constructor/assignment operator with another. Have a 3rd
  function and calls this from all constructors and assignment operators

### Resource Management

#### Item 13: Use objects to manage resources
* Use resource-management objects to manage resources. Acquire resources as
  part of the constructor of these objects and let their destructors return
  the object. This is called RAII
* `std::shared_ptr` and `std::auto_ptr` are 2 common classes for this. Note
  that `shared_ptr` can't break cycles.

#### Item 14: Think carefully about copying of resource-management
* we can prohibit copying / reference-count underlying resource / duplicate or
  deep-copy / transfer ownership (`auto_ptr`)

#### Item 15: Provide access to raw resources
* offer a .get() method to get the underlying resource. (but needs
  clients to invoke this!)
* offer operator overloads like -> * to allow natural usage.
* you may do a operator UnderlyingObject() to allow implicit conversion, but
  this may result in dangling references!  (but same issue results with
  .get(), just that its more explicit to see)

#### Item 16: Use new/delete and new[]/delete[]
* esp be careful when the [] is hidden inside a typedef

#### Item 17: Use separate statements to store new'ed objects into resource-mgmt objects
* That is, dont store them into resource-objects in function arguments.

### Designs and Declarations

#### Item 18: Make interfaces easy to use correctly and hard to use incorrectly
* Keep interfaces consistent. (Eg: name all length functions of all containers the same - like size())
* Behavorial compatibility with built-in-types. If you overload operators, behave as ints do..
* Prevent errors in invoking using new-types, restricting operations on types, constraining object values
* Strive to minimize client-interface responsibiliies (like having to call delete on a ptr received)
  (`shared_ptr` can help here, as well provide the custom deleter function)

#### Item 19:  Treat class design as type design
* How should objects of your new type be created and destroyed?
* How should object initialization differ from object assignment?
* What does it mean for objects of your new type to be passed by value?
* What are the restrictions on legal values for your new type?
* Does your new type fit into an inheritance graph?
* What kind of type conversions are allowed for your new type?
* What operators and functions make sense for the new type?
* What standard functions should be disallowed?
* Who should have access to the members of your new type?
* What is the “undeclared interface” of your new type?
* How general is your new type?
* Is a new type really what you need?

#### Item 20:  Prefer pass by reference-to-const over pass-by-value
* Its more efficient.
* Some exceptions are built-in-types, STL iterators, function-object-types

#### Item 21:  Dont return a reference when you should return a object
* returning reference to local(stack) object is outright wrong.
* reference to heap/static(function) is also wrong
* (item-3: return a const object if assignment to this temporary is to be avoided)

#### Item 22:  Keep data-members private
* protected is not very much encapsulated than public

#### Item 23:  Prefer non-member non-friend functions to member function
* A bit anti-intuitive, but this reduces the amount of code that is to be changed when
   private data-members change.
* You can split these function in multiple header files
* Clients can extend these functions.

#### Item 24:  Declare non-member functions when type conversions should apply to all parameters.
* Rational a = 2 * b; will work by converting 2 to Rational only if
  `const operator*(const Rational &,const Rational&)` is non-member.

#### Item 25: NEEDS RE_READING. Template stuff.

### Implementations

#### Item 26:  Postpone variable definitions as long as possible.
* if sth is needed in a loop, define it within loop (construct/destrcut
  in every iteration) unless the cost of assgn is a lot cheaper than
  construction/destruction.

#### Item 27:  Minimize casting
* Avoid c-sytle and fn-style casts. Fn-style is okay if you are
  explicitly calling constructors. Otherwise use the c++ casts.
* const_cast<T>(expr)
    * removes const. Only cast that can do it.
* dynamic_cast<T>(expr)
    * base to drived. Costly. May do lots of strcmp(). Cannot be performed by c-style cast
* reinterpret_cast<T>(expr)
    * cast ptr to int etc.. May not be portable
* static_const<T>(expr)
    * Safe cast (Probably one cast that is like okay. The above 3 are largely bad).
    * force implicit converstions. Non-const to const, void* to ptr*, int-to-double, ptr-base to ptr-derived (without check)
* Try hiding casts inside functions so that client code is free
  of them.

#### Item 28:  Avoid returning handles to object internals.
* It opens up encapsulation.
* Note the example, where a const member function returns a
  non-const data (as data is not immediately in this object,
  but is pointed to from the object)
* Note the example, where a reference to object is dead, as
  its from a temporary object (an expression result unnamed).
* operator[] typically works by returning reference, but
  this is an exception and not a rule.

#### Item 29:  Strive for exception safe code
* nothrow (doesn't throw anything) is a non-compiler controlled
  guarantee
  throw() -- means, if the fuction throws sth, its a serious
  error (like assert) and the unexpected() will be called.
* basic guarantee -- program is in some valid state after the
                     exception is thrown
  strong guarantee -- program is in prev state before func was
                      called (as if the fn wasn't called)
  nothrow guarantee -- doesn't throw exceptions.
* strong guarantee is typically implemented by copy-and-swap
  technique, but not practical or possible for all functions.
* a function can't guarantee exception safety beyond the
  weakest of the guarantee of the functions that it calls

#### Item 30:  Understand the ins and outs of inling
* Limit most inlining to small, frequently called functions. This
  facilitates debugging and binary upgradability, minimizes potential
  code bloat, and maximizes the chances of greater program speed.
* Don’t declare function templates inline just because they appear in
  header files.

#### Item 31:  Minimize compilation dependencies
* The general idea behind minimizing compilation dependencies is to
  depend on declarations instead of definitions. Two approaches
  based on this idea are Handle classes and Interface classes.
* Library header files should exist in full and declaration-only
  forms.  This applies regardless of whether templates are involved.

### Inheritance and object-oriented design

#### Item 32:  Make sure public inheritance models "is-a" relationship
* Everything that applies to base classes must also apply to
  derived classes, because every derived class object is a base class
  object.
* Remember the penguin-bird-fly problem,
  square-rectangle-change-height problem

#### Item 33:  Avoid hiding inherited names
* Even if one function in derived has same name as a non-virtual
  fn in base (with diff args), it will hide all base class fns with
  same name.
* Names in derived classes hide names in base classes. Under public
  inheritance, this is never desirable.
* To make hidden names visible again, employ 'using' declarations or
  forwarding functions(explicit fns in derived that call Base::fn)

#### Item 34:  Differentiate between inheriting interface and inhering implementation
* Public inheritance always inherits base-class interface
* Pure-virtual -> inherits interface only
    *  (if a default impl is desired on explicit req, impl the
    pure-virtual fn, and use Base::fn_name in derived class
    to choose that impl.)
*  simple-virtual -> inherits interface and an optional implementation
*  non-virtual -> inherits both interface and a mandatory impl.

#### Item 35:  Consider alternatives to virtual functions
* Use NVI(non-virtual interface) idiom.
    * wraps a private virtual function with a non-virtual public interface
      in base-class
    * Derived classes can't call base-virtual function, but can re-implement
      them
    * Base class's non-virtual implementation, can do some pre-stuff,
      post-stuff(like locks/state-assertions,fill-default-values)
      before calling the virtual function. (this is useful if pre-stuff/post-stuff
      are property of base-class impl.)
* Replace virtual functions with fn-pointer data-members
    * But this limits the fn-poiners to only call public interface of the class
    * Might open up encapsulation!
* Use `tr1::function<>` types. This allows assigning fn-pointers with any
  compatible signature type
* Replace virt funs in one hierarchy with vir. functions in antoher hierarchy
  (This is the conventional strategy pattern)

#### Item 36:  Never redefine an inherited non-virtual fn

#### Item 37:  Never redefine an inherited virtual fn's default value
* Default values are statically bound! If needed use the NVI idiom.

#### Item 38:  Model "has-a" or "is-implemented-in-terms-of" using composition
* Composition is diff from is-a (public inheritance)
* Composition is either "has-a" (application domain) or "is-implemented-in-term-of" (impl. domain)

#### Item 39:  Private inheritance
* Rules:
    * derived ref/ptr wont become base ref/ptr
    * public/prot members of base will be priv of derived.
* Private inheritance means is-implemented-in-terms of. It’s usually
  inferior to composition, but it makes sense when a derived class
  needs access to protected base class members or needs to redefine
  inherited virtual functions.
* Unlike composition, private inheritance can enable the empty base
  optimization. This can be important for library developers who strive
  to minimize object sizes.

#### Item 40:  Use multiple inheritance judiciously
* Multiple inheritance is more complex than single inheritance. It can
  lead to new ambiguity issues and to the need for virtual inheritance.
* Virtual inheritance imposes costs in size, speed, and complexity of
  initialization and assignment. It’s most practical when virtual base
  classes have no data.
* Multiple inheritance does have legitimate uses. One scenario involves
  combining public inheritance from an Interface class with
  private inheritance from a class that helps with implementation.

### Templates and Generic Programming

#### Item 41:  Understand implicit and compile-time polymorphism
* Both classes and templates support interfaces and polymorphism.
* For classes, interfaces are explicit and centered on function
  signatures. Polymorphism occurs at runtime through virtual
  functions.
* For template parameters, interfaces are implicit and based on
  valid expressions. Polymorphism occurs during compilation through
  template instantiation and function overloading resolution.

#### Item 42:  Understand the two meanings of typename
* When declaring template parameters, class and typename are
  interchangeable.
* Use typename to identify nested dependent type names, except in
  base class lists or as a base class identifier in a member
  initialization list.

#### Item 43:  Know how to access names in templatized base classes
*  In derived class templates, refer to names in base class templates
   via a “this->” prefix, via using declarations, or via an explicit base
   class qualification.

#### Item 44:  Factor parameter-independant code out of templates
* Templates generate multiple classes and multiple functions, so any
  template code not dependent on a template parameter causes bloat.
* Bloat due to non-type template parameters can often be eliminated
  by replacing template parameters with function parameters or class
  data members.
* Bloat due to type parameters can be reduced by sharing
  implementations for instantiation types with identical binary
  representations.

#### Item 45:   Use member function templates to generate functions that accept all compatible types.
* Let the arg of the member funciton take the other type as a template-arg
  and use normal copy/assignment (like initializer list), to ensure
  only acceptable conversions take place
* If you declare member templates for generalized copy construction
  or generalized assignment, you’ll still need to declare the normal
  copy constructor and copy assignment operator, too.



## More effective c++

### Basics

#### Item 1: Distinguish pointers and references
* References can never be null.
* References cannot re-point to sth else. Hence must be initialized at creation
* References are needed in operator overloading

#### Item 2: Casts
* `static_cast<>()` to be used in most places where c-style casts are needed.
* `const_cast<>()` to cast const away
* `dynamic_cast<()` to cast down a inheriticance
    * will return NULL for pointers
    * will throw exception for references!
* `reinterpret_cast<>()` to perform conversions, whose results are impl. defined
    * most common use-case is to cast one fn-ptr type to another

#### Item 3: Dont treat arrays polymorphically.
* Dont let a array of derived class decay into a base-class pointer
  which is used as array `*(ptr+n)`, as the size of derived class `>` base class

#### Item 4: Avoid gratuitous default constructors if that doesn't make sense.
* However, u cannot do the following if you dont have default constructor
    * Create arrays of the type
        * This can be overcome with allocating the
    * Perhaps not place this in container classes. (however std::vector doesn't
      have that limitation)
    * Another problem is these can't be virtual base classes.

### Operators

#### Item 5: Be wary of user-defined conversions
* implicit conversions are done by
    * single arg constructors
    * implicit type converstion functions   [ operator double() const; ]
        * they have no return type.
* consider writing explicit conversion fns with good names
   eg: `std::string<>::c_str()` , Rational::as_double()
* add explicit keyword to constructors to avoid implicit-type-conversions
    * explicit calling is allowed, static_cast<UrClass>(source_type) is
      also allowed, c-style explicit cast is also allowed

#### Item 6:  Prefix and postfix operators
* Remember to put const for post-fix return type.
* Prefix is cheaper as postfix involves creating a temporary
* Implement postfix in terms of prefix call so that the underlying
  increment activity is the same

#### Item 7:  Never overload &&, || and comma operators
* expr1 && expr2 becomes expr1.operator&&(expr2) for member fn or
    operator&&(expr1, expr2) for global.
    *  This doesn't offer short-circuit - beware!!

#### Item 8:  Understand the different meanings of new and delete
* new operator and delete operator are the normal externally visible
  operators. operator new is internal . See below for relation:
    `new operator`
    * calls operator new to get memory
    * calls constructor on the given memory
    * returns the pointer of corresponding type
* new-operator can't be overloaded. It always does the above 3
  steps. However, operator-new can be overloaded, so that gives
  us flexibility in getting the memory from wherever we need.
* We can't call constructor directly ever. But we have placement
  new that is same as calling constructor on a given memory.
     new(pointer) ClassName(construct_arg1, construct_arg2)
  The syntax (braces and ptr-arg) distinguish normal new from
  placement new.
* Dont overload operator-new globally. This will render ur app
  incompatible with libraries that decide to do same thing.
           (To read further: how to overload operator-new)

## Exceptions

#### Item 9:  Use destructors to prevent resource leaks

#### Item 10: Prevent resource leaks in constructors
* This can be aided by keeping resource holding pointers as
  resource-managing objects so that they are properly
  destructed.

#### Item 11: Prevent exceptions from leaving destructors
* Similar to Item-8 of book1

#### Item 12: Understand how throwing an exception differs from param-pass and virt-func call
* Exceptions are always thrown by a copy-by-value, even
  if catch clause catches it by ptr or reference.
  This is one reason why exceptions are slow
* Copy is done using object's static type (not run-time
  type). If u throw a base-class ptr or ref, what  gets
  thrown is a base-class object.
* throw;  // re-throws what was received.
  throw some_var_name; // will throw a copy of whatever
                          this name is referring.

# CPP Std Lib Book (Packt Pub)
Rainer Grimm (Purchased)

## Ch2: The standard library

### High level concepts

* Sequential Containers
    * std::vector         -- SHOULD BE FIRST CHOICE
    * std::array
    * std::deque
    * std::list
    * std::forward_list
* Associative Containers
    * std::set / set:: multiset
    * std::map / set:: multimap   -- MAP SHOULD BE FIRST CHOICE
    * unordered for all 4 above
* Container Adaptor
    * std::stack
    * std::queue
    * std::priority_queue
* Iterators
* Callables
    * functions
    * function-objects
    * lambda-functions
* Memory Model
* Thread
    * thread-local
* Conditional Variables
* Tasks

### Using

* using declaration: Just brings in one name
    ```
    using std::cout;
    ```
* using directive: Brings everything. Use with care.
    ```
    using std;
    ```
* namespace alias. Shoudn't hide a name though.
    ```
    namespace sysClock= std::chrono::system_clock;
    auto nowFirst= sysClock::now();
    ```
## Ch3: Utility Functions

* min, max, minmax

# Modern C++ Programming Cookbook
Marius Bancila (Packt Pub) (Safari Online)

## Ch1: Learning Modern Core Language Features

### Using auto whenever possible


* auto name = expression
    ```
     auto i = 42;          // int
     d = 42.5;             // double
     auto s = "text";      // char const *
     auto v = { 1, 2, 3 }; // std::initializer_list<int>
                           // Notice its initializer_list and not vector
    ```
* auto name = type-id { expression }
    * This form is like type-changing the literal to someothing
      other than what the literal/expr will get assigned to
    ```
    auto b  = new char[10]{ 0 };            // char*
    auto s1 = std::string {"text"};         // std::string
    auto v1 = std::vector<int> { 1, 2, 3 }; // std::vector<int>
    auto p  = std::make_shared<int>(42);    // std::shared_ptr<int>
    ```
* auto name = lambda-expression
* declare lambda parameters
* declare fn return type

* Benefits of Using auto
    * cant leave a variable uninitialized
    * no implicit conv like `int a=sizeof(..);` type implicit conv
    * less typing for things like iterator
    * consistent coding style.. type is always on right side.
* Gotchas
    * const/volatile is not specified by auto
    * reference type is not specified by auto
    * can't use auto for non-moveable types
    * not for multi-word types (long long, struct foo)

### Creating type aliases and alias templates

* typedef cannot be use for template typedef'ing
* using identifier = type-id
    ```
    using byte    = unsigned char;
    using pbyte   = unsigned char *;
    using array_t = int[10];
    using fn      = void(byte, double);
    ```
* template:
    ```
    template <class T>
    class custom_allocator { /* ... */};

    template <typename T>
    using vec_t = std::vector<T, custom_allocator<T>>;

    vec_t<int>           vi;
    vec_t<std::string>   vs;
    ```
* Dont mix typedef and using

### Understanding uniform initialization

* Brace-initialization is a uniform method for initializing data in C++11. For
  this reason, it is also called uniform initialization
* 2 types of initialization in C++
    ```
    T object {other};   // direct list initialization
    T object = {other}; // copy list initialization
    ```
    * uniform initialization works for both
* Be it:)b
    * standard containers
    * dynamically allocated arrays
    * statically definied arrays
    * built-in types
    * user-defined types
    * user-defined POD types
* Previously we could initialize by
    * direct assignment for built-in types
    * conversion constructor (that takes in one arg of a said type) via assignment.
    * plain initialization without args(def-constructor)
    * explicit constructor arg passing initialization
        * note that a non-zero arg initialization by parenthesis isn't possible.
          Its a fn-declaration.
          ```
            foo f1;           // default initialization
            foo f2(42, 1.2);
            foo f3(42);
            foo f4();         // function declaration
          ```
    * Aggregates could be initiazlied by brace-initialization
* Initialization of standard containers, such as the vector and the map also
  shown above, is possible because all standard containers have an additional
  constructor in C++11 that takes an argument of type std::initializer_list<T>
* The way the initialization using std::initializer_list works is the following:
    * The compiler resolves the types of the elements in the initialization list (all elements must have the same type).
    * The compiler creates an array with the elements in the initializer list.
    * The compiler creates an std::initializer_list<T> object to wrap the previously created array.
    * The std::initializer_list<T> object is passed as an argument to the constructor.
* An initializer list always takes precedence over other constructors where brace-initialization is used
  This gets precedence over other constructors

### Understanding the various forms of non-static member initialization

Use initializer list. You sometimes can't do it when:
* you want use this. Some comipers warn that this is uninitizlied at this point
* you want to refer another member from one member
* test an input, throw an exception and initiazlize after the test

Starting with C++11, non-static data members can be initialized when declared
in the class. This is called default member initialization because it is
supposed to represent initialization with default values. Default member
initialization is intended for constants and for members that are not
initialized based on constructor parameters

* Use default member initialization for providing default values for members of
  classes with multiple constructors that would use a common initializer for
  those members
* Use default member initialization for constants, both static and non-static
* Use the constructor initializer list to initialize members that don't have
  default values, but depend on constructor parameters
* Use assignment in constructors when the other options are not possible
  (examples include initializing data members with pointer this, checking
  constructor parameter values, and throwing exceptions prior to initializing
  members with those values or self-references of two non-static data members).

### Controlling and querying object alignment

2 keywords: alignas alignof

* To control the alignment of a type (both at the class level or data member
  level) or an object, use the alignas specifier:
    ```
    struct alignas(4) foo
    {
      char a;
      char b;
    };
    struct bar
    {
      alignas(2) char a;
      alignas(8) int  b;
    };
    alignas(8)   int a;
    alignas(256) long b[4];
    ```
* To query the alignment of a type, use the alignof operator:
    ```
    auto align = alignof(foo);
    ```

### Using scoped enumerations

* Old style enums are called unscoped.
* new, called scoped enumerations: enum class, enum struct
    * Both are the same.
* what they solve:
    * enumtype1::value1 should be used. Simply doing value1 isn't possible. So, 2 enums can have same value1
    * enums dont implicitly convert to int
    * You can forward-declare enums
      ```
      enum class Codes : unsigned int;
      void print_code(Codes const code) {} /* works! without knowing what all are under the enum */
      ```
### Using override and final for virtual methods

* Always use the virtual keyword when declaring virtual functions in derived
  classes that are supposed to override virtual functions from a base class,
* Always use the override special identifier after the declarator part of a
  virtual function declaration or definition
    * It shows the reader of the code that "this is a virtual method, that is
      overriding a virtual method of the base class0
    * The compiler also knows that it's an override, so it can "check" that you
      are not altering/adding new methods that you think are overrides:
* To ensure that functions cannot be overridden further or classes cannot be
  derived any more, use the final special identifier
    ```
    class Derived2 : public Derived1
    {
        virtual void foo() final {}
    };
    class Derived4 final : public Derived1
    {
        virtual void foo() override {}
    };
    ```
* It should be noted that both override and final keywords are special
  identifiers having a meaning only in a member function declaration or
  definition. They are not reserved keywords and can still be used elsewhere in
  a program as user-defined identifiers:

### Using range-based for loops to iterate on a range






# Questions to check

* c-style cast is this
    (T)expr
  fn-style cast is this
    T(expr)

  Is function style cast a c++ only thing? Is it just syntactic sugar for invoking constructor?

* Understand how virtual base classes work.

* How to throw exceptions of objects that dont permit copying (copy
  constructor is private)

* Wha happens when abstract classes are thrown?
