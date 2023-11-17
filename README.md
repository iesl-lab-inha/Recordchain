
# Recordchain in Go
[Main Publication](https://ieeexplore.ieee.org/abstract/document/9738809)

[Second publication](https://www.mdpi.com/1424-8220/23/21/8762)

A basic implementation of Recordchain in Go

## Usage
Start a new node
```
$ recordchain startnode -port PORT -type DATA_OR_HASH
```

Create a new record
```
$ recordchain addrecord -hash RECORD_HASH -expire EXPIRATION_IN_SECONDS
```

Create a new data
```
$ recordchain adddata -hash DATA_LOCATION -expire EXPIRATION_IN_SECONDS
```

Mine a new block from the Mempool
```
$ recordchain mineblock 
```

Mine a new block using proof-of-work
```
$ recordchain minepow 
```

Print the current version of chain
```
$ recordchain printchain 
```


## TODO
- Wallet management
- Account management
- Coin generation & Incentives
- Transaction handling
- Network optimization


## Requirements
- github.com/boltdb/bolt


## License
[MIT](https://choosealicense.com/licenses/mit/)