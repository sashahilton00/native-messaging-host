[![Build](https://github.com/sashahilton00/native-messaging-host/actions/workflows/build.yml/badge.svg)](https://bit.ly/3djObUY)
[![Coverage](https://img.shields.io/codecov/c/github/sashahilton00/native-messaging-host)](https://bit.ly/2TwjOyb)
[![Dependabot](https://img.shields.io/badge/dependabot-enabled-025e8c?logo=Dependabot)](https://bit.ly/3Li7tqm)
[![GoDev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)][4.1]
[![GoDoc](https://godoc.org/github.com/sashahilton00/native-messaging-host?status.svg)][4.2]
[![License](https://img.shields.io/github/license/sashahilton00/native-messaging-host)][8]

# Native Messaging Host Module for Go

native-messaging-host is a module for sending [native messaging protocol][1]
message marshalled from struct and receiving [native messaging protocol][1]
message unmarshalled to struct. native-messaging-host can auto-update itself
using update URL that response with Google Chrome [update manifest][2],
as well as it provides hook to install and uninstall manifest file to
[native messaging host location][3].

## Installation and Usage

Package documentation can be found on [GoDev][4.1] or [GoDoc][4.2].

Installation can be done with a normal `go get`:

```
$ go get github.com/sashahilton00/native-messaging-host
```

#### Sending Message

```go
messaging := (&host.Host{}).Init()

// host.H is a shortcut to map[string]interface{}
response := &host.H{"key":"value"}

// Write message from response to os.Stdout.
if err := messaging.PostMessage(os.Stdout, response); err != nil {
  log.Fatalf("messaging.PostMessage error: %v", err)
}

// Log response.
log.Printf("response: %+v", response)
```

#### Receiving Message

```go
// Ensure func main returned after calling [runtime.Goexit][5].
defer os.Exit(0)

messaging := (&host.Host{}).Init()

// host.H is a shortcut to map[string]interface{}
request := &host.H{}

// Read message from os.Stdin to request.
if err := messaging.OnMessage(os.Stdin, request); err != nil {
  log.Fatalf("messaging.OnMessage error: %v", err)
}

// Log request.
log.Printf("request: %+v", request)
```

#### Auto Update Configuration

updates.xml example for cross platform executable:

```xml
<?xml version='1.0' encoding='UTF-8'?>
<gupdate xmlns='http://www.google.com/update2/response' protocol='2.0'>
  <app appid='tld.domain.sub.app.name'>
    <updatecheck codebase='https://sub.domain.tld/app.download.all' version='1.0.0' />
  </app>
</gupdate>
```

updates.xml example for individual platform executable:

```xml
<?xml version='1.0' encoding='UTF-8'?>
<gupdate xmlns='http://www.google.com/update2/response' protocol='2.0'>
  <app appid='tld.domain.sub.app.name'>
    <updatecheck codebase='https://sub.domain.tld/app.download.darwin' os='darwin' version='1.0.0' />
    <updatecheck codebase='https://sub.domain.tld/app.download.linux' os='linux' version='1.0.0' />
    <updatecheck codebase='https://sub.domain.tld/app.download.exe' os='windows' version='1.0.0' />
  </app>
</gupdate>
```

```go
// It will do daily update check.
messaging := (&host.Host{
  AppName:   "tld.domain.sub.app.name",
  UpdateUrl: "https://sub.domain.tld/updates.xml", // It follows [update manifest][2]
  Version:   "1.0.0",                              // Current version, it must follow [SemVer][6]
}).Init()
```

#### Install and Uninstall Hooks

```go
// AllowedExts is a list of extensions that should have access to the native messaging host. 
// See [native messaging manifest](https://bit.ly/3aDA1Hv)
messaging := (&host.Host{
  AppName:     "tld.domain.sub.app.name",
  AllowedExts: []string{"chrome-extension://XXX/", "chrome-extension://YYY/"},
}).Init()

...

// When you need to install.
if err := messaging.Install(); err != nil {
  log.Printf("install error: %v", err)
}

...

// When you need to uninstall.
host.Uninstall()
```

#### Syntactic Sugar

You can import client package separately.

```go
import "github.com/sashahilton00/native-messaging-host/client"
```

##### GET call with context

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

resp := client.MustGetWithContext(ctx, "https://domain.tld")
defer resp.Body.Close()
```

##### GET call with tar.gz content

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

client.MustGetAndUntarWithContext(ctx, "https://domain.tld", "/path/to/extract")
```

##### GET call with zip content

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

client.MustGetAndUnzipWithContext(ctx, "https://domain.tld", "/path/to/extract")
```

##### POST call with context

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

resp := client.MustPostWithContext(ctx, "https://domain.tld", "application/json", strings.NewReader("{}"))
defer resp.Body.Close()
```

Contributing
-
If you would like to contribute code to Native Messaging Host repository you can do so
through GitHub by forking the repository and sending a pull request.

If you do not agree to [Contribution Agreement](CONTRIBUTING.md), do not
contribute any code to Native Messaging Host repository.

When submitting code, please make every effort to follow existing conventions
and style in order to keep the code as readable as possible. Please also include
appropriate test cases.

That's it! Thank you for your contribution!

License
-
Copyright (c) 2018 - 2022 Richard Huang.

This utility is free software, licensed under: [Mozilla Public License (MPL-2.0)][8].

Documentation and other similar content are provided under [Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License][9].

[1]: https://bit.ly/3axo5Xv
[2]: https://bit.ly/2vOdAR5
[3]: https://bit.ly/2TuQrMw
[4.1]: http://bit.ly/2Tw22L6
[4.2]: https://bit.ly/2TMGqcj
[5]: https://bit.ly/2Tt4Poo
[6]: https://bit.ly/3cAVAdq
[7]: https://bit.ly/3aDA1Hv
[8]: https://mzl.la/2vLmCye
[9]: https://bit.ly/2SMCRlS
