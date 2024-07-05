package lexer

import "fmt"

/*
 * this file contains definitions like a debugging print statement
 * idk what to comment hers lo
 */

const lexerDebugMode bool = true
func dev_print(msg string) {
    if !lexerDebugMode {
        return
    }

    fmt.Println(msg)
}
