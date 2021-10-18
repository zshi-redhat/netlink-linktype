# Get device link type via netlink (golang)

## Build

```
$ go build main.go
```

## Run

Build the main binary or use existing one in the rep

```
$ ./main <interface-name-1> <interface-name-2> ...
```

For example:

```
# ./main ens1f1v1 ens1f1 ens1f1v0 
device ens1f1v1 link type is: ether
device ens1f1 link type is: ether
device ens1f1v0 link type is: ether
```
