import scrapy
import time
import re
import json
import itertools
from urllib import parse
from spider.items import *
from spider.util import *


class JijiyouxuanSpider(scrapy.Spider):
    name = 'test'
    allowed_domains = ['jijiyouxuan.com']
    start_urls = ['https://www.jijiyouxuan.com/index.php?s=/index/goods/index/id/15360.html']
    attr_group_id = 1
    attr_group_map = {}
    attr_id = 1
    attr_map = {}
    brand_id = 1
    brand_map = {}

    def parse(self, response):
        # 产品id
        r = re.search(r'id/(.*).html', response.url, re.M | re.I)
        product_id = int(r.group(1))

        # 分类
        cat_id = response.xpath('//div[@class="top-navlist am-hide-sm"]/a[last()-1]/@href').re_first(r'category_id/(.*).html')

        # 品牌
        brand_id = response.xpath('//div[@class="brand-model"]//div[@class="btn"]/a/@href').re_first(r'brand_id/(.*).html')
        brand = BrandItem()
        brand['id'] = brand_id
        brand['name'] = response.xpath('//div[@class="brand_name"]/text()').get()
        brand['desc'] = ''
        brand['logo'] = response.xpath('//div[@class="brand-model"]//img/@src').get()
        brand['cover'] = ''
        brand['is_release'] = 1
        brand['created_at'] = int(time.time())
        brand['updated_at'] = int(time.time())
        brand['logo'] = split_img_url(brand['logo'])

        image = ImageItem()
        image['image_url'] = brand['logo']
        yield image

        yield brand

        # 名称
        name = response.xpath('//span[@id="mode_code"]/text()').get().strip()

        # 标题
        title = response.xpath('//div[@class="tb-detail-hd"]/div[1]/text()').get().strip()

        # 副标题
        subtitle = response.xpath('//div[@class="tb-detail-hd"]/p/text()').get()
        if subtitle is None:
            subtitle = ''
        else:
            subtitle = subtitle.strip()

        # 封面
        cover = response.xpath('//div[@class="am-viewport"]//li[1]/img/@src').get()
        cover = split_img_url(cover)
        image = ImageItem()
        image['image_url'] = cover
        yield image

        # 价格
        price = int(response.xpath('//b[@class="goods-price longpress_price"]/text()').get())
        price *= 100  # 转为 分

        # 市场价
        original_price = response.xpath('//div[@class="market-price"]/p[2]/text()').get()
        original_price = int(original_price[1:]) * 100  # 转为 分

        # 销量
        sale_count = response.xpath('//ul[@class="tm-ind-panel num am-hide-sm"]/li[1]//span[@class="tm-count"]/text()').get()

        # 产品图集
        for quote in response.xpath('//ul[@id="thumblist"]/li//img'):
            spuImage = SpuImageItem()
            spuImage['spu_id'] = product_id
            spuImage['img'] = quote.xpath('@src').get()
            spuImage['name'] = ''
            spuImage['is_default'] = 0
            spuImage['img'] = split_img_url(spuImage['img'])

            image = ImageItem()
            image['image_url'] = spuImage['img']
            yield image

            yield spuImage

        attr_group_id = self.attr_group_id
        group_key = '{}-基本信息'.format(cat_id)
        if group_key not in self.attr_group_map:
            attrGroup = AttrGroupItem()
            attrGroup['id'] = self.attr_group_id
            attrGroup['cat_id'] = cat_id
            attrGroup['name'] = '基本信息'
            yield attrGroup
            self.attr_group_map[group_key] = self.attr_group_id
            self.attr_group_id += 1
        else:
            attr_group_id = self.attr_group_map[group_key]

        # 规格属性
        # 第一部分规格属性 通过接口获取 特殊处理
        url = 'https://www.jijiyouxuan.com/index.php?s=/index/goods/skuparam.html&goods_id={}&cat={}&mode_code={}'.format(
            product_id, cat_id, name)
        yield scrapy.Request(url=url, callback=self.parse_attr)
        for quote in response.xpath('//div[@class="am-tabs-bd"]//tr'):
            attr_name = quote.xpath('td[1]/text()').get().strip()
            attr_value = quote.xpath('td[2]/text()').get().strip()
            attr_id = self.attr_id
            attr_key = name + '-' + cat_id
            if attr_key not in self.attr_map:
                attr = AttrItem()
                attr['id'] = self.attr_id
                attr["cat_id"] = cat_id
                attr['name'] = attr_name
                attr['values'] = attr_value
                attr['is_release'] = 1
                attr['type'] = 1
                attr['is_search'] = 0
                attr['created_at'] = int(time.time())
                attr['updated_at'] = int(time.time())
                yield attr
                self.attr_map[attr_key] = self.attr_id
                self.attr_id += 1
                # 绑定主体属性分组
                attrRelGroup = AttrRelGroupItem()
                attrRelGroup['attr_id'] = self.attr_id
                attrRelGroup['group_id'] = attr_group_id
                yield attrRelGroup
            else:  # 当前分类下已存在该属性
                attr_id = self.attr_map[attr_key]

            attrValue = AttrValueItem()
            attrValue['spu_id'] = product_id
            attrValue['attr_id'] = attr_id
            attrValue['attr_name'] = attr_name
            attrValue['attr_value'] = attr_value
            yield attrValue

        # 产品介绍
        images = []
        for quote in response.xpath('//div[@class="detail-content"]/p/img'):
            contentImg = split_img_url(quote.xpath('@src').get())
            images.append(parse_img_url(contentImg))

            image = ImageItem()
            image['image_url'] = contentImg
            yield image
        spu_desc = SpuDescItem()
        spu_desc['spu_id'] = product_id
        spu_desc['content'] = json.dumps(images)
        yield spu_desc

        # 销售属性
        sku_attrs = []
        sku_attr_values = {}
        for quote in response.xpath('//div[@class="theme-options sku-items"]'):
            attr_name = quote.xpath('div[1]/text()').get().strip()
            attr_key = attr_name + '-' + cat_id
            attr_id = self.attr_id
            if attr_key not in self.attr_map:
                attr = AttrItem()
                attr['id'] = self.attr_id
                attr["cat_id"] = cat_id
                attr['name'] = attr_name
                attr['values'] = ''
                attr['is_release'] = 1
                attr['type'] = 0
                attr['is_search'] = 1
                attr['created_at'] = int(time.time())
                attr['updated_at'] = int(time.time())
                yield attr
                self.attr_map[attr_key] = self.attr_id
                self.attr_id += 1
            else:  # 当前分类下已存在该属性
                attr_id = self.attr_map[attr_key]

            sku_attr = []
            for qt in quote.xpath('ul/li/text()'):
                value = qt.get().strip()
                if value == '':
                    continue
                sku_attr.append(value)
                sku_attr_values[value] = {'id': attr_id, 'name': attr_name}
            sku_attrs.append(sku_attr)

        spu_is_many = 0
        if len(sku_attrs) > 0:
            spu_is_many = 1
        # spu
        spu = SpuItem()
        spu['id'] = product_id
        spu['cat_id'] = cat_id
        spu['brand_id'] = brand_id
        spu['name'] = name
        spu['desc'] = ''
        spu['weight'] = 100
        spu['original_price'] = original_price
        spu['is_many'] = spu_is_many
        spu['created_at'] = int(time.time())
        spu['updated_at'] = int(time.time())
        yield spu

        if spu_is_many == 0:  # 单规格
            sku = SkuItem()
            sku['id'] = product_id * 10
            sku['spu_id'] = product_id
            sku['cat_id'] = cat_id
            sku['brand_id'] = brand_id
            sku['name'] = name
            sku['desc'] = ''
            sku['cover'] = cover
            sku['title'] = title
            sku['subtitle'] = subtitle
            sku['price'] = price
            sku['sale_count'] = sale_count
            yield sku

            # sku图集
            for quote in response.xpath('//ul[@id="thumblist"]/li//img'):
                skuImage = SkuImageItem()
                skuImage['sku_id'] = sku['id']
                skuImage['img'] = split_img_url(quote.xpath('@src').get())
                skuImage['is_default'] = 0
                yield skuImage
        else:  # 多规格
            sku_id_start = product_id * 100  # 商品起始id
            # 计算笛卡尔积
            for item in itertools.product(*sku_attrs):
                sku = SkuItem()
                sku['id'] = sku_id_start
                sku['spu_id'] = product_id
                sku['cat_id'] = cat_id
                sku['brand_id'] = brand_id
                sku['name'] = name
                sku['desc'] = ''
                sku['cover'] = cover
                sku['title'] = title + ' ' + ' '.join(item)
                sku['subtitle'] = subtitle
                sku['price'] = price
                sku['sale_count'] = sale_count
                sku['attr_value'] = ','.join(item)
                yield sku

                # sku图集
                for quote in response.xpath('//ul[@id="thumblist"]/li//img'):
                    skuImage = SkuImageItem()
                    skuImage['sku_id'] = sku_id_start
                    skuImage['img'] = split_img_url(quote.xpath('@src').get())
                    skuImage['is_default'] = 0
                    yield skuImage

                # sku 销售属性
                for attr_value in item:
                    skuAttr = SkuAttrItem()
                    skuAttr['sku_id'] = sku_id_start
                    skuAttr['attr_id'] = sku_attr_values[attr_value]['id']
                    skuAttr['attr_name'] = sku_attr_values[attr_value]['name']
                    skuAttr['attr_value'] = attr_value
                    yield skuAttr

                sku_id_start += 1

    # 接口获取规格属性
    def parse_attr(self, response):
        params = parse.parse_qs(parse.urlparse(response.url).query)
        product_id = params['goods_id'][0]
        cat_id = params['cat'][0]

        # 属性分组
        attr_group_id = self.attr_group_id
        group_key = '{}-主体'.format(cat_id)
        if group_key not in self.attr_group_map:
            attrGroup = AttrGroupItem()
            attrGroup['id'] = self.attr_group_id
            attrGroup['cat_id'] = cat_id
            attrGroup['name'] = '主体'
            yield attrGroup
            self.attr_group_map[group_key] = self.attr_group_id
            self.attr_group_id += 1
        else:
            attr_group_id = self.attr_group_map[group_key]

        res = json.loads(response.body)
        if res['code'] == 0:
            for item in res['data']:
                attr_id = self.attr_id
                attr_key = '{}-{}'.format(item['name'], cat_id)
                if attr_key not in self.attr_map:
                    attr = AttrItem()
                    attr['id'] = self.attr_id
                    attr["cat_id"] = cat_id
                    attr['name'] = item['name']
                    attr['values'] = 0 if item['value'] is None else item['value']
                    attr['is_release'] = 1
                    attr['type'] = 1
                    attr['is_search'] = 1
                    attr['created_at'] = int(time.time())
                    attr['updated_at'] = int(time.time())
                    yield attr
                    self.attr_map[attr_key] = self.attr_id
                    self.attr_id += 1
                    # 绑定主体属性分组
                    attrRelGroup = AttrRelGroupItem()
                    attrRelGroup['attr_id'] = self.attr_id
                    attrRelGroup['group_id'] = attr_group_id
                    yield attrRelGroup
                else:  # 当前分类下已存在该属性
                    attr_id = self.attr_map[attr_key]

                attrValue = AttrValueItem()
                attrValue['spu_id'] = product_id
                attrValue['attr_id'] = attr_id
                attrValue['attr_name'] = item['name']
                attrValue['attr_value'] = 0 if item['value'] is None else item['value']
                yield attrValue
