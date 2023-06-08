<?php


namespace App\Models\Ware;


/**
 * 仓储服务基础模型
 * Class BaseModel
 * @package App\Models\Ware
 */
class BaseModel extends \App\Models\BaseModel
{
    protected $conn = parent::CONN_WMS;
}