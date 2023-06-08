<?php


namespace App\Models\Member;


/**
 * 会员服务基础模型
 * Class BaseModel
 * @package App\Models\member
 */
class BaseModel extends \App\Models\BaseModel
{
    protected $conn = parent::CONN_UMS;
}