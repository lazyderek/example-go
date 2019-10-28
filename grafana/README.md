# Grafana 线上服务监控

##  Install

[官方安装方法](https://grafana.com/docs/installation/debian/)

### download and install

```bash
wget https://dl.grafana.com/oss/release/grafana_6.4.3_amd64.deb
sudo apt-get install -y adduser libfontconfig1
sudo dpkg -i grafana_6.4.3_amd64.deb
```

### run grafana
```bash
sudo service grafana-server start
```

### grafana-server访问地址

- localhost:3000

- 帐号/密码：  admin/admin

首次访问grafana页面需要加载一段时间	

