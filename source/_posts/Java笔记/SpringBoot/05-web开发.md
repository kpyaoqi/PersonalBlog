---
title: 05-web开发

date: 2022-05-03	

categories: SpringBoot	

tags: [Java笔记,SpringBoot]
---	

## 静态资源访问

原理：当一个请求进来，先找Controller看能不能处理，不能处理又交给静态资源处理器，静态资源也找不到则响应404页面

要求静态资源放在 called /static (or /public or /resources or /META-INF/resources 类路径下，这是默认配置，需要访问的时候路径为：当前项目根路径/+静态资源名

若需要改变的静态资源路径，需在配置文件配置：

```yaml
spring:
  mvc:
    #这个会导致index.html不能默认访问
    #自定义Favicon：favicon.ico 放在静态资源目录下即可
    #此配置也会导致Favicon功能失效
    static-path-pattern: /res/**

  resources:
    static-locations: [classpath:/yaoqi/]
```

## 请求参数处理

- @PathVariable：标注在形参前面,获取请求路径中的值

- @RequestParam：将请求参数与控制器形参创建映射关系，标注在形参前面，获取请求参数名与形参名相同的

- @RequestHeader：将请求头信息与控制器形参创建映射关系，用法与@RequestParam

- @CookieValue：将Cookie数据与控制器形参创建映射关系，用法与@RequestParam

- @RequestBody：标注在控制器形参位置，请求体会赋值给形

  