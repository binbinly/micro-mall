<?php


namespace App\Admin\Forms;


use App\Models\Market\AppSettingModel;
use GuzzleHttp\Psr7\UploadedFile;
use Illuminate\Http\RedirectResponse;
use Illuminate\Http\Request;
use Exception;
use Illuminate\Support\Facades\Storage;
use Throwable;

class PageSet extends Base
{
    const SESSION_ID = 'app_page_id';

    public $title = '页面数据设置';

    private $id;

    public function __construct($data = [])
    {
        parent::__construct($data);
        $this->id = session()->get(self::SESSION_ID);
    }

    /**
     * @param Request $request
     * @return RedirectResponse
     * @throws Throwable
     */
    public function handle(Request $request)
    {
        parent::handle($request);

        $data = $request->all();
        $res = [];
        if (in_array($data['type'], [AppSettingModel::TYPE_THREE_ADV, AppSettingModel::TYPE_SWIPER])) {
            $disk = Storage::disk('minio');
            foreach ($data['data'] as $file) {
                if ($file instanceof UploadedFile)
                $res[] = $disk->put(config('admin.upload.directory.image'), $file);
            }
        } else if ($data['type'] == AppSettingModel::TYPE_NAV){
            $disk = Storage::disk('minio');
            foreach ($data['data'] as $item) {
                if ($item['_remove_'] == 1) {
                    continue;
                }
                $res[] = [
                    'title' => $item['title'],
                    'icon' => $disk->put(config('admin.upload.directory.image'), $item['icon'])
                ];
            }
        } else if ($data['type'] == AppSettingModel::TYPE_ONE_ADV) {
            $disk = Storage::disk('minio');
            $res = [
                'title' => $data['data']['title'],
                'cover' => $disk->put(config('admin.upload.directory.image'), $data['data']['cover'])
            ];
        } else if ($data['type'] == AppSettingModel::TYPE_GOODS_LIST) {
            $res = $data['data'];
        }
        if (!$res) {
            return $this->error();
        }
        AppSettingModel::query()->where('id', $this->id)->update([
            'data' => json_encode($res)
        ]);

        return $this->success();
    }

    /**
     * Build a form here.
     */
    public function form()
    {
        $this->hidden('type')->value($this->data()->type);
        switch ($this->data()->type) {
            case AppSettingModel::TYPE_SWIPER:
            case AppSettingModel::TYPE_THREE_ADV:
                $this->multipleImage('data', '轮播图配置')->removable();
                break;
            case AppSettingModel::TYPE_NAV:
                $this->table('data', '图标ICON配置', function ($form) {
                    $form->text('title', '标题');
                    $form->image('icon', 'ICON');
                });
                break;
            case AppSettingModel::TYPE_ONE_ADV:
                $this->embeds('data', '单图广告配置', function ($form) {
                    $form->text('title', '标题');
                    $form->image('cover', '封面图');
                });
                break;
            case AppSettingModel::TYPE_GOODS_LIST:
                $this->text('data', '商品列表路由');
                break;
        }
    }

    /**
     * @return array|mixed
     * @throws Exception
     */
    public function data()
    {
        return AppSettingModel::query()->findOrFail($this->id);
    }
}