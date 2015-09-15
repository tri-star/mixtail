mixtail
===================================

![release](http://img.shields.io/github/release/tri-star/mixtail.svg?style=flat-square)
![license](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)

# Overview
Mixtail watches multiple log files( or command output) and output them in single console.
This command is useful for watching file(or command output) across many servers.


## Usage
At first, create a config file(YAML) like following.
Config file creation is also available by running "mixtail -example".

```
input:
  xxx-log:
    type: ssh
    host: 
      - 192.168.1.10
      - 192.168.1.11
      - 192.168.1.12
      - 192.168.1.13
      - 192.168.1.14
      - 192.168.1.15
    # if you don't want define user,pass or identity here,
    # you can define them in external file, see "Advanced usage" section.
    user: user_name
    pass: password
    # identity: /path/to/key
    command: tail -f /tmp/test.log
  yyy-log:
    type: ssh
    host: 192.168.1.10
    user: user_name
    pass: password
    # identity: /path/to/key
    command: tail -f /tmp/another.log | grep "error"
```

Then, running mixtail command with config file name.

```
mixtail config.yml
```

The command starts wathcing according to config file,
and output all output into a single console.

e.g.

```
[192.168.1.10: xxx-log] XX.XX.XX.XX - - [26/Aug/2015:07:26:57 +0900] "GET /robots.txt HTTP/1.1" ...
[192.168.1.11: xxx-log] XX.XX.XX.XX - - [26/Aug/2015:07:27:02 +0900] "GET /xxx/xxx/xxx HTTP/1.1" ...
[192.168.1.10: xxx-log] XX.XX.XX.XX - - [26/Aug/2015:07:27:17 +0900] "POST /xxx/xxx/xxx HTTP/1.1" ...
[192.168.1.12: xxx-log] XX.XX.XX.XX - - [26/Aug/2015:07:27:29 +0900] "GET /xxx/xxx/xxx HTTP/1.1" ...
[192.168.1.13: xxx-log] XX.XX.XX.XX - - [26/Aug/2015:07:32:18 +0900] "GET /xxx/xxx/xxx HTTP/1.1" ...
[192.168.1.10: yyy-log] error: xxxxxxxx
[192.168.1.10: xxx-log] XX.XX.XX.XX - - [26/Aug/2015:07:32:26 +0900] "GET /xxx/xxx/xxx HTTP/1.1" ...
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

### Advanced usage

#### Define credential info at external config file.
If you don't want define user,pass or identity in config.yml,
you can define them in external config file(./.mixtail.yml or ~/.mixtail.yml).

Example:
```
cat > ./.mixtail.yml <<EOT
default:
  user: username
  identity: path/to/key
EOT
# if you have defined above settings,
# you can ommit user,pass or identity in config.yml.
cat > ./config.yml <<EOT
input:
  xxx-log:
    type: ssh
    host: 
      - 192.168.1.10
      - 192.168.1.11
    command: tail -f /tmp/test.log
EOT
mixtail config.yml
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

* Make user ID, password, identity could define external file(~/.mixtail.yml) or pass via command line. 
* Add option for prepend time to each log line.
* Keyword highlighting.
* Support more input types.

## Contribution

1. Fork ([https://github.com/tri-star/mixtail/fork](https://github.com/tri-star/mixtail/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Create new Pull Request

## Licence

[MIT](https://github.com/tri-star/mixtail/blob/master/LICENSE)

## Author

[tri-star](https://github.com/tri-star)
