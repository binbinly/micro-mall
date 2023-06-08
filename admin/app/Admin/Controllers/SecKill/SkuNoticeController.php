<?php

namespace App\Admin\Controllers\SecKill;

use App\Admin\Controllers\BaseController;
use App\Models\Market\SeckillSessionModel;
use App\Models\Market\SeckillSkuNoticeModel;
use Encore\Admin\Grid;

class SkuNoticeController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '订阅通知';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new SeckillSkuNoticeModel());
        $grid->disableExport();
        $grid->disableCreateButton();
        $grid->model()->orderBy('id', 'desc');

        $grid->column('id', 'ID');
        $grid->column('sku_id', '商品id');
        $grid->column('member_id', '会员id');
        $grid->column('session_id', '场次')->using(SeckillSessionModel::getAll());
        $grid->column('send_at', '发送时间')->datetime();
        $grid->column('type', '通知方式');
        $this->setGridTimeView($grid, true);

        return $grid;
    }
}
