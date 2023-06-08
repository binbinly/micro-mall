<?php


namespace App\Models\Market;


/**
 * app设置
 * Class AppSettingModel
 * @package App\Models
 */
class AppSettingModel extends BaseModel
{
    const TYPE_SWIPER = 1;
    const TYPE_NAV = 2;
    const TYPE_THREE_ADV = 3;
    const TYPE_ONE_ADV = 4;
    const TYPE_GOODS_LIST = 5;

    const PAGE_HOME = 1;
    const PAGE_SEARCH = 2;

    public static $typeLabel = [
        self::TYPE_SWIPER => '轮播图',
        self::TYPE_NAV => '图标',
        self::TYPE_THREE_ADV => '三图广告',
        self::TYPE_ONE_ADV => '单图广告',
        self::TYPE_GOODS_LIST => '商品列表'
    ];

    public static $pageLabel = [
        self::PAGE_HOME => '首页',
        self::PAGE_SEARCH => '搜索页'
    ];

    protected $table = 'sms_app_setting';

    public function getDataAttribute($value)
    {
        try{
            return json_decode($value, true);
        }catch (\Exception $e) {
            return $value;
        }
    }

    public function setDataAttribute($value)
    {
        if (is_array($value)) {
            if (isset($value['title']) || isset($value['router'])) {
                $this->attributes['data'] = json_encode($value);
            } else {
                $this->attributes['data'] = json_encode(array_values($value));
            }
        } else {
            $this->attributes['data'] = $value;
        }
    }
}