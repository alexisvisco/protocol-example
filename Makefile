IMAGE_VERSION := 0.0.9
BUF_VERSION := 0.56.0
IMG_NAME := test-protocol
FULL_IMAGE := docker.io/library/$(IMG_NAME)
EXEC := docker run --platform linux/amd64 --volume "`pwd`:/workspace" --workdir /workspace $(FULL_IMAGE)

build-docker-image:
	docker build --build-arg buf_version=$(BUF_VERSION) --platform linux/amd64 -t $(IMG_NAME) .


run: build-docker-image
	rm -rf gen && mkdir gen
	$(EXEC) "buf generate --template buf.gen.yaml \
&& buf build --as-file-descriptor-set -o gen/descriptor.json \
&& chmod 777 -R gen"
