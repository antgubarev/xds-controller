build-echo:
	eval $(minikube docker-env) && \
	docker build -t xds-echo -f ./deployments/echo.Dockerfile .

frd-echo:
	kubectl port-forward pod/echo 8081:80

frd-admin:
	kubectl port-forward pod/echo 9901:9901

redeploy:
	kubectl delete po/echo && \
	kubectl apply -f deployments/echo.yaml
