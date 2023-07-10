<?php

namespace App\Admin\Actions\Post;

use App\Models\Ware\PurchaseModel;
use Encore\Admin\Actions\RowAction;
use Illuminate\Http\Request;

/**
 * 领取采购单
 * Class RelAttr
 * @package App\Admin\Actions\Post
 */
class Draw extends RowAction
{
    public $name = '领取';

    public function handle(PurchaseModel $model, Request $request)
    {
        if ($model->status != PurchaseModel::STATUS_USE) {
            return $this->response()->error('状态非法');
        }
        $model->status = PurchaseModel::STATUS_RECEIVED;
        $model->saveOrFail();
        return $this->response()->success('领取成功')->refresh();
    }

    public function dialog(): void
    {
        $this->confirm('请确认是否领取该采购单？');
    }
}
