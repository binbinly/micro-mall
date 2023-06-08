<?php

namespace App\Admin\Actions\Post;

use App\Models\BaseModel;
use App\Models\Ware\PurchaseDetailModel;
use App\Models\Ware\PurchaseModel;
use Encore\Admin\Actions\BatchAction;
use Illuminate\Database\Eloquent\Collection;
use Illuminate\Http\Request;

class PurchaseSubmit extends BatchAction
{
    public $name = '提交采购需求';

    public function handle(Collection $collection, Request $request)
    {
        $id = intval($request->input('purchase_id', 0));
        $ids = [];
        foreach ($collection as $model) {
            if ($model instanceof PurchaseDetailModel && $model->status == PurchaseDetailModel::STATUS_INIT) {
                $ids[] = $model->id;
            }
        }
        if (!$ids) {
            return $this->response()->success('没有找到匹配记录哦');
        }
        if ($id) {//绑定已存在的采购单
            PurchaseDetailModel::query()->whereIn('id', $ids)->update([
                'purchase_id' => $id,
                'status' => PurchaseDetailModel::STATUS_USE
            ]);
        } else {//生成新的采购单
            BaseModel::transaction(BaseModel::CONN_WMS);
            try{
                $pur = new PurchaseModel();
                $pur->priority = 1;
                $pur->ware_id = $collection[0]->ware_id;
                $pur->saveOrFail();

                PurchaseDetailModel::query()->whereIn('id', $ids)->update([
                    'purchase_id' => $pur->id,
                    'status' => PurchaseDetailModel::STATUS_USE
                ]);
                BaseModel::commit(BaseModel::CONN_WMS);
            }catch (\Exception $e) {
                BaseModel::rollBack(BaseModel::CONN_WMS);
                return $this->response()->error('失败了')->refresh();
            } catch (\Throwable $e) {
                BaseModel::rollBack(BaseModel::CONN_WMS);
                return $this->response()->error('失败了')->refresh();
            }
        }

        return $this->response()->success('提交成功')->refresh();
    }

    public function form(){
        $this->select('purchase_id', '选择采购单')->options(PurchaseModel::getAssigned());
    }
}