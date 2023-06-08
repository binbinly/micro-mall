<?php

namespace App\Admin\Forms\Spu;

use App\Admin\Services\SpuService;
use App\Models\Product\AttrModel;
use Encore\Admin\Widgets\StepForm;
use Illuminate\Http\Request;

class Sku extends StepForm
{
    /**
     * The form title.
     *
     * @var string
     */
    public $title = '规格属性';

    /**
     * Handle the form request.
     *
     * @param Request $request
     *
     * @return \Illuminate\Http\RedirectResponse
     */
    public function handle(Request $request)
    {
        $service = new SpuService();
        $suc = $service->create($this->all(), $err);
        if ($suc) {
            admin_success('添加成功');
            $this->clear();
            return redirect('/product/spu');
        }
        admin_error('创建失败,' . $err);
        return back();
    }

    /**
     * Build a form here.
     */
    public function form()
    {
        list($data, $title) = $this->getSaleAttrName();
        $this->sku('sku', '销售属性设置')->data($data)->title($title);
    }

    protected function getSaleAttrName()
    {
        $basic = session()->get('steps.basic', []);
        if (!$basic) {
            return false;
        }
        $catId = $basic['cat_id'];
        if (!$catId) {
            return false;
        }
        $title = $basic['name'];
        $attrs = AttrModel::getAttrByCatID($catId, AttrModel::TYPE_SALE);
        return [array_map(function ($val) {
            return [
                'id' => $val['id'],
                'name' => $val['name'],
                'values' => explode(',', $val['values'])
            ];
        }, $attrs), $title];
    }
}
