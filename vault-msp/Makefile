docker-lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.33.0 golangci-lint run -v

vault-server : 
	docker run --cap-add=IPC_LOCK -e 'VAULT_DEV_ROOT_TOKEN_ID=myroot' -e 'VAULT_DEV_LISTEN_ADDRESS=127.0.0.1:8200' --network host --name vault-server -d vault

docker-prune:
	docker system prune 
	docker volume prune 
