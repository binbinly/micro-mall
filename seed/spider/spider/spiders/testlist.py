import scrapy
import re
import json
from lxml import etree


class TestlistSpider(scrapy.Spider):
    name = 'testlist'
    allowed_domains = ['jijiyouxuan.com']
    start_urls = ['https://www.jijiyouxuan.com/index.php?s=/index/search/index/category_id/1221.html']

    def parse(self, response):
        r = re.search(r'category_id/(.*).html', response.url, re.M | re.I)
        cat_id = int(r.group(1))
        print('cat_id', cat_id)

        # 商品列表是异步从接口获取的
        product_list_api = 'https://www.jijiyouxuan.com/index.php?s=/index/search/goodlistnew.html'

        # 三级商品分类
        for quote in response.xpath('//div[@id="iiii_header"]/div'):
            root_cat_id = quote.xpath('a[@class="head_nal_itemdt header_fldhout header_fldhtxt"]/@href').re_first(
                r'category_id/(.*).html')
            # 进入分类商品列表
            yield scrapy.FormRequest(
                url=product_list_api,
                formdata={"category_id": str(root_cat_id)},
                headers={"X-Requested-With": "XMLHttpRequest"},
                callback=self.product_list)
            return

    def product_list(self, response):
        res = json.loads(response.body)
        html = etree.HTML(res['data']['data'])
        for quote in html.xpath('//a'):
            print('detail url', quote.get('href'))
            yield scrapy.Request(url=quote.get('href'), callback=self.product_detail)

    def product_detail(self, response):
        print('url', response.url)
