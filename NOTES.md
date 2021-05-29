# Course notes

- Golang strings are immutable, so we can use `strings.Builder` to mutate a string while building it.
    ```golang
    var sb strings.Builder
    sb.WriteString("Hello,")
    sb.WriteString(" World!")
    fmt.Println(sb.String())
    ```