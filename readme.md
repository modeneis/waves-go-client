# Waves REST API - Golang Client

# Supported:
* GET /transactions/info/{id}
Return transaction data by transaction ID.
Usage:
```
wavesClient := client.NewTransactionsService()
transaction, res, err := wavesClient.GetTransactionsInfoID(expected.ID)
```

* GET /blocks/seq/{from}/{to}
Return block data at the given height range
Usage:
```
wavesClient := client.NewBlocksService()
blocks, res, err := wavesClient.GetBlocksSeqFromTo(333,334)
```

* GET /blocks/at/{height}
Return block data at the given height
Usage:
```
wavesClient := client.NewBlocksService()
block, res, err := wavesClient.GetBlocksAtHeight(333)
```

* GET /blocks/last
Return the last block data
Usage:
```
wavesClient := client.NewBlocksService()
block, res, err := wavesClient.GetBlocksLast()
```

## Get module:
`go get -u github.com/modeneis/waves-go-client`


## LICENSE
```
Copyright (c) 2018 Thomas Modeneis

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
