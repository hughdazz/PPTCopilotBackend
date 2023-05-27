FROM golang:1.20

WORKDIR /home/tmp
# 安装mysql-client
RUN apt-get update && apt-get install -y lsb-release && wget https://repo.mysql.com//mysql-apt-config_0.8.24-1_all.deb && export DEBIAN_FRONTEND=noninteractive && dpkg -i mysql-apt-config_0.8.24-1_all.deb && apt-get update && apt-get install -y mysql-client

WORKDIR /home/app

# 预先复制/缓存go.mod以预先下载依赖项，并且仅在后续构建中重新下载它们（如果它们发生变化）
COPY go.mod go.sum ./
# 下载bee工具以及依赖
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn
RUN go install github.com/beego/bee/v2@latest && go mod download && go mod verify

COPY . .

# 连接mysql并创建数据库now_db
RUN -h host.docker.internal -P3307 -padmin -e "create database now_db"

CMD ["bee", "run"]