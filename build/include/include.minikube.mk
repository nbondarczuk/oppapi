#
# Targets for handling minikube testing environment
#

#
# Minikube parameters
#
NODES=1
CNI=calico
MEMORY=2000
CPUS=2
CONTAINER_RUNTIME=docker

#
# Deployment image naming conveention: it must be in sych with deployment yaml file.
#
LATEST_IMAGE=$(REPOSITORY)/$(TARGET):latest

minikube/start:
	minikube start -n $(NODES) --cni=$(CNI) --memory $(MEMORY) --cpus $(CPUS) container-runtime=$(CONTAINER_RUNTIME)
	minikube addons enable ingress
	minikube addons enable ingress-dns
	minikube addons enable registry
	eval $(minikube docker-env)
	docker login
	make minikube/deploy/build

minikube/delete:
	minikube delete

minikube/status:
	minikube status

minikube/deploy/build:
	minikube image build -t $(LATEST_IMAGE) .

minikube/deploy:
	make -C build/deploy create

minikube/deploy/delete:
	make -C build/deploy delete

minikube/deploy/show:
	make -C build/deploy show

minikube/deploy/tunnel:
	minikube tunnel

minikube/deploy/test:
	make -C build/deploy test

minikube/deploy/clean:
	minikube image rm $(LATEST_IMAGE)

minikube/help:
	@echo
	@echo '*** Minikube utility targets ***'
	@echo
	@echo 'Usage:'
	@echo '    make minikube/start          Start test cluster'
	@echo '    make minikube/delete         Delete test cluster'
	@echo '    make minikube/status         Show test cluster status'
	@echo '    make minikube/deploy/build   Build deployment image(s)'
	@echo '    make minikube/deploy         Create deployment'
	@echo '    make minikube/deploy/delete  Delete deployment'
	@echo '    make minikube/deploy/show    Show deployment objects'
	@echo '    make minikube/deploy/tunnel  Make tunnel to access the service'
	@echo '    make minikube/deploy/test    Test deployment pinging it health check'
	@echo '    make minikube/deploy/clean   Clean deployment images'
