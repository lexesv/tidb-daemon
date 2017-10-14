# tidb-daemon
demonizing the tidb-server

### Usage:
1. Copy daemon.go to $GOPATH/src/github.com/pingcap/tidb/tidb-server
2. In Makefile replace all strings  "tidb-server/main.go" to "tidb-server/*.go"
3. Run "make"

```sh
./tidb-server -d -store memory -P 4001
netstat -tpln | grep 4001
```
