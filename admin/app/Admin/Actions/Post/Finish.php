<?php

namespace App\Admin\Actions\Post;

use App\Models\BaseModel;
use App\Models\Ware\PurchaseDetailModel;
use App\Models\Ware\PurchaseModel;
use App\Models\Ware\WareSkuModel;
use Encore\Admin\Actions\RowAction;
use Illuminate\Http\Request;

/**
 * 完成采购
 * Class RelAttr
 * @package App\Admin\Actions\Post
 */
class Finish extends RowAction
{
    public $name = '完成采购';

    public function handle(PurchaseModel $model, Request $request)
    {
        if ($model->status != PurchaseModel::STATUS_RECEIVED) {
            return $this->response()->error('状态非法');
        }
        $model->status = PurchaseModel::STATUS_FINISH;

        BaseModel::transaction(BaseModel::CONN_WMS);
        try {
            $model->saveOrFail();

            //修改采购需求状态
            PurchaseDetailModel::query()->where('purchase_id', $model->id)->update([
                'status' => PurchaseDetailModel::STATUS_FINISH
            ]);

            //采购入库
            $list = PurchaseDetailModel::query()->select(['sku_id', 'sku_num', 'ware_id'])->where('purchase_id', $model->id)->get()->toArray();
            foreach ($list as $item) {
                $sku = WareSkuModel::query()->where('sku_id', $item['sku_id'])
                    ->where('ware_id', $item['ware_id'])->first();
                if ($sku instanceof WareSkuModel) {//已存在修改库存
                    WareSkuModel::query()->where('sku_id', $item['sku_id'])
                        ->where('ware_id', $item['ware_id'])->increment('stock', $item['sku_num']);
                    continue;
                }
                //新增
                WareSkuModel::create([
                    'sku_id' => $item['sku_id'],
                    'ware_id' => $item['ware_id'],
                    'stock' => $item['sku_num']
                ]);
            }
            BaseModel::commit(BaseModel::CONN_WMS);
        } catch (\Exception $e) {
            BaseModel::rollBack(BaseModel::CONN_WMS);
            return $this->response()->error('失败了j,'.$e->getMessage())->refresh();
        } catch (\Throwable $e) {
            BaseModel::rollBack(BaseModel::CONN_WMS);
            return $this->response()->error('失败了j,'.$e->getMessage())->refresh();
        }
        return $this->response()->success('已完成')->refresh();
    }

    public function dialog(): void
    {
        $this->confirm('请确认是否已完成采购？');
    }
}
