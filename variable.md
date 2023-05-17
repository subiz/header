# List of variable

### Account information
`account.*`

#### Examples
`user.fullname` -> `Pham Kieu Thanh`
`user.emails[0]` -> `thanhpk@live.com`
`user.phone`

### User attribute
`user.<attribute-key>`
#### Examples
user.fullname
user.emails[0]
user.phone


`agent.<attribute-key>`

### Facebook Page
`fanpage.*`


Hello {{#user.fullname}}{{user.fullname}}{{/user.fullname}{{^user.fullname}}quý khách{{/user.fullname}}
You have juse won {{value}} dollars!
