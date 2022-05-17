docker run -dit --env BG_COLOR="blue" -p 5050:3000 gcr.io/tetratelabs/color-app:1.0.0
docker run -dit --env BG_COLOR="green" -p 5000:3000 gcr.io/tetratelabs/color-app:1.0.0
envoy -c front-proxy-multiple-services.yaml
