pwpls
===================

pwpls is a simple password generator.

It allows you to create passwords that conform to specific guidelines.

Some password guidelines require _n_ special characters,
_x_ uppercase letters, a minimum length of _y_, and so on.

pwpls allows you to conform to these guidelines in a dead-simple way.

## Usage

```shell
# Password with a length of 20, 3 special characters, and 5 uppercase
# letters. The shorthand flag notation is the same as
# pwpls --length=20 --special=3 --uppercase=5

eric@archbox ~/ $ pwpls -l 20 -s 3 -u 5
```

## Notes

It'll return an error message if you try to do something impossible. For
example, asking pwpls for a password with a length of 5 that contains
10 special characters will look like this:

```shell
eric@archbox ~/ $ pwpls -l 5 -s 20
pwpls: special + uppercase + digits should be <= length
```

## How It Works (Crypto Notes)

By default it uses the OS' PRNG to scoop up _n_ bytes, where _n_ is the length
of password you want.

It then encodes the random bytes using hexidecimal encoding (to convert the
random bytes to printable characters). It then randomly truncates the encoded
buffer to the desired length (hexidecimal encodes twice as large as its input)
and then randomly loops over the truncated buffer, swapping out characers
until the password conforms to the provided specifications.

It randomly loops by using the xorshift 4096 algorithm (seeded with a random
64-bit prime) to generate a random number _x_ and using _x_ % len(_buffer_)
as the position of the character to replace.

The tables of characters used to replace other characters are defined as

```go
var specialTable = table{
	'~', '`', '!', '@',
	'#', '$', '%', '^',
	'&', '*', '(', ')',
	'_', '+', '[', ']',
	'{', '}', '|', '\\',
	'"', '\'', ';', ':',
	'.', '>', ',', '<',
	'/', '?',
}

var upperTable = table{
	'A', 'B', 'C', 'D',
	'E', 'F', 'G', 'H',
	'I', 'J', 'K', 'L',
	'M', 'N', 'O', 'P',
	'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X',
	'Y', 'Z',
}

var lowerTable = table{
	'a', 'b', 'c', 'd',
	'e', 'f', 'g', 'h',
	'i', 'j', 'k', 'l',
	'm', 'n', 'o', 'p',
	'q', 'r', 's', 't',
	'u', 'v', 'q', 'x',
	'y', 'z',
}

var digitTable = table{
	'0', '1', '2', '3',
	'4', '5', '6', '7',
	'8', '9',
}
```

## License

The Unilicense: Public domain.
