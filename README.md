# logrotate

```go
func ExampleOpen() {
	var f = logrotate.Open("./logs.txt", 500)
	defer f.Close()
	for i := 0; i < 30; i++ {
		f.Write([]byte(time.Now().String() + "\n"))
	}
	fmt.Scanln()
}
```

## Reference

[Logging Architecture](https://kubernetes.io/docs/concepts/cluster-administration/logging/#how-nodes-handle-container-logs)
