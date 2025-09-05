Description
===========
digital_museum is an app that shows you a random art piece from the Metropolitan Mueseum of Arts
It uses [The Metropolitan Museum of Art Collection API](https://metmuseum.github.io/) to request a random art piece and parses the description of the requested art piece to display them.

Usage
=====
1. Clone this repo
``` bash
$ git clone https://github.com/snufflo/digital_museum.git
```

2. In the cloned directory, run `main.go` with
``` bash
$ go run main.go
```

- if it fails, you might need to run the following command beforehand:
``` bash
$ go install && go mod tidy
```
