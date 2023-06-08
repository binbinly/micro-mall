# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


# useful for handling different item types with a single interface
import os
import datetime
import pymysql
from spider.items import *
from spider.util import *
import logging
import scrapy
from scrapy.pipelines.images import ImagesPipeline


# see: https://www.osgeo.cn/scrapy/topics/media-pipeline.html
class MinioPipeline(ImagesPipeline):

    def get_media_requests(self, item, info):
        if item.get('image_url'):
            return scrapy.Request(url=item["image_url"])

    def item_completed(self, results, item, info):
        image_paths = [x["path"] for ok, x in results if ok]
        if image_paths:
            item["image_url"] = image_paths[0]
        return item

    def file_path(self, request, response=None, info=None, *, item=None):
        return parse_img_url(request.url)


class ProductPipeline:

    def __init__(self, db_host, db_port, db_user, db_pwd, db_name):
        self.sub_dir = datetime.date.today().strftime('%Y%m%d')
        self.db_host = db_host
        self.db_port = db_port
        self.db_user = db_user
        self.db_pwd = db_pwd
        self.db_name = db_name

    @classmethod
    def from_crawler(cls, crawler):
        return cls(
            db_host=os.environ.get("MYSQL_HOST", "127.0.0.1"),
            db_pwd=os.environ.get("MYSQL_PWD", "root"),
            db_name=os.environ.get("MYSQL_DB_NAME", "mall_pms"),
            db_port=os.environ.get("MYSQL_PORT", "3306"),
            db_user=os.environ.get("MYSQL_USER", "root")
        )

    def open_spider(self, spider):
        self.db = pymysql.connect(
            host=self.db_host,
            user=self.db_user,
            password=self.db_pwd,
            db=self.db_name,
            charset="utf8mb4",
            cursorclass=pymysql.cursors.DictCursor,
        )
        self.cursor = self.db.cursor()

    def close_spider(self, spider):
        self.db.close()

    def process_item(self, item, spider):
        try:
            if isinstance(item, CatItem):
                self.send_db(item, 'pms_category', item['id'])
            elif isinstance(item, AttrGroupItem):
                self.send_db(item, 'pms_attr_group', item['id'])
            elif isinstance(item, AttrItem):
                self.send_db(item, 'pms_attr', item['id'])
            elif isinstance(item, SpuItem):
                self.send_db(item, 'pms_spu', item['id'])
            elif isinstance(item, BrandItem):
                item['logo'] = parse_img_url(item['logo'])
                if item['cover'] != '':
                    item['cover'] = parse_img_url(item['cover'])

                self.send_db(item, 'pms_brand', item['id'])
            elif isinstance(item, SkuItem):
                item['cover'] = parse_img_url(item['cover'])
                self.send_db(item, 'pms_sku', item['id'])
            elif isinstance(item, AttrValueItem):
                sql = 'select * from pms_attr_value where spu_id=%s and `attr_id`=%s' % (
                    item['spu_id'], item['attr_id'])
                self.cursor.execute(sql)
                if not self.cursor.fetchone():
                    self.save_db(item, 'pms_attr_value')
            elif isinstance(item, SpuImageItem):
                item['img'] = parse_img_url(item['img'])
                sql = 'select * from pms_spu_image where spu_id=%s and `img`="%s"' % (item['spu_id'], item['img'])
                self.cursor.execute(sql)
                if not self.cursor.fetchone():
                    self.save_db(item, 'pms_spu_image')
            elif isinstance(item, SkuImageItem):
                item['img'] = parse_img_url(item['img'])
                sql = 'select * from pms_sku_image where sku_id=%s and `img`="%s"' % (item['sku_id'], item['img'])
                self.cursor.execute(sql)
                if not self.cursor.fetchone():
                    self.save_db(item, 'pms_sku_image')
            elif isinstance(item, SkuAttrItem):
                sql = 'select * from pms_sku_attr where sku_id=%s and `attr_id`=%s' % (item['sku_id'], item['attr_id'])
                self.cursor.execute(sql)
                if not self.cursor.fetchone():
                    self.save_db(item, 'pms_sku_attr')
            elif isinstance(item, SpuDescItem):
                sql = 'select * from pms_spu_desc where spu_id=%s' % item['spu_id']
                self.cursor.execute(sql)
                if not self.cursor.fetchone():
                    self.save_db(item, 'pms_spu_desc')
            elif isinstance(item, AttrRelGroupItem):
                sql = 'select * from pms_attr_rel_group where attr_id=%s and group_id=%s' % (
                    item['attr_id'], item['group_id'])
                self.cursor.execute(sql)
                if not self.cursor.fetchone():
                    self.save_db(item, 'pms_attr_rel_group')
            else:
                return
        except Exception as e:
            logging.error(item)
            logging.error(e)
        return item

    def send_db(self, item, table, pri_id):
        sql = 'select * from %s where id=%s' % (table, pri_id)
        self.cursor.execute(sql)
        if not self.cursor.fetchone():
            return self.save_db(item, table)

    def save_db(self, item, table):
        keys = item.keys()
        values = tuple(item.values())
        fields = ",".join(['`' + v + '`' for v in keys])
        temp = ",".join(["%s"] * len(keys))
        sql = "INSERT INTO {} ({}) VALUES ({})".format(table, fields, temp)
        self.cursor.execute(sql, values)
        return self.db.commit()
