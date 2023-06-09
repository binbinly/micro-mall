<?php

namespace App\Admin\Extensions;

use Encore\Admin\Form\Field;
use Illuminate\Support\Facades\Storage;

class FormMediaField extends Field
{
    protected $view = 'narwhalformmedia::mediayel_field';

    protected static $css = [
        'vendor/yelphp/narwhalformmedia/mediayel_field.css'
    ];

    protected static $js = [
        'vendor/yelphp/narwhalformmedia/mediayel_field.js'
    ];

    protected int $limit = 1;
    protected bool $remove = false;
    protected string $type = 'img';
    protected string $path = "./";

    /**
     * @param int $limit
     * @return $this
     */
    public function limit(int $limit = 1): static
    {
        $this->limit = $limit;
        return $this;
    }

    public function remove(): static
    {
        $this->remove = true;
        return $this;
    }

    public function path($path): static
    {
        $this->path = $path;
        return $this;
    }

    public function render()
    {
        $disk = config('admin.upload.disk');

        $storage = Storage::disk($disk);
        $this->addVariables([
            'limit' => $this->limit,
            'rootpath' => $storage->url($this->path),
            'remove' => $this->remove,
        ]);

        $name = $this->column;
        $limit = $this->limit;

        $rootpath = $storage->url($this->path);
        $remove = $this->remove;

        // 初始化
        $this->script = "
            if(!window.Demo{$name}){
                window.Demo{$name} = new MediaYelDemo('{$name}',{$limit},'{$rootpath}',{$remove}+'','{$this->type}');
                Demo{$name}.Run();
            }
            Demo{$name}.init();
            Demo{$name}.getdata();//获取数据
            ";

        return parent::render();
    }

}
