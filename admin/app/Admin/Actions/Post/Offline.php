<?php

namespace App\Admin\Actions\Post;

use App\Admin\Services\SpuService;
use App\Models\BaseModel;
use App\Models\Product\SpuModel;
use Encore\Admin\Actions\RowAction;
use Illuminate\Http\Request;

/**
 * 商品下架
 * Class RelAttr
 * @package App\Admin\Actions\Post
 */
class Offline extends RowAction
{
    public $name = '下架';

    public function handle(SpuModel $model, Request $request)
    {
        if ($model->status != SpuModel::STATUS_ONLINE) {
            return $this->response()->error('状态非法');
        }

        BaseModel::transaction(BaseModel::CONN_PMS);
        try{
            $model->status = SpuModel::STATUS_OFFLINE;
            $model->saveOrFail();

            SpuService::offline($model->id);
            BaseModel::commit(BaseModel::CONN_PMS);
        }catch (\Exception $e) {
            BaseModel::rollBack(BaseModel::CONN_PMS);
            return $this->response()->error('失败了,'.$e->getMessage())->refresh();
        } catch (\Throwable $e) {
            BaseModel::rollBack(BaseModel::CONN_PMS);
            return $this->response()->error('失败了,'.$e->getMessage())->refresh();
        }
        return $this->response()->success('下架成功')->refresh();
    }

    public function dialog(): void
    {
        $this->confirm('请确认是否要下架该商品？');
    }
}
