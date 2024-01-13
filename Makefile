# chmod data dir to 777
chmod_data:
	sudo chmod -R 777 data

# docker build gomodd
gomodd:
	cd deploy/gomodd && docker build -t gomodd .

# start project env
docker_env:
	docker compose -p my_project -f docker-compose-env.yml up -d

# start project
docker_project:
	docker compose -p my_project up -d

# start docker
start: chmod_data gomodd docker_env docker_project
	echo "Docker start success."

# update docker
update: docker_env docker_project
	echo "Docker update success."

# restart docker
restart:
	docker compose -p my_project restart

# Stop and remove containers, networks
down:
	docker compose -p my_project down

# 生成 api 业务代码 ， 进入"服务/cmd/api/desc"目录下，执行下面命令
api:
	cd apps/$(svc)/cmd/api/desc && goctl api go -api *.api -dir ../  --style=go_zero

# 生成 rpc 业务代码 ， 进入"服务/cmd/rpc/pb"目录下，执行下面命令
#    去除proto中的json的omitempty
#    mac: sed -i "" 's/,omitempty//g' *.pb.go
#    linux: sed -i 's/,omitempty//g' *.pb.go
rpc:
	cd apps/$(svc)/cmd/rpc/pb && \
	goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=go_zero && \
	sed -i 's/,omitempty//g' *.pb.go

# 生成 model 业务代码 ， 进入"deploy/script/mysql/"目录下，执行下面命令
model:
	cd deploy/script/mysql/ && ./genModel.sh $(dbname) $(tables) && cp -r model/ ../../../apps/$(svc)/
api2:
	cd apps/usercenter/cmd/api/desc && goctl api go -api *.api -dir ../  --style=go_zero
model2:
	cd deploy/script/mysql/ && ./genModel.sh gamex user,user_auth && cp -r model/ ../../../apps/usercenter/
