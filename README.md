跟着这篇博客做的练习
> [Go Web 框架 Gin 实践](https://segmentfault.com/a/1190000013297625#articleHeader5)
> 
> 作者：煎鱼
> 
> 项目地址：https://github.com/EDDYCJY/go-gin-example

项目结构
```
gin-blog/
├── conf
├──  middleware
├── models
├── pkg
├── routers
└──  runtime 
```
- conf：用于存储配置文件
- middleware：应用中间件
- models：应用数据库模型
- pkg：第三方包
- routers 路由逻辑处理
- runtime 应用运行时数据

构建镜像
```
docker image build -t blog .
```
运行容器
```
docker container run -d --rm -d -p 8000:8000 blog
```
两阶段式构建，减少镜像大小，并指定用户运行
```
docker image build -f Dockerfile-new -t blog .
```