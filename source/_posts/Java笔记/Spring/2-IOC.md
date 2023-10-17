---
title: 2-IOC

date: 2022-04-13	

categories: Spring	

tags: [Java笔记,Spring]
---	

#### IOC操作Bean管理

##### 1、基于 xml方式创建对象

（1）在 spring 配置文件中，使用 bean 标签，标签里面添加对应属性，就可以实现对象创建

（2）在 bean 标签有很多属性，介绍常用的属性

​			id 属性：唯一标识

​			class 属性：类全路径（包类路径）

（3）创建对象时候，默认也是执行无参数构造方法完成对象创建

##### 2、基于 xml方式注入属性

​    DI：依赖注入，就是注入属性

##### 3、第一种注入方式：使用 set 方法进行注入

（1）创建类，定义属性和对应的 set 方法

```java
public class User {
 //创建属性
 private String Name;
 private String Age;
 //创建属性对应的 set 方法
 public void setName(String Name) {
 this.Name = Name;
 }
 public void setAge(String Age) {
 this.Age = Age;
 } 
}
```

（2）在 spring 配置文件配置对象创建，配置属性注入

```xml
<bean id="User" class="com.yaoqi.bean.User">
 <!--使用 property 完成属性注入
     name：类里面属性名称
     value：向属性注入的值
 -->
 <property name="Name" value="张三"></property>
 <property name="Age" value="1"></property>
</bean> 
```

##### 4、第二种注入方式：使用有参数构造进行注入

（1）创建类，定义属性，创建属性对应有参数构造方法

```java
/使用有参数构造注入/使用有参数构造注入/
public class User {
 //属性
 private String name;
 private String age;
 //有参数构造
 public User(String name,String age) {
 this.name = name;
 this.age = age;
 } 
}
```

（2）在 spring 配置文件中进行配置

```xml
<!--3 有参数构造注入属性--> 
<bean id="User" class="com.yaoqi.bean.User">
 <constructor-arg name="name" value="张三"></constructor-arg>
 <constructor-arg name="age" value="18"></constructor-arg>
</bean> 
```

#### 注入其他类型属性:

1、字面量

（1）null 值

```xml
<property name="属性名">
 	<null/>
</property> 
```

（2）属性值包含特殊符号

-   把<>进行转义 < >

-   把带特殊符号内容写到 CDATA


```xml
<property name="属性名">
    <!--设置值为<<深圳>>-->
 <value><![CDATA[<<深圳>>]]></value>
</property> 
```

2、注入属性-外部 bean

（1）创建两个类 service 类和 dao 类 

（2）在 service 调用 dao 里面的方法

（3）在 spring 配置文件中进行配置

```java
public class UserService {
 //创建 UserDao 类型属性，生成 set 方法
 private UserDao userDao;
 public void setUserDao(UserDao userDao) {
	 this.userDao = userDao;
 }
    
 public void add() {
     System.out.println("service add...............");
     userDao.update();
 } 
}
```

```xml
<bean id="userService" class="com.yaoqi.service.UserService">
 	<property name="userDao" ref="userDaoImpl"></property>
</bean> 
<bean id="userDaoImpl" class="com.yaoqi.dao.UserDaoImpl"></bean> 
```

3、注入属性-内部 bean

（1）一对多关系：部门和员工

（2）在实体类之间表示一对多关系，员工表示所属部门，使用对象类型属性进行表示

```java
//部门类
public class Dept {
 private String dname;
 public void setDname(String dname) {
	 this.dname = dname;
 } 
}

//员工类
public class Emp {
     private String ename;
     private String gender;
     private Dept dept;
    
     public void setDept(Dept dept) {
     	this.dept = dept;
     }
     public void setEname(String ename) {
    	 this.ename = ename;
     }
     public void setGender(String gender) {
    	 this.gender = gender; 
     } 
}
```

（3）在 spring 配置文件中进行配置

```xml
<bean id="emp" class="com.yaoqi.bean.Emp">
 <!--设置两个普通属性-->
 <property name="ename" value="张三"></property>
 <property name="gender" value="男"></property>
 <!--设置对象类型属性-->
 <property name="dept">
 	<bean id="dept" class="com.yaoqi.bean.Dept">
     	<property name="dname" value="销售部"></property>
    </bean>
 </property>
</bean> 
```

4、注入属性-级联赋值 

（1）第一种写法 

```xml
<bean id="emp" class="com.yaoqi.bean.Emp">
 <!--设置两个普通属性-->
 <property name="ename" value="张三"></property>
 <property name="gender" value="男"></property>
 <!--级联赋值-->
 <property name="dept" ref="dept"></property>
</bean> 

<bean id="dept" class="com.yaoqi.bean.Dept">
 <property name="dname" value="销售部"></property>
</bean> 
```

（2）第二种写法

```xml
<bean id="emp" class="com.yaoqi.bean.Emp">
 <!--设置两个普通属性-->
 <property name="ename" value="张三"></property> 
 <property name="gender" value="男"></property>
 <!--级联赋值-->
 <property name="dept.dname" value="销售部"></property>
</bean> 
```

####  注入集合属性:

1、注入数组类型属性

2、注入 List 集合类型属性

3、注入 Map 集合类型属性

（1）创建类，定义数组、list、map、set 类型属性，生成对应 set 方法

```java
public class Stu {
 //1 数组类型属性
 private String[] courses;
 //2 list 集合类型属性
 private List<String> list;
 //3 map 集合类型属性
 private Map<String,String> maps;
 //4 set 集合类型属性
 private Set<String> sets;

 public void setSets(Set<String> sets) {
 	this.sets = sets;
 }

 public void setCourses(String[] courses) {
 	this.courses = courses;
 }

 public void setList(List<String> list) {
 	this.list = list;
 }

 public void setMaps(Map<String, String> maps) {
 	this.maps = maps;
 } 
}
```

（2）在 spring 配置文件进行配置

```xml
<bean id="stu" class="com.yaoqi.bean.Stu">
 <!--数组类型属性注入-->
 <property name="courses">
     <array>
         <value>A</value>
         <value>B</value>
     </array>
 </property>

 <!--list 类型属性注入-->
 <property name="list">
     <list> 
         <value>A</value>
         <value>B</value>
     </list>
 </property>

 <!--map 类型属性注入-->
 <property name="maps">
     <map>
         <entry key="A" value="a"></entry>
         <entry key="B" value="b"></entry>
     </map>
 </property>

 <!--set 类型属性注入-->
 <property name="sets">
     <set>
         <value>A</value>
         <value>B</value>
     </set>
 </property>
</bean> 
```

4、在集合里面设置对象类型值

```xml
<!--创建多个class对象--> 
<bean id="class1" class="com.yaoqi.bean.Class">
 	<property name="cname" value="语文"></property>
</bean> 
<bean id="class2" class="com.yaoqi.bean.Class">
 	<property name="cname" value="数学"></property>
</bean>

<!--注入 list 集合类型，值是对象--> 
<property name="classList">
     <list>
         <ref bean="class1"></ref>
         <ref bean="class2"></ref>
     </list>
</property> 
```

#### 工厂Bean（FactoryBean）：

1、Spring 有两种类型 bean，一种普通 bean，另外一种工厂 bean（FactoryBean） 

2、普通 bean：在配置文件中定义 bean 类型就是返回类型

3、工厂 bean：在配置文件定义 bean 类型可以和返回类型不一样

第一步 创建类，让这个类作为工厂 bean，实现接口 FactoryBean

第二步 实现接口里面的方法，在实现的方法中定义返回的 bean 类型

```java
public class MyBean implements FactoryBean<Class> {
 //定义返回 bean
 @Override
 public Course getObject() throws Exception {
     Course course = new Course();
     course.setCname("英语");
     return course;
 }

 @Override
 public Class<?> getObjectType() {
	 return null;
 }

 @Override
 public boolean isSingleton() {
	 return false;
 } 
}
```

```xml
<bean id="myBean" class="com.yaoqi.factorybean.MyBean">
</bean>
```

```java
@Test
public void test3() {
     ApplicationContext context = new ClassPathXmlApplicationContext("bean3.xml");
     Course course = context.getBean("myBean", Course.class);
     System.out.println(course);
}
```

#### Bean作用域：

scope:指对象的作用范围，取值如下：

| 取值范围       | 说明                                                         |
| -------------- | ------------------------------------------------------------ |
| singleton      | 默认值，单例的                                               |
| prototype      | 多例的                                                       |
| request        | WEB 项目中，Spring 创建一个 Bean 的对象，将对象存入到 request 域中 |
| session        | WEB 项目中，Spring 创建一个 Bean 的对象，将对象存入到 session 域中 |
| global session | WEB 项目中，应用在 Portlet 环境，如果没有 Portlet 环境那么globalSession 相当于 session |

#### Bean 生命周期：

1. 执行无参数构造创建Bean实例

2. 调用set方法设置属性值

   —初始之前执行的方法

3. 执行初始化的方法

   —初始之后执行的方法

4. 获取创建Bean实例对象

5. 执行销毁的方法

（1）创建类，实现接口 BeanPostProcessor，创建后置处理器

```java
public class MyBeanPost implements BeanPostProcessor {
 @Override
 public Object postProcessBeforeInitialization(Object bean, String beanName) throws BeansException {
     System.out.println("这是在初始化之前执行的方法");
     return bean;
 }

 @Override
 public Object postProcessAfterInitialization(Object bean, String beanName) throws BeansException {
     System.out.println("这是在初始化之后执行的方法");
     return bean;
 } 
}
```

```xml
<!--配置后置处理器--> 
<bean id="myBeanPost" class="com.yaoqi.bean.MyBeanPost"></bean>
```

#### 外部属性文件：

把外部 properties 属性文件引入到 spring 配置文件中引入 context 名称空间

![image-20220628173255831](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Java笔记/Spring/img/image-20220628173255831.png) 

```xml
<beans xmlns="http://www.springframework.org/schema/beans"  
       xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
       xmlns:p="http://www.springframework.org/schema/p" 
       xmlns:util="http://www.springframework.org/schema/util" 
       	xmlns:context="http://www.springframework.org/schema/context" 
       xsi:schemaLocation="http://www.springframework.org/schema/beans 
http://www.springframework.org/schema/beans/spring-beans.xsd 
http://www.springframework.org/schema/util 
http://www.springframework.org/schema/util/spring-util.xsd
<!--context空间名称-->                          
http://www.springframework.org/schema/context 
http://www.springframework.org/schema/context/spring-context.xsd"> 
```

 在 spring 配置文件使用标签引入外部属性文件

```xml
<context:property-placeholder location="classpath:jdbc.properties"/>
<!--配置连接池--> 
<bean id="dataSource" class="com.alibaba.druid.pool.DruidDataSource">
     <property name="driverClassName" value="${prop.driverClass}"></property>
     <property name="url" value="${prop.url}"></property>
     <property name="username" value="${prop.userName}"></property>
     <property name="password" value="${prop.password}"></property>
</bean>
```

#### Spring注解

@Component	 使用在类上用于实例化Bean
@Controller	    使用在web层类上用于实例化Bean
@Service			 使用在service层类上用于实例化Bean
@Repository	  使用在dao层类上用于实例化Bean
@Autowired	   使用在字段上用于根据类型依赖注入
@Qualifier	      结合@Autowired一起使用用于根据名称进行依赖注入
@Resource	     相当于@Autowired+@Qualifier，按照名称进行注入
@Value	            注入普通属性
@Scope			   标注Bean的作用范围

@Configuration	用于指定当前类是一个 Spring 配置类，当创建容器时会从该类上加载注解
@Bean	使用在方法上，标注将该方法的返回值存储到 Spring 容器中
@PropertySource	用于加载.properties 文件中的配置
@Import	用于导入其他配置类

使用注解进行开发时，需要在applicationContext.xml中配置组件扫描，作用是指定哪个包及其子包下的Bean需要进行扫描以便识别使用注解配置的类、字段和方法，或者使用@ComponentScan注解，用于指定 Spring 在初始化容器时要扫描的包。 

```xml
<!--注解的组件扫描-->
<context:component-scan base-package="com.yaoqi"></context:component-scan>
```
