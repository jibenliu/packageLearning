# 生成私钥
openssl genrsa -out server.key 2048
# 生成证书
openssl req -new -x509 -key server.key -out server.crt -days 3650
# 只读权限
chmod 400 server.key