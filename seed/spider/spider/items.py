# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html

from scrapy import Field, Item


# 商品三级分类
class CatItem(Item):
    id = Field()
    name = Field()
    parent_id = Field()
    level = Field()
    is_release = Field()


# 品牌
class BrandItem(Item):
    id = Field()
    name = Field()
    desc = Field()
    logo = Field()
    cover = Field()
    is_release = Field()
    created_at = Field()
    updated_at = Field()


# spu信息
class SpuItem(Item):
    id = Field()
    cat_id = Field()
    brand_id = Field()
    name = Field()
    desc = Field()
    weight = Field()
    status = Field()
    is_many = Field()
    original_price = Field()
    created_at = Field()
    updated_at = Field()


# spu介绍
class SpuDescItem(Item):
    spu_id = Field()
    content = Field()


# spu图集
class SpuImageItem(Item):
    spu_id = Field()
    name = Field()
    img = Field()
    is_default = Field()


# sku信息
class SkuItem(Item):
    id = Field()
    spu_id = Field()
    cat_id = Field()
    brand_id = Field()
    name = Field()
    desc = Field()
    cover = Field()
    title = Field()
    subtitle = Field()
    price = Field()
    sale_count = Field()
    attr_value = Field()


# sku销售属性
class SkuAttrItem(Item):
    sku_id = Field()
    attr_id = Field()
    attr_name = Field()
    attr_value = Field()


# 规格属性
class AttrValueItem(Item):
    spu_id = Field()
    attr_id = Field()
    attr_name = Field()
    attr_value = Field()


# sku图集
class SkuImageItem(Item):
    sku_id = Field()
    img = Field()
    is_default = Field()


# 所有属性
class AttrItem(Item):
    id = Field()
    cat_id = Field()
    name = Field()
    values = Field()
    is_release = Field()
    is_search = Field()
    type = Field()
    created_at = Field()
    updated_at = Field()


# 属性分组
class AttrGroupItem(Item):
    id = Field()
    cat_id = Field()
    name = Field()


# 属性分组关联属性
class AttrRelGroupItem(Item):
    attr_id = Field()
    group_id = Field()


# 所有图片
class ImageItem(Item):
    image_url = Field()
