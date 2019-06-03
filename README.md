# kube-local-proxy
Local Proxy Server (Grpc + HTTP) for cross contexts

![kelp](https://i.ibb.co/jWcSsp6/kisspng-seaweed-kelp-clip-art-5b3902b62c96b9-9221605215304629021827.jpg "kelp")

## Install

```
$ git clone https://github.com/lukluk/kube-local-proxy.git $GOPATH/src/github/lukluk/kube-local-proxy
$ cd $GOPATH/src/github/lukluk/kube-local-proxy
$ go get github.com/google/tcpproxy
$ go build -o bin/kube-local-proxy 
$ cp bin/* /usr/local/bin/

```

## How to use

1. Create kube-local-proxy configuration file `$HOME/.klp.cfg`
2. Example kube-local-proxy conf :
   ```
   # for HTTP service
   # host = context/app-name:port

   staging-web-api.localhost = STAGING-WEB/front-web-api:8080
   cms.localhost = STAGING-CMS/cms-web-app:8080

   # for GRPC service
   # host = context/app-name:port:package (package of proto file)

   localhost = STAGING-BACKEND/foo-service:8080:com.foo.package
   localhost = STAGING-BACKEND/bar-service:8080:com.bar.package
   ```
Done .

## Starting Proxy
```
$ klp
kube-local-proxy ready!
```

## Listen ports

80 for HTTP 

443 for GRPC 

## Dependencies 

"github.com/google/tcpproxy"