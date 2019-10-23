# Nu Plugin Salad

This is an example sink plugin (based off of [nushell-plugin-pokemon](https://github.com/vsoch/nushell-plugin-pokemon)
but adopted to output salad puns instead of pokemons. Instead of Python, we are trying to use 
the Go library [nu-plugin](https://github.com/vsoch/nu-plugin).
You can look at [main.go](main.go) for how much simpler it is to use!
Instructions for building are below.

## Build

If you have Go installed locally, you can use the Makefile to build the plugin

```bash
$ make
go build -o nu_plugin_salad
```

You can then start nu, and interact with your pun generator! (You of course
cannot see the color here)

```bash
> help salad
Make a punny salad, with fork or spoon.

Usage:
  > salad {flags} 

flags:
  --color <String>
  --fork
  --spoon
```

Generate a fork

```bash
> salad --fork

 Fork off, already.  

          _________________  .========
         [_________________>< :======
                             '======== 
```
Or a spoon!

```bash
> salad --spoon

 They call me the cereal kiler.  

          ___           .-""-.
         /   '''---...-'.'  '\\
         \___...---"""-._-.__//
                         '---'

         

> salad --spoon 

 They say I'm rather spoontaneous!  


        __________        .-"""-.
       /          ''''---' .'    \
       \__________....---. '.    /
                          '-...-'
         
```

You can add a color, although it won't show up in this README

```bash
> salad --spoon --color red


 See you spoon!  

           ________   .==.
          [________>c((_  )
                      '=='

```

## Test With Docker

Once you are happy, you can install the plugin with Nushell easily via Docker.
Here we build the container using first GoLang to compile, and then
copying the binary into quay.io/nushell/nu-base in /usr/local/bin.
We do this so that the plugin is discovered. So first, build the container:

```bash
$ docker build -t vanessa/nu-plugin-salad .
```

Then shell inside and run nu!

```bash
$ docker run -it vanessa/nu-plugin-salad
# nu
```

Once inside, you can also use `nu -l trace` to confirm that nu found your plugin.


## Logging

Don't forget to check out the logging file at `/tmp/nushell-plugin-salad.log` if you 
need to debug!

```bash
/tmp> cat /tmp/nu_plugin_salad.log
```
