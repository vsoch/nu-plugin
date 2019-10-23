# Nu Plugin Hello

This is an example "hello" sink plugin that uses the Go library [nu-plugin](https://github.com/vsoch/nu-plugin).
You can look at [main.go](main.go) for how much simpler it is to use!
Instructions for building are below.

## Build

If you have Go installed locally, you can use the Makefile to build the plugin

```bash
$ make
go build -o nu_plugin_hello
```

Then you can run nushell, here is inspecting help:

```bash
> help hello
A friendly plugin

Usage:
  > hello {flags} 

flags:
  --excited
  --name <String>
```

Without arguments, it just says hello
```bash
> hello 
Hello
```

You can give it your name!

```bash
> hello --name Vanessa
Hello Vanessa
```
Or tell it to be excited:

```bash
> hello --name Vanessa --excited
Hello Vanessa!
```


## Test With Docker

Once you are happy, you can install the plugin with Nushell easily via Docker.
Here we build the container using first GoLang to compile, and then
copying the binary into quay.io/nushell/nu-base in /usr/local/bin.
We do this so that the plugin is discovered. So first, build the container:

```bash
$ docker build -t vanessa/nu-plugin-hello .
```

Then shell inside - the default entrypoint is already the nushell.

```bash
$ docker run -it vanessa/nu-plugin-hello
```

Once inside, you can run `nu`, or use `nu -l trace` to confirm that nu found your plugin.

```bash
/code(add/circleci)> nu -l trace
```

And proceed with usage as we did before. The plugin will output logs to
`/tmp/nu-plugin-hello.log`
