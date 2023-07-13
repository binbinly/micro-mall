import scrapy
from spider.util import *


class JijiyouxuanSpider(scrapy.Spider):
    name = 'jijiyouxuan'
    allowed_domains = ['jijiyouxuan.com']
    start_urls = ['https://www.jijiyouxuan.com/index.php?s=/index/brand/brandlist.html']
    # 商品列表是异步从接口获取的
    product_list_api = 'https://www.jijiyouxuan.com/index.php?s=/index/search/goodlistnew.html'

    # 以品牌街为入口
    def parse(self, response):
        # 全站点通用属性分组
        yield build_attr_group(1, '基础信息')
        yield build_attr_group(2, '生产信息')

        # 品牌
        for quote in response.xpath('//div[@class="brand_item"]'):
            brand = BrandItem()
            brand['id'] = quote.xpath('div/div[@class="brand_item_R_btn"]/a/@href').re_first(
                r'brand_id/(.*).html')
            brand['name'] = quote.xpath('div/div[@class="brand_item_R_title"]/p/text()').get()
            brand['desc'] = quote.xpath('div/div[@class="brand_item_R_text"]/p/text()').get()
            brand['logo'] = split_img_url(quote.xpath('div/div[@class="brand_item_R_logo"]/img/@src').get())
            brand['cover'] = split_img_url(quote.xpath('div[@class="brand_item_img"]/img/@src').get())
            if brand['desc'] is None:
                brand['desc'] = ''

            image = ImageItem()
            image['image_url'] = brand['logo']
            yield image
            image = ImageItem()
            image['image_url'] = brand['cover']
            yield image

            yield brand

        # 三级商品分类
        for quote in response.xpath('//div[@id="iiii_header"]/div'):
            # 根分类
            cat = CatItem()
            root_cat_id = quote.xpath('a[@class="head_nal_itemdt header_fldhout header_fldhtxt"]/@href').re_first(
                r'category_id/(.*).html')
            cat['id'] = root_cat_id
            cat['name'] = quote.xpath('a[@class="head_nal_itemdt header_fldhout header_fldhtxt"]/text()').get()
            cat['parent_id'] = 0
            cat['level'] = 1
            yield cat

            for qt in quote.xpath('div//ul/li'):
                # 二级分类
                cat = CatItem()
                cid = qt.xpath('a/@href').re_first(r'category_id/(.*).html')
                cat['id'] = cid
                cat['name'] = qt.xpath('a/p/text()').get()
                cat['parent_id'] = root_cat_id
                cat['level'] = 2
                yield cat

                for q in qt.xpath('div/a'):
                    # 三级分类
                    cid_child = q.xpath('@href').re_first(r'category_id/(.*).html')
                    cat = CatItem()
                    cat['id'] = cid_child
                    cat['name'] = q.xpath('text()').get()
                    cat['parent_id'] = cid
                    cat['level'] = 3
                    yield cat

                    # 进入分类商品列表
                    yield scrapy.FormRequest(
                        url=self.product_list_api,
                        formdata={"category_id": str(cid_child)},
                        headers={"X-Requested-With": "XMLHttpRequest"},  # 此header头必须，否则获取不到数据
                        callback=self.product_list)
                    # break

    def product_list(self, response):
        res = json.loads(response.body)
        if res['data'] == '' or 'list' not in res['data'] or len(res['data']['list']) == 0:
            return
        for item in res['data']['list']:
            yield scrapy.Request(url=item['detail_url'], callback=self.product_detail)

    def product_detail(self, response):
        data = parse_detail_spu(response)

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
