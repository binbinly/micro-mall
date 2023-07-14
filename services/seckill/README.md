### 秒杀（高并发）系统关注问题
- 服务单一职责，独立部署
  秒杀服务即使自己扛不住压力，挂掉，也不会影响其他
- 秒杀连接加密
  防止恶意攻击，模拟秒杀请求，防止链接暴露，提前秒杀商品
- 库存预热+快速扣减
  秒杀要读多写少，无需每次实时校验库存，我们库存预热，放到redis中，信号量控制秒杀请求
- 动静分离
  使用cdn
- 恶意请求拦截
  获取非法攻击请求进行拦截，网关层实现
- 流量错峰
  使用各种手段，将流量分担到更大宽度的时间点，比如验证码，加入购物车
- 限流&熔断&降级
  前端限流 + 后端限流
  限制次数，限制总量，快速失败降级运行，熔断隔离防止雪崩
- 队列削峰
  所有秒杀成功的请求，进入队列，慢慢创建订单，扣减库存即可

# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [warehouse/warehouse.proto](#warehouse_warehouse-proto)
    - [SkuStockLockReq](#warehouse-SkuStockLockReq)
    - [SkuStockLockReq.SkuNumsEntry](#warehouse-SkuStockLockReq-SkuNumsEntry)
    - [SkuStockReply](#warehouse-SkuStockReply)
    - [SkuStockUnlockReq](#warehouse-SkuStockUnlockReq)
    - [SkusReq](#warehouse-SkusReq)
    - [SkusStockReply](#warehouse-SkusStockReply)
    - [SkusStockReply.SkuNumsEntry](#warehouse-SkusStockReply-SkuNumsEntry)
  
    - [Warehouse](#warehouse-Warehouse)
  
- [Scalar Value Types](#scalar-value-types)



<a name="warehouse_warehouse-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## warehouse/warehouse.proto
仓储服务,spu,sku库存管理


<a name="warehouse-SkuStockLockReq"></a>

### SkuStockLockReq
锁定库存请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [int64](#int64) |  | 订单id |
| order_no | [string](#string) |  | 订单号 |
| consignee | [string](#string) |  | 收货人 |
| phone | [string](#string) |  | 收货人手机号 |
| address | [string](#string) |  | 收货地址 |
| note | [string](#string) |  | 订单备注 |
| sku_nums | [SkuStockLockReq.SkuNumsEntry](#warehouse-SkuStockLockReq-SkuNumsEntry) | repeated | sku_id =&gt; 锁库存数量 |






<a name="warehouse-SkuStockLockReq-SkuNumsEntry"></a>

### SkuStockLockReq.SkuNumsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [int32](#int32) |  |  |






<a name="warehouse-SkuStockReply"></a>

### SkuStockReply
库存数量响应结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| num | [int32](#int32) |  | 库存数量 |






<a name="warehouse-SkuStockUnlockReq"></a>

### SkuStockUnlockReq
解锁库存请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [int64](#int64) |  | 订单id |
| finish | [bool](#bool) |  | 订单是否已完成 |






<a name="warehouse-SkusReq"></a>

### SkusReq
多个sku id


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [int64](#int64) | repeated | sku ids |






<a name="warehouse-SkusStockReply"></a>

### SkusStockReply
多sku库存数量,map结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_nums | [SkusStockReply.SkuNumsEntry](#warehouse-SkusStockReply-SkuNumsEntry) | repeated | sku_id =&gt; 库存数量 |






<a name="warehouse-SkusStockReply-SkuNumsEntry"></a>

### SkusStockReply.SkuNumsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [int32](#int32) |  |  |





 

 

 


<a name="warehouse-Warehouse"></a>

### Warehouse
仓储服务接口定义

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetSkuStock | [.common.SkuIDReq](#common-SkuIDReq) | [SkuStockReply](#warehouse-SkuStockReply) | 获取sku的库存数量 |
| BatchSkusStock | [SkusReq](#warehouse-SkusReq) | [SkusStockReply](#warehouse-SkusStockReply) | 批量获取sku的库存数量 |
| SKuStockLock | [SkuStockLockReq](#warehouse-SkuStockLockReq) | [.google.protobuf.Empty](#google-protobuf-Empty) | 锁定sku库存 |
| SkuStockUnlock | [SkuStockUnlockReq](#warehouse-SkuStockUnlockReq) | [.google.protobuf.Empty](#google-protobuf-Empty) | 解锁sku库存 |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

