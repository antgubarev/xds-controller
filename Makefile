start:
	docker-compose up -d

stop:
	docker-compose down

ping1:
	curl http://localhost:8081/?echo_env_body=HOSTNAME

ping2:
	curl http://localhost:8082/?echo_env_body=HOSTNAME
