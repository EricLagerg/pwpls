package main

type table []byte

func (t table) in(b byte) bool { return t.find(b) > -1 }
func (t table) get() byte      { return t[next(len(t))] }

func (t table) remove(b byte) {
	if i := t.find(b); i > -1 {
		t[i] = '\x00'
	}
}

func (t table) find(b byte) int {
	for i := range t {
		if t[i] == b {
			return i
		}
	}
	return -1
}

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
