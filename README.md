### Build Image

```shell
# Build image
docker buildx build --platform linux/arm64 -t internet-watcher-arm:0.1 -f Dockerfile.arm --load  .
# Tag image for push to registry
docker tag internet-watcher-arm:0.1 <my_registry_domain>/internet-watcher-arm:0.1
# Push image to registry
docker push <my_registry_domain>/internet-watcher-arm:0.1
```
