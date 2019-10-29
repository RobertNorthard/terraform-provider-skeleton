
all: build test

build:
		go get -v
		go build -o terraform-provider-coinbase

test:
		terraform init example/
		terraform plan example/
		terraform apply example/