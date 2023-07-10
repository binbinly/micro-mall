<?php


namespace App\Models\Product;


/**
 * spu模型
 * Class SpuModel
 * @package App\Models\Product
 * @property int id
 * @property int cat_id
 * @property int brand_id
 * @property string name
 * @property string desc
 * @property int weight
 * @property int status
 */
class SpuModel extends BaseModel
{
    const STATUS_INIT = 0;
    const STATUS_ONLINE = 1;
    const STATUS_OFFLINE = 2;

    public static array $statusLabel = [
        self::STATUS_INIT => '草稿',
        self::STATUS_ONLINE => '上线',
        self::STATUS_OFFLINE => '下线'
    ];

    protected $table = 'pms_spu';
}
