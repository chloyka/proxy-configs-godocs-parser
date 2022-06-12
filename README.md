## Install

```bash
go install github.com/chloyka/proxy-configs-godocs-parser
```

## What is this?

This is a simple godocs parser that makes it easy to declare and collect instructions for proxy api gateway
Now supports only https://github.com/chloyka/simple-proxy-gateway format

## Why?

Because writing a configs for proxy gateways (f.e. Krakend/Lura) is a little not comfortable and boring

This little console util solves this problem

## How can i use this?

Firstly we need to annotate our endpoints with proxy annotations

| Annotation      | Purpose                                                 |
|-----------------|---------------------------------------------------------|
| @ExternalPath   | External path to call from http                         |
| @InternalPath   | Internal path to call inside your awesome service       |
| @ExternalMethod | External method to call from http                       |
| @InternalMethod | Internal method used for this route inside your service |

```go
// @ExternalPath   /my-awesome-service/health-status
// @ExternalMethod GET
// @InternalPath   /health
// @InternalMethod GET
func Health(context *gin.Context) {
    context.Status(status, 418)
}
```

Now you can call 
```bash
proxy-configs-godocs-parser -host=localhost -port=8080 -scheme=http -outputFile=proxy-config.json
```
To generate your json schema
It doesn't matter what server package do you use

### Args
- host (Doesnt have default value. You should define it yourself)
- port (Default value is 80)
- scheme (Default value is http)
- outputFile (Default value is proxy-config.json)

## Docker example
```dockerfile
FROM your_image as build_configs
ARG SERVICE_HOST
ARG SERVICE_PORT
ARG SERVICE_SCHEME

RUN go install github.com/chloyka/proxy-configs-godocs-parser
RUN proxy-configs-godocs-parser -host=$SERVICE_HOST -port=$SERVICE_PORT -scheme=$SERVICE_PORT
# RUN cat proxy-config.json
```

