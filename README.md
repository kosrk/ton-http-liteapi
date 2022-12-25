# REST API for TON liteserver
The service provides the following lite api methods from [lite_api.tl](https://github.com/ton-blockchain/ton/blob/master/tl/generate/scheme/lite_api.tl) file (by default):
* liteServer.getMasterchainInfo
* liteServer.getMasterchainInfoExt
* liteServer.getTime
* liteServer.getBlock
* liteServer.getState
* liteServer.getBlockHeader
* liteServer.sendMessage
* liteServer.getAccountState
* liteServer.getShardInfo
* liteServer.getAllShardsInfo
* liteServer.getTransactions
* liteServer.listBlockTransactions
* liteServer.getBlockProof
* liteServer.getConfigAll
* liteServer.getShardBlockProof

The service also supports the generation of additional endpoints. See: [generator.go](api/generator.go)

### TL requests and responses are encoded in json:
* int to number
* long to string
* int256 to hex string
* bytes to base64 string

#### Request example:

See: [example.http](api/example.http)

`POST {{url}}/v1/lite_server_get_block_header`

```json
{
  "id": {
    "workchain": 0,
    "shard": "-9223372036854775808",
    "seqno": 5456001,
    "root_hash": "DC9806B841A03090300A2D5736A0E0FB4BA2D49D5832CF70EA7F78FF1C3BF451",
    "file_hash": "726374D979990AFD5FF1B9C9AFEBC779128EB080CE847FCA7CDBF1BA72E86984"
  },
  "mode": 0
}
```

#### Response body example:

```json
{
  "id": {
    "workchain": 0,
    "shard": "-9223372036854775808",
    "seqno": 5456001,
    "root_hash": "dc9806b841a03090300a2d5736a0e0fb4ba2d49d5832cf70ea7f78ff1c3bf451",
    "file_hash": "726374d979990afd5ff1b9c9afebc779128eb080ce847fca7cdbf1ba72e86984"
  },
  "mode": 0,
  "header_proof": "te6ccgECCAEAAZYACUYD3JgGuEGgMJAwCi1XNqDg+0ui1J1YMs9w6n94/xw79FEAFwEkEBHvVar////9AgMEBQKgm8ephwAAAACEAQBTQIEAAAAAAAAAAAAAAAAAAAAAAGNhD0IAAAURMF0+AAAABREwXT4Iknhx/gAA660ARStcAEUmEcQAAAADAAAAAAAAAC4GByhIAQFxs/gEKCmri/9jYB98y9qzYBwGP/Uu/WPbH5w69CPilwABKEgBAepb8ef5HgKOfund1xZaHZTtlX0jWEEhxAySkImqMreAABYoSAEB66XRh7pyKoNuVbnXBbwM0vJ4LPXokSMw9Iho7kFIbGAAEgCYAAAFETBN+8QARStcd7Vi/ID9SPKWVODxBnbVBzEPIcf5o/tHzeagqKghXGiSsvSA7RObX6BaSWqWEzuWSTmBgqQdux8nlSK0ltUscACYAAAFETBN+8EAU0CAH/jAObl+3jc4BwGZ9R59jgWqgpxL+jMp6SbuxN8JE6Abq/5AVIzjUTEp+SM4hxMC6OEfLoFC2LbvijAc4Cn84Q=="
}
```

### Configurable parameters
| ENV variable           | Description                                                         |
|------------------------|---------------------------------------------------------------------|
| `LITESERVER`           | IP and port of lite server in format `host:port`                    |
| `LITESERVER_KEY`       | public key of lite server as base64 string                          |
| `API_HOST`             | host for REST API, example `localhost:8081`, default `0.0.0.0:8081` |

### Docker image
```shell
docker pull kosrk/ton-http-liteapi:latest
```
or use [Makefile](Makefile)