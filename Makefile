all:
	@echo "available commands: run, build, ssh, enable"

run:
	go run cli/start-server/main.go -c $$PWD/cv-data.yaml -p $$PWD/CV-NicolasVillanueva.pdf

build:
	docker build -t apivitae .

deploy: build
	docker run -dp 80:80 apivitae

ssh:
	gcloud compute --project "api-vitae" ssh --zone "us-east1-b" "vita-1"

enable:
	eval $(docker-machine env vita-1)