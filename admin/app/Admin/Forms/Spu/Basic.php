<?php

namespace App\Admin\Forms\Spu;

use App\Models\Product\BrandModel;
use App\Models\Product\CategoryModel;
use Encore\Admin\Widgets\StepForm;
use Illuminate\Http\Request;

class Basic extends StepForm
{
    /**
     * The form title.
     *
     * @var string
     */
    public $title = '基本信息';

    /**
     * Handle the form request.
     *
     * @param Request $request
     *
     * @return \Illuminate\Http\RedirectResponse
     */
    public function handle(Request $request)
    {
        $this->clear();

        return $this->next($request->all());
    }

    /**
     * Build a form here.
     */
    public function form()
    {
        $this->text('name', '商品名称')->required();
        $this->text('desc', '商品描述');
        $this->select('cat_id', '商品分类')->options(CategoryModel::selectOptions())->required();
        $this->select('brand_id', '选择品牌')->options(BrandModel::getAll())->required();
        $this->number('weight', '商品重量/g')->required();
        $this->number('integral', '设置积分')->default(0);
        $this->number('growth', '成长值')->default(0);
        $this->photo('images','商品介绍')->limit(6)->remove()->required();
        $this->photo('banners','图集')->limit(6)->remove()->required();
    }
}
