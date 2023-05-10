# Backend

本项目作为PPTCopilot的后端使用，提供了docker的部署方式

## docker部署


how-to-use:

```bash
git clone https://github.com/hughdazz/PPTCopilot-backend
cd PPTCopilot-backend/
# 构建镜像
docker build -t backend-app .
# backend-app或者其他任意名字
# 运行容器
docker run -it --name running-app -p 8080:8080 --add-host=host.docker.internal:host-gateway backend-app
# 运行mysql容器，将默认的3306接口映射到本地机3307接口
docker run -it --name some-mysql -p 3307:3306 -e MYSQL_ROOT_PASSWORD=admin -d mysql
```

经过上列步骤，现在你的宿主机里存在两个容器running-app,some-mysql
开发过程中running-app通过访问宿主机的3307端口进行数据库操作，运行以下命令可以了解更多
```bash
mysql -h host.docker.internal -P3307 -p
```

## 项目结构

标准的MVC框架，script中存了一些mysql脚本