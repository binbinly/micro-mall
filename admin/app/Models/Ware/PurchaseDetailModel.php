<?php


namespace App\Models\Ware;


/**
 * 库存采购详情
 * Class PurchaseDetailModel
 * @package App\Models\Ware
 * @property int id
 * @property int status
 */
class PurchaseDetailModel extends BaseModel
{
    //状态[0新建，1已分配，2正在采购，3已完成，4采购失败]
    const STATUS_INIT = 0;
    const STATUS_USE = 1;
    const STATUS_ING = 2;
    const STATUS_FINISH = 3;
    const STATUS_ERR = 4;

    public static $statusLabel = [
        self::STATUS_INIT => '新建',
        self::STATUS_USE => '已分配',
        self::STATUS_ING => '采购中',
        self::STATUS_FINISH => '已完成',
        self::STATUS_ERR => '采购失败'
    ];

    protected $table = 'wms_purchase_detail';
}