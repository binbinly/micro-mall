<?php


namespace App\Admin\Common;


use Encore\Admin\Grid;

/**
 * 自定义列表尾
 * Trait Header
 * @package AdminBase\Traits
 */
trait Footer
{
    /**
     * 列表尾定义
     * @param Grid $grid
     */
    protected function footer(Grid &$grid)
    {
        $grid->footer(function ($query) {
            return $this->footerBody($query);
        });
    }

    /**
     * 控制器覆盖操作-自定义列表尾部
     * @param $query
     * @return mixed
     */
    abstract protected function footerBody($query);
}