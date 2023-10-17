---
title: 1-Mybatis

date: 2022-03-29	

categories: Mybatis	

tags: [Java笔记,Mybatis]
---	

# Mybatis（基于Java的持久层的框架）

## 特性：

- 具有定制化SQL、存储过程、高级映射的持久层框架
- 避免几乎所有JDBC代码的手动设置参数及获取参数集
- 使用简单的XML或注解用于配置，可以将接口和Java的POJO映射成数据库的记录
- 半自动的ORM框架（框架格式：用元数据【XML格式】描述对象与关系映射的细节）

## 快速入门

1.引入依赖

```xml
<dependencies> 
    <!-- Mybatis核心 --> 
    <dependency> 
        <groupId>org.mybatis</groupId> 
        <artifactId>mybatis</artifactId> 
        <version>3.5.7</version> 
    </dependency> 
    <!-- MySQL驱动 --> 
    <dependency> 
        <groupId>mysql</groupId> 
        <artifactId>mysql-connector-java</artifactId> 
        <version>5.1.3</version> 
    </dependency> 
</dependencies>
```

2.创建配置文件

```xml
<?xml version="1.0" encoding="UTF-8" ?> 
<!DOCTYPE configuration PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
        "http://mybatis.org/dtd/mybatis-3-config.dtd">
<configuration> 
    <!--引入properties文件，此时就可以${属性名}的方式访问属性值-->
    <properties resource="jdbc.properties"></properties>
    <settings> 
        <!--将表中字段的下划线自动转换为驼峰-->

        <setting name="mapUnderscoreToCamelCase" value="true"/> 
        <!--开启延迟加载-->
        <setting name="lazyLoadingEnabled" value="true"/>
    </settings>
    <typeAliases> 
        <!--设置别名-->
        <package name="com.yao.mybatis.bean" alias="abc"/>
    </typeAliases>
    <!--设置连接数据库的环境-->
    <environments default="development">
        <environm id="development">
            <transactionManager type="JDBC"/>
            <dataSource type="POOLED">
                <property name="driver" value="com.mysql.jdbc.Driver"/>
                <property name="url" value="jdbc:mysql://localhost:3306/MyBatis"/>
                <property name="username" value="root"/>
                <property name="password" value="root"/>
            </dataSource>
        </environm>
    </environments> 
    <!--引入映射文件-->
    <mappers>
        <mapper resource="mappers/Mapper.xml"/>
    </mappers>
</configuration>
```

3.创建mapper接口

```java
public interface Mapper {
    int addUser();
    User getUserById(@Param("id") int id);
    User mohuSelect(@Param("String") String field)
}
```

4.创建Mybatis映射文件

```xml
<?xml version="1.0" encoding="UTF-8" ?> <!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="com.yaoqi.mybatis.mapper.Mapper"> 
<!--
	<sql语句开头单词 id="对应接口方法名"> 
        sql语句
    </sql语句开头单词>
-->
    <insert id="addUser"> 
        insert into user values (1, '张三', 23, '女') 
    </insert>
    <select id="getUserById" resultType="User"> 
        select * from user where id = #{id} 
    </select>
    <!--模糊查询-->
    <select id="mohuSelect" resultType="User"> 
        select * from user where username like "%"#{field}"%" 
    </select>
</mapper>
```

5. 功能测试

```java
//读取MyBatis的核心配置文件 
InputStream is = Resources.getResourceAsStream("config.xml"); 
//创建SqlSessionFactoryBuilder对象 
SqlSessionFactoryBuilder sqlSessionFactoryBuilder = new SqlSessionFactoryBuilder(); 
//通过核心配置文件所对应的字节输入流创建工厂类SqlSessionFactory，生产SqlSession对象 
SqlSessionFactory sqlSessionFactory = sqlSessionFactoryBuilder.build(is); 
//创建SqlSession对象，此时通过SqlSession对象所操作的sql都必须手动提交或回滚事务 
//SqlSession sqlSession = sqlSessionFactory.openSession(); 
//创建SqlSession对象，此时通过SqlSession对象所操作的sql都会自动提交 
SqlSession sqlSession = sqlSessionFactory.openSession(true); 
//通过代理模式创建UserMapper接口的代理实现类对象 
UserMapper userMapper = sqlSession.getMapper(Mapper.class); 
//调用Mapper接口中的方法，就可根据Mapper的全类名匹配元素文件，通过调用的方法名匹配 映射文件中的SQL标签，并执行标签中的SQL语句 
int result = Mapper.insertUser(); 
//sqlSession.commit();
System.out.println("结果："+result); 
```

