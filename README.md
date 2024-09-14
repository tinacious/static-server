# static-server

`static-server` is a simple command line app written in Go that serves your current working directory as a static server.

- [Installation](#installation)
- [Usage](#usage)
  - [Custom port](#custom-port)
- [Uninstall](#uninstall)


## Installation

Install the CLI tool:

```sh
go install github.com/tinacious/static-server@latest
```

## Usage

Run the `static-server` command in any directory you'd like to serve as a web server.

If you have the following directory `test`:

```
test
├── index.html
└── style.css
```

You can do:

```sh
cd test
static-server
```

This should serve the contents on a random port. 


### Custom port

You can configure the port by passing it as an environment variable. For example, to run on port 1337 run:

```sh
PORT=1337 static-server
```


## Uninstall

Open `~/go/bin` and delete `static-server`.
