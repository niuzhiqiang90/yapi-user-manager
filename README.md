# yapi-user-manager 
yapi-user-manager is a simple, easy and fast command line tool for managing yapi users.

English | [简体中文](./README-zh.md)

> [YApi](https://github.com/YMFE/yapi) is an efficient, easy-to-use and powerful api management platform designed to provide more elegant interface management services for developers, products and testers. It helps developers to create, publish and maintain APIs easily. YApi also provides an excellent interactive experience for users, and developers can manage interfaces by simply using the interface data writing tools and simple click operations provided by the platform.

## 1. Why do this?
The downside is that there is no entry point for the administrator to add users, and after our private deployment, the entry point is usually closed for security reasons, so if you want to add a new account, you need to
1. modify the startup configuration `"closeRegister":false`
2. restart the service
3. register the user
4. modify the configuration to close the user registration again `"closeRegister":true` 
6. restart the service again

As we can see, a user creation operation, which requires restarting the service twice, and seriously affect the normal use of the user, which is obviously not (very) very (not) convenient (reasonable).
If you can login the server where the Yapi database is located, or you have Yapi database privileges, it becomes very convenient to manage users directly in the database. It would be great if the database statement that spells out the creation of users could be eliminated and the above function could be achieved with a single command.

## 2. Install
Note that the configuration in the config/config.yaml file needs to be modified according to the actual situation, no matter which of the following methods is used.

### 2.1 Source code
Golang environment is required.

```
git clone https://github.com/niuzhiqiang90/yapi-user-manager.git
cd yapi-user-manager 
go run main.go add user -u name -e xxx@xxx.xxx
```

### 2.2 Binary
Download directly from [here](https://github.com/niuzhiqiang90/yapi-user-manager/releases).
```
tar -zxvf yapi-user-manager -linux-<version>.tar.gz
cd yapi-user-manager -linux-<version>
chmod +x yapi-user-manager 
```

## 3. Usage

### 3.1 Add user
```
yapi-user-manager add user -u name -e xxx@xxx.xxx
```
Output
```
Add user successfully.
Username: xxx
Account: xxx@xxx.xxx
Password: 1234qwer!@#$
Please change your password after login.
```

### 3.2 Block user
```
yapi-user-manager block user -e xxx@xxx.xxx
```

### 3.3 UnBlock user
```
yapi-user-manager unblock user -e xxx@xxx.xxx
```

### 3.4 Delete user
```
yapi-user-manager delete user -e xxx@xxx.xxx
```
