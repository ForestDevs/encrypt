package encrypt_decrypt

//go:generate protoc --proto_path=proto proto/*proto --go-grpc_out=. --go_out=. --grpc-gateway_out=. --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out=:swagger
