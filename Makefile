IMG:=apivitae
VMNAME:=vita-1

all:
	@echo "available commands: run, build, ssh, enable\nImage name: $(IMG)\nVM name: $(VMNAME)"

run:
	go run cli/start-server/main.go -c $$PWD/cv-data.yaml -p $$PWD/CV-NicolasVillanueva.pdf

build:
	docker build -t $(IMG) .

deploy: build
	docker run -dp 80:80 $(IMG)

ssh:
	gcloud compute --project "api-vitae" ssh --zone "us-east1-b" "$(VMNAME)"

enable:
	eval $(docker-machine env $(VMNAME))

export: build
	docker save -o $(IMG)-export $(IMG)

import:
	docker load -i $(IMG)-export