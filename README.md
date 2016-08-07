## go-shorten

Running a URL redirection service doesn't need to be difficult. go-shorten takes the pain out of setting one up and maintaining it. Just edit the JSON configuration file or use the built-in web interface to manage the redirects.

### Building go-shorten

Building the application couldn't be easier. Assuming you have Go installed, just run:

    go get github.com/nathan-osman/go-shorten

That's it. The executable will be built and installed to `$GOPATH/bin`. There are no runtime dependencies.

### Launching go-shorten

The first time you lauch go-shorten, a configuration file will be created in the current directory (named `config.json`). This file is used to customize the behavior of the application. To try the program out locally, change the value of "addr" to ":8000".

Now start the application with a single argument - the config file:

    ./go-shorten config.json

Now open a browser to [http://localhost:8000](http://localhost:8000) and enter "admin" for the username and "passw0rd" for the password when prompted.

**Be sure to change the password in the config file if you use this application in production.**

### Running in Docker

go-shorten can be easily run with Docker.

[TODO]
