package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	flag "github.com/EricLagerg/pflag"
)

const (
	version = `pwpls 1.0
Copyright (C) 2015 Eric Lagergren
License: MIT <https://opensource.org/licenses/MIT>
This is free software.
There is NO WARRANTY, to the extent permitted by law.
`

	help = `Usage: wc [option]...

Generate a password using specific criteria including required
characters (2 uppercase, 1 special character, etc.), password
creation algorithm, and more.

  -c, --common		reject "common" passwords (default = true)
  -d, --digits		number of digits (default = random)
  -e, --exclude		exclude specific special characters (default = "")
  -i, --include-special	include special characters (random)
  -l, --length		required password length (default = 8)
  -w, --lower		number of lowercase characters (default = random)
  -n, --number		number of passwords to print
  -s, --special		number of special characters (default = 0)
  -t, --type 		type of algorithm to use (default = OS' PRNG)
            		  other types include (case independent):
            		  * xorshift [xs] (4096)
            		  * mersenne_twister [mt]
  -u, --uppercase	number of uppercase characters (default = random)

Report pwpls bugs to ericscottlagergren@gmail.com
pwpls home page: <https://github.com/EricLagerg/pwpls>
`
)

var (
	alg            = flag.StringP("alg", "a", "random", "")
	common         = flag.BoolP("common", "c", true, "")
	digits         = flag.IntP("digits", "d", 0, "")
	exclude        = flag.StringP("exclude", "e", "", "")
	includeSpecial = flag.BoolP("include-special", "i", false, "")
	length         = flag.IntP("length", "l", 8, "")
	lower          = flag.IntP("lower", "w", 0, "")
	number         = flag.IntP("number", "n", 1, "")
	special        = flag.IntP("special", "s", -1, "")
	upper          = flag.IntP("uppercase", "u", -1, "")

	vers = flag.BoolP("version", "v", false, "")

	logger = log.New(os.Stderr, "pwpls: ", 0)
)

type AlgFn func() string

var knownAlgorithms = map[string]AlgFn{
	"random":           randAlg,
	"xorshift":         xorshiftAlg,
	"xs":               xorshiftAlg,
	"mersenne_twister": mersenneAlg,
	"mt":               mersenneAlg,
}

func exit(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
	os.Stderr.WriteString("\n")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s", help)
		os.Exit(1)
	}
	flag.Parse()

	if *vers {
		fmt.Printf("%s", version)
		return
	}

	if *includeSpecial && *special > -1 {
		exit("Cannot use --include-special/-i and --special/-s at the same time")
	}

	if fix(*special) < 0 ||
		fix(*upper) < 0 ||
		*digits < 0 ||
		*lower < 0 ||
		*length < 0 {
		exit("special, uppercase, digits, and length must be positive integers")
	}

	if *special+*upper+*digits+*lower > *length {
		exit("special + uppercase + digits should be <= length")
	}

	if *exclude != "" {
		for i := range *exclude {
			specialTable.remove((*exclude)[i])
		}
	}

	if *includeSpecial {
		*special = int(next(*length - (*upper + *digits + *lower)))
	}

	if *upper < 0 {
		n := *length - (*special + *digits + *lower)
		if n > 0 {
			*upper = int(next(n))
		}
	}

	fn := randAlg
	ok := false

	if *alg != "" {
		if fn, ok = knownAlgorithms[strings.ToLower(*alg)]; !ok {
			exit("unknown algorithm %q", *alg)
		}
	}

	for i := *number; i > 0; i-- {
		fmt.Println(fn())
	}
}

func fix(x int) int {
	if x < 0 {
		return 0
	}
	return x
}
