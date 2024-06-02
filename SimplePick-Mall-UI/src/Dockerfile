# 使用一个基础镜像作为构建环境
FROM node:16 as build

# 设置工作目录
WORKDIR /app

# 复制 package.json 和 package-lock.json 到工作目录
COPY package*.json ./

# 安装依赖
RUN npm install

# 复制应用程序源代码到工作目录
COPY . .

# 构建Vue应用程序
RUN npm run build

# 使用另一个基础镜像作为运行环境
FROM nginx:latest

# 复制构建后的应用程序到Nginx的默认静态文件目录
COPY --from=build /app/dist /usr/share/nginx/html

# 更新Nginx配置文件，将默认端口更改为5173
RUN sed -i 's/80/5173/g' /etc/nginx/conf.d/default.conf

# 暴露容器的5173端口
EXPOSE 5173

# 启动Nginx服务
CMD ["nginx", "-g", "daemon off;"]