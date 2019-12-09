#elasticsearch对接canal
    利用go-canal api 获取数据库表结构，在应用层程序生成mapping文件
    对于多表关联问题
        1.在应用层做数据组装(先查出目标索引后如果其中有关联关系则在根据标记字段去查关联对象)，缺陷：数据处理量大的情况下会有延迟
        2.采用canal的适配器，缺点写xml蓝瘦
    
## [go-echarts](https://go-echarts.github.io/go-echarts/)
## [pyecharts](http://pyecharts.org/#/zh-cn/)