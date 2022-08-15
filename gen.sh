protoc -I . --proto_path=./api/ \
    --go_out ./gen/ --go_opt paths=import \
    --go-grpc_out ./gen/ --go-grpc_opt paths=import \
    --grpc-gateway_out ./gen/ \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=import \
    --grpc-gateway_opt generate_unbound_methods=true \
    --openapiv2_out ./swaggerui/ \
    --openapiv2_opt logtostderr=true \
    --openapiv2_opt merge_file_name=swagger.json \
    --openapiv2_opt allow_merge=true \
    api/*.proto