# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/health.proto](#proto_health-proto)
    - [HealthCheckRequest](#warehouse-HealthCheckRequest)
    - [HealthCheckResponse](#warehouse-HealthCheckResponse)
  
    - [HealthCheckResponse.ServingStatus](#warehouse-HealthCheckResponse-ServingStatus)
  
    - [Health](#warehouse-Health)
  
- [proto/warehouse.proto](#proto_warehouse-proto)
    - [SkuReq](#warehouse-SkuReq)
    - [SkuStockLockReq](#warehouse-SkuStockLockReq)
    - [SkuStockLockReq.SkuNumEntry](#warehouse-SkuStockLockReq-SkuNumEntry)
    - [SkuStockReply](#warehouse-SkuStockReply)
    - [SkuStockUnlockReq](#warehouse-SkuStockUnlockReq)
    - [SpuReq](#warehouse-SpuReq)
    - [SpuStockReply](#warehouse-SpuStockReply)
    - [SpuStockReply.SkuNumEntry](#warehouse-SpuStockReply-SkuNumEntry)
  
    - [Warehouse](#warehouse-Warehouse)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto_health-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/health.proto



<a name="warehouse-HealthCheckRequest"></a>

### HealthCheckRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service | [string](#string) |  |  |






<a name="warehouse-HealthCheckResponse"></a>

### HealthCheckResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [HealthCheckResponse.ServingStatus](#warehouse-HealthCheckResponse-ServingStatus) |  |  |





 


<a name="warehouse-HealthCheckResponse-ServingStatus"></a>

### HealthCheckResponse.ServingStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| SERVING | 1 |  |
| NOT_SERVING | 2 |  |
| SERVICE_UNKNOWN | 3 |  |


 

 


<a name="warehouse-Health"></a>

### Health


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Check | [HealthCheckRequest](#warehouse-HealthCheckRequest) | [HealthCheckResponse](#warehouse-HealthCheckResponse) |  |
| Watch | [HealthCheckRequest](#warehouse-HealthCheckRequest) | [HealthCheckResponse](#warehouse-HealthCheckResponse) stream |  |

 



<a name="proto_warehouse-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/warehouse.proto
仓储服务,spu,sku库存管理


<a name="warehouse-SkuReq"></a>

### SkuReq
sku库存请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_id | [int64](#int64) |  | sku_id |






<a name="warehouse-SkuStockLockReq"></a>

### SkuStockLockReq
锁定库存请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [int64](#int64) |  | 订单id |
| spu_id | [int64](#int64) |  | spu_id |
| order_no | [string](#string) |  | 订单号 |
| consignee | [string](#string) |  | 收货人 |
| phone | [string](#string) |  | 收货人手机号 |
| address | [string](#string) |  | 收货地址 |
| note | [string](#string) |  | 订单备注 |
| sku_num | [SkuStockLockReq.SkuNumEntry](#warehouse-SkuStockLockReq-SkuNumEntry) | repeated | sku_id =&gt; 锁库存数量 |






<a name="warehouse-SkuStockLockReq-SkuNumEntry"></a>

### SkuStockLockReq.SkuNumEntry



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
| spu_id | [int64](#int64) |  | spu_id |
| finish | [bool](#bool) |  | 订单是否已完成 |






<a name="warehouse-SpuReq"></a>

### SpuReq
spu库存请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spu_id | [int64](#int64) |  | spu_id |
| sku_ids | [int64](#int64) | repeated | sku_id |






<a name="warehouse-SpuStockReply"></a>

### SpuStockReply
spu下所有sku库存数量,map结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_num | [SpuStockReply.SkuNumEntry](#warehouse-SpuStockReply-SkuNumEntry) | repeated | sku_id =&gt; 库存数量 |






<a name="warehouse-SpuStockReply-SkuNumEntry"></a>

### SpuStockReply.SkuNumEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [int32](#int32) |  |  |





 

 

 


<a name="warehouse-Warehouse"></a>

### Warehouse
仓储服务接口定义

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetSkuStock | [SkuReq](#warehouse-SkuReq) | [SkuStockReply](#warehouse-SkuStockReply) | 获取sku的库存数量 |
| GetSpuStock | [SpuReq](#warehouse-SpuReq) | [SpuStockReply](#warehouse-SpuStockReply) | 获取spu下所有sku的库存数量 |
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

