# sup 👋
Simple UDP app that always responds with 'sup'

## Usage

**server**
```
$ PORT=8080 go run main.go
```

Example output
```
sup is running...
```

**client**
```
$ go run client/main.go
```

Example output
```
Send: hey
Resp: sup
```
