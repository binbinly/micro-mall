## Scrapy 爬虫 填充商品数据

### 创建虚拟环境 venv

### 安装依赖包
```bash
# 到处依赖包
pip freeze > requirements.txt
# 安装依赖包
pip install -r requirements.txt
```
### 使用方法
>源站点元素可能会变化，xpath 路径可能要更新
```python
# mysql 配置获取系统环境变量,位于文件 pipelines.py 中
db_host=os.environ.get("CHAT_DB_HOST", "127.0.0.1"),
db_pwd=os.environ.get("CHAT_DB_PWD", "root"),
db_name=os.environ.get("MYSQL_DB_NAME", "mall_pms"),
db_port=os.environ.get("CHAT_DB_PORT", "3306"),
db_user=os.environ.get("CHAT_DB_USER", "root")
```
```bash
# 测试详情页数据获取
scrapy crawl example
# 启动
scrapy crawl jijiyouxuan
```
### 其他
[文档地址](https://www.osgeo.cn/scrapy/topics/media-pipeline.html)