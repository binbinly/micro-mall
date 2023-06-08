<?php


namespace App\Models\Market;

/**
 * 营销服务基础模型
 * Class BaseModel
 * @package App\Models\Product
 */
class BaseModel extends \App\Models\BaseModel
{
    protected $conn = parent::CONN_SMS;
}