## Scrapy 爬虫 填充商品数据

### 使用方法
```bash
scrapy crawl jijiyouxuan
```

### Notice
- 存储使用minio 时 需要单独安装扩展包，根据提示，否则 MinioPipeline 无法开启
- 源站点元素可能会变化，xpath 路径可能要更新

[文档地址](https://www.osgeo.cn/scrapy/topics/media-pipeline.html)