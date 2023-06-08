<?php

namespace App\Admin\Services;

use App\Models\Market\ConfigModel;
use Encore\Admin\Facades\Admin;
use Throwable;

/**
 * 配置相关
 * Class ConfigService
 * @package AdminCommon\Services
 */
class ConfigService
{
    /**
     * 配置修改保存
     * @param $data
     * @return bool
     * @throws Throwable
     */
    public static function save($data)
    {
        $ret = false;
        foreach ($data as $name => $value) {
            $ret = self::saveOnw($name, $value);
        }
        return $ret;
    }

    /**
     * 添加一条配置
     * @param string $name
     * @param $value
     * @return bool
     */
    public static function saveOnw(string $name, $value)
    {
        $model = ConfigModel::query()->where('name', $name)->first();
        if ($model instanceof ConfigModel) {
            if ($model->value == $value) return true;
            $model->update_by = Admin::user()->id;
        } else {
            $model = new ConfigModel();
            $model->name = $name;
            $model->create_by = Admin::user()->id;
        }
        $model->value = $value;
        return $model->save();
    }
}