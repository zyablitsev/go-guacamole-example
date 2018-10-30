guacd_host=$(or $(GUACD_HOST),127.0.0.1)
guacd_port=$(or $(GUACD_PORT),4822)
ssh_host=$(or $(SSH_HOST),127.0.0.1)
ssh_port=$(or $(SSH_PORT),22)

all:
	docker rm -f guacd 2>/dev/null || true
	docker run --name guacd -d -p 4822:4822 guacamole/guacd
	GUACD_HOST=$(guacd_host) \
	GUACD_PORT=$(guacd_port) \
	SSH_HOST=$(ssh_host) \
	SSH_PORT=$(ssh_port) \
	SSH_USER=$(SSH_USER) \
	SSH_PASSWORD=$(SSH_PASSWORD) \
	go run .
