# BashHereDoc

This project is a very simple and tiny library for parsing bash-like (bourne-again shell) heredoc syntax.
Heredoc (Bash Here Documents) is bash built-in functionality. You can read more about it [here](https://tldp.org/LDP/abs/html/here-docs.html)
Also this project includes a CLI executable _bhd_ to extract doc here contents.

##bhd
_bhd_ is a tool like a grep to extract bash here document contents

###Usage
```
Usage of ./bhd [file1, file2, ...]:
    -a    list all HereDocument contents
    -t string
        get here document by token. If token is not provided, the first heredoc occurrence is returned
        
    If no arguments are provided, standard input will be used instead
```

###How to build bhd
Just in the project directory type:
`$ make build`

Or directly with go:
`$ go build -o bhd cmd/main.go` 

