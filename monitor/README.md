## a monitor
基于prometheus+grafana线上监控服务

prometheus: 高维度的时序数据库
grafana:  可视化数据监控服务

prometheus除了可以提供监控数据之外，其实还支持数据可视化的功能，但页面比较粗糙，所以这里使用grafana去监控服务。
而grafana除了提供可视化页面之外，还可以自己连接数据源拿数据，如：mysql,elastic-search,prometheus等，
但缺点也很明显，它只能去数据库拿数据，如果想从自己的项目拿数据的话，就不行了，所以这里需要prometheus去收集数据，grafana数据可视化。


### 简介
大概的流程流图如下：

![iamge](./1574243273(1).png)

上面的monitor是我们自己的监控服务，它会定时去收集其他服务（如：nsq、es等服务）的状态信息，
之后等prometheus定时去monitor拉取数据，并存储到数据库，grafana再到prometheus拿数据，并做可视化。

被监控的服务和存储数据的的服务都可以自己实现，以插件形式实现监控和上传。monitor可以新增其他的监控对象，也可以新增上传对象。


### telegraf
我的monitor服务是参照github上的开源项目telegraf写的，[传送门](https://github.com/influxdata/telegraf)

telegraf可以监控多种服务，并能向远端服务推送监控数据。被监控服务（input）和接收数据的服务（output），在telelgraf里面都只是一个插件，可以自己配置插件。
同时，telegraf能支持多种插件。telegraf的代码比较复杂，我没有看完，只参考了部分代码，做的也比较简单。


### 安装prometheus+grafana

```
# 安装grafana
sudo apt install grafana

# 安装prometheus
sudo apt install prometheus
```
或者使用docker来安装，能安装成功就可以了，只是拿来试一下水。

### monitor



