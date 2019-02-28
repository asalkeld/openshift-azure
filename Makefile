COMMIT=$(shell git rev-parse --short HEAD)$(shell [[ $$(git status --porcelain) = "" ]] && echo -clean || echo -dirty)
CLUSTER_VERSION=$(shell awk '/^clusterVersion: /{ print $$2 }' <pluginconfig/pluginconfig-311.yaml)
# if we are on master branch we should always use dev tag
$(info CLUSTER_VERSION is ${CLUSTER_VERSION})
ifeq ($(CLUSTER_VERSION),v0.0)
  TAG := dev
else
  TAG := ${CLUSTER_VERSION}
endif
$(info TAG set to ${TAG})
LDFLAGS="-X main.gitCommit=$(COMMIT)"
E2E_IMAGE ?= quay.io/openshift-on-azure/e2e-tests:$(TAG)
AZURE_CONTROLLERS_IMAGE ?= quay.io/openshift-on-azure/azure-controllers:$(TAG)
ETCDBACKUP_IMAGE ?= quay.io/openshift-on-azure/etcdbackup:$(TAG)
METRICSBRIDGE_IMAGE ?= quay.io/openshift-on-azure/metricsbridge:$(TAG)
SYNC_IMAGE ?= quay.io/openshift-on-azure/sync:$(TAG)
STARTUP_IMAGE ?= quay.io/openshift-on-azure/startup:$(TAG)
CANARY_IMAGE ?= quay.io/openshift-on-azure/canary:$(TAG)

ALL_BINARIES = azure-controllers e2e-bin etcdbackup sync metricsbridge startup canary
ALL_IMAGES = $(addsuffix -image, $(ALL_BINARIES))
ALL_PUSHES = $(addsuffix -push, $(ALL_BINARIES))

IMAGEBUILDER = ${GOPATH}/bin/imagebuilder

# all is the default target to build everything
all: clean build $(ALL_BINARIES)

version:
	echo ${TAG}

build: generate
	go build ./...

clean:
	rm -f coverage.out e2e $(ALL_BINARIES)

test: unit e2e

generate:
	go generate ./...

create:
	timeout 1h ./hack/create.sh ${RESOURCEGROUP}

delete:
	./hack/delete.sh ${RESOURCEGROUP}
	rm -rf _data

azure-controllers: generate
	go build -ldflags ${LDFLAGS} ./cmd/$@

azure-controllers-image: azure-controllers $(IMAGEBUILDER)
	imagebuilder -f images/azure-controllers/Dockerfile -t $(AZURE_CONTROLLERS_IMAGE) .

azure-controllers-push: azure-controllers-image
	docker push $(AZURE_CONTROLLERS_IMAGE)

e2e-bin: generate
	go test -ldflags ${LDFLAGS} -tags e2e -c -o ./e2e ./test/e2e

e2e-bin-image: e2e-bin $(IMAGEBUILDER)
	imagebuilder -f images/e2e/Dockerfile -t $(E2E_IMAGE) .

e2e-bin-push: e2e-bin-image
	docker push $(E2E_IMAGE)

recoveretcdcluster: generate
	go build -ldflags ${LDFLAGS} ./cmd/recoveretcdcluster

etcdbackup: generate
	go build -ldflags ${LDFLAGS} ./cmd/etcdbackup

etcdbackup-image: etcdbackup $(IMAGEBUILDER)
	imagebuilder -f images/etcdbackup/Dockerfile -t $(ETCDBACKUP_IMAGE) .

etcdbackup-push: etcdbackup-image
	docker push $(ETCDBACKUP_IMAGE)

metricsbridge:
	go build -ldflags ${LDFLAGS} ./cmd/metricsbridge

metricsbridge-image: metricsbridge $(IMAGEBUILDER)
	imagebuilder -f images/metricsbridge/Dockerfile -t $(METRICSBRIDGE_IMAGE) .

metricsbridge-push: metricsbridge-image
	docker push $(METRICSBRIDGE_IMAGE)

sync: generate
	go build -ldflags ${LDFLAGS} ./cmd/sync

sync-image: sync $(IMAGEBUILDER)
	imagebuilder -f images/sync/Dockerfile -t $(SYNC_IMAGE) .

sync-push: sync-image
	docker push $(SYNC_IMAGE)

all-image: $(ALL_IMAGES)

all-push: $(ALL_PUSHES)

startup: generate
	go build -ldflags ${LDFLAGS} ./cmd/startup

startup-image: startup $(IMAGEBUILDER)
	imagebuilder -f images/startup/Dockerfile -t $(STARTUP_IMAGE) .

startup-push: startup-image
	docker push $(STARTUP_IMAGE)

canary:
	go build -ldflags ${LDFLAGS} ./cmd/$@

canary-image: canary $(IMAGEBUILDER)
	$(IMAGEBUILDER) -f images/canary/Dockerfile -t $(CANARY_IMAGE) .

canary-push: canary-image
	docker push $(CANARY_IMAGE)

verify:
	./hack/validate-generated.sh
	go vet ./...
	./hack/verify-code-format.sh
	./hack/validate-util.sh
	go run ./hack/validate-imports/validate-imports.go cmd hack pkg test
	go run ./hack/lint-addons/lint-addons.go -n

unit: generate
	go test ./... -coverprofile=coverage.out -covermode=atomic
ifneq ($(ARTIFACT_DIR),)
	mkdir -p $(ARTIFACT_DIR)
	cp coverage.out $(ARTIFACT_DIR)
endif

cover: unit
	go tool cover -html=coverage.out

codecov: unit
	./hack/codecov-report.sh

upgrade:
	./hack/upgrade-e2e.sh release-test-${TAG}-${COMMIT} ${SOURCE}

e2e:
	FOCUS="\[AzureClusterReader\]|\[CustomerAdmin\]|\[EndUser\]\[Fake\]" TIMEOUT=60m ./hack/e2e.sh

e2e-prod:
	FOCUS="\[Default\]\[Real\]" TIMEOUT=70m ./hack/e2e.sh

e2e-etcdbackuprecovery:
	FOCUS="\[EtcdRecovery\]\[Fake\]" TIMEOUT=70m ./hack/e2e.sh

e2e-keyrotation:
	FOCUS="\[KeyRotation\]\[Fake\]" TIMEOUT=70m ./hack/e2e.sh

e2e-reimagevm:
	FOCUS="\[ReimageVM\]\[Fake\]" TIMEOUT=10m ./hack/e2e.sh

e2e-scaleupdown:
	FOCUS="\[ScaleUpDown\]\[Fake\]" TIMEOUT=30m ./hack/e2e.sh

e2e-forceupdate:
	FOCUS="\[ForceUpdate\]\[Fake\]" TIMEOUT=70m ./hack/e2e.sh

e2e-vnet:
	FOCUS="\[Vnet\]\[Real\]" TIMEOUT=70m ./hack/e2e.sh

$(IMAGEBUILDER):
	docker pull registry.access.redhat.com/rhel7:latest
	go get -u github.com/openshift/imagebuilder/cmd/imagebuilder

.PHONY: clean verify generate unit create delete upgrade e2e all-image all-push $(ALL_PUSHES) $(ALL_IMAGES)
