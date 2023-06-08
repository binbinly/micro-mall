<?php

namespace App\Console\Commands;

use App\Admin\Common\Format;
use App\Models\Market\SeckillSessionModel;
use App\Models\Market\SeckillSkuModel;
use App\Models\Product\SkuModel;
use Illuminate\Console\Command;
use Illuminate\Support\Facades\Redis;

class Seckill extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'seckill:online';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = '秒杀商品上线';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        //清除上一次数据
        Redis::del('seckill:sessions');
        Redis::del('seckill:sku:session');

        //加载明日所有场次及商品入缓存
        $start = strtotime(date('Y-m-d', strtotime('+1 day')));
        $end = strtotime(date('Y-m-d', strtotime('+2 day'))) - 1;
        $sessions = SeckillSessionModel::query()->select(['id', 'name', 'start_at', 'end_at'])
            ->where('is_release', 1)->whereBetween('start_at', [$start, $end])
            ->whereBetween('end_at', [$start, $end])->get()->toArray();
        if (!$sessions) {
            $this->output->error('无上架场次');
            return;
        }
        $sessions = array_column($sessions, null, 'id');
        //场次入缓存
        Redis::hmset('seckill:sessions', array_map(function ($val){
            $val['start_at'] = strtotime($val['start_at']);
            $val['end_at'] = strtotime($val['end_at']);
            return json_encode($val);
        }, $sessions));

        $skus = [];
        $secSkuList = SeckillSkuModel::query()->whereIn('session_id', array_keys($sessions))->get()->toArray();
        $skuList = SkuModel::query()->whereIn('id', array_column($secSkuList, 'sku_id'))->get()->toArray();
        $skuMap = array_column($skuList, null, 'id');
        foreach ($secSkuList as $sku) {
            //为秒杀商品生成key
            $key = $this->uuid();
            $skus[$sku['session_id']][] = [
                'id' => $sku['sku_id'],
                'price' => Format::amountToPenny($sku['price']),
                'count' => $sku['count'],
                'limit' => $sku['limit'],
                'title' => $skuMap[$sku['sku_id']]['title'],
                'original_price' => $skuMap[$sku['sku_id']]['price'],
                'cover' => $skuMap[$sku['sku_id']]['cover'],
                'start_at' => strtotime($sessions[$sku['session_id']]['start_at']),
                'end_at' => strtotime($sessions[$sku['session_id']]['end_at']),
                'key' => $key,
            ];
            Redis::hset('seckill:sku:session', $sku['sku_id'], $sku['session_id']);
            //预热库存
            Redis::incrBy('seckill:stock:' . $key, $sku['count']);
            Redis::expire('seckill:stock:' . $key, 100000);
        }
        //秒杀商品入缓存
        foreach ($skus as $sessionId => $val) {
            Redis::hmset('seckill:skus:'.$sessionId, array_map(function ($v){
                return json_encode($v);
            }, array_column($val, null, 'id')));
            Redis::expire('seckill:skus:'.$sessionId, 100000);
        }
        $this->output->success('success');
    }

    protected function uuid()
    {
        return md5(uniqid(mt_rand(), true));
    }
}
