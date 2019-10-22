# Nu Plugin Len

This is an example plugin (based off of [nushell-plugin-len](https://github.com/vsoch/nushell-plugin-len)
that aims to do the same thing, but using the Go library [nu-plugin](https://github.com/vsoch/nu-plugin)
instead. You can look at [main.go](main.go) for how much simpler it is to use!
Instructions for building are below.

## Build

If you have Go installed locally, you can use the Makefile to build the plugin

```bash
$ make
go build -o nu_plugin_len
```

## Test Without Nu

It's possible to test without nushell by giving json objects to the binary.
After you run it and press enter, here are some functions to interact:

### Config

```bash
$ ./nu_plugin_len
{"method":"config"}
{"jsonrpc":"2.0","method":"response","params":{"Ok":{"name":"len","usage":"Return the length of a string","positional":[],"rest_positional":null,"named":{},"is_filter":true}}}
```

### Start and End Filter

```bash
$ ./nu_plugin_len
{"method":"begin_filter"}
{"jsonrpc":"2.0","method":"response","params":{"Ok":[]}}

{"method":"end_filter"}
{"jsonrpc":"2.0","method":"response","params":{"Ok":[]}}
```

### Calculate Length

```bash
$ ./nu_plugin_len
{"method":"begin_filter"}
{"jsonrpc":"2.0","method":"response","params":{"Ok":[]}}

{"method":"filter", "params": {"item": {"Primitive": {"String": "oogabooga"}}}}
{"jsonrpc":"2.0","method":"response","params":{"Ok":[{"Ok":{"Value":{"item":{"Primitive":{"Int":9}},"tag":{"anchor":null,"span":{"end":0,"start":0}}}}}]}}

{"method":"end_filter"}
{"jsonrpc":"2.0","method":"response","params":{"Ok":[]}}
```

The above filter works because the script generates a tag (under params) if the
filter stream doesn't provide one, however with nushell we actually grab
the tag group from the stream and pass it forward. 

```bash
$ ./nu_plugin_len
{"method":"filter", "params": {"item": {"Primitive": {"String": "oogabooga"}}, "tag":{"anchor":null,"span":{"end":10,"start":12}}}}

{"jsonrpc":"2.0","method":"response","params":{"Ok":[{"Ok":{"Value":{"item":{"Primitive":{"Int":9}},"tag":{"anchor":null,"span":{"end":10,"start":12}}}}}]}}
```

In the above we see the custom start and end are passed forward. Both examples (with and without)
will work here, done so that you can test locally without nushell.

## Order of Operations

Keep in mind that when nushell finds the plugin on the path, it's going to:

 - discover the plugin by way of being on the path
 - call the "config" method to get metadata and register it
 - call begin_filter when it's invoked
 - call filter during invokation
 - call end_filter to finish up

That's a very high level description, see [the plugin page](https://github.com/nushell/contributor-book/blob/master/en/plugins.md) to read more about discovery and usage.


## Logging

Note that since I'm going to be using the plugin in a container, I don't
mind logging to a temporary file at `/tmp/nu-plugin-len.log`. If you want
to remove this, remove the logger.* snippets from [main.go](main.go) along
with the "log" import.


## Test With Nu

Once you are happy, you can install the plugin with Nushell easily via Docker.
Here we build the container using first GoLang to compile, and then
copying the binary into quay.io/nushell/nu-base in /usr/local/bin.
We do this so that the plugin is discovered. So first, build the container:

```bash
$ docker build -t vanessa/nu-plugin-len .
```

Then shell inside - the default entrypoint is already the nushell.

```bash
$ docker run -it vanessa/nu-plugin-len
```

Once inside, you can use `nu -l trace` to confirm that nu found your plugin.
Here we see that it did!

```bash
/code(add/circleci)> nu -l trace
...
 TRACE nu::cli > Trying "/usr/local/bin/nu_plugin_len"
 TRACE nu::cli > processing response (176 bytes)
 TRACE nu::cli > response: {"jsonrpc":"2.0","method":"response","params":{"Ok":{"name":"len","usage":"Return the length of a string","positional":[],"rest_positional":null,"named":{},"is_filter":true}}}

 TRACE nu::cli > processing Signature { name: "len", usage: "Return the length of a string", positional: [], rest_positional: None, named: {}, is_filter: true }
 TRACE nu::data::config > config file = /root/.config/nu/config.toml
```

You can also (for newer versions of nu > 0.2.0) use help to see the command:

```bash
/code(master)> help len
Return the length of a string

Usage:
  > len 

/code(master)> 
```

Try out calculating the length of something! First, calculate the length of a simple
string (piped from echo):

```bash
/code(master)> echo four | len
━━━━━━━━━━━
 <unknown> 
───────────
         4 
━━━━━━━━━━━
```

Now let's do a more complex example. Here we are in a directory with one file named "myname" that is empty.

```
/tmp/test> ls
━━━━━━━━┯━━━━━━┯━━━━━━━━━━┯━━━━━━┯━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━━━━
 name   │ type │ readonly │ size │ accessed       │ modified 
────────┼──────┼──────────┼──────┼────────────────┼────────────────
 myname │ File │          │  —   │ 41 seconds ago │ 41 seconds ago 
━━━━━━━━┷━━━━━━┷━━━━━━━━━━┷━━━━━━┷━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━
```

Try listing, getting the name, and calculating the length.

```bash
/tmp/test> ls | get name | len
━━━━━━━━━━━
 <unknown> 
───────────
         6 
━━━━━━━━━━━
```

or test it out with debug.

```bash
ls | get name | len | debug

/tmp/test> ls | get name | len | debug
Tagged { tag: Tag { anchor: None, span: Span { start: 0, end: 2 } }, item: Primitive(Int(BigInt { sign: Plus, data: BigUint { data: [6] } })) }
━━━━━━━━━━━
 <unknown> 
───────────
         6 
━━━━━━━━━━━
```

Add another file to see the table get another row

```bash
touch four
```
```bash
/tmp/test> ls | get name | len 
━━━┯━━━━━━━━━━━
 # │ <unknown> 
───┼───────────
 0 │         4 
 1 │         6 
━━━┷━━━━━━━━━━━
```

Mind you, I'm not a wizard Go Programmer, but I'd like the community to 
at least have an example to start with! Please contribute to this plugin to make
it better!

## Docker Hub

If you don't want to build but just want to play with the plugin,
you can pull directly from Docker Hub

```bash
$ docker pull vanessa/nu-plugin-len
$ docker run -it vanessa/nu-plugin-len
```

Don't forget to check out the logging file at `/tmp/nushell-plugin-len.log` if you 
need to debug!

```bash
/tmp> cat /tmp/nu_plugin_len.log
nu_plugin_len 2019/10/16 15:01:02 Request for config map[jsonrpc:2.0 method:config params:[]]
nu_plugin_len 2019/10/16 15:01:16 Request for begin filter map[jsonrpc:2.0 method:begin_filter params:map[args:map[named:<nil> positional:<nil>] name_tag:map[anchor:<nil> span:map[end:19 start:16]]]]
nu_plugin_len 2019/10/16 15:01:16 Request for filter map[jsonrpc:2.0 method:filter params:map[item:map[Primitive:map[String:nu_plugin_len.log]] tag:map[anchor:<nil> span:map[end:2 start:0]]]]
nu_plugin_len 2019/10/16 15:01:16 Request for end filter map[jsonrpc:2.0 method:end_filter params:[]]
```
