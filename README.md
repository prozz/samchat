# samchat
Simple tcp chat server using [mobster](http://github.com/prozz/mobster) library.

```
go get github.com/prozz/mobster
go build
./samchat -port=<port>
```

Now you may connect with:
```
telnet 127.0.0.1 <port>
```

First message to server have to be auth packet in following format:
```
a <user> <room>
```
(be quick, server disconnects quite fast when no auth packet sent)

After auth packet is sent, any message send to server will be broadcasted to all connected users in same room. Can be easily tested with few telnet sessions opened.

Enjoy!
