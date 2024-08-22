App : go run ./main.go --file=sample/response-small.json

Input Json 
```json
{
  "success": true,
  "code": 200,
  "message": "Client successfully retrieved",
  "data": {
    "client_id": 15634,
    "client_current_rate": {
      "client_custom_rate_id": 4924,
      "client_status_custom_rate": "nearly_expired"
    },
    "client_renewal_rate": {
      "client_custom_rate_id": 4982,
      "client_status_custom_rate": "requested",
      "client_status_updated_at": null
    },
    "client_meta": {
      "is_return_config": false,
      "return_config_limit": 0,
      "return_config_reason_ids": "",
      "insurance_type": ""
    }
  }
}
```

Output root.md file:
```md

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| success  | bool  | true | |  | 
| code  | float64  | 200 | |  | 
| message  | string  | Client successfully retrieved | |  | 
| data  | object  | Table data object | |  | 

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
### Table data object    
    

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| client_renewal_rate | object | Table client_renewal_rate object | |  |  |
| client_meta | object | Table client_meta object | |  |  |
| client_id | float64 | 15634 | |  |  |
| client_current_rate | object | Table client_current_rate object | |  |  |

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
### Table client_renewal_rate object    
    

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| client_status_updated_at |  | <nil> | |  |  |
| client_custom_rate_id | float64 | 4982 | |  |  |
| client_status_custom_rate | string | requested | |  |  |

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
### Table client_current_rate object    
    

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| client_custom_rate_id | float64 | 4924 | |  |  |
| client_status_custom_rate | string | nearly_expired | |  |  |

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
### Table client_current_rate object    
    

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| client_status_custom_rate | string | nearly_expired | |  |  |
| client_custom_rate_id | float64 | 4924 | |  |  |

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
### Table client_meta object    
    

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| is_return_config | bool | false | |  |  |
| return_config_limit | float64 | 0 | |  |  |
| return_config_reason_ids | string |  | |  |  |
| insurance_type | string |  | |  |  |

```

Output Preview:

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| success  | bool  | true | |  | 
| code  | float64  | 200 | |  | 
| message  | string  | Client successfully retrieved | |  | 
| data  | object  | Table data object | |  | 

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
### Table data object


| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| client_renewal_rate | object | Table client_renewal_rate object | |  |  |
| client_meta | object | Table client_meta object | |  |  |
| client_id | float64 | 15634 | |  |  |
| client_current_rate | object | Table client_current_rate object | |  |  |

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
### Table client_renewal_rate object


| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| client_status_updated_at |  | <nil> | |  |  |
| client_custom_rate_id | float64 | 4982 | |  |  |
| client_status_custom_rate | string | requested | |  |  |

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
### Table client_current_rate object


| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| client_custom_rate_id | float64 | 4924 | |  |  |
| client_status_custom_rate | string | nearly_expired | |  |  |

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
### Table client_current_rate object


| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| client_status_custom_rate | string | nearly_expired | |  |  |
| client_custom_rate_id | float64 | 4924 | |  |  |

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
### Table client_meta object


| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| is_return_config | bool | false | |  |  |
| return_config_limit | float64 | 0 | |  |  |
| return_config_reason_ids | string |  | |  |  |
| insurance_type | string |  | |  |  |
