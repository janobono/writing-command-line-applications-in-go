# 05 adding version information

```bash
go build -ldflags="-X main.version=1.2.3"
```

```bash
go build -ldflags="-X main.version=$(git rev-parse --short HEAD)"
```
