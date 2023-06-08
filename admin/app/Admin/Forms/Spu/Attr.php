<?php

namespace App\Admin\Forms\Spu;

use App\Models\Product\AttrGroupModel;
use App\Models\Product\AttrRelGroupModel;
use Encore\Admin\Widgets\Form;
use Encore\Admin\Widgets\StepForm;
use Illuminate\Http\Request;

class Attr extends StepForm
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
        return $this->next($request->all());
    }

    /**
     * Build a form here.
     */
    public function form()
    {
        list($groups, $attrs) = $this->getAttrName();
        if (!$attrs){
            $this->html('请先设置分类属性');
            return;
        }
        foreach ($attrs as $groupId => $attr) {
            $this->fieldset($groups[$groupId], function (Form $form) use ($attr) {
                foreach ($attr as $item) {
                    $form->hidden('id['.$item['id'].']')->default($item['name']);
                    $form->tags('attr-' . $item['id'], $item['name'])->options(explode(',', $item['values']))->required();
                }
            });
        }
    }

    /**
     * 获取当前分类下的所有分组和下面的属性
     * @return array|false
     */
    protected function getAttrName()
    {
        $basic = session()->get('steps.basic', []);
        if (!$basic) {
            return false;
        }
        $catId = $basic['cat_id'];
        if (!$catId) {
            return false;
        }
        //当前分下的所有分组
        $groups = AttrGroupModel::getGroupByCatId($catId);
        //对应的所有分组属性
        $attrs = AttrRelGroupModel::getBaseAttrs(array_keys($groups));
        return [$groups, $this->attrGroup($attrs)];
    }

    /**
     * 对所有属性分组
     * @param $attrs
     * @return array
     */
    protected function attrGroup($attrs) {
        $res = [];
        foreach ($attrs as $item) {
            $res[$item['group_id']][] = [
                'id' => $item['attr']['id'],
                'name' => $item['attr']['name'],
                'values' => $item['attr']['values'],
                'is_many' => $item['attr']['is_many'],
                'is_show' => $item['attr']['is_show']
            ];
        }
        return $res;
    }
}
