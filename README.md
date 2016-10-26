# golang service public model

### 版本说明

####1.1

修正了sms配置参数的BUG， 支持变更签名，当渠道号为1时使用默认配置

    [ALIDAYU]
    url = ""
    name = ""
    key = ""
    secert = ""
    sign = ""
    template = ""

    [ALIDAYU:10]
    url = ""
    name = ""
    key = ""
    secert = ""
    sign = ""
    template = ""

#### install

go get -u github.com/nikugame/public

#### use

import (
    "github.com/nikugame/public/xxxx"
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


#### 配置文件说明

需要定义在执行目录下创建 conf文件夹，配置文件名为settings

##### mysql配置

    [Mysql]
    host = "127.0.0.1"  
    port = "3306"
    user = "root"
    password = "123456"
    dbname = "test"
    open = 100 ;最大连接数
    idle = 10  ;连接池 

##### redis 配置

    [Redis]
    host = "10.0.0.207"
    password = ""
    port = "6379"
    mode = true   ;true为使用密码访问，false使用空密码
    idel = 5 ;连接池存活
    active = 500 ;连接池最大链接数
    timeout = 180 ; 超时时间（秒)

##### sms 配置

v1.1 版本更新

    [SMS]
    type = 2  ;type: 1, 北纬；2，阿里大于；3，希奥

    [BEIWEI]
    url = ""
    sn = ""
    pwd = ""
    ext = ""
    content = ""


    [ALIDAYU]
    url = ""
    name = ""
    key = ""
    secert = ""
    sign = ""
    template = ""

    [ALIDAYU:10]
    url = ""
    name = ""
    key = ""
    secert = ""
    sign = ""
    template = ""

    [XIAO]
    url = ""
    uid = ""
    cid = ""
    pwd = ""
    content = ""



v1.0 版本更新

    [SMS]
    type = 1  ;type: 1, 北纬；2，阿里大于；3，希奥

    [BEIWEI]
    url = ""
    sn = ""
    pwd = ""
    ext = ""
    C0 = ""  //渠道对应的短信模板
    C1 = ""
    C12 = ""

    [ALIDAYU]
    url = ""
    name = ""
    key = ""
    secert = ""
    sign = ""
    C0 = "" //渠道对应的短信模板
    C1 = ""
    C12 = ""

    [XIAO]
    url = ""
    uid = ""
    cid = ""
    pwd = ""
    C0 = "" //渠道对应的短信模板
    C1 = ""
    C36 = ""
