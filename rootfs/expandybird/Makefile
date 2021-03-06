# Copyright 2016 The Kubernetes Authors All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
DOCKER_REGISTRY := gcr.io
PREFIX := $(DOCKER_REGISTRY)/$(PROJECT)
IMAGE ?= expandybird
TAG ?= git-$(shell git rev-parse --short HEAD)
FULL_IMAGE := $(PREFIX)/$(IMAGE)

.PHONY: container
container: binary expansion
	docker build -t $(FULL_IMAGE):latest -f Dockerfile .
	docker tag $(FULL_IMAGE):latest $(FULL_IMAGE):$(TAG)

.PHONY: push
push: container
ifeq ($(DOCKER_REGISTRY),gcr.io)
	gcloud docker push $(PREFIX)/$(IMAGE):$(TAG)
else
	docker push $(PREFIX)/$(IMAGE):$(TAG)
endif

expansion:
	cp -R ../../expansion ./opt

binary:
	cp ../../bin/linux-amd64/expandybird ./bin
