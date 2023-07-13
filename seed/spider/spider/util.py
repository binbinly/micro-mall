import re
import requests
import json
from spider.items import *


# 解析url地址
def parse_img_url(img_url):
    index = img_url.find('images')
    if index > -1:
        return img_url[index:]
    else:
        return img_url


# 去除url ? 后的字符
def split_img_url(img_url):
    if img_url is None:
        return ""
    index = img_url.find('?')
    if index > -1:
        return img_url[0:index]
    return img_url


def parse_detail_sku(response):
    data = {"download": []}
    # spu id
    data['spu_id'] = int(response.xpath('//a[@id="copyurl"]/@data-gid').get())
    # sku_id
    data['sku_id'] = int(response.xpath('//a[@id="copyurl"]/@data-bid').get())
    # 分类
    data['cat_id'] = int(response.xpath('//div[@class="top-navlist am-hide-sm"]/a[last()-1]/@href').re_first(
        r'category_id/(.*).html'))

    data['brand_id'] = int(response.xpath('//div[@class="brand-model"]//a/@href').re_first(r'brand_id/(.*).html'))
    data['brand_name'] = response.xpath('//div[@class="brand_name"]/text()').get().strip()
    data['brand_logo'] = split_img_url(response.xpath('//div[@class="brand-model"]//img/@src').get())
    data['download'].append(data['brand_logo'])

    # sku销售属性
    data['sku_sale_attr'] = response.xpath(
        '//div[@class="theme-signin-left"]//*[contains(@class,"selected")]/@data-value').get().strip()

    # 名称
    data['sku_name'] = response.xpath('//span[@id="mode_code"]/text()').get().strip()
    # 标题
    sku_title = response.xpath('string(//div[@class="detail-title am-margin-bottom-xs"])').get()
    data['sku_title'] = sku_title.strip().replace('\t', '').replace('\r', '').replace('\n', '')
    names = data['sku_title'].split(' ')
    data['spu_name'] = names[0]
    # 副标题
    subtitle = response.xpath('//div[@class="tb-detail-hd"]/p/text()').get()
    if subtitle is None:
        data['sku_subtitle'] = ''
    else:
        data['sku_subtitle'] = subtitle.strip()
    # 价格
    price = int(response.xpath('(//b[@class="goods-price longpress_price"]/@data-original-price)[2]').get())
    data['price'] = price * 100  # 转为 分
    # 市场价
    original_price = response.xpath('//span[@class="price"]/text()').get()
    data['original_price'] = int(original_price) * 100  # 转为 分
    # 销量
    sale_count = response.xpath('(//span[@class="tm-count"]/text())[1]').get()
    data['sale_count'] = int(sale_count)

    # 产品图集
    data['covers'] = []
    for quote in response.xpath('//ul[@id="thumblist"]/li//img'):
        image_url = split_img_url(quote.xpath('@src').get())
        data['covers'].append(image_url)
        data['download'].append(image_url)
    data['cover'] = data['covers'][0]
    data['sku_sale_other_attrs'] = [0]  # 数组长度为一，多规格判断

    # 产品详情图
    data['images'] = []
    for quote in response.xpath('//div[@class="detail-content"]/p/img'):
        image_url = split_img_url(quote.xpath('@src').get())
        data['images'].append(parse_img_url(image_url))
        data['download'].append(image_url)
    # 详情附加图
    for quote in response.xpath('//div[@class="am-tabs-bd"]/img'):
        image_url = split_img_url(quote.xpath('@src').get())
        data['images'].append(parse_img_url(image_url))
        data['download'].append(image_url)

    return data


# 解析响应体数据
def parse_detail_spu(response):
    # 需要下载的图片地址
    data = parse_detail_sku(response)

    # 其他销售属性url
    data['sku_sale_other_attrs'] = []
    current_url = response.xpath(
        '//div[@class="theme-signin-left"]//*[contains(@class,"selected")]/@data-reload').get().strip()
    for quote in response.xpath('//div[@class="theme-signin-left"]//ul/a'):
        url_a = quote.xpath('@href').get()
        # 过滤当前的url
        if url_a is None or url_a == current_url:
            continue
        data['sku_sale_other_attrs'].append(url_a)
    for quote in response.xpath('//div[@class="theme-signin-left"]//ul/li'):
        url_li = quote.xpath('@data-reload').get()
        # 过滤当前的url
        if url_li is None or url_li == current_url:
            continue
        data['sku_sale_other_attrs'].append(url_li)

    # 属性一，接口获取
    data['attrs'] = get_attr(data['spu_id'], data['sku_name'])
    # 属性二，页面获取
    data['attrs2'] = []
    for quote in response.xpath('//table[@class="biao2"]//tr'):
        attr = {'name': quote.xpath('td[1]/text()').get().strip(), 'value': quote.xpath('td[2]/text()').get().strip()}
        data['attrs2'].append(attr)

    return data


# 通过接口获取基础规格属性
def get_attr(spu_id, sku_name):
    url = 'https://www.jijiyouxuan.com/index.php?s=/index/goods/skuparam.html&goods_id={}&mode_code={}'.format(
        spu_id, sku_name)
    response = requests.get(url)
    if response.status_code == 200:
        data = json.loads(response.text)
        if data['code'] == 0:
            return data['data']
    return []


# 构建品牌数据
def build_brand_item(data):
    brand = BrandItem()
    brand['id'] = data['brand_id']
    brand['name'] = data['brand_name']
    brand['logo'] = data['brand_logo']
    brand['cover'] = ''
    return brand


# 构建spu数据
def build_spu_item(data):
    spu = SpuItem()
    spu['id'] = data['spu_id']
    spu['cat_id'] = data['cat_id']
    spu['brand_id'] = data['brand_id']
    spu['name'] = data['spu_name']
    spu['original_price'] = data['original_price']
    spu['is_many'] = 1 if len(data['sku_sale_other_attrs']) > 0 else 0
    return spu


def build_spu_image_item(spu_id, url):
    spu_image = SpuImageItem()
    spu_image['spu_id'] = spu_id
    spu_image['img'] = url
    return spu_image


def build_sku_image_item(sku_id, url):
    sku_image = SkuImageItem()
    sku_image['sku_id'] = sku_id
    sku_image['img'] = url
    return sku_image


def build_spu_desc_item(data):
    spu_desc = SpuDescItem()
    spu_desc['spu_id'] = data['spu_id']
    spu_desc['content'] = json.dumps(data['images'])
    return spu_desc


def build_sku_item(data):
    sku = SkuItem()
    sku['id'] = data['sku_id']
    sku['spu_id'] = data['spu_id']
    sku['cat_id'] = data['cat_id']
    sku['brand_id'] = data['brand_id']
    sku['name'] = data['sku_name']
    sku['cover'] = data['cover']
    sku['title'] = data['sku_title']
    sku['subtitle'] = data['sku_subtitle']
    sku['price'] = data['price']
    sku['sale_count'] = data['sale_count']
    sku['attr_value'] = data['sku_sale_attr']
    return sku


def build_sku_attr_item(data):
    sku_attr = SkuAttrItem()
    sku_attr['sku_id'] = data['sku_id']
    sku_attr['attr_id'] = 1
    sku_attr['attr_name'] = '规格'
    sku_attr['attr_value'] = data['sku_sale_attr']
    return sku_attr


def build_attr_item(cat_id, attr_id, attrs):
    attr = AttrItem()
    attr['id'] = attr_id
    attr["cat_id"] = cat_id
    attr['name'] = attrs['name']
    attr['values'] = attrs['value']
    if isinstance(attrs['value'], str) and len(attrs['value']) < 10:
        attr['is_search'] = 1
    return attr


def build_attr_rel_group_item(attr_id, group_id):
    attr_rel_group = AttrRelGroupItem()
    attr_rel_group['attr_id'] = attr_id
    attr_rel_group['group_id'] = group_id
    return attr_rel_group


def build_attr_value_item(spu_id, attr_id, attrs):
    attr_value = AttrValueItem()
    attr_value['spu_id'] = spu_id
    attr_value['attr_id'] = attr_id
    attr_value['attr_name'] = attrs['name']
    attr_value['attr_value'] = attrs['value']
    return attr_value


def build_attr_group(gid, name):
    attr_group = AttrGroupItem()
    attr_group['id'] = gid
    attr_group['cat_id'] = 0
    attr_group['name'] = name
    return attr_group
