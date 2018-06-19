uselessd
========

A useless server


## Why??

A common use case is to create a reusable library and an application that consumes it, and host both on GitHub. The application called ["uselessd"](https://github.com/helmedeiros/uselessd) will consumes a this trivial library called ["useless"](https://github.com/helmedeiros/useless).


## Code Layout
The app and both libraries live on GitHub, each in its own repository. `$GOPATH` is the root of the project - each of your GitHub repos will be checked out several folders below `$GOPATH`.

Your code layout would look like this:
```
$GOPATH/
    src/
        github.com/
            helmedeiros/
                useless/
                    .git/
                    useless.go
                    useless_test.go
                    README.md
                uselessd/
                    .git/
                    uselessd.go
                    uselessd_test.go
                    README.md
```
Each folder under `src/github.com/helmedeiros/` is the root of a separate git checkout.

## Setup the Workspace
Let's assume we are starting from scratch. Initialize the two new repositories on GitHub, using the "Initialize this repository with a README" option so your repos can be cloned immediately. Then setup the project like this:

```
export GOPATH=~/go # GOPATH = Go workspace
cd $GOPATH
mkdir -p src/github.com/helmedeiros
cd src/github.com/helmedeiros
git clone git@github.com:helmedeiros/useless.git
git clone git@github.com:helmedeiros/uselessd.git
```

## Applications
An application - Go code that will be compiled into an executable command - always defines package main with a `main() function`.

So `uselessd.go` looks like this:

```
package main

import (
	"net/http"

	"golang.org/x/net/websocket"
	"github.com/helmedeiros/useless"
)

func main() {
	http.Handle("/useless", websocket.Handler(func(ws *websocket.Conn) {
		ws.Write([]byte(useless.Gorillaz()))
	}))
	http.ListenAndServe(":3000", nil)
}

```

## Dependencies
Your project will probably depend on some existing packages. The application above depends upon `golang.org/x/net/websocket`. You can install all dependencies by running `"go get -v ./..."` from the root of your workspace. The `"go get"` command is similar to `"go install"` in that it will attempt to build and install all packages in the workspace (technically, all packages matched by `"./..."`), except that it will also examine their dependencies and download (and install) any that are missing first.

## Build
During development you can build the `useless` library by itself with the command `"go build ...useless"`. You could also give the full path to the package name, `"go build github.com/helmedeiros/useless"`.

To compile `uselessd.go` and its dependencies into an executable, use the command `"go build ...uselessd"`.
