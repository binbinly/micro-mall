<?php

namespace App\Admin\Controllers\Market;

use App\Admin\Controllers\BaseController;
use App\Models\Market\CouponMemberModel;
use Encore\Admin\Grid;

class CouponMemberController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '用户优惠券';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new CouponMemberModel());
        $grid->model()->orderBy('id', 'desc');
        $grid->disableExport();

        $grid->column('id', 'ID');
        $grid->column('member_id', '会员id');
        $grid->column('coupon_id', '优惠券id');
        $grid->column('nickname', '会员昵称');
        $grid->column('get_type', '获取方式')->using(CouponMemberModel::$getTypeLabel);
        $grid->column('status', '状态')->using(CouponMemberModel::$statusLabel);
        $grid->column('used_at', '使用时间');
        $grid->column('order_id', '订单id');
        $grid->column('order_no', '订单号');
        $this->setGridTimeView($grid);

        return $grid;
    }
}
