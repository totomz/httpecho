# httpecho
A simple HTTP-echo container to debug Kubernetes ingress

# How to use
HTTPEcho listen to the port `8080` for incoming HTTP GET requests on any path, and returns the request URI and all the headers. 

Check the `httpecho.yaml` for the resource definition, and then `kubectl apply -f httpecho.yaml`  

The container image is [totomz84/httpecho](https://hub.docker.com/r/totomz84/httpecho)

## Test upload
It is possible to test uploads by sending form-data to the `/upload` endpoint: 
```bash
curl -X POST  -F file=@Syleak_Brand.zip  http://localhost:8080/upload
```