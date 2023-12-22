# asn

Links:
* https://obj-sys.com/asn1tutorial/node2.html
* https://www.oss.com/asn1/resources/asn1-made-simple/introduction.html
    * This one is really good.

Online decoders:
https://www.marben-products.com/cgi-bin/asn1tools/free-online-asn1-decoder.pl

* ASN.1 (Abstract Syntax Notation One).
* The Basic Encoding Rules (BER),

# Standards


ITU-T X.680	ISO/IEC 8824-1
ITU-T X.681	ISO/IEC 8824-2
ITU-T X.682	ISO/IEC 8824-3
ITU-T X.683	ISO/IEC 8824-4
ITU-T X.690	ISO/IEC 8825-1
ITU-T X.691	ISO/IEC 8825-2

# Module

* The fundamental unit of ASN.1 is the module

```asn
InventoryList {1 2 0 0 6 1} DEFINITIONS ::=  -- ModuleReference: InventoryList, Object Identifier, DEFINITIONS is a keyword
  BEGIN
    {
      ItemId ::= SEQUENCE                    -- this is a type assignment. Begins with a Capital letter
      {
        partnumber IA5String,                -- componenet of a sequeunce. identifiers. Begin with a lower case
        quantity INTEGER,
        wholesaleprice REAL,
        saleprice REAL
      }
      StoreLocation ::= ENUMERATED
      {
        Baltimore (0),
        Philadelphia (1),
        Washington (2)
      }
    }
  END

-- Another one
MyShopPurchaseOrders DEFINITIONS AUTOMATIC TAGS ::= BEGIN

    PurchaseOrder ::= SEQUENCE {
        dateOfOrder DATE,
        customer    CustomerInfo,
        items       ListOfItems
    }

    CustomerInfo ::= SEQUENCE {
        companyName    VisibleString (SIZE (3..50)),      -- constratints
        billingAddress Address,
        contactPhone   NumericString (SIZE (7..12))
    }

    Address::= SEQUENCE {
        street  VisibleString (SIZE (5 .. 50)) OPTIONAL,
        city    VisibleString (SIZE (2..30)),
        state   VisibleString (SIZE(2) ^ FROM ("A".."Z")),
        zipCode NumericString (SIZE(5 | 9))
    }

    ListOfItems ::= SEQUENCE (SIZE (1..100)) OF Item

    Item ::= SEQUENCE {
        itemCode        INTEGER (1..99999),
        color           VisibleString ("Black" | "Blue" | "Brown"),
        power           INTEGER (110 | 220),
        deliveryTime    INTEGER (8..12 | 14..19),
        quantity        INTEGER (1..1000),
        unitPrice       REAL (1.00 .. 9999.00),
        isTaxable       BOOLEAN
    }
END

```

* Type Assignment
    * name of type, symbol - `::=`(read, is defined as), and the type

* value assignment
    * value reference always begins with lowercase
    ```asn
    gadget  ItemId  ::=             -- name of value, type of value and ::=
    {
        partnumber      "7685B2",
        quantity        73,
        wholesaleprice  13.50,
        saleprice       24.95
    )
    ```

# Asn syntax

* Begin with 2 hypens `--`
* names cant have `: ; = , < . ( ) [ ] ' "`

# Built-in types

* simple, structured (futher has simple/structured)
* user-defined types - made of simple/structured.
* ASN.1 also has another category of types called useful, which provide
  standard definitions for a small number of commonly used types.

## Simple types

BOOLEAN           | 1  |  Model logical, two-state variable values
INTEGER           | 2  |  Model integer variable values
BIT STRING        | 3  |  Model binary data of arbitrary length
OCTET STRING      | 4  |  Model binary data whose length is a multiple of eight
NULL              | 5  |  Indicate effective absence of a sequence element
OBJECT IDENTIFIER | 6  |  Name information objects
REAL              | 9  |  Model real variable values
ENUMERATED        | 10 |  Model values of variables with at least three states
CHARACTER STRING  | *  |  Models values that are strings of characters from a specified characterset

* Type BOOLEAN takes values TRUE and FALSE. Usually, the type reference for
  BOOLEAN describes the true state. For example: `Female ::= BOOLEAN` is
  preferable to `Gender ::= BOOLEAN`.

* Type INTEGER takes any of the infinite set of integer values. Its syntax is
  similar to programming languages such as C or Pascal. It has an additional
  notation that names some of the possible values of the integer. For example,

  ```asn
   ColorType ::= INTEGER
       {
          red      (0)
          white    (1)
          blue     (2)
       }
  ```
   indicates that the `ColorType` is an INTEGER and its values 0, 1, and 2 are
   named red, white, and blue, respectively. The ColorType COULD ALSO HAVE ANY
   of the other valid integer values, such as 4 or -62.

See https://obj-sys.com/asn1tutorial/node10.html

## Structed types

Structured Types  |  Tag |  Typical Use
SEQUENCE          |  16  |  Models an ordered collection of variables of different type
SEQUENCE OF       |  16  |  Models an ordered collection of variables of the same type
SET               |  17  |  Model an unordered collection of variables of different types
SET OF            |  17  |  Model an unordered collection of variables of the same type
CHOICE            |  *   |  Specify a collection of distinct types from which to choose one type
SELECTION         |  *   |  Select a component type from a specified CHOICE type
ANY               |  *   |  Enable an application to specify the type
Note: ANY is a deprecated ASN.1 Structured Type. It has been replaced with X.680 Open Type.

```asn
  AirlineFlight  ::=  SEQUENCE
     {
      airline   IA5String,
      flight    NumericString,
      seats     SEQUENCE
                    {
                     maximum   INTEGER,
                     occupied  INTEGER,
                     vacant    INTEGER
                    },
      airport   SEQUENCE
                    {
                     origin             IA5String,
                     stop1       [0]    IA5String  OPTIONAL,  -- the [0] is necessary as consequetive optional components are of same type
                     stop2       [1]    IA5String  OPTIONAL,
                     destination        IA5String
                    },
       crewsize ENUMERATED
                    {
                      six    (6),
                      eight  (8),
                      ten    (10)
                    },
       cancel   BOOLEAN    DEFAULT FALSE  -- DEFAULT is like optional. Not sent unless value is TRUE
      }.

```

# constrints

* Single value constraint

    ```asn
    color VisibleString ("Black" | "Blue" | "Brown")
    power INTEGER (110 | 220)
    ```
* restricting values to a particular range in value and alphabet range constraint
    ```asn
    quantity INTEGER (1..1000)
    deliveryTime INTEGER (8..12 | 14..19)
    unitPrice REAL (1.00..9999.00)
    State ::= VisibleString SIZE(2) ^ FROM ("A".."Z"))     -- note ^ is intersection(or logical-and)
    ```
* restricting the length of values in size constraint:
    ```asn
    contactPhone NumericString (SIZE (7..12))
    ListOfItems ::= SEQUENCE (SIZE (1..100)) OF Item
    zipCode NumericString (SIZE (5 | 9))
    ```


