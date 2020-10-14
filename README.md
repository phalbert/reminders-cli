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

To pass arguments, we can use one of the following forms

```bash
--id=3

--id 3

 -id=3

 -id 3
```

Ways of parsing commands

1. [Using `os.Args`](docs/os-args.md)
2. [Using flag sets](docs/flag-set.md)