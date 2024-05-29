package main

import (
	"fmt"
	"kogab-interpreter/internal/lexer"
	"kogab-interpreter/internal/token"
)

func main() {
	var input string = `
let five = 5;
let ten = 10;
let add = fun(x, y) {
	x + y;
};

let result = add(five, ten);

if (result != 15) {
	return false;
} else {
	result = result / 2;
	result = result * 3;
	return true;
}

10 == 10;
10 != 9;
10 < 9;
10 > 9;
10 <= 10;
10 >= 10;

let reversedBool = !true;
let anotherBool = !false;
`

	var l = lexer.New(input)

	for {
		var tok = l.NextToken()
		fmt.Printf("%+v\n", tok)

		if tok.Type == token.EOF {
			break
		}
	}
}
