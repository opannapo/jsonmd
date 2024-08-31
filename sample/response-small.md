

#### sample/response-small

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
| hobbies  | array  | Table : .data.hobbies | |  | 
| nik  | object  | Table : .data.nik | |  | 
| meta  | object  | Table : .data.meta | |  | 
| image  | object  | Table : .data.image | |  | 
| address  | array  | Table : .data.address | |  | 

#### .data.meta

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| is_return_config  | bool  | false | |  | 
| return_config_limit  | float64  | 0 | |  | 
| return_config_reason_ids  | string  |  | |  | 
| insurance_type  | string  |  | |  | 
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

#### .data.address

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| name  | string  | home 1 | |  | 
| address  | string  | JL.Selamat Dunia Akhirat Rt/Rw 001/002 | |  | 

#### .data.hobbies

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
|   | string  | Ibadah | |  | 

#### .data.nik

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| number  | string  | 1001212121212112121 | |  | 
| name  | string  | Opan Napo ST. MsI | |  | 
| expired_at  | string  | 17 Aug 2075 | |  | 