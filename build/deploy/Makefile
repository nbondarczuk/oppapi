#
# Targets managing life cycle of the k8s objects. They are generic enouh so
# that they can be handled by Minikube deploy and Kubernetes as well (bold assumption).
#

SRC=oppapi.yaml
PORT=30000
IP=127.0.0.1
DSN=oppapi.info

create:
	kubectl apply -f $(SRC)

delete:
	kubectl delete -f $(SRC)

show:
	kubectl get ingress,svc,pods -o wide

test: test-service test-ingress

test-service:
	curl http://${IP}:${PORT}/health -vvv
	@echo

test-ingress:
	curl --resolve "${DSN}:80:${IP}" -i http://${DSN}/health -vvv
	@echo
