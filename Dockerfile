FROM golang:1.20

WORKDIR /home/tmp
# 安装mysql-client
# 换源 
RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN apt-get update && apt-get install -y lsb-release && wget https://repo.mysql.com//mysql-apt-config_0.8.24-1_all.deb && export DEBIAN_FRONTEND=noninteractive && dpkg -i mysql-apt-config_0.8.24-1_all.deb && apt-get update && apt-get install -y mysql-client

WORKDIR /home/app

# 预先复制/缓存go.mod以预先下载依赖项，并且仅在后续构建中重新下载它们（如果它们发生变化）
COPY go.mod go.sum ./
# 下载bee工具以及依赖
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn
RUN go install github.com/beego/bee/v2@latest && go mod download && go mod verify

COPY . .

# 如果没有docker-compose未传递，使用默认值host.docker.internal
ARG MYSQL_HOST=host.docker.internal
ARG MYSQL_PORT=3306
ENV MYSQL_HOST=${MYSQL_HOST}
ENV MYSQL_PORT=${MYSQL_PORT}

# 如果arg server_ip不为空，则替换配置文件中的server_ip
ARG SERVER_IP
# 运行env.py传递参数
RUN python3 env.py $SERVER_IP


CMD ["bee", "run"]