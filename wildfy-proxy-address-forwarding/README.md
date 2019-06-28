# wildfy-proxy-address-forwarding

Code for article about proxy-address-forwarding in WildFly

## Running

```
gradle :build && docker-compose up --build
```


## Exploring

```
curl \
     http://docker:8080/proxy-address-forwarding-demo/api/request-info

{
   "servletRequest":{
      "contextPath":"/proxy-address-forwarding-demo",
      "localAddr":"172.21.0.2",
      "localName":"172.21.0.2",
      "localPort":8080,
      "parameterMap":{

      },
      "pathInfo":"/request-info",
      "remoteAddr":"192.168.99.1",
      "remoteHost":"192.168.99.1",
      "remotePort":51357,
      "requestURI":"/proxy-address-forwarding-demo/api/request-info",
      "requestURL":"http://docker:8080/proxy-address-forwarding-demo/api/request-info",
      "serverName":"docker"
   },
   "uriInfo":{
      "absolutePath":"http://docker:8080/proxy-address-forwarding-demo/api/request-info",
      "baseUri":"http://docker:8080/proxy-address-forwarding-demo/api/",
      "path":"/request-info",
      "queryParameters":{

      },
      "requestUri":"http://docker:8080/proxy-address-forwarding-demo/api/request-info"
   }
}

```

```

curl -H 'X-Forwarded-Host: evil.com' \
     http://docker:8080/proxy-address-forwarding-demo/api/request-info


{
   "servletRequest":{
      "contextPath":"/proxy-address-forwarding-demo",
      "localAddr":"evil.com",
      "localName":"evil.com",
      "localPort":0,
      "parameterMap":{

      },
      "pathInfo":"/request-info",
      "remoteAddr":"192.168.99.1",
      "remoteHost":"192.168.99.1",
      "remotePort":51396,
      "requestURI":"/proxy-address-forwarding-demo/api/request-info",
      "requestURL":"http://evil.com/proxy-address-forwarding-demo/api/request-info",
      "serverName":"evil.com"
   },
   "uriInfo":{
      "absolutePath":"http://evil.com/proxy-address-forwarding-demo/api/request-info",
      "baseUri":"http://evil.com/proxy-address-forwarding-demo/api/",
      "path":"/request-info",
      "queryParameters":{

      },
      "requestUri":"http://evil.com/proxy-address-forwarding-demo/api/request-info"
   }
}

```


```
curl -H 'X-Forwarded-Host: evil.com' \
	 -H 'X-Forwarded-Proto: https' \
     http://docker:8080/proxy-address-forwarding-demo/api/request-info


{
   "servletRequest":{
      "contextPath":"/proxy-address-forwarding-demo",
      "localAddr":"evil.com",
      "localName":"evil.com",
      "localPort":0,
      "parameterMap":{

      },
      "pathInfo":"/request-info",
      "remoteAddr":"192.168.99.1",
      "remoteHost":"192.168.99.1",
      "remotePort":51411,
      "requestURI":"/proxy-address-forwarding-demo/api/request-info",
      "requestURL":"https://evil.com/proxy-address-forwarding-demo/api/request-info",
      "serverName":"evil.com"
   },
   "uriInfo":{
      "absolutePath":"https://evil.com/proxy-address-forwarding-demo/api/request-info",
      "baseUri":"https://evil.com/proxy-address-forwarding-demo/api/",
      "path":"/request-info",
      "queryParameters":{

      },
      "requestUri":"https://evil.com/proxy-address-forwarding-demo/api/request-info"
   }
}
```