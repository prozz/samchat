# samchat
simple tcp chat server using mobster framework

```
go build
./samchat -port=<port>
```

now you may do
```
telnet 127.0.0.1 <port>
```

first message to server have to be auth packet in following format
```
a <user> <room>
```
be quick, server disconnects quite fast when no auth packet sent

after auth packet any message send to server will be broadcasted to any connected user in same room. can be easily tested with few telnet sessions opened.

enjoy!
