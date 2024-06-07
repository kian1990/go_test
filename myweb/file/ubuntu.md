---
date: 2024-03-21
authors:
  - kian
categories:
  - 文档
---

# **Ubuntu**

## 配置APT源
```bash
cat <<EOF >/etc/apt/sources.list
deb https://mirrors.ustc.edu.cn/ubuntu/ jammy main restricted
deb https://mirrors.ustc.edu.cn/ubuntu/ jammy-updates main restricted
deb https://mirrors.ustc.edu.cn/ubuntu/ jammy universe
deb https://mirrors.ustc.edu.cn/ubuntu/ jammy-updates universe
deb https://mirrors.ustc.edu.cn/ubuntu/ jammy multiverse
deb https://mirrors.ustc.edu.cn/ubuntu/ jammy-updates multiverse
deb https://mirrors.ustc.edu.cn/ubuntu/ jammy-backports main restricted universe multiverse
deb http://security.ubuntu.com/ubuntu/ jammy-security main restricted
deb http://security.ubuntu.com/ubuntu/ jammy-security universe
deb http://security.ubuntu.com/ubuntu/ jammy-security multiverse
EOF

apt update && apt install -y wget
```
<!-- more -->

## 安装SSH
```bash
apt install -y openssh-server
sed -i "s/#Port 22/Port 30022/g" /etc/ssh/sshd_config
sed -i "s/#PermitRootLogin prohibit-password/PermitRootLogin yes/g" /etc/ssh/sshd_config
sed -i "s/#PasswordAuthentication yes/PasswordAuthentication yes/g" /etc/ssh/sshd_config
systemctl enable --now ssh
```

## 安装MySQL5.7
[官方文档](https://dev.mysql.com/doc/mysql-secure-deployment-guide/5.7/en/secure-deployment-post-install.html)
```bash
apt install -y libaio1 libncurses5
wget https://downloads.mysql.com/archives/get/p/23/file/mysql-5.7.44-linux-glibc2.12-x86_64.tar.gz
tar zxvf mysql-5.7.44-linux-glibc2.12-x86_64.tar.gz
mv mysql-5.7.44-linux-glibc2.12-x86_64 /opt/mysql
mkdir /opt/mysql/{data,mysql-files,etc,log}

cat <<EOF >/opt/mysql/etc/my.cnf
[mysqld]
datadir=/opt/mysql/data
socket=/tmp/mysql.sock
port=33306
log-error=/opt/mysql/log/localhost.localdomain.err
user=root
secure_file_priv=/opt/mysql/mysql-files
bind-address=0.0.0.0
EOF

cat <<EOF >/usr/lib/systemd/system/mysqld.service
[Unit]
Description=MySQL Server
Documentation=man:mysqld(7)
Documentation=http://dev.mysql.com/doc/refman/en/using-systemd.html
After=network.target
After=syslog.target

[Install]
WantedBy=multi-user.target

[Service]
User=root
Group=root
Type=forking
PIDFile=/opt/mysql/mysqld.pid

# Disable service start and stop timeout logic of systemd for mysqld service.
TimeoutSec=0

# Start main service
ExecStart=/opt/mysql/bin/mysqld --defaults-file=/opt/mysql/etc/my.cnf --daemonize --pid-file=/opt/mysql/mysqld.pid $MYSQLD_OPTS

# Use this to switch malloc implementation
EnvironmentFile=-/etc/sysconfig/mysql

# Sets open_files_limit
LimitNOFILE = 5000
Restart=on-failure
RestartPreventExitStatus=1
PrivateTmp=false
EOF

/opt/mysql/bin/mysqld --defaults-file=/opt/mysql/etc/my.cnf --initialize
systemctl enable --now mysqld
cat /opt/mysql/log/localhost.localdomain.err |grep password
mysql -uroot -p
## 修改默认密码
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY 'yourpassword';
mysql> use mysql;
## 允许外部访问
mysql> update user set host = '%' where user = 'root';
mysql> FLUSH PRIVILEGES;

## 配置MySQL环境变量
cat <<EOF >/etc/profile.d/mysql.sh
export MYSQL_HOME=/opt/mysql
export CLASSPATH=.:$MYSQL_HOME/lib
export PATH=$PATH:$MYSQL_HOME/bin
EOF

source /etc/profile.d/mysql.sh
```

## 安装MySQL8.2
[官方文档](https://dev.mysql.com/doc/mysql-secure-deployment-guide/8.0/en/secure-deployment-post-install.html)
```bash
apt install -y libaio1 libncurses5
wget https://downloads.mysql.com/archives/get/p/23/file/mysql-8.2.0-linux-glibc2.28-x86_64.tar.xz
tar zxvf mysql-8.2.0-linux-glibc2.28-x86_64.tar.gz
mv mysql-8.2.0-linux-glibc2.28-x86_64 /opt/mysql
mkdir /opt/mysql/{data,mysql-files,etc,log}

cat <<EOF >/opt/mysql/etc/my.cnf
[mysqld]
datadir=/opt/mysql/data
socket=/tmp/mysql.sock
port=33306
log-error=/opt/mysql/log/localhost.localdomain.err
user=root
secure_file_priv=/opt/mysql/mysql-files
bind-address=0.0.0.0
EOF

cat <<EOF >/usr/lib/systemd/system/mysqld.service
[Unit]
Description=MySQL Server
Documentation=man:mysqld(8)
Documentation=http://dev.mysql.com/doc/refman/en/using-systemd.html
After=network.target
After=syslog.target

[Install]
WantedBy=multi-user.target

[Service]
User=root
Group=root

# Have mysqld write its state to the systemd notify socket
Type=notify

# Disable service start and stop timeout logic of systemd for mysqld service.
TimeoutSec=0

# Start main service
ExecStart=/opt/mysql/bin/mysqld --defaults-file=/opt/mysql/etc/my.cnf $MYSQLD_OPTS 

# Use this to switch malloc implementation
EnvironmentFile=-/etc/sysconfig/mysql

# Sets open_files_limit
LimitNOFILE = 10000
Restart=on-failure
RestartPreventExitStatus=1

# Set environment variable MYSQLD_PARENT_PID. This is required for restart.
Environment=MYSQLD_PARENT_PID=1
PrivateTmp=false
EOF

/opt/mysql/bin/mysqld --defaults-file=/opt/mysql/etc/my.cnf --initialize
systemctl enable --now mysqld
cat /opt/mysql/log/localhost.localdomain.err |grep password
mysql -uroot -p
## 修改默认密码
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY 'yourpassword';
mysql> use mysql;
## 允许外部访问
mysql> update user set host = '%' where user = 'root';
mysql> FLUSH PRIVILEGES;

## 配置MySQL环境变量
cat <<EOF >/etc/profile.d/mysql.sh
export MYSQL_HOME=/opt/mysql
export CLASSPATH=.:$MYSQL_HOME/lib
export PATH=$PATH:$MYSQL_HOME/bin
EOF

source /etc/profile.d/mysql.sh
```

## 安装RabbitMQ
[官方文档](https://www.rabbitmq.com/docs/install-debian)
```bash
apt install -y curl gnupg apt-transport-https
curl -1sLf "https://keys.openpgp.org/vks/v1/by-fingerprint/0A9AF2115F4687BD29803A206B73A36E6026DFCA" | sudo gpg --dearmor | sudo tee /usr/share/keyrings/com.rabbitmq.team.gpg > /dev/null
curl -1sLf https://github.com/rabbitmq/signing-keys/releases/download/3.0/cloudsmith.rabbitmq-erlang.E495BB49CC4BBE5B.key | sudo gpg --dearmor | sudo tee /usr/share/keyrings/rabbitmq.E495BB49CC4BBE5B.gpg > /dev/null
curl -1sLf https://github.com/rabbitmq/signing-keys/releases/download/3.0/cloudsmith.rabbitmq-server.9F4587F226208342.key | sudo gpg --dearmor | sudo tee /usr/share/keyrings/rabbitmq.9F4587F226208342.gpg > /dev/null

sudo tee /etc/apt/sources.list.d/rabbitmq.list <<EOF
## Provides modern Erlang/OTP releases
##
deb [signed-by=/usr/share/keyrings/rabbitmq.E495BB49CC4BBE5B.gpg] https://ppa1.novemberain.com/rabbitmq/rabbitmq-erlang/deb/ubuntu jammy main
deb-src [signed-by=/usr/share/keyrings/rabbitmq.E495BB49CC4BBE5B.gpg] https://ppa1.novemberain.com/rabbitmq/rabbitmq-erlang/deb/ubuntu jammy main

# another mirror for redundancy
deb [signed-by=/usr/share/keyrings/rabbitmq.E495BB49CC4BBE5B.gpg] https://ppa2.novemberain.com/rabbitmq/rabbitmq-erlang/deb/ubuntu jammy main
deb-src [signed-by=/usr/share/keyrings/rabbitmq.E495BB49CC4BBE5B.gpg] https://ppa2.novemberain.com/rabbitmq/rabbitmq-erlang/deb/ubuntu jammy main

## Provides RabbitMQ
##
deb [signed-by=/usr/share/keyrings/rabbitmq.9F4587F226208342.gpg] https://ppa1.novemberain.com/rabbitmq/rabbitmq-server/deb/ubuntu jammy main
deb-src [signed-by=/usr/share/keyrings/rabbitmq.9F4587F226208342.gpg] https://ppa1.novemberain.com/rabbitmq/rabbitmq-server/deb/ubuntu jammy main

# another mirror for redundancy
deb [signed-by=/usr/share/keyrings/rabbitmq.9F4587F226208342.gpg] https://ppa2.novemberain.com/rabbitmq/rabbitmq-server/deb/ubuntu jammy main
deb-src [signed-by=/usr/share/keyrings/rabbitmq.9F4587F226208342.gpg] https://ppa2.novemberain.com/rabbitmq/rabbitmq-server/deb/ubuntu jammy main
EOF

apt update
apt install -y erlang-base \
erlang-asn1 erlang-crypto erlang-eldap erlang-ftp erlang-inets \
erlang-mnesia erlang-os-mon erlang-parsetools erlang-public-key \
erlang-runtime-tools erlang-snmp erlang-ssl \
erlang-syntax-tools erlang-tftp erlang-tools erlang-xmerl

apt install -y rabbitmq-server --fix-missing

cat <<EOF >/etc/rabbitmq/rabbitmq.conf
loopback_users = none
EOF

rabbitmqctl enable_feature_flag all
rabbitmq-plugins enable rabbitmq_management
systemctl enable --now rabbitmq-server
```

## 安装Nginx
[官方文档](https://nginx.org/en/linux_packages.html)
```bash
apt install curl gnupg2 ca-certificates lsb-release ubuntu-keyring
curl https://nginx.org/keys/nginx_signing.key | gpg --dearmor | tee /usr/share/keyrings/nginx-archive-keyring.gpg >/dev/null
echo "deb [signed-by=/usr/share/keyrings/nginx-archive-keyring.gpg] http://nginx.org/packages/ubuntu `lsb_release -cs` nginx" | tee /etc/apt/sources.list.d/nginx.list
apt update
apt install -y nginx

cat <<EOF >/etc/nginx/conf.d/defaults.conf
upstream your_front {
  server 127.0.0.1:38080;
}

upstream your_back {
  server 127.0.0.1:8080;
}

server {
  listen 80;
  server_name frontdomain.com;

  location /yourstatic {
    root /opt/static/;
  }

  location ~* \.(txt)$ {
    root /opt/static/weixin;
  }

  location / {
    proxy_http_version 1.1;
    proxy_redirect off;
    proxy_pass http://your_front;
  }
  access_log /var/log/nginx/your_front.log;
}

server {
  listen 80;
  server_name backdomain.com;

  location ~^/favicon.ico$ {
    root /opt/tomcat/apps/your_back;
  }

  location / {
    proxy_http_version 1.1;
    proxy_redirect     off;
    proxy_pass http://your_back;
  }
  access_log /var/log/nginx/your_back.log;
}
EOF

systemctl enable --now nginx
```

## 安装PHP8.3
[PHP文档](https://php.watch/)
[phpMyAdmin文档](https://docs.phpmyadmin.net/zh-cn/latest/setup.html)
```bash
add-apt-repository ppa:ondrej/php
apt install -y php8.3-fpm
systemctl enable --now php8.3-fpm
mkdir /opt/www

cat <<EOF >/opt/www/phpinfo.php
<?php
phpinfo();
?>
EOF

## nginx 配置
server {
  listen 80;
  server_name localhost;
  root /opt/www;
  access_log /var/log/nginx/defaults.log;
  location / {
    autoindex on;
    autoindex_exact_size on;
    autoindex_localtime on;
  }
  location ~ \.php$ {
    include fastcgi_params;
    fastcgi_intercept_errors on;
    fastcgi_pass unix:/run/php/php8.3-fpm.sock;
    fastcgi_param SCRIPT_FILENAME $document_root/$fastcgi_script_name;
  }
}

## 配置phpMyAdmin
wget https://files.phpmyadmin.net/phpMyAdmin/5.2.1/phpMyAdmin-5.2.1-all-languages.zip
unzip phpMyAdmin-5.2.1-all-languages.zip
mv phpMyAdmin-5.2.1-all-languages /opt/www/phpMyAdmin
apt install -y php8.3-mysql php8.3-bz2 php8.3-zip php8.3-mbstring
## 注意不能使用localhost，要使用127.0.0.1
```

## 安装Tomcat
[JDK下载](https://www.oracle.com/cn/java/technologies/downloads/archive/)
[Tomcat下载](https://tomcat.apache.org/)
```bash
wget https://dlcdn.apache.org/tomcat/tomcat-8/v8.5.99/bin/apache-tomcat-8.5.99.tar.gz
tar zxvf apache-tomcat-8.5.99.tar.gz
mv apache-tomcat-8.5.99 /opt/tomcat
rm -rf /opt/tomcat/webapps/*
mkdir /opt/tomcat/apps
## 官网下载JRE1.8
https://www.java.com/en/download/manual.jsp jre-8u401-linux-x64.tar.gz
tar zxvf jre-8u401-linux-x64.tar.gz
mv jre1.8.0_401 /opt/jre

## 配置JAVA环境变量
cat <<EOF >/etc/profile.d/java.sh
export JAVA_HOME=/opt/jre
export CLASSPATH=.:$JAVA_HOME/lib
export PATH=$PATH:$JAVA_HOME/bin
EOF

source /etc/profile.d/java.sh

## 配置用户权限
vim /opt/tomcat/conf/tomcat-users.xml
## 添加下面字段
  <role rolename="tomcat"/>
  <role rolename="role1"/>
  <role rolename="manager-script"/>
  <role rolename="manager-gui"/>
  <role rolename="manager-status"/>
  <role rolename="admin-gui"/>
  <role rolename="admin-script"/>
  <user username="root" password="root" roles="manager-gui,manager-script,tomcat,admin-gui,admin-script"/>

## 允许外部访问
vim /opt/tomcat/webapps/manager/META-INF/context.xml
## 注释下面这行，webapps下其他目录也需要修改
<!--   <Valve className="org.apache.catalina.valves.RemoteAddrValve"
         allow="127\.\d+\.\d+\.\d+|::1|0:0:0:0:0:0:0:1" /> -->

## 配置根目录访问应用
## Host标签添加一行
      <Host name="localhost"  appBase="webapps"
            unpackWARs="true" autoDeploy="true">
        <Valve className="org.apache.catalina.valves.AccessLogValve" directory="logs"
               prefix="localhost_access_log" suffix=".txt"
               pattern="%h %l %u %t &quot;%r&quot; %s %b" />
            <Context docBase="/opt/tomcat/webapps/jenkins" path="/" reloadable="true"/>            
      </Host>
```

## 安装Redis
[官方文档](https://redis.io/docs/install/install-redis/install-redis-on-linux)
```bash
sudo apt install -y lsb-release curl gpg
curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list
apt update
apt install -y redis
sed -i "s/bind 127.0.0.1/# bind 127.0.0.1/g" /etc/redis/redis.conf
sed -i "s/6379/31013/g" /etc/redis/redis.conf
sed -i "s/# requirepass foobared/requirepass 123456/g" /etc/redis/redis.conf
sed -i "s/notify-keyspace-events \"\"/notify-keyspace-events Ex/g" /etc/redis/redis.conf
systemctl enable --now redis-server
```

## 安装Squid
```bash
apt install -y squid

cat <<EOF >/etc/squid/squid.conf
http_port 3128
http_port 3129 intercept
cache_mem 64 MB
maximum_object_size 10 MB
cache_dir ufs /var/spool/squid 100 16 256
access_log /var/log/squid/access.log
acl localnet src 192.168.0.0/24
http_access allow localnet
http_access deny all
EOF

systemctl enable --now squid
```

## 安装Privoxy
```bash
apt install -y privoxy

cat <<EOF >/etc/privoxy/config
user-manual /usr/share/doc/privoxy/user-manual
confdir /etc/privoxy
logdir /var/log/privoxy
actionsfile match-all.action
actionsfile default.action
actionsfile user.action
filterfile default.filter
filterfile user.filter
logfile logfile
toggle 1
enable-remote-toggle 0
enable-remote-http-toggle 0
enable-edit-actions 0
enforce-blocks 0
buffer-limit 4096
enable-proxy-authentication-forwarding 0
forwarded-connect-retries  0
accept-intercepted-requests 0
allow-cgi-request-crunching 0
split-large-forms 0
keep-alive-timeout 5
tolerate-pipelining 1
socket-timeout 300
permit-access 192.168.0.0/24
deny-access www.example.com
listen-address 0.0.0.0:8118
#forward-socks5t / 127.0.0.1:1080 .
EOF

systemctl enable --now privoxy
```

## 安装Shadowsocks-rust
```bash
wget https://github.com/shadowsocks/shadowsocks-rust/releases/download/v1.18.1/shadowsocks-v1.18.1.x86_64-unknown-linux-musl.tar.xz
tar Jxvf shadowsocks-v1.18.1.x86_64-unknown-linux-musl.tar.xz -C /usr/local/bin

## 服务端
cat <<EOF >/usr/local/etc/ssserver.json
{
  "server":"::",
  "server_port":31000,
  "password":"Xjj91oCesEfu2qwbNcMx6ELOoXV3qzYnYKFspgu5CIQ=",
  "timeout":60,
  "method":"2022-blake3-chacha20-poly1305",
  "mode":"tcp_and_udp",
  "fast_open":false,
  "ipv6_only": false,
  "ipv6_first": true
}
EOF

cat <<EOF >/usr/lib/systemd/system/ssserver.service
[Unit]
Description=Shadowsocks-rust Default Server Service
Documentation=https://github.com/shadowsocks/shadowsocks-rust
After=network.target

[Service]
Type=simple
User=root
Group=root
LimitNOFILE=32768
ExecStart=/usr/local/bin/ssserver -c /usr/local/etc/ssserver.json

[Install]
WantedBy=multi-user.target
EOF

systemctl enable --now ssserver

## 客户端
cat <<EOF >/usr/local/etc/sslocal.json
{
  "server":"your_ip",
  "server_port":31000,
  "password":"Xjj91oCesEfu2qwbNcMx6ELOoXV3qzYnYKFspgu5CIQ=",
  "timeout":60,
  "method":"2022-blake3-chacha20-poly1305",
  "mode":"tcp_and_udp",
  "fast_open":false,
  "local_address":"0.0.0.0",
  "local_port":1080
}
EOF

cat <<EOF >/usr/lib/systemd/system/sslocal.service
[Unit]
Description=Shadowsocks-rust Default Server Service
Documentation=https://github.com/shadowsocks/shadowsocks-rust
After=network.target

[Service]
Type=simple
User=root
Group=root
LimitNOFILE=32768
ExecStart=/usr/local/bin/sslocal -c /usr/local/etc/sslocal.json

[Install]
WantedBy=multi-user.target
EOF

systemctl enable --now sslocal

cat <<EOF >proxy.sh
export http_proxy=http://localhost:1080
export https_proxy=http://localhost:1080
EOF

cat <<EOF >unproxy.sh
unset http_proxy
unset https_proxy
EOF

## 启用代理
source proxy.sh
## 禁用代理
source unproxy.sh
```

## 安装Gitlab
```bash
apt install -y curl openssh-server ca-certificates postfix
curl -sS https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.deb.sh | sudo bash
apt install -y gitlab-ce
gitlab-ctl reconfigure
systemctl enable --now gitlab-runsvdir
## 修改Gitlab登录用户root的密码，最少8位密码
gitlab-rails console
user = User.where(id: 1).first
user.password = 'password'
user.password_confirmation = 'password'
user.save!
exit
## 修改域名
vim /etc/gitlab/gitlab.rb
gitlab-ctl reconfigure
```

### 查看使用的端口
```bash
netstat -tunlp
tcp        0      0 127.0.0.1:9229          0.0.0.0:*               LISTEN      2304/gitlab-workhor 
tcp        0      0 127.0.0.1:9236          0.0.0.0:*               LISTEN      2295/gitaly          
tcp        0      0 0.0.0.0:80              0.0.0.0:*               LISTEN      2317/nginx: master  
tcp        0      0 127.0.0.1:9121          0.0.0.0:*               LISTEN      2347/redis_exporter 
tcp        0      0 127.0.0.1:9090          0.0.0.0:*               LISTEN      2357/prometheus     
tcp        0      0 127.0.0.1:9093          0.0.0.0:*               LISTEN      2372/alertmanager   
tcp        0      0 127.0.0.1:9100          0.0.0.0:*               LISTEN      2335/node_exporter  
tcp        0      0 127.0.0.1:9187          0.0.0.0:*               LISTEN      2382/postgres_expor 
tcp        0      0 127.0.0.1:9168          0.0.0.0:*               LISTEN      2342/ruby           
tcp        0      0 127.0.0.1:8080          0.0.0.0:*               LISTEN      1884/puma 6.4.0 (un 
tcp        0      0 127.0.0.1:8082          0.0.0.0:*               LISTEN      1915/sidekiq_export 
tcp        0      0 127.0.0.1:8092          0.0.0.0:*               LISTEN      1913/sidekiq 7.1.6  
tcp        0      0 127.0.0.1:8150          0.0.0.0:*               LISTEN      1818/gitlab-kas     
tcp        0      0 127.0.0.1:8151          0.0.0.0:*               LISTEN      1818/gitlab-kas     
tcp        0      0 127.0.0.1:8153          0.0.0.0:*               LISTEN      1818/gitlab-kas     
tcp        0      0 127.0.0.1:8154          0.0.0.0:*               LISTEN      1818/gitlab-kas     
tcp        0      0 127.0.0.1:8155          0.0.0.0:*               LISTEN      1818/gitlab-kas     
tcp        0      0 0.0.0.0:8060            0.0.0.0:*               LISTEN      2317/nginx: master  
tcp6       0      0 :::9094                 :::*                    LISTEN      2372/alertmanager   
udp6       0      0 :::9094                 :::*                                2372/alertmanager
```

## 安装Prometheus
[官方文档](https://prometheus.io/download)
```bash
cd /opt/prometheus

cat <<EOF >/usr/lib/systemd/system/prometheus.service
[Unit]
Description=Prometheus Server
Documentation=https://prometheus.io/docs/introduction/overview/
After=network.target

[Service]
Type=simple
ExecStart=/opt/prometheus/prometheus --config.file=/opt/prometheus/prometheus.yml --web.listen-address=:39090
Restart=always

[Install]
WantedBy=multi-user.target
EOF

cat <<EOF >/usr/lib/systemd/system/alertmanager.service
[Unit]
Description=Alertmanager
After=network.target

[Service]
User=root
ExecStart=/opt/prometheus/alertmanager/alertmanager --config.file=/opt/prometheus/alertmanager/alertmanager.yml --cluster.listen-address="0.0.0.0:39094" --web.listen-address=:39093

[Install]
WantedBy=default.target
EOF

cat <<EOF >/usr/lib/systemd/system/blackbox_exporter.service
[Unit]
Description=Blackbox Exporter
After=network.target

[Service]
User=root
ExecStart=/opt/prometheus/blackbox_exporter/blackbox_exporter --config.file=/opt/prometheus/blackbox_exporter/blackbox.yml --web.listen-address=:39115

[Install]
WantedBy=default.target
EOF

cat <<EOF >/usr/lib/systemd/system/consul_exporter.service
[Unit]
Description=Consul Exporter
After=network.target

[Service]
User=root
ExecStart=/opt/prometheus/consul_exporter/consul_exporter --web.listen-address=:39107

[Install]
WantedBy=default.target
EOF

cat <<EOF >/usr/lib/systemd/system/graphite_exporter.service
[Unit]
Description=Graphite Exporter
After=network.target

[Service]
User=root
ExecStart=/opt/prometheus/graphite_exporter/graphite_exporter --graphite.listen-address=":39109" --web.listen-address=:39108

[Install]
WantedBy=default.target
EOF

cat <<EOF >/usr/lib/systemd/system/memcached_exporter.service
[Unit]
Description=Memcached Exporter
After=network.target

[Service]
User=root
ExecStart=/opt/prometheus/memcached_exporter/memcached_exporter --web.listen-address=:39150

[Install]
WantedBy=default.target
EOF

cat <<EOF >/usr/lib/systemd/system/node_exporter.service
[Unit]
Description=Node Exporter
After=network.target

[Service]
User=root
ExecStart=/opt/prometheus/node_exporter/node_exporter --web.listen-address=:39100

[Install]
WantedBy=default.target
EOF

cat <<EOF >/usr/lib/systemd/system/statsd_exporter.service
[Unit]
Description=Statsd Exporter
After=network.target

[Service]
User=root
ExecStart=/opt/prometheus/statsd_exporter/statsd_exporter --statsd.listen-udp=":39125" --statsd.listen-tcp=":39125" --web.listen-address=:39102

[Install]
WantedBy=default.target
EOF

vim /opt/prometheus/prometheus.yml
# 添加监控
# my global config
global:
  scrape_interval: 15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.
  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.
  # scrape_timeout is set to the global default (10s).

# Alertmanager configuration
alerting:
  alertmanagers:
    - static_configs:
        - targets:
           - alertmanager:39093

# Load rules once and periodically evaluate them according to the global 'evaluation_interval'.
rule_files:
  # - "first_rules.yml"
  # - "second_rules.yml"

# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
  - job_name: "prometheus"

    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.

    static_configs:
      - targets: ["192.168.254.100:39090"]

  - job_name: 'blackbox_exporter'
    static_configs:
      - targets: ['192.168.254.100:39115']
  - job_name: 'consul_exporter'
    static_configs:
      - targets: ['192.168.254.100:39107']
  - job_name: 'graphite_exporter'
    static_configs:
      - targets: ['192.168.254.100:39108']
  - job_name: 'memcached_exporter'
    static_configs:
      - targets: ['192.168.254.100:39150']
  - job_name: 'node_exporter'
    static_configs:
      - targets: ['192.168.254.100:39100']
  - job_name: 'statsd_exporter'
    static_configs:
      - targets: ['192.168.254.100:39102']

systemctl enable --now prometheus
systemctl enable --now alertmanager
systemctl enable --now blackbox_exporter
systemctl enable --now consul_exporter
systemctl enable --now graphite_exporter
systemctl enable --now memcached_exporter
systemctl enable --now node_exporter
systemctl enable --now statsd_exporter

systemctl restart prometheus
```

## 安装Grafana
[http://localhost:3000](http://localhost:3000)
```bash
apt install -y apt-transport-https software-properties-common wget
mkdir -p /etc/apt/keyrings/
wget -q -O - https://apt.grafana.com/gpg.key | gpg --dearmor | tee /etc/apt/keyrings/grafana.gpg > /dev/null
echo "deb [signed-by=/etc/apt/keyrings/grafana.gpg] https://apt.grafana.com stable main" | tee -a /etc/apt/sources.list.d/grafana.list
apt update
apt install -y grafana
systemctl enable --now grafana-server
```

## 安装Hadoop2.10.2
[JDK下载](https://www.oracle.com/java/technologies/downloads/archive)
[Hadoop下载](https://archive.apache.org/dist/hadoop/common)
[http://localhost:50070](http://localhost:50070)
```bash
wget https://archive.apache.org/dist/hadoop/common/hadoop-2.10.2/hadoop-2.10.2.tar.gz
tar zxvf hadoop-2.10.2.tar.gz
mv hadoop-2.10.2 /opt/hadoop
tar zxvf jdk-8u391-linux-x64.tar.gz -C /opt/

vim /opt/hadoop/etc/hadoop/hadoop-env.sh
JAVA_HOME=/opt/jdk1.8.0_391

vim /opt/hadoop/etc/hadoop/hdfs-site.xml
<configuration>
    <property>
        <name>dfs.replication</name>
        <value>1</value>
    </property>
</configuration>

vim /opt/hadoop/etc/hadoop/core-site.xml
<configuration>
    <property>
        <name>fs.defaultFS</name>
        <value>hdfs://localhost:9000</value>
    </property>
    <property>
        <name>hadoop.proxyuser.root.hosts</name>
        <value>*</value>
    </property>
    <property>
        <name>hadoop.proxyuser.root.users</name>
        <value>*</value>
    </property>
    <property>
        <name>hadoop.tmp.dir</name>
        <value>/opt/hadoop/data</value>
    </property>
</configuration>

## hadoop3.3.6需要加入以下配置允许root运行
vim /opt/hadoop/sbin/start-dfs.sh
vim /opt/hadoop/sbin/stop-dfs.sh
# limitations under the License.
HDFS_DATANODE_USER=root
HADOOP_SECURE_DN_USER=hdfs
HDFS_NAMENODE_USER=root
HDFS_SECONDARYNAMENODE_USER=root

vim /opt/hadoop/sbin/start-yarn.sh
vim /opt/hadoop/sbin/stop-yarn.sh
# limitations under the License.
YARN_RESOURCEMANAGER_USER=root
HADOOP_SECURE_DN_USER=yarn
YARN_NODEMANAGER_USER=root

ssh-keygen -t rsa -P ""
cat $HOME/.ssh/id_rsa.pub >> $HOME/.ssh/authorized_keys

/opt/hadoop/bin/hdfs namenode -format
/opt/hadoop/sbin/start-all.sh
/opt/hadoop/sbin/stop-all.sh
```

## 安装Hive2.3.9
[guava下载](https://repo1.maven.org/maven2/com/google/guava/guava)
[mysql-connector-java下载](https://repo1.maven.org/maven2/mysql/mysql-connector-java)
[Hive下载](https://dlcdn.apache.org/hive)
[http://localhost:10002](http://localhost:10002)
```bash
wget https://dlcdn.apache.org/hive/hive-2.3.9/apache-hive-2.3.9-bin.tar.gz
tar zxvf apache-hive-2.3.9-bin.tar.gz
mv apache-hive-2.3.9-bin /opt/hive
mv /opt/hive/conf/hive-env.sh.template /opt/hive/conf/hive-env.sh
vim /opt/hive/conf/hive-env.sh
HADOOP_HOME=/opt/hadoop

## 本地模式
cat <<EOF >/opt/hive/conf/hive-site.xml
<?xml version="1.0"?>
<?xml-stylesheet type="text/xsl" href="configuration.xsl"?>
<configuration>
<property>
  <name>hive.metastore.warehouse.dir</name>
  <value>/user/hive/warehouse</value>
</property>
<property>
  <name>hive.metastore.local</name>
  <value>true</value>
</property>
<property>
  <name>javax.jdo.option.ConnectionURL</name>
  <value>jdbc:mysql://localhost/metastore?createDatabaseIfNotExist=true</value>
</property>
<property>
  <name>javax.jdo.option.ConnectionDriverName</name>
  <value>com.mysql.cj.jdbc.Driver</value>
</property>
<property>
  <name>javax.jdo.option.ConnectionUserName</name>
  <value>root</value>
</property>
<property>
  <name>javax.jdo.option.ConnectionPassword</name>
  <value>root</value>
</property>
<property>
  <name>hive.strict.checks.cartesian.product</name>
  <value>false</value>
</property>
<property>
  <name>hive.support.quoted_identifiers</name>
  <value>none</value>
</property>
<property>
  <name>hive.support.subqueries</name>
  <value>true</value>
</property>
</configuration>
EOF

## 远程模式
cat <<EOF >/opt/hive/conf/hive-site.xml
<?xml version="1.0"?>
<?xml-stylesheet type="text/xsl" href="configuration.xsl"?>
<configuration>
<property>
  <name>hive.metastore.warehouse.dir</name>
  <value>/user/hive/warehouse</value>
</property>
<property>
  <name>javax.jdo.option.ConnectionURL</name>
  <value>jdbc:mysql://localhost:3306/metastore?createDatabaseIfNotExist=true</value>
</property>
<property>
  <name>javax.jdo.option.ConnectionDriverName</name>
  <value>com.mysql.cj.jdbc.Driver</value>
</property>
<property>
  <name>javax.jdo.option.ConnectionUserName</name>
  <value>root</value>
</property>
<property>
  <name>javax.jdo.option.ConnectionPassword</name>
  <value>root</value>
</property>
<property>
  <name>hive.metastore.local</name>
  <value>false</value>
</property>
<property>
  <name>hive.metastore.uris</name>
  <value>thrift://localhost:9083</value>
</property>
<property>
  <name>hive.strict.checks.cartesian.product</name>
  <value>false</value>
</property>
<property>
  <name>hive.support.quoted_identifiers</name>
  <value>none</value>
</property>
<property>
  <name>hive.support.subqueries</name>
  <value>true</value>
</property>
</configuration>
EOF

mysql -uroot -p
mysql> create database metastore;

/opt/hadoop/bin/hadoop fs -mkdir -p /user/hive/warehouse
/opt/hadoop/bin/hadoop fs -chmod -R 777 /tmp
/opt/hadoop/bin/hadoop fs -chmod -R 777 /user
wget https://repo1.maven.org/maven2/mysql/mysql-connector-java/8.0.30/mysql-connector-java-8.0.30.jar
cp mysql-connector-java-8.0.30.jar /opt/hive/lib
/opt/hive/bin/schematool -dbType mysql -initSchema -verbose
/opt/hive/bin/hive
/opt/hive/bin/hiveserver2
/opt/hive/bin/beeline -u jdbc:hive2://localhost:10000

## 创建启动脚本
mkdir /opt/hive/logs

vim /opt/hive/start.sh
#!/bin/bash
ps -ef|grep 'hiveserver2'|grep -v 'grep'|awk '{print $2}'|xargs kill -9
nohup /opt/hive/bin/hiveserver2 >> /opt/hive/logs/hiveserver2.log 2>&1 &
tail -f /opt/hive/logs/hiveserver2.log

vim /opt/hive/stop.sh
#!/bin/bash
ps -ef|grep 'hiveserver2'|grep -v 'grep'|awk '{print $2}'|xargs kill -9

chmod +x /opt/hive/*.sh
```

## 安装Zookeeper
[官方文档](https://dlcdn.apache.org/zookeeper)
```bash
wget https://dlcdn.apache.org/zookeeper/zookeeper-3.8.4/apache-zookeeper-3.8.4-bin.tar.gz
tar zxvf apache-zookeeper-3.8.4-bin.tar.gz
mv apache-zookeeper-3.8.4-bin /opt/zookeeper
mv /opt/zookeeper/conf/zoo_sample.cfg /opt/zookeeper/conf/zoo.cfg

vim /opt/zookeeper/conf/zoo.cfg
dataDir=/opt/zookeeper/data
4lw.commands.whitelist=mntr,conf,ruok

mkdir /opt/zookeeper/data
/opt/zookeeper/bin/zkServer.sh start
/opt/zookeeper/bin/zkServer.sh stop
```

## 安装Hbase
[官方文档](https://dlcdn.apache.org/hbase)
[http://localhost:16010](http://localhost:16010)
```bash
wget https://dlcdn.apache.org/hbase/2.5.7/hbase-2.5.7-bin.tar.gz
tar zxvf hbase-2.5.7-bin.tar.gz
mv hbase-2.5.7-bin /opt/hbase
mkdir /opt/hbase/data

vim /opt/hbase/conf/hbase-env.sh
export JAVA_HOME=/opt/jdk1.8.0_391
export HBASE_MANAGES_ZK=false

cat <<EOF >/opt/hbase/conf/hbase-site.xml
<?xml version="1.0"?>
<?xml-stylesheet type="text/xsl" href="configuration.xsl"?>
<configuration>
  <property>
    <name>hbase.rootdir</name>
    <value>hdfs://localhost:9000/hbase</value>
  </property>
  <property>
    <name>hbase.zookeeper.quorum</name>
    <value>localhost</value>
  </property>
  <property>
   <name>hbase.zookeeper.property.clientPort</name>
   <value>2181</value>
  </property>
  <property>
    <name>hbase.cluster.distributed</name>
    <value>true</value>
  </property>
  <property>
    <name>hbase.tmp.dir</name>
    <value>/opt/hbase/data</value>
  </property>
  <property>
    <name>hbase.unsafe.stream.capability.enforce</name>
    <value>false</value>
  </property>
</configuration>
EOF

/opt/hbase/bin/start-hbase.sh
/opt/hbase/bin/stop-hbase.sh
```

## 安装Solr
[官方文档](https://solr.apache.org/downloads.html)
[http://localhost:8983](http://localhost:8983)
```bash
tar zxvf solr-8.11.3.tgz
mv solr-8.11.3 /opt/solr

vim /opt/solr/bin/solr.in.sh
SOLR_JAVA_HOME="/opt/jdk1.8.0_391"
ZK_HOST="localhost:2181"
SOLR_RECOMMENDED_OPEN_FILES=65000
SOLR_RECOMMENDED_MAX_PROCESSES=65000
SOLR_ULIMIT_CHECKS=false

/opt/solr/bin/solr start -force
/opt/solr/bin/solr stop -force
```

## 安装Sqoop1.4.7
[Sqoop下载](https://archive.apache.org/dist/sqoop)
```bash
wget https://archive.apache.org/dist/sqoop/1.4.7/sqoop-1.4.7.bin__hadoop-2.6.0.tar.gz
tar zxvf sqoop-1.4.7.bin__hadoop-2.6.0.tar.gz
mv sqoop-1.4.7.bin__hadoop-2.6.0 /opt/sqoop
mv /opt/sqoop/conf/sqoop-env-template.sh /opt/sqoop/conf/sqoop-env.sh

vim /opt/sqoop/conf/sqoop-env.sh
export HADOOP_COMMON_HOME=/opt/hadoop
export HADOOP_MAPRED_HOME=/opt/hadoop/share/hadoop/mapreduce
export HBASE_HOME=/opt/hbase
export HIVE_HOME=/opt/hive
export ZOOCFGDIR=/opt/zookeeper/conf
export HIVE_CONF_DIR=/opt/hive/conf

wget https://repo1.maven.org/maven2/mysql/mysql-connector-java/8.0.30/mysql-connector-java-8.0.30.jar
cp mysql-connector-java-8.0.30.jar /opt/sqoop/lib
cp /opt/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-* /opt/sqoop/lib
cp -r /opt/hive/lib/* /opt/sqoop/lib
/opt/hadoop/bin/hadoop fs -mkdir -p /opt
/opt/hadoop/bin/hdfs dfs -copyFromLocal /opt/sqoop hdfs://localhost:9000/opt/sqoop
/opt/hadoop/bin/hadoop fs -rm -r /opt

## 测试导入MySQL数据库
/opt/sqoop/bin/sqoop list-tables \
--connect jdbc:mysql://localhost:3306/test \
--username root \
--password root

/opt/sqoop/bin/sqoop import \
--connect jdbc:mysql://localhost:3306/test \
--username root \
--password root \
--table web \
--delete-target-dir \
--hive-import \
--hive-table web \
--hive-overwrite \
--create-hive-table \
--fields-terminated-by ',' \
-m 1
```
