package kysscript

import "fmt"

/* this file contains definitions like a debugging print statement
 * idk what to comment hers lo
 */

 const DebugMode bool = true
 func lexer_dev_print(msg string) {
	 if !DebugMode {
		 return
	 }
		
	 fmt.Println(msg)
 }

 func parser_debug_print(msg string) {
	 if !DebugMode {
		 return
	 }
		
	 fmt.Println(msg)
 }
