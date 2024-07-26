docker run --add-host host.docker.internal:host-gateway \
    --network host \
    kong/deck:v1.37.0 gateway dump \
    --kong-addr http://host.docker.internal:8001 \
    >> kong/declarative/kong.yaml
