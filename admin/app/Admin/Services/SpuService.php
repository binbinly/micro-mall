<?php


namespace App\Admin\Services;


use App\Admin\Common\Format;
use App\Models\BaseModel;
use App\Models\Market\SpuBoundsModel;
use App\Models\Product\AttrValueModel;
use App\Models\Product\BrandModel;
use App\Models\Product\CategoryModel;
use App\Models\Product\SkuAttrModel;
use App\Models\Product\SkuImageModel;
use App\Models\Product\SkuModel;
use App\Models\Product\SpuDescModel;
use App\Models\Product\SpuImageModel;
use App\Models\Product\SpuModel;
use App\Models\Ware\WareSkuModel;
use Elasticsearch;

/**
 * spu商品服务类
 * Class SpuService
 * @package App\Admin\Services
 */
class SpuService
{

    /**
     * 创建商品
     * @param $data
     * @param string $err
     * @return bool
     */
    public function create($data, &$err = '')
    {
        BaseModel::transaction(BaseModel::CONN_PMS);
        try {
            //1，保存spu基本信息 pms_spu
            $spuId = $this->saveSpu($data);

            //2，保存spu的描述图片 pms_spu_desc
            $this->saveSpuDesc($data, $spuId);

            //3，保存spu的图片集 pms_spu_images
            $banners = json_decode($data['basic']['banners'], true);
            $images = $this->parseSpuImage($banners, $spuId, 'spu_id');
            if (!SpuImageModel::query()->insert($images)) {//回滚
                BaseModel::rollBack(BaseModel::CONN_PMS);
                $err = '批量插入spu图集失败';
                return false;
            }

            //4，保存spu的规格参数 pms_attr_value
            $baseAttrs = $this->parseSpuAttr($data, $spuId);
            if (!AttrValueModel::query()->insert($baseAttrs)) {//回滚
                BaseModel::rollBack(BaseModel::CONN_PMS);
                $err = '批量插入spu规格属性失败';
                return false;
            }

            //5，保存spu的积分信息 sms_spu_bounds
            $this->saveSpuBounds($data, $spuId);

            //6，保存当前spu对应的所有sku信息
            //6.1，sku的基本信息 pms_sku
            //6.2，sku的图片信息 pms_sku_images
            //6.3，sku的销售属性信息 pms_sku_attr
            //6.4，sku的优惠，满减信息 sms_sku_ladder sms_sku_full sms_member_price
            $skuData = json_decode($data['sku']['sku'], true);
            $skuImage = [];
            $skuAttrs = [];
            if ($skuData['type'] == 'many') {// 多规格
                foreach ($skuData['sku'] as $item) {
                    $sku = new SkuModel();
                    $sku->spu_id = $spuId;
                    $sku->cat_id = $data['basic']['cat_id'];
                    $sku->brand_id = $data['basic']['brand_id'];
                    $sku->name = $item['name'];
                    $sku->desc = $data['basic']['desc'];
                    $sku->title = $item['title'];
                    $sku->subtitle = $item['subtitle'];
                    $sku->cover = $banners[0];
                    $sku->price = Format::amountToPenny($item['price']);
                    $sku->saveOrFail();

                    //TODO sku信息暂时只有基本信息,图片集直接取spu图集
                    $skuImage = array_merge($skuImage, $this->parseSpuImage($banners, $sku->id, 'sku_id'));

                    foreach (array_keys($item) as $val) {
                        if (is_numeric($val)) {//sku属性
                            $skuAttrs[] = [
                                'sku_id' => $sku->id,
                                'attr_id' => $val,
                                'attr_name' => $skuData['attrIds'][$val],
                                'attr_value' => $item[$val]
                            ];
                        }
                    }
                }
                if (!SkuImageModel::query()->insert($skuImage)) {//回滚
                    BaseModel::rollBack(BaseModel::CONN_PMS);
                    $err = '批量插入spu规格属性失败';
                    return false;
                }
                if (!SkuAttrModel::query()->insert($skuAttrs)) {//回滚
                    BaseModel::rollBack(BaseModel::CONN_PMS);
                    $err = '批量插入spu规格属性失败';
                    return false;
                }
            } else {//单规格
                $err = '暂不支持单规格';
                BaseModel::rollBack(BaseModel::CONN_PMS);
                return false;
            }

            BaseModel::commit(BaseModel::CONN_PMS);
            return true;
        } catch (\Exception $e) {
            BaseModel::rollBack(BaseModel::CONN_PMS);
            $err = $e->getMessage();
//            throw new \Exception($e);
            return false;
        } catch (\Throwable $e) {
            BaseModel::rollBack(BaseModel::CONN_PMS);
            $err = $e->getMessage();
//            throw new \Exception($e);
            return false;
        }
    }

    /**
     * 商品上线至 es
     * @param $id
     */
    public static function online($id)
    {
        //spu信息
        $spu = SpuModel::query()->findOrFail($id);
        //品牌信息
        $brand = BrandModel::query()->findOrFail($spu->brand_id);
        //分类信息
        $cat = CategoryModel::query()->findOrFail($spu->cat_id);
        //spu下所有可被搜索的规格属性
        $attrs = AttrValueModel::query()->where('spu_id', $id)->whereHas('attr', function ($query) {
            $query->where('is_search', 1);
        })->get();
        //获取spu下所有sku
        $skus = SkuModel::query()->where('spu_id', $id)->get();
        //库存
        $stocks = WareSkuModel::query()->whereIn('sku_id', array_column($skus->toArray(), 'id'))->get()->toArray();
        $stocks = array_column($stocks, null, 'sku_id');

        $skuAttrs = [];
        if (count($attrs) > 0) {
            /** @var AttrValueModel $attr */
            foreach ($attrs as $attr) {
                $skuAttrs[] = [
                    'attrId' => $attr->attr_id,
                    'attrName' => $attr->attr_name,
                    'attrValue' => $attr->attr_value
                ];
            }
        }
        //组合数据添加至 es
        $data['index'] = 'product';
        /** @var SkuModel $sku */
        foreach ($skus as $sku) {
            $data['body'][] = array(
                'create' => array( #注意create也可换成index
                    '_id' => $sku->id
                ),
            );
            $data['body'][] = [
                'skuId' => $sku->id,
                'spuId' => $id,
                'skuTitle' => $sku->title,
                'skuPrice' => $sku->price,
                'skuImg' => $sku->cover,
                'saleCount' => $sku->sale_count,
                'hasStock' => isset($stocks[$sku->id]) ? $stocks[$sku->id]['stock'] - $stocks[$sku->id]['stock_lock'] > 0 : false,
                'hotScore' => 0,
                'brandId' => $sku->brand_id,
                'catId' => $sku->cat_id,
                'catName' => $cat->name,
                'brandName' => $brand->name,
                'brandImg' => $brand->cover,
                'attrs' => $skuAttrs
            ];
        }
        $ret = Elasticsearch::bulk($data);
        if (!$ret['errors']) {
            return true;
        }
        return false;
    }

    /**
     * 商品下线 删除 es 中的数据
     * @param $id
     */
    public static function offline($id)
    {
        //获取spu下所有sku
        $skus = SkuModel::query()->where('spu_id', $id)->pluck('id')->toArray();
        $data['index'] = 'product';
        foreach ($skus as $skuId) {
            $data['body'][] = array(
                'delete' => array( #注意create也可换成index
                    '_id' => $skuId
                ),
            );
        }
        $ret = Elasticsearch::bulk($data);
        if (!$ret['errors']) {
            return true;
        }
        return false;
    }

    /**
     * 报错spu信息
     * @param $data
     * @return int
     * @throws \Throwable
     */
    protected function saveSpu($data)
    {
        $spu = new SpuModel();
        $spu->name = $data['basic']['name'];
        $spu->desc = $data['basic']['desc'];
        $spu->cat_id = $data['basic']['cat_id'];
        $spu->brand_id = $data['basic']['brand_id'];
        $spu->weight = $data['basic']['weight'];
        $spu->status = SpuModel::STATUS_INIT;
        $spu->saveOrFail();
        return $spu->id;
    }

    /**
     * 报错spu描述图片
     * @param $data
     * @param $spuId
     * @throws \Throwable
     */
    protected function saveSpuDesc($data, $spuId)
    {
        $spuDesc = new SpuDescModel();
        $spuDesc->spu_id = $spuId;
        $spuDesc->content = $data['basic']['images'];
        $spuDesc->saveOrFail();
    }

    /**
     * 保存spu积分信息
     * @param $data
     * @param $spuId
     * @throws \Throwable
     */
    protected function saveSpuBounds($data, $spuId)
    {
        $bounds = new SpuBoundsModel();
        $bounds->spu_id = $spuId;
        $bounds->grow_bounds = $data['basic']['growth'] ?? 0;
        $bounds->buy_bounds = $data['basic']['integral'] ?? 0;
        $bounds->saveOrFail();
    }

    /**
     * 格式化spu图集
     * @param $banners
     * @param $spuId
     * @param $key
     * @return array
     */
    protected function parseSpuImage($banners, $spuId, $key)
    {
        if (!$banners) return [];
        $images = [];
        foreach ($banners as $i => $url) {
            $isDefault = 0;
            if ($i == 0) {
                $isDefault = 1;
            }
            $images[] = [
                $key => $spuId,
                'img' => $url,
                'is_default' => $isDefault,
                'sort' => $i
            ];
        }
        return $images;
    }

    /**
     * 格式化spu规格属性结构
     * @param $data
     * @param $spuId
     * @return array
     */
    protected function parseSpuAttr($data, $spuId)
    {
        $baseAttrs = [];
        $baseAttrNames = $data['attr']['id'];
        foreach ($data['attr'] as $key => $attr) {
            if ($key == 'id') {
                continue;
            }
            $attrId = intval(substr($key, -1, 1));
            $baseAttrs[] = [
                'spu_id' => $spuId,
                'attr_id' => $attrId,
                'attr_name' => $baseAttrNames[$attrId] ?? '',
                'attr_value' => join(',', array_filter($attr))
            ];
        }
        return $baseAttrs;
    }
}