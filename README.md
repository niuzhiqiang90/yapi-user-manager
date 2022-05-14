# yapi-user-operator
[YApi 是高效、易用、功能强大的 api 管理平台](https://github.com/YMFE/yapi)。
有一点点美中不足的是管理员也没有添加用户的入口，而在我们私有化部署之后，出于安全方面考虑，通常会关闭用户注册入口，因此如果想要添加新的账号，则需要
1. 修改启动配置 `"closeRegister":false`
2. 重启服务
3. 用户注册
4. 再次修改配置关闭用户注册 `"closeRegister":true`
6. 再次重启服务

上述步骤可见，一个创建用户操作，需要重启两次服务。这显然不太方(合)便(理)。
如果可以登录Yapi数据库所在的服务器，或者说有Yapi数据库权限，直接在数据库中进行用户管理就会变得非常方便。出于运维的职业习惯，如果能够省去拼创建用户的数据库语句，能够通过一条命令实现上述功能的话，那就更棒了。
