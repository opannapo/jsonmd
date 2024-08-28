

#### result

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| success  | bool  | true | |  | 
| code  | float64  | 200 | |  | 
| message  | string  | Create user successfully | |  | 
| data  | object  | Table : .data | |  | 

#### .data

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| id  | float64  | 15634 | |  | 
| name  | string  | opan napo | |  | 
| birthdate  | string  | 17 Aug 1945 | |  | 
| nik  | object  | Table : .data.nik | |  | 
| meta  | object  | Table : .data.meta | |  | 
| image  | object  | Table : .data.image | |  | 
| address  | array  | Table : .data.address | |  | 

#### .data.nik

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| number  | string  | 1001212121212112121 | |  | 
| name  | string  | Opan Napo ST. MsI | |  | 
| expired_at  | string  | 17 Aug 2075 | |  | 

#### .data.meta

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| is_return_config  | bool  | false | |  | 
| return_config_limit  | float64  | 0 | |  | 
| insurance_type  | string  |  | |  | 
| return_config_reason_ids  | string  |  | |  | 
| status  | object  | Table : .data.meta.status | |  | 

#### .data.meta.status

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| is_enable  | bool  | true | |  | 
| child_status  | object  | Table : .data.meta.status.child_status | |  | 

#### .data.meta.status.child_status

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| status  | bool  | false | |  | 
| name  | string  | child1 | |  | 

#### .data.image

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| width  | float64  | 200 | |  | 
| height  | float64  | 200 | |  | 
| url  | string  | images/0001.jpg | |  | 