# Go playground

| Branch |                                                                                   Pipeline                                                                                   |                                                                                Code coverage                                                                                 |                                          Test report                                           |                                 SonarCloud                                 |
|:------:|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------:|
|  main  | [![pipeline status](https://gitlab.com/ShowMeYourCodeYouTube/go-playground/badges/main/pipeline.svg)](https://gitlab.com/ShowMeYourCodeYouTube/go-playground/-/commits/main) | [![coverage report](https://gitlab.com/ShowMeYourCodeYouTube/go-playground/badges/main/coverage.svg)](https://gitlab.com/ShowMeYourCodeYouTube/go-playground/-/commits/main) | [link](https://showmeyourcodeyoutube.gitlab.io/go-playground/test-report/coverage-report.html) | [link](https://sonarcloud.io/organizations/showmeyourcodeyoutube/projects) |

---

Get familiar with Go!

https://go.dev/

https://play.golang.com/p/lC0mTi8gxbU

Most code samples are based on:
- https://go.dev/tour/list
- https://gobyexample.com/

## Getting started

1. Install Go.
2. Run `go mod tidy`.
3. Run `main.go`.

## Go

Golang, also known as Go, is a statically typed and compiled programming language created by Google.
It is compiled to binary and a Go binary does not need system dependencies such as Go tooling to run on a new system.
Putting these executables in an executable filepath on your own system will allow you to run the program from anywhere on your system.

A compiled language is a programming language whose implementations
are typically compilers (translators that generate machine code from source code),
and not interpreters (step-by-step executors of source code, where no pre-runtime translation takes place).

Static typing is a typing system where variables are bound to a data type during compilation.
Once a variable is assigned a data type it remains unchanged throughout the program's execution.

> Go compiles to a target OS and CPU architecture.
The result is statically linked so the binary will run almost anywhere that those two values hold true.
[Reference](https://www.quora.com/Is-the-Golang-write-once-run-everywhere-language-like-Java)

When you run a command like go build, Go uses the current platform’s GOOS and GOARCH to determine how to build the binary.
```
go env GOOS GOARCH
```
Running this on Windows will print: `windows amd64`.

[Building Go Applications for Different Operating Systems](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures)

### Concurrency

Concurrency refers to the ability of a computer system to perform multiple tasks simultaneously.

Goroutines and Channels are a lightweight built-in feature for 
managing concurrency and communication between several functions executing at the same time.

#### Goroutines

A goroutine is a lightweight thread managed by the Go runtime. [Docs](https://go.dev/tour/concurrency/1)

#### Channels

A Channel is a built-in data structure that allows Goroutines to communicate and synchronize their activities. 
Channels provide a way to send and receive values between Goroutines in a safe and efficient way, 
without the need for locks or other synchronization primitives.

Channels can be used to communicate between Goroutines by sending and receiving values. 
The <- operator is used to send and receive values on a channel. 
For example, to send a value on a channel, you would write `channel <- value`. 
To receive a value from a Channel, you would write `value := <-channel`.

#### Benefits of Goroutines/Channels

- Goroutines are extremely lightweight, requiring only a few kilobytes of memory compared to the several megabytes required by traditional threads. This means that Go programs can create and manage a large number of Goroutines without incurring significant memory overhead, leading to improved performance and scalability.
- using Goroutines in Go also simplifies synchronization between concurrent activities. Unlike traditional thread synchronization mechanisms like locks and semaphores, Goroutines can be synchronized using channels, which are simpler and more intuitive to use. Channels help to avoid race conditions and deadlocks by ensuring that data is safely shared between Goroutines.
