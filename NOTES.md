# Course notes

- Golang strings are immutable, so we can use `strings.Builder` to mutate a string while building it.
    ```golang
    var sb strings.Builder
    sb.WriteString("Hello,")
    sb.WriteString(" World!")
    fmt.Println(sb.String())
    ```
- binary is base 2; decimal is base 10; hex is base 16
- the maximum number of a base system is n-1; so in binary the maximum number is 1 (2-1)
- in hex we represent numbers greater than 9 with letters. e.g 10=A, 11=B, 12=C ... 15=F, because 15 (16-1) is the maximum number of base 16
- little endian/big endian https://learnmeabitcoin.com/technical/little-endian
- Golang has builtin functions to convert numbers to different bases:
    ```golang
    fmt.Printf("%b\n", 10) // prints 1010 (base2)
	fmt.Printf("%X\n", 10) // prints A (base16)
    ```