generate:
	protoc -I models/ models/*.proto --go_out=plugins=grpc,import_path=models:models

build: generate
	go build -o bin/auth ./service_auth
	go build -o bin/msin ./service_main