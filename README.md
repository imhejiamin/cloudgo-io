# cloudgo-io：处理 web 程序的输入与输出

## 1、概述

设计一个web小应用，展示静态文件服务、js 请求支持、模板输出、表单处理、Filter 中间件设计等方面的能力。（不需要数据库支持）

## 2、任务要求

编程web应用程序cloudgo-io。请在项目 README.MD 给出完成任务的证据！

### 基本要求

1)、支持静态文件服务

2)、支持简单 js 访问

3)、提交表单，并输出一个表格

4)、对 /unknown 给出开发中的提示，返回码 5xx

## 3、实验结果与测试

说明：本次实验代码使用的是"github.com/codegangsta/negroni"中提供的框架。

参考博客：http://blog.csdn.net/pmlpml/article/details/78539261 。

实验过程：

（1）启动监听端口，默认监听端口为8080，开始监听。

![1](https://github.com/imhejiamin/cloudgo-io/blob/master/test_images/1.png)

（2）打开浏览器进入到localhost：8080端口，浏览器显示出首页信息；shell监听正常，有返回码和详细信息。

![2](https://github.com/imhejiamin/cloudgo-io/blob/master/test_images/2.png)

（3）js访问。支持js访问，能够显示出hello.js内容。

![4](https://github.com/imhejiamin/cloudgo-io/blob/master/test_images/4.png)

（5）静态文件服务。能够按照路径显示静态文件并访问静态文件。

![5](https://github.com/imhejiamin/cloudgo-io/blob/master/test_images/5.png)

（6）提交表单并输出一个表格。

用/login提示输入用户名和密码，从而提交一个用户和密码的表单。点击login输出一个表格。

![6](https://github.com/imhejiamin/cloudgo-io/blob/master/test_images/6.png)

![7](https://github.com/imhejiamin/cloudgo-io/blob/master/test_images/7.png)

（7）对 /unknown 给出开发中的提示，返回码 5xx（这里用了501，表示未完成的页面）。

![9](https://github.com/imhejiamin/cloudgo-io/blob/master/test_images/9.png)





