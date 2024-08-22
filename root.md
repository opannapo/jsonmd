
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
