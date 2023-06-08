<?php

namespace App\Admin\Common {
    interface Constant
    {
        //普通开关样式
        const SWITCH = [
            'on' => ['value' => 1, 'text' => '开', 'color' => 'success'],
            'off' => ['value' => 0, 'text' => '关', 'color' => 'danger'],
        ];

        //状态开关样式
        const STATUS_SWITCH = [
            'on' => ['value' => 1, 'text' => '启用', 'color' => 'success'],
            'off' => ['value' => 2, 'text' => '冻结', 'color' => 'danger']
        ];

        //状态bool定义
        const STATUS_BOOL = [1 => true, 2 => false];

        //默认label样式
        const LABEL = [0 => 'primary', 1 => 'success', 2 => 'danger'];
    }
}