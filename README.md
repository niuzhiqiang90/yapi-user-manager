# yapi-user-operator
yapi-user-operator is a simple, easy and fast command line tool for managing yapi users.

[中文](./README-zh.md)

> [YApi](https://github.com/YMFE/yapi) is an efficient, easy-to-use and powerful api management platform designed to provide more elegant interface management services for developers, products and testers. It helps developers to create, publish and maintain APIs easily. YApi also provides an excellent interactive experience for users, and developers can manage interfaces by simply using the interface data writing tools and simple click operations provided by the platform.

## 1. Why do this?
The downside is that there is no entry point for the administrator to add users, and after our private deployment, the entry point is usually closed for security reasons, so if you want to add a new account, you need to
1. modify the startup configuration `"closeRegister":false`
2. restart the service
3. register the user
4. modify the configuration to close the user registration again `"closeRegister":true` 
6. restart the service again

As we can see from the above steps, a single user creation operation requires two times service restart. This is obviously not very convenient (reasonable).
If you can login to the server where the Yapi database is situated, or if you have access to the Yapi database, it would be very convenient to manage users directly in the database. It would be great if we could eliminate the need to spell out the database statements to create users and achieve the above function with a single command.


## 2. Usage
Note that the configuration in the config/config.yaml file needs to be modified according to the actual situation, no matter which of the following methods is used.

### 2.1 Source code
Golang environment is required.

```
git clone https://github.com/niuzhiqiang90/yapi-user-operator.git
cd yapi-user-operator
go run main.go add user -u xxx@xxx.com
```
Output
```
Add user successfully.
Account: xxx@xxx.com
Password: 1234qwer!@#$
Please change your password after login
```

### 2.2 Binary
Download directly from [here](https://github.com/niuzhiqiang90/yapi-user-operator/releases).
```
tar -zxvf yapi-user-operator-linux-<version>.tar.gz
cd yapi-user-operator-linux-<version>
chmod +x yapi-user-operator
./yapi-user-operator add user -u xxx@xxx.com
```
Output
```
Add user successfully.
Account: xxx@xxx.com
Password: 1234qwer!@#$
Please change your password after login.
```


