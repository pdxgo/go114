# Preemption

From the release notes - 

```
 This release improves the performance of most uses of defer to incur almost zero overhead compared to calling the deferred function directly. As a result, defer can now be used in performance-critical code without overhead concerns. 

 Goroutines are now asynchronously preemptible. As a result, loops without function calls no longer potentially deadlock the scheduler or significantly delay garbage collection. This is supported on all platforms except windows/arm, darwin/arm, js/wasm, and plan9/*. 
 ```

 ## Running the example

With Go 1.14 installed, just do a simple `go run main.go` and notice when the program outputs `f1`. 

You should see a new line printed out every tenth of a second with `f2` at the end of the line, then after a few seconds a line that prints `f1`, followed by a number. 

 Now, let's get the previous version of Go installed on your system.

 ```
 go get golang.org/dl/go1.13.8
 go1.13.8 download
 ```

This lets you run the older versions of Go and means that you should be able to see what the impact is for any of your programs between versions.

Now run the previous example - 

```
go1.13.8 run main.go
```

What happens? 

Do the empty loops block?

If you are using Windows, what happens when you run this?
 
## What is the cost?

According to the blog post, it looks like they accomplish this by utilizing a mix of signals and syscalls in the runtime. If your programs use the `syscall` package, you may want to evaluate them and mitigate extra failed `EINTR` errors by retrying the calls. 


## Acknowledgements

All credit for this code example goes to Eli Bendersky (@elibendersky on Twitter, or check out his [blog](https://eli.thegreenplace.net/))