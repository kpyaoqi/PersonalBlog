---
title: mysql知识点

date: 2023-07-02	

categories: 数据库	

tags: [数据库]
---	

# MySQL(关系型数据库)

MySQL是目前最流行的[关系型数据库管理系统](https://baike.baidu.com/item/关系型数据库管理系统/696511)之一，关系数据库将数据保存在不同的表中，而不是将所有数据放在一个大仓库内，增加了速度并提高了灵活性，MySQL所使用的 SQL 语言是用于访问数据库的最常用标准化语言。

## SQL语句分类：

- DQL（数据查询语言）: 查询语句，凡是select语句都是DQL。
- DDL（数据定义语言）：create drop alter，对表结构的增删改。
- DML（数据操作语言）：insert delete update，对表当中的数据进行增删改。
- TCL（事务控制语言）：commit提交事务，rollback回滚事务。
- DCL（数据控制语言）: grant授权、revoke撤销权限等。

基本语句：

- 查看有哪些数据库：show databases;
- 查看当前使用的数据库中有哪些表：show tables;
- 查看表结构：desc 表名;

## DQL（数据查询语言）

select 字段名1 （’as‘） 别名,字段名2,字段名3,.... from 表名 where 条件;

——需要起别名时as可以省略

条件分类：等于大于小于、空值（is null）、和（and）、或者（or）、范围（bteween ... and ...）、集合（in）、模糊查询（like，%代表任意多个字符，_代表任意1个字符）、排序（order by 字段值 desc; 默认是升序，asc表示升序，desc表示降序）

**分组函数**：count 计数、sum 求和、avg 平均值、max 最大值、min 最小值

例：查询某字段的数据总数量 ”select count (字段名) from 表名;“

注意：分组函数不可直接使用在where子句当中，因为分组函数是在where执行之后执行

注意：count（*）：统计总记录条数

​			count（字段值）：统计字段中不为NULL的数据总数量

**group by 和 having**：

group by ： 按照某个字段或者某些字段进行分组。
having : having是对分组之后的数据进行再次过滤。

例：找出每个部门的平均薪资，要求显示薪资大于5000的数据

​		select deptno,avg(sal) from emp group by deptno having avg(sal) > 5000;	

**distinct:**关键字，只能出现在字段最前面，去重。

**执行顺序：**select(5)  ...   from(1)   ...   where(2)	...   group by(3)   ...   having(4)	...   order by(6)	...

### 连接查询

——在表的连接查询中有一个笛卡尔积现象，就是查询条件没有要求时，查询出来的记录数量是两张表记录数量的乘积。

**分类**：内连接：两张表是平等的，满足两张表之间条件的所有记录才能查询出来

​			外连接：两张表有主次之分，主表的数据需要全部查询出来，当副表中的数据没有与主表相匹配，副表会将数据转化为Null与主表匹配

​			自连接：自连接是一种特殊的内连接，它是指互相连接的两张表在物理上为同一张表，但在逻辑上可以分为两张表

例子：

内连接：查询每个员工的部门名称，显示员工名和部门名

```sql
select 
	e.ename,d.dname
from
	emp e
join
	dept d
on
	e.deptno = d.deptno;
```

外连接：查询员工的领导,显示员工名和领导名（右连接，即就算有些员工的bossno在领导表中不存在也显示出来，领导字段为null）

```sql
select 
	a.ename '员工', b.ename '领导'
from
	boss b
right join
	emp a
on
	a.bossno = b.empno;
```

### 子查询：select语句中嵌套select语句

前面说过分组函数不可直接使用在where子句当中，此时就需要用到子查询

例：查询工资大于平均工资的员工

​		select * from emp where sal > (select avg(sal) from emp)

### limit

取出查询数据中的部分数据(limit是mysql特有的，其他数据库中没有，不通用)，重点用于分页查询

例：取出工资前5名的员工
		select ename,sal from emp order by sal desc limit 0, 5;

注意：limit是sql语句最后执行的的一个步骤

## DDL（数据定义语言）

新建表的语句：

create table 表名(
			字段名1 数据类型,
			字段名2 数据类型,
			字段名3 数据类型,
			....
);

MySQL当中常见的字段的数据类型,( )中的类型相当于在Java中的类型：
	int			 整数型(int)
	bigint	   长整型(long)
	float		 浮点型(float double)
	char		 定长字符串(String)
	varchar	可变长字符串(StringBuffer/StringBuilder)
	date		 日期类型 (java.sql.Date)
	BLOB		Binary Large OBject (Object)，二进制大对象，存储图片、视频等流媒体信
	CLOB		Character Large OBject(Object),字符大对象,存储较大文本，比如，可以存储4G的字符串

——**char和varchar的区别：**当某个字段中的数据长度不发生改变的时候，是定长的，例如：性别、生日等都是采用char；当一个字段的数据长度不确定，例如：姓名、地址等都是采用varchar。

**添加列：**ALTER TABLE 表名 ADD 列名 数据类型;

**修改表的列名**：ALTER TABLE 表名CHANGE 原列名 新列名 数据类型(xx);

**修改数据字段类型：**ALTER TABLE 表名 MODIFY 列名 数据类型(xx);

**删除一列：**ALTER TABLE 表名 DROP 列名;

**删除表：**DROP TABLE 表名;

## 约束（*）

在创建表的时候，可以给表的字段添加不同的约束，填写在字段名后面，添加约束可以保证表中数据的合法性、有效性、完整性。

常见的约束：

- 非空约束(not null)：约束的字段不能为NULL
- 唯一约束(unique)：约束的字段不能重复
- 主键约束(primary key)：约束的字段既不能为NULL，也不能重复
- 外键约束(foreign key)

——（primary key auto_increment）：主键自增，从1开始，以1递增。

示例：创建一张id为主键并且自增的用户表

​			create table t_user(
​         		 id int primary key auto_increment,  
​	  			username varchar(255)
​       	 );

注意：一张表只能有一个主键

## DML（数据操作语言）

insert语句插入数据

​	语法格式：

​		insert into 表名(字段名1,字段名2,字段名3,....) values(值1,值2,值3,....)

​		要求：字段的数量和值的数量相同，并且数据类型要对应相同。

注意：当一条insert语句执行成功之后，表格当中必然会多一行记录。即使多的这一行记录当中某些字段是NULL，后期也没有办法在执行insert语句插入数据了，只能使用update进行更新

update 表名 set 字段名1=值1,字段名2=值2... where 条件;

delete from 表名 where 条件;

## TCL（事务控制语言）

**事务**：一个完整的业务逻辑单元，不可再分，比如：银行账户，从A账户向B账户转账100元，需要执行两条update语句，A账户-100，B账户+100。 以上两t条	DML语句必须同时成功，或者同时失败，不允许出现一条成功，一条失败.想要保证以上的两条DML语句同时成功或者同时失败，那么就要使用数据库的"事务机制"。

 事务包括四大特性（ACID）：

- 原子性：事务是最小的工作单元，不可再分。
- 一致性：必须保证多条DML语句同时成功或者同时失败。
- 隔离性：事务A与事务B之间具有隔离。
- 持久性：最后数据应该持久化到硬盘中，事务才能结束。

 事务隔离性存在隔离级别，理论上隔离级别包括4个：

​       第一级别：读未提交(read uncommitted)
​	       对方事务还没有提交，我们当前事务可以读取到对方未提交的数据。
​	       读未提交存在脏读现象：表示读到了脏数据（读到未提交的数据）。
​	   第二级别：读已提交(read committed)
​	       对方事务提交之后的数据我方可以读取到。
​	       读已提交存在的问题是：不可重复读（在多次读取范围内读到数据不一致）。
​	   第三级别：可重复读(repeatable read)
​	       这种隔离级别解决了：不可重复读问题。
​	       这种隔离级别存在的问题是：读取到的数据是幻读（在多次操作读取范围内发现数据不一致问题）。
​	   第四级别：序列化读/串行化读
​	       解决了所有问题。
​	       效率低，需要事务排队。

查看隔离级别：select @@tx_isolation;

设置隔离级别：set session|global transaction isolation level 隔离级别;

示例：

```sql
-- 关闭事务自动提交
set autocommit=0;

-- 开启事务 
start transaction;

update bal set balance = balance - 100 where balno = 'A';
update bal set balance = balance + 100 where balno = 'B';

-- 提交事务 
commit;

-- 事务回滚
rollback;
```

