# golang service public model

#### install

go get -u github.com/nikgame/public

#### use

import (
    "github.com/nikgame/public/xxxx"
)

#### models

##### config

配置文件加载类，copy的beego config方法，只实现了ini文件的加载

##### db

mysql连接池

##### redis

redis连接池

##### sms

短信发送模块，实现了北纬通信、阿里大于的短信发送方法

##### proto

protobuffer的协议文件
