<?php

namespace App\Console\Commands;

use App\Admin\Services\SpuService;
use Illuminate\Console\Command;
use Elasticsearch;

class Test extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'test:test';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Command description';

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
        SpuService::offline(11);
    }

    protected function esCreate(){
        $params['index'] = 'paopao';
        $params['type'] = 'test';
        for($i = 21; $i <= 30; $i ++) {
            $params['body'][]=array(
                'create' => array( #注意create也可换成index
                    '_id'=>$i
                ),
            );

            $params['body'][]=array(
                'aa'=>$i
            );
        }
        $return = Elasticsearch::index($params);
        print_r($return);
    }
}
