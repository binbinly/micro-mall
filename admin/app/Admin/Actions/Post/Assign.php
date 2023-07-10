<?php

namespace App\Admin\Actions\Post;

use App\Models\Ware\PurchaseModel;
use Encore\Admin\Actions\RowAction;
use Illuminate\Http\Request;

/**
 * 分配采购单
 * Class RelAttr
 * @package App\Admin\Actions\Post
 */
class Assign extends RowAction
{
    public $name = '分配';

    public function handle(PurchaseModel $model, Request $request)
    {
        if ($model->status != PurchaseModel::STATUS_INIT) {
            return $this->response()->error('状态非法');
        }
        $model->status = PurchaseModel::STATUS_USE;
        $model->assignee_id = $request->input('id');
        $model->assignee_name = $request->input('name');
        $model->phone = $request->input('phone');
        $model->saveOrFail();
        return $this->response()->success('分配成功')->refresh();
    }

    public function form(): void
    {
        $this->text('id', '采购人id')->required();
        $this->text('name', '采购人')->required();
        $this->mobile('phone', '手机号')->required();
    }
}
