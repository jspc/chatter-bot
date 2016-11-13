chatter-bot
==

`chatter-bot` is a very simple irc bot that, based on a number of commands, will do some stuff.

I'm not sure what it'll do; I'm more curious to see whether I can hook it into my daily workflow and whether or not this idea of [chatops](http://blogs.atlassian.com/2016/01/what-is-chatops-adoption-guide/) is worth looking at in depth.

| who       | what |
|-----------|------|
| dockerhub | https://hub.docker.com/r/jspc/chatter-bot/   |
| circleci  | https://circleci.com/gh/jspc/chatter-bot   |
| licence   | MIT   |


Compiling
--

`chatter-bot` is written in go. Thus: the usual go build:

```bash
go build
go test -v
./chatter-bot
```

Running
--

Currently there is nothing particularly complicated in running the tool:

```bash
/chatter-bot -help
Usage of ./chatter-bot:
  -c string
        Channel with which to connect (default "#control-room")
  -h string
        IRC daemon to connect to (default "localhost:6667")
  -n string
        Nick with which to connect (default "rosie")
  -s string
        Location of user scripts (default "/scripts")
  -u string
        User to listen to (default "jspc")
```

The room and user are worth looking at.

In [`ircrouter.go`](https://github.com/jspc/chatter-bot/blob/master/ircrouter.go#L56-L63) we determine whether `chatter-bot` should try to do anything based on whether the command is a ['Safe Command'](https://github.com/jspc/chatter-bot/blob/master/ircrouter.go#L11) (which anybody can run) or whether the command has come from `User to listen to` via a private message (`query` window, for instance).


Docker
--

You can find instructions for docker in the table above.

Plugins/ Tasks
--

Plugins are written in javascript. For a sample, please see the scripts directory. Ultimately; we pass context from chatter-bot (See `scripts.go`) in an object called `Context`. We execute the script and expect a string at the end called `Output`.

The filename determines how it is called. `echo.js` will be run when `chatter-bot` receives the command `echo`.


Licence
--

MIT License

Copyright (c) 2016 jspc

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
