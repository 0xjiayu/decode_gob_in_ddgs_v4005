# decode_gob_in_ddgs_v4005
**Decode gob encoded Seed Nodes data in malware sample of DDG botnet v4005**

- [hubs_dump.go](hubs_dump.go) is the decoding program written in Go
- [hubs_gob.raw](hubs_gob.raw) is the gob encoded raw data dumped from malware sample(MD5: 638061d2a06ebdfc82a189cf027d8136)
- [ddgs_v4005.log](ddgs_v4005.log) is the debug log while ddgs running.

```
$ go run hubs_dump.go -f hubs_gob.raw
```
