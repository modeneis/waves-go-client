# Waves REST API - Golang Client

# Supported:
* GET /transactions/info/{id}
* GET /blocks/seq/{from}/{to}
* GET /blocks/at/{height}
* GET /blocks/last


## Get module:
`go get -u github.com/modeneis/waves-go-client`


## Usage:
```
wavesClient := client.NewTransactionsService()
transaction, res, err := wavesClient.GetTransactionsInfoID(expected.ID)
```