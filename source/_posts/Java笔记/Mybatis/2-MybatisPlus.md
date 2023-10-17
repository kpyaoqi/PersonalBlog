---
title: 2-MybatisPlus

date: 2022-04-01	

categories: Mybatis	

tags: [Java笔记,Mybatis]
---	

# MybatisPlus

## 特性：

- **无侵入**：只做增强不做改变，引入它不会对现有工程产生影响，如丝般顺滑
- **损耗小**：启动即会自动注入基本 CURD，性能基本无损耗，直接面向对象操作
- **强大的 CRUD 操作**：内置通用 Mapper、通用 Service，仅仅通过少量配置即可实现单表大部分 CRUD 操作，更有强大的条件构造器，满足各类使用需求
- **支持 Lambda 形式调用**：通过 Lambda 表达式，方便的编写各类查询条件，无需再担心字段写错
- **支持主键自动生成**：支持多达 4 种主键策略（内含分布式唯一 ID 生成器 - Sequence），可自由配置，完美解决主键问题
- **支持 ActiveRecord 模式**：支持 ActiveRecord 形式调用，实体类只需继承 Model 类即可进行强大的 CRUD 操作
- **支持自定义全局通用操作**：支持全局通用方法注入（ Write once, use anywhere ）
- **内置代码生成器**：采用代码或者 Maven 插件可快速生成 Mapper 、 Model 、 Service 、 Controller 层代码，支持模板引擎，更有超多自定义配置等您来使用
- **内置分页插件**：基于 MyBatis 物理分页，开发者无需关心具体操作，配置好插件之后，写分页等同于普通 List 查询
- **分页插件支持多种数据库**：支持 MySQL、MariaDB、Oracle、DB2、H2、HSQL、SQLite、Postgre、SQLServer 等多种数据库
- **内置性能分析插件**：可输出 SQL 语句以及其执行时间，建议开发测试时启用该功能，能快速揪出慢查询
- **内置全局拦截插件**：提供全表 delete 、 update 操作智能分析阻断，也可自定义拦截规则，预防误操作

## 快速入门

1.导入依赖

```xml
<!-- 数据库驱动 -->
<dependency>
    <groupId>mysql</groupId>
    <artifactId>mysql-connector-java</artifactId>
</dependency> 
<!-- lombok -->
<dependency>
    <groupId>org.projectlombok</groupId>
    <artifactId>lombok</artifactId>
</dependency> 
<!-- mybatis-plus --> 
<dependency>
    <groupId>com.baomidou</groupId>
    <artifactId>mybatis-plus-boot-starter</artifactId>
    <version>版本号</version>
</dependency>
```

2.操作对象

```java
@Data
@AllArgsConstructor
@NoArgsConstructor
public class User {
    private Long id;
    private String name;
    private Integer age;
    private String email;
}
```

3.配置数据源

```yaml
spring.datasource.username=root 
spring.datasource.password=root 
spring.datasource.url=jdbc:mysql://localhost:3306/mybatis_plus?useSSL=false&useUnicode=true&characterEncoding=utf-8&serverTimezone=GMT%2B8 
spring.datasource.driver-class-name=com.mysql.cj.jdbc.Driver
```

4.mapper接口

```java
// 在对应的Mapper上面继承基本的类BaseMapper 
@Repository
public interface UserMapper extends BaseMapper<User> { 
    //CRUD操作都已经编写完成了 
}
```

5.扫描包（主启动类上添加注解@MapperScan("com.yaoqi.mapper")）

6.测试

```java
@Autowired 
private UserMapper userMapper; 

@Test 
void Tests() { 
    // 参数是一个 Wrapper，条件构造器，这里我们先不用 null 
    // 查询全部用户 
    List<User> users = userMapper.selectList(null); 
    users.forEach(System.out::println); 
} 
```

## 扩展

1.主键生成设置：实体类字段上添加@TableId（type=IdType.AUTO）

2.自动生成时间：

```java
// 字段添加填充内容 
@TableField(fill = FieldFill.INSERT) 
private Date createTime; 

@TableField(fill = FieldFill.INSERT_UPDATE) 
private Date updateTime; 
```

```java
@Component
public class tHandler implements MetaObjectHandler {
    @Override
    public void insertFill(MetaObject metaObject) {
        this.setFieldValByName("createTime", new Date(), metaObject);
        this.setFieldValByName("updateTime", new Date(), metaObject);
    }

    @Override
    public void updateFill(MetaObject metaObject) {
        this.setFieldValByName("updateTime", new Date(), metaObject);
    }
}
```

3.设置乐观锁：

数据库增加Version字段

```java
@Version
private Integer version; 
```

```java
@MapperScan("com.yaoqi.mapper")
@EnableTransactionManagement
@Configuration
public class MyBatisPlusConfig {
    @Bean
    public OptimisticLockerInterceptor optimisticLockerInterceptor() {
        return new OptimisticLockerInterceptor();
    }
}
```

4.分页查询

```java
@Bean 
public PaginationInterceptor paginationInterceptor() { 
	return new PaginationInterceptor(); 
}
```

```java
@Test
public void testPage() {
    // 参数一：当前页
    // 参数二：页面大小 
    Page<User> page = new Page<>(2, 5);
    List<User> users=userMapper.selectPage(page, null);
    users.forEach(System.out::println);
}
```

5.逻辑删除（没有从数据库删除，通过一个变量来让数据失效）

在数据库新增一个deleted字段

```java
@TableLogic //逻辑删除 
private Integer deleted; 
```

```java
@Bean 
public ISqlInjector sqlInjector() { 
	return new LogicSqlInjector(); 
}
```

```yml
mybatis-plus.global-config.db-config.logic-delete-value=1 
mybatis-plus.global-config.db-config.logic-not-delete-value=0
```

执行的是更新操作不是删除操作！

## **条件构造器（Wrapper）**

```java
QueryWrapper<User> wrapper = new QueryWrapper<>(); 
wrapper.查询方式(参数1，参数2); 
Mapper.selectCount(wrapper);
```



![image-20220715104450923](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Java笔记/Mybatis/img/image-20220715104450923.png) 
