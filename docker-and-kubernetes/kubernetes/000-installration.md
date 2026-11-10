# Linux Kubernetes Installation

## *.Ubuntu Install

#### Can't connecting kubetnetes api

```bash
geonwoo@geonwoo-work:~/projects/org-main/main-infra-manage/docker$ kubectl get pods
E1110 08:36:27.879096   30365 memcache.go:265] "Unhandled Error" err="couldn't get current server API group list: Get \"http://localhost:8080/api?timeout=32s\": dial tcp 127.0.0.1:8080: connect: connection refused"
E1110 08:36:27.879288   30365 memcache.go:265] "Unhandled Error" err="couldn't get current server API group list: Get \"http://localhost:8080/api?timeout=32s\": dial tcp 127.0.0.1:8080: connect: connection refused"
E1110 08:36:27.880565   30365 memcache.go:265] "Unhandled Error" err="couldn't get current server API group list: Get \"http://localhost:8080/api?timeout=32s\": dial tcp 127.0.0.1:8080: connect: connection refused"
E1110 08:36:27.880725   30365 memcache.go:265] "Unhandled Error" err="couldn't get current server API group list: Get \"http://localhost:8080/api?timeout=32s\": dial tcp 127.0.0.1:8080: connect: connection refused"
E1110 08:36:27.881935   30365 memcache.go:265] "Unhandled Error" err="couldn't get current server API group list: Get \"http://localhost:8080/api?timeout=32s\": dial tcp 127.0.0.1:8080: connect: connection refused"
The connection to the server localhost:8080 was refused - did you specify the right host or port?
```

#### Get Kubectl config data

```bash
geonwoo@geonwoo-work:/usr/bin$ kubectl config view
apiVersion: v1
clusters: null
contexts: null
current-context: ""
kind: Config
users: null
```
