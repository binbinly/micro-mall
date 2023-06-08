<?php


namespace App\Admin\Common;


class Format
{
    /**
     * 格式化成分
     * @param $value
     * @return int
     */
    public static function amountToPenny($value)
    {
        return intval($value * 100);
    }

    /**
     * 格式化为元
     * @param $value
     * @param bool $format 是否格式化为数字字符串
     * @return float|int
     */
    public static function amountToYuan($value, $format = false)
    {
        if ($value == 0) return 0;
        if ($format) {
            return number_format(round($value / 100, 2));
        } else {
            return round($value / 100, 2);
        }
    }

    /**
     * 汇总
     * @param bool $amount 是否为金额
     * @param bool $ret 直接返回
     * @return \Closure
     */
    public static function totalView($amount = false, $ret = false)
    {
        return function ($val) use ($amount, $ret) {
            if ($ret) {
                return "<span class='text-danger text-bold'>总计: {$val} </span>";
            }
            $val = Format::amountToYuan($val);
            return "<span class='text-danger text-bold'>总计: {$val} </span>";
        };
    }

    /**
     * 格式化二维数组成map结构
     * @param $list
     * @param $key
     * @param $value
     * @return array|false
     */
    public static function formatColumn(array $list, $key, $value)
    {
        return array_combine(array_column($list, $key), array_column($list, $value));
    }
}