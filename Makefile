IMG:=nicovillanueva/vitae
VMNAME:=vita-1
REMOTE_ENV:=DOCKER_TLS_VERIFY="1" DOCKER_HOST="tcp://34.73.0.147:2376" DOCKER_CERT_PATH="/Users/nico/.docker/machine/machines/vita-1" DOCKER_MACHINE_NAME="vita-1"

# Locally Go-run the server
run: swagger
	go run cli/main.go start -c $$PWD/cv-data.yaml -p $$PWD/latex/nsv-cv.pdf

# Generate the docs and build the image
build: swagger tex
	docker build -t $(IMG) .

# Generate docs, build and push image
release: build
	docker push $(IMG)

# Run the image remotely
deploy:
	$(REMOTE_ENV) docker pull $(IMG) ; $(REMOTE_ENV) docker rm -f vitae ;\
	$(REMOTE_ENV) docker run -dp 80:8080 --restart unless-stopped --name vitae $(IMG)

ssh:
	gcloud compute --project "api-vitae" ssh --zone "us-east1-b" "$(VMNAME)"

# Generate the docs
swagger:
	@echo ">> generating swagger"
	swag init -g server/server.go -o server/docs/

# Clean aux files
clean:
	@echo ">> cleaning LaTeX auxiliary files"
	rm latex/*.aux latex/*.log latex/*.out

# Generate Latex files
tex:
	@echo ">> compiling LaTeX to CV"
	cd latex ; pdflatex nsv-cv.tex