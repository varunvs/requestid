# Description
A plugin for go-kit to generate request id and store it in context. Also, write the request id in the response as well. It will look for X-Request-Id header and use that as the request id. If it is not available, it'll generate one.

It does not use UUID for generating the id. A random string is generated using the built-in rand.Read function and is encoded as a string.

The obtained request id is stored in context as "requestID" key.

# Usage
Use it as a ServerOption while building the server. For example:
```
    import kithttp "github.com/go-kit/kit/transport/http"
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(util.RequestIDToContext()),
		kithttp.ServerAfter(util.RequestIDToResponse()),
	}

```