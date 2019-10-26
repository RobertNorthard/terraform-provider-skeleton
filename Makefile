
all: build test

build:
		go build -o terraform-provider-skeleton

test:
		terraform init example/
		terraform plan example/