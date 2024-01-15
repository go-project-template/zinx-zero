# docker build gomodd
gomodd:
	cd deploy/gomodd && docker build -t gomodd .

# start project run env
docker_env:
	docker compose -p example -f docker-compose-env.yml up -d

# start project
docker_project:
	docker compose -p example up -d

# start docker
start: docker_env docker_project
	echo "Start success."

# start all
all: gomodd docker_env docker_project
	echo "Start success."
