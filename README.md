Golang xtask demo
=====================

Inspired by [cargo-xtask](https://github.com/matklad/cargo-xtask) and use xtask for your Go project scripts.

# Get started

Create a `xtask` directory in your project root, and create a `main.go` file in the `xtask` directory.

```shell
go run xtask/main.go hello
```

# Standard tasks

* `go run xtask/main.go --help`: list all available tasks, and output as following

```
Tasks:
  setup - environment setup
  hello - hello task
```

# Tools & Libraries

* Command Line - Awesome Go: https://awesome-go.com/command-line/

# References

* Use task.go for your Go project scripts: https://dev.to/jcbhmr/use-taskgo-for-your-go-project-scripts-2cm4
* cargo xtask: https://github.com/matklad/cargo-xtask