List of cpp stl containers
----------------------------

vector       ->   auto-growing array on one side.
deque        ->   auto-growing array on both sides
list         ->   linked list
set          ->   ordered keys
map          ->   ordered (key,value)s
pair         ->   holder for a pair
forward_list -> singly linked list

I/O streams
-----------

Open a file handling errors:
~~~~~~~~~~~~~~~~~~~~~~~~~~~

int main () {
  ifstream file;
  file.exceptions ( ifstream::failbit | ifstream::badbit );
  try {
    file.open ("test.txt");
    while (!file.eof()) file.get();
  }
  catch (const ifstream::failure& e) {
    cout << "Exception opening/reading file";
  }

  file.close();

  return 0;
}

Line based reading
~~~~~~~~~~~~~~~~~

#include <iftream>
#include <sstream>
#include <string>

ifstream infile('filename.txt');
std::string line;
while (std::getline(infile, line))
{
    std::istringstream iss(line);
    int a, b;
    if (!(iss >> a >> b)) { break; } // error

    // process pair (a,b)
}

Extracting tokens of a string
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

std::string s = "scott>=tiger>=mushroom";
std::string delimiter = ">=";

size_t pos = 0;
std::string token;
while ((pos = s.find(delimiter)) != std::string::npos) {
    token = s.substr(0, pos);
    std::cout << token << std::endl;
    s.erase(0, pos + delimiter.length());
}
std::cout << s << std::endl;


Printing format control
~~~~~~~~~~~~~~~~~~~~~~~

oldwidth=cout.width(newWidth);
/* or */
cout<<std::setw(newWidth)<<i;

Detailed
--------

Commonish-stuff
~~~~~~~~~~~~~~~

container.size()    /* gets len of vector deque, list, set, map */
string.length()     /* string along has a length! Not sure of other container, but this is O(1). They are synonyms */

Iterating through a container
+++++++++++++++++++++++++

C++11
for (auto it = begin (vector); it != end (vector); ++it) {
      it->doSomething ();
}
/* not sure where it works */
for (auto & element : vector) {
      element.doSomething ();
}

//If you need both index and value.
for (int i = 0; i < v.size(); i++) {
    // .. use v[i]
    // .. You can also use v[i-1], v[i+1] with careful checks too.
}

Max_elem
++++++++

max = *max_element(v.begin(), v.end())
min = *min_element(v.begin(), v.end())

swap
++++

swap(ref1,ref2)

Reversing a container:
++++++++++++++++++++++

/* the container in place */
reverse(v.begin(), v.end());

Removing from a vector as you iterate
++++++++++++++++++++++++++++++++++++++

//see if remove_if cuts for you
bool unary_predicate(val_type);
v.erase (remove_if(v.begin(), v.end(), unary_predicate),
         v.end());

Sort
++++

#include <algorithm>
/* sorts the container in place */
sort(v.begin(), v.end())
sort(numbers.begin(), numbers.end(), std::greater<int>());

bool myCustomFn(TypeOfObject &lhs, TypeOfObject &rhs);
sort(numbers.begin(), numbers.end(), myCustomFn);

Rotate
+++++++

/* left rotate by n */
rotate(v.begin(),v.begin()+n,v.end())

Get typenames
+++++++++++++

Container::value_type   /* gets value-type for vector, set, deque, list, set
                           gets the pair-type for map */
Container::iterator     /* iterator type */

Accumulate
+++++++++++

std::vector<int> v{1, 2, 3, 4, 5, 6, 7, 8, 9, 10};

int sum = std::accumulate(v.begin(), v.end(), 0);  /* (begin,end,init-value,operator-fn) */

int product = std::accumulate(v.begin(), v.end(), 1, std::multiplies<int>());

std::string s = std::accumulate(std::next(v.begin()), v.end(),
                                    std::to_string(v[0]), // start with first element
                                    [](std::string a, int b) {
                                        return a + '-' + std::to_string(b);
                                    });

upper-bound/lower-bound
+++++++++++++
get-next/higher/value

/* returns iter to first element, that is more than a */
upper_bound(v.begin(), v.end(), a);

/* returns iter to first element, that is more or equal to a */
lower_bound(v.begin(), v.end(), a);

fill/populate/generate/assign
+++++++++++++++++++++++++++++

#include<numeric>
std::vector<int> v(10);
std::iota(v.begin(), v.end(), 0);  // add value++ to every element.

std::vector<int> v2(10);
int i = 0;
std::generate(v2.begin(), v2.end(), [&i](){return i++; });

Some common member functions
++++++++++++++++++

v.clear()        /* Clear the container */
v.front()        /* get value of front element  -- undefined behavior on empty containers */
v.back()         /* get value of last element -- undefined behavior on empty containers  */


char stuff
-----------

tolower  // Convert uppercase letter to lowercase (function )
toupper

character classification routines
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

isalnum, isalpha, isascii, isblank,
iscntrl, isdigit, isgraph, islower,
isprint, ispunct, isspace, isupper,
isxdigit


strings
--------

str.length()
str += "string to append";
str.substr(begin_pos, count=<>);  /* if count is skipped, the size() is assumed) */

str[i]   /* gets char at position */

constructing:
string();                                                  // default. empty string
string (const string& str);                                // copy construct.
string (const string& str, size_t pos, size_t len = npos); // create a substring from another str
string (const char* s);                                    // create from c-string
string (const char* s, size_t n);                          // create from c-string restricted to size
string (size_t n, char c);                                 // create prefilled of size
template <class InputIterator>                             // create from another range
string  (InputIterator first, InputIterator last);

/* assign can be used to create a string with new values */
string &assign (....);  // all construct options

npos
~~~~

static const to let know the max possible value. (used in string function to signal end of string)


trim a string
~~~~~~~~~~~~

s.erase(s.find_last_not_of(" \n\r\t")+1);

string trim(const string& str)
{
    size_t first = str.find_first_not_of(' ');
    if (string::npos == first)
    {
        return "";
    }
    size_t last = str.find_last_not_of(' ');
    return str.substr(first, (last - first + 1));
}

Math
------

#include <cstdlib>
int abs(int);
long abs(long);
long long abs(long long);

int max(int, int);

int min(min, min);


pair
~~~~

p.first()    /* gets reference to first */
p.second()      /* gets reference to second */

make_pair(v1,v2)  /* create a pair from 2 variables */

vector
~~~~~~

vector<Type> v(100, initValue); /* intializes v with a size of 100 and each element to initValue(optional) */
v.reserve(100);                 /* Grabs space for 100 w/0 having to reallocate */

v.push_back(value);  /* adds at end */
v.pop_back(value);   /* removes from end */

/* int fun_taking_ref_to_type(Type &t) */
for_each (v.begin(),v.end(), fun_taking_ref_to_type);

/* initializing 2D vector */
vector<vector<int> > 2dv(M, vector<int>(N,-1));


/* clone/copy a vector - simple use copy constructor or assignment operator */
vector<int> aCopy(orig);
aCopy = anotherOrig;

/* erase the 3rd element. Note the +(n-1) */
vector.erase(begin()+2);

/* append one vector to another */
target.insert(target.end(), other.begin(), other.end());

Static define a vector
+++++++++++++++++++++

const vector<int> primes({2, 3, 5, 7, 11});


map
~~~

v = m[key]   /* to access the value at and to assign */
             /* if key doesn't exist, a default value is created!
              *   Thus m[key]++ will auto-matically either add value or set to 1 resp if key already exist/not */
m[key] = v;

User the [] if you want to insert/assign w/o caring if its already there. Note that the default value gets created if
it doesn't exist. You can use the m.insert() if you want to check for exclusive-insert
pair<it,bool> = m.insert(pair<K,V> kv)


//key exists in map
m.find(key) == m.end()

removed = m.erase(key)   /* returns 0/1 (erase-cnt) if key was deleted */
m.erase(iterator);
erase(begin, end);

map's iterator is a iterator to a pair. it->first is the key, it->second is the value

unordered_map
~~~~~~~~~~~~~

unordered_map<key_type,value_type,hash_functor,equality_functor>

#include<functional>
std::hash<int>()  // int functor.

set
~~~

//initialize
struct CompareStructure
{
  bool operator() (const Type &lhs, const Type &rhs) const {
    //whatever..
  }
};
set<Type,CompareStructure> s;

or

set<Type,bool(*)(const Type &,const Type &)> s(&fn_name);

pair<it, bool> r = s.insert(val);
0/1 = s.erase(val);

s.find(val) == s.end()


multimap
~~~~~~~~~

m.insert(pair<key,val>(v));

m.count(key_val)  /* returns the number of values against this key */

pair<Iter, Iter> range = mm.equal_range("Group1");
int total = accumulate(range.first, range.second, 0);

list
~~~~~

l.remove(val)   /* O(n) removes elements whose value matches */
l.erase(iterator)  /* O(1) .. removes specified element */
l.insert(iterator, value) /* inserts at iterator, whetever is prevly pointed by interator comes after current insertion */

bitset
~~~~~~

* biggest bummer is the size of bit set is compile time defined
* bitset doesn't have iterators.

bitset<SIZE>  bs;


permutation
~~~~~~~~~~~

#include<algorithm>

// Rearranges v into the next lexicographical permutation. Returns 0 if its already greater
bool next_permutation(v.begin(),v.end())

There is no inherent next_combination, but we can use a bool-set to get next_combination
See: http://stackoverflow.com/a/9430993/2587153


Remove while iterating
---------------------

for (auto it = m.cbegin(); it != m.cend() /* not hoisted */; /* no increment */)
{
  if (must_delete)
  {
    m.erase(it++);    // or "it = m.erase(it)" since C++11
  }
  else
  {
    ++it;
  }
}


https://www.think-cell.com/en/career/talks/pdf/think-cell_talk_memorymodel.pdf
