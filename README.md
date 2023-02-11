# iron

Iron is an much better reverse shell manager than delltaxa/ilander it is async and connects via a golang-nc client

# Commands

show -> displays all connections which where gathered by an async listener
mkexploit <addr> -> generate a payload to connect to this revere shell manager (payloads from revshells.com)
exploit <addr> or just exploit -> waits for an connection from an client

# Installation

```bash
git clone https://github.com/delltaxa/iron
cd iron/
```

# Usage:

```bash
go run .
```

or 

```bash
go run . 0.0.0.0:12345
```
