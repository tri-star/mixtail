mixtail
===================================

![release](http://img.shields.io/github/release/tri-star/mixtail.svg?style=flat-square)
![license](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)

# Overview
Mixtail watches multiple log files( or command output) and output them in single console.

## Description
This command is useful for watching file(or command output) across many servers.

The command takes a YAML file like following. 

## Usage
At first, create a config file(YAML) like following.
Config file creation is also available by running "mixtail -example".

```
input:
  xxx-log-hostA:
    type: ssh
    host: example-a.com
    user: user_name
    identity: /path/to/key
    command: tail -f /tmp/test.log
  xxx-log-hostB:
    type: ssh
    host: example-b.com
    user: user_name
    identity: /path/to/key
    command: tail -f /tmp/test.log
```

Then, running mixtail command with config file name.

```
mixtail config.yml
```

The command starts wathcing according to config file,
and output all output into a single console.

e.g.

```
[xxx-log-hostA] XX.XX.XX.XX - - [26/Aug/2015:07:26:57 +0900] "GET /robots.txt HTTP/1.1" ...
[xxx-log-hostA] XX.XX.XX.XX - - [26/Aug/2015:07:27:02 +0900] "GET /xxx/xxx/xxx HTTP/1.1" ...
[xxx-log-hostB] XX.XX.XX.XX - - [26/Aug/2015:07:27:17 +0900] "POST /xxx/xxx/xxx HTTP/1.1" ...
[xxx-log-hostA] XX.XX.XX.XX - - [26/Aug/2015:07:27:29 +0900] "GET /xxx/xxx/xxx HTTP/1.1" ...
[xxx-log-hostB] XX.XX.XX.XX - - [26/Aug/2015:07:32:18 +0900] "GET /xxx/xxx/xxx HTTP/1.1" ...
[xxx-log-hostB] XX.XX.XX.XX - - [26/Aug/2015:07:32:22 +0900] "POST /xxx/xxx/xxx HTTP/1.1" ...
[xxx-log-hostA] XX.XX.XX.XX - - [26/Aug/2015:07:32:26 +0900] "GET /xxx/xxx/xxx HTTP/1.1" ...
```


### General usage
General usage is:

```
mixtail -example > config.yml
# modify config.yml as you want.
mixtail config.yml
```

Available options are following:

```
mixtail [options] config-file-path

options:
  --example: Print an example of config file.
  --version: Show version.
  --help:    Show this help.
```
 
## Install

### Binary installation

work in progress...


### From source code

Build instruction: 

```
$ go get -d github.com/tri-star/mixtail
$ cd $GOPATH/github.com/tri-star/mixtail
$ make install
```

By default, `mixtail` is installed on $GOPATH/bin.


## Road map

* Improve config file format, especially for watching many servers. 
  (e.g. make "host" parameter could take host name array.)
* Output filtering and keywork high lighting.
* Output logging.
* Support more input types.

## Contribution

1. Fork ([https://github.com/tri-star/mixtail/fork](https://github.com/tri-star/mixtail/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Create new Pull Request

## Licence

[MIT](https://github.com/tri-star/tigrep/blob/master/LICENSE)

## Author

[tri-star](https://github.com/tri-star)
