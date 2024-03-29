# Markdown Notes

[[_TOC_]]

# Links for reference

https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet

# Headings

```
# H1
## H2
### H3
#### H4
##### H5
###### H6

Alt-H1
======

Alt-H2
------
```

# Text stuff

* Emphasis, aka italics
    * surround with `*asterisks*` or `_underscores_`.

* Strong emphasis, aka bold
    * surrenout with double `**asterisks**` or `__underscores__`.

* Combined emphasis
    * surround with both `**asterisks and _underscores_**`.

* Strikethrough uses two tildes. ~~Scratch this.~~

---
**NOTE**

It works with almost all markdown flavours (the below blank line matters).

---


## Code

* `monospace` or `code`
    * Have backticks around it.

* Blocks with language mentioned
    * Have 3 backtick blocks

```javascript
var s = "JavaScript syntax highlighting";
alert(s);
```

```python
s = "Python syntax highlighting"
print s
```

```
No language indicated, so no syntax highlighting.
But let's throw in a <b>tag</b>.
```
## BlockQuotes

> Blockquotes are very handy in email to emulate reply text.
> This line is part of the same quote.

Quote break.

> This is a very long line that will still be quoted properly when it wraps. Oh
> boy let's keep writing to make sure this is long enough to actually wrap for
> everyone. Oh, you can *put* **Markdown** into a blockquote.

## RAW HTML

Just type is as is. A Definition list is most useful.

<dl>
  <dt>Definition list</dt>
  <dd>Is something people use sometimes.</dd>

  <dt>Markdown in HTML</dt>
  <dd>Does *not* work **very** well. Use HTML <em>tags</em>.</dd>
</dl>

## Comments:

[//]: # (This is a comment, it will not be included. Note everything)
[//]: # (the colon, space, hash, paren, and comments are only one line)
[//]: # (long)

# Lists

1. First ordered list item
2. Another item
  * Unordered sub-list.
1. Actual numbers don't matter, just that it's a number
  1. Ordered sub-list
4. And another item.

   You can have properly indented paragraphs within list items. Notice the
   blank line above, and the leading spaces (at least one, but we'll use three
   here to also align the raw Markdown).

   To have a line break without a paragraph, you will need to use two trailing
   spaces.  
   Note that this line is separate, but within the same paragraph.  
   (This is contrary to the typical GFM line break behaviour, where trailing
   spaces are not required.)

* Unordered list can use asterisks
    * lists under lists just indent below. No double asteriks
- Or minuses
+ Or pluses

# Links

[I'm an inline-style link](https://www.google.com)

[I'm an inline-style link with title](https://www.google.com "Google's Homepage")

[I'm a reference-style link][Arbitrary case-insensitive reference text]

[I'm a relative reference to a repository file](../blob/master/LICENSE)

[You can use numbers for reference-style link definitions][1]

Or leave it empty and use the [link text itself].

URLs and URLs in angle brackets will automatically get turned into links.
http://www.example.com or <http://www.example.com> and sometimes
example.com (but not on Github, for example).

Some text to show that the reference links can follow later.

[arbitrary case-insensitive reference text]: https://www.mozilla.org
[1]: http://slashdot.org
[link text itself]: http://www.reddit.com

# Images

Here's our logo (hover to see the title text):

Inline-style: 
![alt text](https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Logo Title Text 1")

Reference-style: 
![alt text][logo]

[logo]: https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Logo Title Text 2"


# Tables

Colons can be used to align columns.

| Tables        | Are           | Cool  |
| ------------- |:-------------:| -----:|
| col 3 is      | right-aligned | $1600 |
| col 2 is      | centered      |   $12 |
| zebra stripes | are neat      |    $1 |

There must be at least 3 dashes separating each header cell.
The outer pipes (|) are optional, and you don't need to make the 
raw Markdown line up prettily. You can also use inline Markdown.

Markdown | Less | Pretty
--- | --- | ---
*Still* | `renders` | **nicely**
1 | 2 | 3

# text blocks

---
**NOTE**

Should render in most markdown flavors

---


# Rules

3 or more hyphens/asteriks/underscore


# Table of contents

* Add this `[[_TOC_]]` at the top after the first heading
