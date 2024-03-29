# IU Common Go

A Go module containing a collection of carefully-maintained packages designed to enable rapid implementation of microservice codepaths.

## Installation

Since this is a private module, additional auth steps must be completed to `go get` it.

1. Ensure your [.netrc file](https://www.gnu.org/software/inetutils/manual/html_node/The-_002enetrc-file.html) is configured with an access token to use for gitlab.innovationup.stream.

**What/where is .netrc?**
>The .netrc file contains login and initialization information used by the auto-login process. It generally resides in the user’s home directory

```
machine gitlab.innovationup.stream
login <gitlab email>
password <gitlab personal access token>
```

2. Set the GOPRIVATE go env var so Golang bypasses module proxy servers and downloads directly from this Gitlab.

```shell
$ go env -w GOPRIVATE=gitlab.innovationup.stream
$ go get github.com/innovation-upstream/iu-common-go
```
