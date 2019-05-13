IMG:=nicovillanueva/vitae
VMNAME:=vita-1
REMOTE_ENV:=DOCKER_TLS_VERIFY="1" DOCKER_HOST="tcp://34.73.0.147:2376" DOCKER_CERT_PATH="/Users/nico/.docker/machine/machines/vita-1" DOCKER_MACHINE_NAME="vita-1"

run: swagger
	go run cli/main.go start -c $$PWD/cv-data.yaml -p $$PWD/CV-NicolasVillanueva.pdf

build: swagger
	docker build -t $(IMG) .

release: build
	docker push $(IMG)

deploy:
	$(REMOTE_ENV) docker run -dp 8080:80 $(IMG)

ssh:
	gcloud compute --project "api-vitae" ssh --zone "us-east1-b" "$(VMNAME)"

swagger: clear-swag
	@echo ">> generating swagger"
	swag init -g server/server.go -o server/docs/

clear-swag:
	@echo ">> clearing generated swagger"
	rm -r server/docs/

tex:
	@echo ">> compiling LaTeX to CV"
	cd latex ; pdflatex cv.tex