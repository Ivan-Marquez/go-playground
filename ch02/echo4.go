package ch02

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

// Echo4 function variation on the earlier echo command takes
// two optional flags: -n causes echo to omit the trailing
// newline that would normally be printed, and -s sep causes
// it to separate the output arguments by the contents of
// the string sep instead of the default single space.
func Echo4() {
	flag.Parse()

	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}

/*
 * The function flag.Bool creates a new flag variable of type bool.
 * It takes three arguments: the name of the flag ("n"), the variable’s
 * default value (false), and a message that will be printed if the user
 * provides an invalid argument, an invalid flag, or -h or -help.
 * Similarly, flag.String takes a name, a default value, and a message,
 * and creates a string variable.
 *
 * The variables sep and n are pointers to the flag variables, which must
 * be accessed indirectly as *sep and *n. When the program is run, it must
 * call flag.Parse before the flags are used, to update the flag variables
 * from their default values. The non-flag arguments are available from
 * flag.Args() as a slice of strings. If flag.Parse encounters an error,
 * it prints a usage message and calls os.Exit(2) to terminate the program.
 */
