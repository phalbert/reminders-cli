# Reminders

## CLI SWITCH COMMANDS

```bash
./bin/client create

./bin/client edit

./bin/client fetch

./bin/client delete

./bin/client health
```

### What does the command look like?

```bash
./bin/client delete --id=3
```

1. cmd -> `./bin/client`
2. sub cmd -> `delete` 
3. args -> `--id=3`

Since we are using golang, we will utilise `os.Args` which is a slice which gets populated whenever we run a go program. 

> os.Args[0] is the path to the program, and os.Args[1:] holds the arguments to the program. 

```golang
package main
import (
    "fmt"
    "os"
)
func main() {
    // os.Args provides access to raw command-line arguments.
    // Note that the first value in this slice is the path to the program, 
    // and os.Args[1:] holds the arguments to the program.

    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    // You can get individual args with normal indexing.
    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}
```

To pass arguments, we can use one of the following forms

```bash
--id=3

--id 3

 -id=3

 -id 3
```
