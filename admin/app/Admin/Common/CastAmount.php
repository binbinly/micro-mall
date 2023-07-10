<?php


namespace App\Admin\Common;


use Illuminate\Contracts\Database\Eloquent\CastsAttributes;
use Illuminate\Database\Eloquent\Model;

class CastAmount implements CastsAttributes
{
    public function get(Model $model, string $key, mixed $value, array $attributes)
    {
        if ($value == 0) return 0;
        return round($value / 100, 2);
    }

    public function set(Model $model, string $key, mixed $value, array $attributes)
    {
        return intval($value * 100);
    }
}
