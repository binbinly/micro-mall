<?php

namespace App\Admin\Actions\Post;

use App\Models\Product\AttrGroupModel;
use App\Models\Product\AttrModel;
use App\Models\Product\AttrRelGroupModel;
use Encore\Admin\Actions\Action;
use Illuminate\Http\Request;

/**
 * 属性分组关联属性
 * 1，当前分组只能关联自己所属分类的所有属性
 * 2，当前分组只能关联别的分组没有关联过的属性
 * Class RelAttr
 * @package App\Admin\Actions\Post
 */
class RelAttr extends Action
{
    protected $selector = '.rel-attr';

    public $name = '关联属性';

    protected int $groupId;

    public function __construct($groupId = 0)
    {
        $this->groupId = $groupId;
        parent::__construct();
    }

    public function handle(Request $request)
    {
        $attrIds = array_filter($request->input('attr_id'));
        $groupId = intval($request->input('group_id'));
        if (count($attrIds) == 0 || !$groupId) {
            return $this->response()->success('未操作...');
        }
        $data = array_map(function ($val) use($groupId) {
            return ['attr_id' => $val, 'group_id' => $groupId];
        }, $attrIds);
        if (AttrRelGroupModel::query()->insert($data)) {
            return $this->response()->success('添加成功')->refresh();
        }
        return $this->response()->error('添加失败');
    }

    public function form(): void
    {
        //当前分组模型
        $group = AttrGroupModel::query()->find($this->groupId);
        if (! ($group instanceof AttrGroupModel)) {
            return;
        }
        //当前分组只能关联自己所属分类的所有属性
        $attrs = AttrModel::getAttrMapByCatID($group['cat_id']);
        //当前分类下的所有分组id
        $groupIds = AttrGroupModel::getGroupIdByCatId($group['cat_id']);
        //当前分类下的其他分组
//        $groupId = $this->groupId;
//        $groupIds = array_filter($groupIds, function ($val) use($groupId) {
//            return $val != $groupId;
//        });
        //这些分组关联的属性id
        $attrIds = AttrRelGroupModel::getAttrByGroupIds($groupIds);
        //从当前分类的属性中移除这些属性
        if ($attrIds) {
            foreach ($attrs as $id => $name) {
                if (in_array($id, $attrIds)) {
                    unset($attrs[$id]);
                }
            }
        }

        if (count($attrs) == 0) {
            $this->checkbox('attr_id', '属性')->disable()->help('没有可以关联的属性');
        } else {
            $this->hidden('group_id')->value($this->groupId);
            $this->checkbox('attr_id', '属性')->options($attrs);
        }
    }

    public function html(): string
    {
        return <<<HTML
        <a class="btn btn-sm btn-info rel-attr">添加关联</a>
HTML;
    }
}
