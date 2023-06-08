<?php


namespace App\Admin\Forms;


use App\Admin\Services\ConfigService;
use App\Models\Market\ConfigModel;
use App\Models\Product\CategoryModel;
use Encore\Admin\Form\NestedForm;
use Illuminate\Http\RedirectResponse;
use Illuminate\Http\Request;
use Exception;
use Throwable;

class HomeCat extends Base
{
    public $title = '首页推荐分类';

    /**
     * @param Request $request
     * @return RedirectResponse
     * @throws Throwable
     */
    public function handle(Request $request)
    {
        parent::handle($request);

        $data = $request->all();
        $list = [];
        foreach($data[ConfigModel::KEY_HOME_CAT] as $item) {
            if ($item['_remove_'] == 1) {
                continue;
            }
            $list[] = [
                'id' => intval($item['id']),
                'name' => $item['name']
            ];
        }

        ConfigService::save([ConfigModel::KEY_HOME_CAT => json_encode($list)]);

        return $this->success();
    }

    /**
     * Build a form here.
     */
    public function form()
    {
        $this->table(ConfigModel::KEY_HOME_CAT, '首页推荐分类', function (NestedForm $form){
            $form->select('id', '选择分类')->options(CategoryModel::parentAll());
            $form->text('name', '分类名');
        });
    }

    /**
     * @return array|mixed
     * @throws Exception
     */
    public function data()
    {
        return [ConfigModel::KEY_HOME_CAT => ConfigModel::init(ConfigModel::KEY_HOME_CAT, ConfigModel::TYPE_JSON)];
    }
}