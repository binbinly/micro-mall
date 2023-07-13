import scrapy
from spider.util import *


class ExampleSpider(scrapy.Spider):
    name = "example"
    allowed_domains = ["jijiyouxuan.com"]
    start_urls = ["http://jijiyouxuan.com/index.php?s=/index/goods/index/id/159324.html"]

    def parse(self, response):
        data = parse_detail_spu(response)
        print('data', data)
        return

        yield build_brand_item(data)
        yield build_spu_item(data)
        yield build_spu_desc_item(data)
        yield build_sku_item(data)
        yield build_sku_attr_item(data)

        # spu产品图集
        for url in data['covers']:
            yield build_spu_image_item(data['spu_id'], url)
            yield build_sku_image_item(data['sku_id'], url)

        attr_id = data['spu_id'] * 100
        for attrs in data['attrs']:
            yield build_attr_item(data['cat_id'], attr_id, attrs)
            yield build_attr_rel_group_item(attr_id, 1)
            yield build_attr_value_item(data['spu_id'], attr_id, attrs)
            attr_id += 1

        for attrs in data['attrs2']:
            yield build_attr_item(data['cat_id'], attr_id, attrs)
            yield build_attr_rel_group_item(attr_id, 2)
            yield build_attr_value_item(data['spu_id'], attr_id, attrs)
            attr_id += 1

        # 下载所有图片
        for url in data['download']:
            image = ImageItem()
            image['image_url'] = url
            yield image

        # 获取其他sku
        for url in data['sku_sale_other_attrs']:
            yield scrapy.Request(url=url, callback=self.sku_other)

    def sku_other(self, response):
        data = parse_detail_sku(response)

        yield build_spu_item(data)
        yield build_spu_desc_item(data)
        yield build_sku_item(data)
        yield build_sku_attr_item(data)

        # spu产品图集
        for url in data['covers']:
            yield build_spu_image_item(data['spu_id'], url)
            yield build_sku_image_item(data['sku_id'], url)
            image = ImageItem()
            image['image_url'] = url
            yield image
