---
title: 3-AOP

date: 2022-04-19	

categories: Spring	

tags: [Java笔记,Spring]
---	

### AOP 的作用及其优势

作用：在程序运行期间，在不修改源码的情况下对方法进行功能增强

优势：减少重复代码，提高开发效率，并且便于维护

### AOP 的底层实现

实际上，AOP 的底层是通过 Spring 提供的的动态代理技术实现的。在运行期间，Spring通过动态代理技术动态的生成代理对象，代理对象方法执行时进行增强功能的介入，在去调用目标对象的方法，从而完成功能的增强。

### AOP 的动态代理技术

第一种 有接口情况，使用 JDK 动态代理（创建接口实现类代理对象，增强类的方法）

第二种 没有接口情况，使用 CGLIB 动态代理（创建子类的代理对象，增强类的方法）

（1）创建接口，定义方法

```java
public interface UserDao {
 public int add(int a,int b);
}
```

（2）创建接口实现类，实现方法

```java
public class UserDaoImpl implements UserDao {
 @Override
 public int add(int a, int b) {
	 return a+b;
 }
}
```

（3）使用 Proxy 类创建接口代理对象

```java
public class JDKProxy {
 public static void main(String[] args) {
 //创建接口实现类代理对象
     Class[] interfaces = {UserDao.class};
     UserDaoImpl userDao = new UserDaoImpl();
     UserDao dao = (UserDao)Proxy.newProxyInstance(JDKProxy.class.getClassLoader(), interfaces, new UserDaoProxy(userDao)); 
     int result = dao.add(1, 2);
     System.out.println("result:"+result);
 } 
}

//创建代理对象代码
class UserDaoProxy implements InvocationHandler {
     //1 把创建的是谁的代理对象，把谁传递过来
     //有参数构造传递
     private Object obj;
     public UserDaoProxy(Object obj) {
     this.obj = obj;
 }

 //增强的逻辑
 @Override
 public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
     System.out.println("方法之前执行....");
     //被增强的方法执行
     Object result = method.invoke(obj, args);
     System.out.println("方法之后执行....");
     return result;
 } 
}
```

### AOP 的相关术语

常用的术语如下：
Target（目标对象):代理的目标对象
Proxy （代理):一个类被 AOP 织入增强后，就产生一个结果代理类
Joinpoint（连接点):所谓连接点是指那些被拦截到的点。在spring中,这些点指的是方法，因为spring只支持方法类型的连接点
Pointcut（切入点):所谓切入点是指我们要对哪些 Joinpoint 进行拦截的定义
Advice（通知/ 增强):所谓通知是指拦截到 Joinpoint 之后所要做的事情就是通知
Aspect（切面):是切入点和通知（引介）的结合
Weaving（织入):是指把增强应用到目标对象来创建新的代理对象的过程。spring采用动态代理织入，而AspectJ采用编译期织入和类装载期织入

### AOP 操作（AspectJ 注解）

1、创建类，在类里面定义方法

```java
public class User {
 public void add() {
	 System.out.println("add starting...");
 } 
}
```

2、创建增强类（编写增强逻辑）

（1）在增强类里面，创建方法，让不同方法代表不同通知类型

```java
//增强的类
public class UserProxy {
 public void before() {
 	System.out.println("before starting...");
 } 
}
```

3、进行通知的配置

（1）在 spring 配置文件中，开启注解扫描

```xml
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans" 
 xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
 xmlns:context="http://www.springframework.org/schema/context" 
 xmlns:aop="http://www.springframework.org/schema/aop" 
 xsi:schemaLocation="http://www.springframework.org/schema/beans 
http://www.springframework.org/schema/beans/spring-beans.xsd 
 http://www.springframework.org/schema/context 
http://www.springframework.org/schema/context/spring-context.xsd 
 http://www.springframework.org/schema/aop 
http://www.springframework.org/schema/aop/spring-aop.xsd">

 <!-- 开启注解扫描 -->
 <context:component-scan base-package="com.yaoqi.package"></context:component-scan> 
```

（2）使用注解创建 User 和 UserProxy 对象

（3）在增强类上面添加注解 @Aspect

```java
//增强的类
@Component
@Aspect //生成代理对象
public class UserProxy {
```

（4）在 spring 配置文件中开启生成代理对象

<!-- 开启 Aspect 生成代理对象--> 

```xml
<aop:aspectj-autoproxy></aop:aspectj-autoproxy> 
```

4、配置不同类型的通知

（1）在增强类的里面，在作为通知方法上面添加通知类型注解，使用切入点表达式配置

```java
@Component
@Aspect 
public class UserProxy {
 //前置通知
 @Before(value = "execution(*com.yaoqi.package.User.add(..))")
 public void before() {
	 System.out.println("before.........");
 }

 //后置通知
 @AfterReturning(value = "execution(*com.yaoqi.package.User.add(..))")
 public void afterReturning() {
	 System.out.println("afterReturning.........");
 }

 //最终通知
 @After(value = "execution(*com.yaoqi.package.User.add(..))")
 public void after() {
 	System.out.println("after.........");
 }

 //异常通知
 @AfterThrowing(value = "execution(*com.yaoqi.package.User.add(..))")
 public void afterThrowing() {
 	System.out.println("afterThrowing.........");
 }

 //环绕通知
 @Around(value = "execution(*com.yaoqi.package.User.add(..))")
 public void around(ProceedingJoinPoint proceedingJoinPoint) throws Throwable {
     System.out.println("环绕之前.........");
     //被增强的方法执行
     proceedingJoinPoint.proceed();
     System.out.println("环绕之后.........");
 } 
}
```

5、相同的切入点抽取

```java
//相同切入点抽取
@Pointcut(value = "execution(*com.yaoqi.package.User.add(..))")
public void commondemo() {
}

//前置通知
@Before(value = "commondemo()")
public void before() {
 System.out.println("before.........");
}
```

6、有多个增强类多同一个方法进行增强，设置增强类优先级

（1）在增强类上面添加注解 @Order(数字类型值)，数字类型值越小优先级越高

```java
@Component
@Aspect
@Order(1)
public class PersonProxy
```

7、完全使用注解开发 

（1）创建配置类，不需要创建 xml 配置文件 

```java
@Configuration
@ComponentScan(basePackages = {"com.yaoqi"})
@EnableAspectJAutoProxy(proxyTargetClass = true)
public class ConfigAop {
}
```

#### AOP 操作（AspectJ 配置文件）

1、创建两个类，增强类和被增强类，创建方法

2、在 spring 配置文件中创建两个类对象

```xml
<bean id="book" class="com.yaoqi.Book"></bean> 
<bean id="bookProxy" class="com.yaoqi.BookProxy"></bean> 
```

3、在 spring 配置文件中配置切入点

```xml
<!--配置 aop 增强--> 
<aop:config>
 <!--切入点-->
 <aop:pointcut id="p" expression="execution(*com.yaoqi..Book.buy(..))"/>
 <!--配置切面-->
 <aop:aspect ref="bookProxy">
 <!--增强作用在具体的方法上-->
     <aop:before method="before" pointcut-ref="p"/>
 </aop:aspect>
</aop:config>


```

