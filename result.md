

#### result

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| code  | float64  | 200 | |  | 
| message  | string  | Client successfully retrieved | |  | 
| data  | object  | Table : .data | |  | 
| success  | bool  | true | |  | 

#### .data

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| client_meta  | object  | Table : .data.client_meta | |  | 
| child  | array  | Table : .data.child | |  | 
| client_id  | float64  | 15634 | |  | 
| client_current_rate  | object  | Table : .data.client_current_rate | |  | 
| client_renewal_rate  | object  | Table : .data.client_renewal_rate | |  | 

#### .data.client_current_rate

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| client_custom_rate_id  | float64  | 4924 | |  | 
| client_status_custom_rate  | string  | nearly_expired | |  | 

#### .data.client_renewal_rate

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| client_custom_rate_id  | float64  | 4982 | |  | 
| client_status_custom_rate  | string  | requested | |  | 
| client_status_updated_at  |   | <nil> | |  | 

#### .data.client_meta

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| return_config_reason_ids  | string  |  | |  | 
| insurance_type  | string  |  | |  | 
| status  | object  | Table : .data.client_meta.status | |  | 
| is_return_config  | bool  | false | |  | 
| return_config_limit  | float64  | 0 | |  | 

#### .data.client_meta.status

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| is_enable  | bool  | true | |  | 
| child_status  | object  | Table : .data.client_meta.status.child_status | |  | 

#### .data.client_meta.status.child_status

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| status  | bool  | false | |  | 
| name  | string  | child1 | |  | 