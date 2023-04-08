# 配置了golang的镜像
# docker build -t backend-app .
# 
FROM golang:1.20

WORKDIR /home/tmp
# 安装mysql-client
RUN apt-get update && apt-get install -y lsb-release && wget https://repo.mysql.com//mysql-apt-config_0.8.24-1_all.deb && export DEBIAN_FRONTEND=noninteractive && dpkg -i mysql-apt-config_0.8.24-1_all.deb && apt-get update && apt-get install -y mysql-client

WORKDIR /home/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# 翻译: 预先复制/缓存go.mod以预先下载依赖项，并且仅在后续构建中重新下载它们（如果它们发生变化）
COPY go.mod go.sum ./
# 下载bee工具以及依赖
RUN go install github.com/beego/bee/v2@latest && go mod download && go mod verify

COPY . .

# mysql -h host.docker.internal -P3307 -p

# 运行项目
# CMD ["bee run"]

