<?php

namespace App\Admin\Actions\Post;

use App\Models\BaseModel;
use App\Models\Market\CouponModel;
use Encore\Admin\Actions\RowAction;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Redis;

/**
 * 生成优惠券
 * Class CouponCreate
 * @package App\Admin\Actions
 */
class CouponRelease extends RowAction
{
    public $name = '发布优惠券';

    public function handle(Model $model, Request $request)
    {
        //已经生成数量
        $coupon = CouponModel::query()->findOrFail($model->id);
        if ($coupon->is_release == 1) {
            return $this->response()->error('不可重复发行哦');
        }
        $coupon->is_release = 1;
        $coupon->publish_count = $request->input('num');
        BaseModel::transaction();
        try{
            Redis::incrby('coupon_num:'.$model->id, $coupon->publish_count);
            $coupon->saveOrFail();
            BaseModel::commit();
        }catch (\Exception $e) {
            BaseModel::rollBack();
            return $this->response()->error('发布失败，'.$e->getMessage());
        } catch (\Throwable $e) {
            BaseModel::rollBack();
            return $this->response()->error('发布失败，'.$e->getMessage());
        }
        return $this->response()->success('操作成功')->refresh();
    }

    public function form($model)
    {
        $this->text('num', '发行数量')->default($model->num)->required()->rules('required|numeric');
    }

}