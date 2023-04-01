### Inventory

#### Features

<ol>
    <li>Register Admin (user)</li>
    <li>Create Suppliers</li>
    <li>Create customers</li>
    <li>Create categories</li>
    <li>Create products</li>
    <li>Create purchases</li> 
</ol>

#### Requirements

#### Installation

##### Admin(user)

| Field Name | Data Type   | Default |
| ---------- | ----------- | ------- |
| id         | int         |         |
| first_name | varchar(50) | N/A     |
| last_name  | varchar(50) | N/A     |
| email      | varchar(50) | N/A     |
| username   | varchar(50) | N/A     |
| type       | varchar(20) | admin   |
| is_active  | boolean     | true    |


##### Supplier table:


| Field Name | Data Type   | Default |
| ---------- | ----------- | ------- |
| id         | int         |         |
| name       | varchar(50) | N/A     |

##### Customer table:


| Field Name | Data Type   | Default |
| ---------- | ----------- | ------- |
| id         | int         |         |
| name       | varchar(50) | N/A     |

##### Category table:


| Field Name | Data Type   | Default |
| ---------- | ----------- | ------- |
| id         | int         |         |
| name       | varchar(50) | N/A     |


##### Product table:


| Field Name | Data Type   | Default |
| ---------- | ----------- | ------- |
| id         | int         |         |
| category_id| int         |         |
| name       | text        | N/A     |
| description| text        | N/A     |
| price      | float       | N/A     |

##### Purchase table: 


| Field Name | Data Type   | Default |
| ---------- | ----------- | ------- |
| id         | int         |         |
| supplier_id| int         |         |
| product_id | int         |         |
| quantity   | int         | N/A     |
| description| text        | N/A     |
| price      | float       | N/A     |
| total      | float       | N/A     |

##### Stock table: 


| Field Name | Data Type   | Default |
| ---------- | ----------- | ------- |
| id         | int         |         |
| product_id | int         |         |
| quantity   | int         | N/A     |


##### Sell table: 

| Field Name | Data Type   | Default |
| ---------- | ----------- | ------- |
| id         | int         |         |
| customer_id| int         |         |
| total      | float       | N/A     |

##### Sell product table: 


| Field Name | Data Type   | Default |
| ---------- | ----------- | ------- |
| id         | int         |         |
| sell_id    | int         |         |
| product_id | int         |         |
| quantity   | int         | N/A     |
| price      | float       | N/A     |
| total      | float       | N/A     | 