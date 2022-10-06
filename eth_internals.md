# Content

<br />

### :phone: RPC API
* [eth_blockNumber](#ethblocknumber)

<br />
<br />

# :phone: RPC API

### eth_blockNumber

<br />

This api endpoint is not reading levelDB files directly. It's using preload header from HeaderChain.currentHeader 
which was set in BlockChain.loadLastState().

<br />

| Variable                          | Value                    | File                  |
| --------------------------------- | -------------------------| --------------------- |
| rawdb.headBlockKey                | []byte("LastBlock")      | ethapi/api.go         |

<br />

| Function                          | File                            |
| --------------------------------- | ------------------------------- |
| BlockChainAPI.BlockNumber         | ethapi/api.go                   |
| BlockChain.loadLastState          | core/blockchain.go              |
| rawdb.ReadHeadBlockHash           | rawdb/accessor_chain.go         |
| BlockChain.GetBlockByHash         | core/blockchain.go              |
| HeaderChain.GetBlockNumber        | core/headerchain.go             |
| BlockChain.GetBlock               | core/blockchain_reader.go       |
| rawdb.ReadBlock                   | rawdb/accessor_chain.go         |



