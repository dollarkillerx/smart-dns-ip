Generate:
	@echo 'Build GRPC'
	protoc -I generate/ generate/smart_dns_ip/*.proto --go_out=plugins=grpc:generate/smart_dns_ip/.
