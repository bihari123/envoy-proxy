docker run --rm  \
  -v $(pwd)/envoy.yaml:/envoy.yaml \
  -v$(pwd)/:/etc/envoy/  \
  -p 80:8080 -p 443:8443 -p 8001:8001  \
  envoyproxy/envoy-dev:f84dc815efe6f3ad78fa6ee0de85d7a47f80fbe8 -c /envoy.yaml;
  
 docker run -d katacoda/docker-http-server; 
 docker run -d katacoda/docker-http-server;


