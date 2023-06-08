<?php


namespace App\Models\Product;

/**
 * 产品服务基础模型
 * Class BaseModel
 * @package App\Models\Product
 */
class BaseModel extends \App\Models\BaseModel
{
    protected $conn = parent::CONN_PMS;
}