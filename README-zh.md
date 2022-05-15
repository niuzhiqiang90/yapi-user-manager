# yapi-user-operator
yapi-user-operator是一个简单、方便、快捷的管理yapi用户的命令行工具。

[English](./README.md)

> [YApi 是高效、易用、功能强大的 api 管理平台](https://github.com/YMFE/yapi),旨在为开发、产品、测试人员提供更优雅的接口管理服务。可以帮助开发者轻松创建、发布、维护 API，YApi 还为用户提供了优秀的交互体验，开发人员只需利用平台提供的接口数据写入工具以及简单的点击操作就可以实现接口的管理。

## 1. 缘由
有一点点美中不足的是管理员也没有添加用户的入口，而在我们私有化部署之后，出于安全方面考虑，通常会关闭用户注册入口，因此如果想要添加新的账号，则需要
1. 修改启动配置 `"closeRegister":false`
2. 重启服务
3. 用户注册
4. 再次修改配置关闭用户注册 `"closeRegister":true`
6. 再次重启服务

正如我们所见，一个创建用户操作，需要重启两次服务。这显然不太方(合)便(理)。
如果可以登录Yapi数据库所在的服务器，或者说有Yapi数据库权限，直接在数据库中进行用户管理就会变得非常方便。如果能够省去拼创建用户的数据库语句，通过一条命令实现上述功能的话，那就更棒了。


## 2. 用法
注意，不管使用如下哪种方式都需根据实际情况，修改config/config.yaml文件中的配置。

### 2.1 源码
需要go环境
```
git clone https://github.com/niuzhiqiang90/yapi-user-operator.git
cd yapi-user-operator
go run main.go add user -u xxx@xxx.com
```
输出
```
Add user success
Account: xxx@xxx.com
Password: 1234qwer!@#$
Please change your password after login
```

### 2.2 二进制包
从[这里](https://github.com/niuzhiqiang90/yapi-user-operator/releases)下载二进制包。
```
tar -zxvf yapi-user-operator-linux-<version>.tar.gz
cd yapi-user-operator-linux-<version>
chmod +x yapi-user-operator
./yapi-user-operator add user -u xxx@xxx.com
```
输出
```
Add user success
Account: xxx@xxx.com
Password: 1234qwer!@#$
Please change your password after login.
```

## 3. 待改进
1. 用户表中的_id值，参照[Yapi源码](https://github.com/YMFE/yapi/blob/master/server/models/base.js)，发现创建用户时_id并不是每次加1的顺序增长。  
目前使用随机生成的100以内整数。
```
    if (this.isNeedAutoIncrement() === true) {
      this.schema.plugin(autoIncrement.plugin, {
        model: this.name,
        field: this.getPrimaryKey(),
        startAt: 11,
        incrementBy: yapi.commons.rand(1, 10)
      });
    }
```

2. 使用了固定的密码，需用户登录后修改密码。
3. 安全的删除用户。



