<?php


namespace App\Models\Market;


use App\Admin\Common\Format;
use Exception;

/**
 * 配置
 * Class AppConfigModel
 * @package App\Models
 * @property string name
 * @property string value
 * @property string desc
 * @property string create_by
 * @property string update_by
 */
class ConfigModel extends BaseModel
{
    const KEY_HOME_CAT = 'app_home_cat'; //首页分类
    const KEY_PAY_LIST = 'app_pay_list';    //支付列表

    protected $table = 'sms_config';

    protected $fillable = ['name', 'value', 'create_by', 'update_by'];

    const TYPE_STRING = 1;
    const TYPE_JSON = 2;
    const TYPE_COLUMN = 3;

    /**
     * 配置加载
     * @param null $name
     * @param integer $format
     * @return array|mixed
     * @throws Exception
     */
    public static function init($name = null, $format = self::TYPE_STRING)
    {
        $list = self::query()->pluck('value', 'name')->toArray();
        if ($name) {
            if ($format == self::TYPE_STRING) {
                return $list[$name] ?? '';
            } elseif ($format == self::TYPE_JSON) {
                if (!isset($list[$name]) || !$list[$name]) return [];
                return json_decode($list[$name], true) ?: [];
            } elseif ($format == self::TYPE_COLUMN) {
                $data = $list[$name];
                if (!is_array($list[$name])) {
                    $data = json_decode($list[$name], true);
                }
                return Format::formatColumn($data, 'id', 'name');
            }
            return $list[$name] ?? '';
        }
        return $list ?: [];
    }

    /**
     * 分类map结构
     * @return array|false
     * @throws Exception
     */
    public static function cat()
    {
        return Format::formatColumn(self::init(self::KEY_HOME_CAT, self::TYPE_JSON), 'id', 'name');
    }
}