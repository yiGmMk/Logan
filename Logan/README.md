# loggan

## 使用docker-compose部署

部署前将LoganSite下的.env.development.example拷贝一份为.env.development
该文件为配置前端服务依赖的后端服务地址

部署使用命令

 ```sh
 docker-compose up -d
 ```

### 后端服务

本地访问地址:<http://localhost:8008/logan-web>

### 前端服务

本地访问地址:<http://localhost:3000/#/web-list>

## 参考

- maven构建应用镜像 <https://developer.aliyun.com/article/65274>
- awesome-compose,样例配置<https://github.com/docker/awesome-compose>
