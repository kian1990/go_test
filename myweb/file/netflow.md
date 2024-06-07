---
date: 2024-04-29
authors:
  - kian
categories:
  - 文档
---

# **Netflow**
<!-- more -->

## 安装依赖
```bash
apt install libpcap-dev libtool rrdtool librrd-dev flex libsocket6-perl librrdtool-oo-perl nginx php5.6-fpm php5.6-mbstring php5.6-gd php5.6-xml
```

## 安装 fprobe
```bash
https://sourceforge.net/projects/fprobe/
tar jxvf fprobe-1.1.tar.bz2
cd fprobe-1.1
./configure
make
make install
fprobe -i ens33 127.0.0.1:9995
tcpdump -i lo port 9995 -v
```

## 安装 nfdump
```bash
git clone https://github.com/phaag/nfdump.git
cd nfdump
./autogen.sh
./configure --enable-nfprofile --enable-nftrack --with-rrdpath=/usr/bin
make
make install
ldconfig
nfdump -V
```

## 安装 nfsen
```bash
git clone https://github.com/phaag/nfsen.git
cd nfsen
cp etc/nfsen-dist.conf etc/nfsen.conf

vim etc/nfsen.conf

$BASEDIR = "/opt/nfsen";
$HTMLDIR = "/opt/nfsen/www/";
$USER    = "www-data";
$WWWUSER  = "www-data";
$WWWGROUP = "www-data";
%sources = (
    'source' => { 'port' => '9995', 'col' => '#0000ff', 'type' => 'netflow' },
);
@plugins = (
[ 'live', 'PortTracker'], 
);

./install.pl etc/nfsen.conf

vim /etc/nginx/conf.d/nfsen.conf

server {
    listen 80;
    server_name localhost;
    access_log /var/log/nginx/nfsen.log;
    root /opt/nfsen/www;
    index nfsen.php;
    location ~ \.php$ {
      include fastcgi_params;
      fastcgi_intercept_errors on;
      fastcgi_pass unix:/run/php/php5.6-fpm.sock;
      fastcgi_param SCRIPT_FILENAME $document_root/$fastcgi_script_name;
    }
}

/opt/nfsen/bin/nfsen start

vim /usr/lib/systemd/system/nfsen.service

[Unit]
Description=NfSen Service
After=network.target

[Service]
Type=forking
PIDFile=/opt/nfsen/var/run/nfsend.pid
ExecStart=/opt/nfsen/bin/nfsen start
ExecStop=/opt/nfsen/bin/nfsen stop
Restart=on-abort

[Install]
WantedBy=multi-user.target

systemctl enable --now nfsen
systemctl enable --now nginx

cp contrib/PortTracker/PortTracker.pm  /opt/nfsen/plugins/
cp contrib/PortTracker/PortTracker.php  /opt/nfsen/www/plugins/

vim /opt/nfsen/plugins/PortTracker.pm
my $PORTSDBDIR = "/opt/nfsen/data/ports-db"

mkdir -p /opt/nfsen/data/ports-db
chown www-data:www-data /opt/nfsen/data/ports-db
nftrack -I -d /opt/nfsen/data/ports-db
/opt/nfsen/bin/nfsen reload
chown -R www-data:www-data /opt/nfsen/data/ports-db
```

## 安装 hoststats
```bash
https://sourceforge.net/projects/hoststats/
tar zxvf hoststats-1.1.5
cd hoststats-1.1.5
./install-libnfdump.sh
./configure
make
./install.sh
chown www-data.www-data -R /opt/nfsen/data/hoststats

vim /etc/ld.so.conf
/usr/lib64

ldconfig
/opt/nfsen/data/hoststats/hoststats start

vim /usr/lib/systemd/system/hoststats.service

[Unit]
Description=Hoststats Service
After=network.target

[Service]
Type=forking
ExecStart=/opt/nfsen/data/hoststats/hoststats start
ExecStop=/opt/nfsen/data/hoststats/hoststats stop
Restart=on-abort

[Install]
WantedBy=multi-user.target

systemctl enable --now hoststats
```

