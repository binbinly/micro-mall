<?php


namespace App\Admin\Common;


use Encore\Admin\Grid;

/**
 * 搜索
 * Trait Search
 * @package AdminBase\Traits
 */
trait Search
{
    /**
     * 筛选
     * @param Grid $grid
     */
    protected function search(Grid &$grid): void
    {
        $grid->filter(function (Grid\Filter $filter) {
            $filter->disableIdFilter();
            $this->filter($filter);
        });
    }

    /**
     * 控制器覆盖操作-自定义筛选
     * @param Grid\Filter $filter
     */
    abstract protected function filter(Grid\Filter &$filter);
}
