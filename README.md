## forego

[Foreman](https://github.com/ddollar/foreman) in Go.

This project was forked from [ddollar/forego](https://github.com/ddollar/forego).

It does not accept issues because its sole purpose is to have a working and maintained version of forego for [nginx-proxy/nginx-proxy](https://github.com/nginx-proxy/nginx-proxy).

##### Compile from Source

    $ go get -u github.com/nginx-proxy/forego

### Usage

    $ cat Procfile
    web: bin/web start -p $PORT
    worker: bin/worker queue=FOO

    $ forego start
    web    | listening on port 5000
    worker | listening to queue FOO

Use `forego help` to get a list of available commands, and `forego help
<command>` for more detailed help on a specific command.

### License

Apache 2.0 &copy; 2015 David Dollar
