#### 需要启动
<!-- grpc 服务 -->
C:\Users\Administrator\Desktop\blog_srvs\blog_srv>go run main.go
<!-- web api 服务 -->
C:\Users\Administrator\Desktop\mxshop-api\user-web>go run main.go


## 查看端口信息
netstat -anp |grep 3000

## 停止服务
kill -9 PID

## 后台运行程序
nohup ./main &