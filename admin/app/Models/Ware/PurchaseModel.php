<?php


namespace App\Models\Ware;


/**
 * 库存采购模型
 * Class PurchaseModel
 * @package App\Models\Ware
 * @property int id
 * @property int assignee_id
 * @property string assignee_name
 * @property string phone
 * @property int priority
 * @property int status
 * @property int ware_id
 * @property int amount
 */
class PurchaseModel extends BaseModel
{
    const STATUS_INIT = 0;
    const STATUS_USE = 1;
    const STATUS_RECEIVED = 2;
    const STATUS_FINISH = 3;
    const STATUS_ERR = 4;

    public static $statusLabel = [
        self::STATUS_INIT => '新建',
        self::STATUS_USE => '已分配',
        self::STATUS_RECEIVED => '已领取',
        self::STATUS_FINISH => '已完成',
        self::STATUS_ERR => '采购失败'
    ];

    protected $table = 'wms_purchase';

    /**
     * 获取可以被分配的采购单
     * @return array
     */
    public static function getAssigned(){
        return self::query()->whereIn('status', [self::STATUS_INIT, self::STATUS_USE])->pluck('priority', 'id')->toArray();
    }
}