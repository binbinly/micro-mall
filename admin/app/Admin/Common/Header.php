<?php


namespace App\Admin\Common;


use Encore\Admin\Grid;

/**
 * 自定义列表头
 * Trait Header
 * @package AdminBase\Traits
 */
trait Header
{
    /**
     * 列表头定义
     * @param Grid $grid
     */
    protected function header(Grid &$grid)
    {
        $grid->header(function ($query) {
            return $this->headerBody($query);
        });
    }

    /**
     * 控制器覆盖操作-自定义列表头部
     * @param $query
     * @return mixed
     */
    abstract protected function headerBody($query);
}