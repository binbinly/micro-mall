<?php

namespace App\Admin\Actions;

use Encore\Admin\Actions\RowAction;

class Href extends RowAction
{
    public $name;

    /**
     * 跳转地址
     * @var string
     */
    public string $url;

    /**
     * 是否开启新窗口
     * @var bool
     */
    public bool $blank = true;

    public function __construct($name, $url)
    {
        $this->name = $name;
        $this->url = $url;
        parent::__construct();
    }

    public function render(): string
    {
        return "<a href='{$this->url}'>{$this->name}</a>";
    }
}
