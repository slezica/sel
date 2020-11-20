# sel

**NOTE**: I'm the process of replacing https://github.com/slezica/sel-python-old with this Golang implementation, but
it's still buggy and incomplete. Go there for a more functional tool.

`sel` is an inline field selection and table transformation tool
that aims to replace `cut`.


## Installation

Just copy the right executable from the `bin` directory into your path.


## Usage

The `sel` tool takes each line from `stdin`, splits it into fields, and prints back a selection of those fields.

The simplest use is selecting a single field from a line of columns:

    $ echo a b c d e | sel 3
    c

Field indices begin at `1`.


#### Selecting multiple fields

You can select multiple fields by passing a list of selectors: 

    $ echo a b c d e | sel 3 5
    c e

The selectors don't have to be in order. You can use `sel` to rearrange columns:

    $ echo a b c d e | sel 4 1
    d a


#### Selecting fields from the end

You can use negative indices to select fields counting from the end:

    $ echo a b c d e | sel -1 -3
    e c


#### Selecting field range

`sel` understands field ranges, in python style:

    $ echo "a b c d e" | sel 2:4
    b c d

    $ echo "a b c d e" | sel 2:-2
    b c d

This includes unbounded ranges:

    $ echo "a b c d e" | sel 3:
    c d e

    $ echo "a b c d e" | sel :2
    a b

    $ echo "a b c d e" | sel :
    a b c d e


#### Splitting fields using a custom separator

By default, `sel` splits the input on whitespace. It can also use a custom regular expression:

    $ cat users.csv
    1241,Bob
    3192,MitM
    3255,Alice

    $ cat users.csv | sel --split , 1
    1241
    3192
    3255
    
    $ echo 1a2b3c4d | sel --split [a-z] 2:3
    2 3


#### Joining fields using a custom string

By default, `sel` prints selected fields separated by a single space character. You can change this:

    $ echo a b c | sel --join - 2 3
    b-c