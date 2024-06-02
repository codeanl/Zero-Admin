# 使用基础镜像，此处以Node.js为例
FROM node:20

# 设置工作目录
WORKDIR /app

# 将本地应用程序文件复制到容器中
COPY package.json package-lock.json /app/

# 安装应用程序依赖
RUN npm install

# 复制应用程序源代码到容器中
COPY . /app

# 构建应用程序（根据实际需要执行构建命令）
RUN npm run build

# 暴露容器的端口（根据实际需要修改端口号）
EXPOSE 5173

# 启动应用程序
CMD [ "npm", "run", "start" ]