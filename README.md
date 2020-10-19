# Generating User-Agent

Simple example:
```go
package main
import "github.com/Seven45/user_agent"

func main() {
    var ua = user_agent.GenerateUserAgent("chrome", "windows")
    print(ua)
}
```
---
Supported generating for browsers: 
- chrome;
- firefox;
- opera

Supported generating for OS:
- windows;
- macos;
- linux

When you pass unknown arguments, user-agent is generated randomly.
