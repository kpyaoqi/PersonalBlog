---
title: YAML文件详解

date: 2022-05-09	

categories: SpringBoot	

tags: [Java笔记,SpringBoot]
---	

# YAML 简介

YAML，即 ”YAML Ain’t a Markup Language“（YAML 不是一种标记语言）的递归缩写，YAML 意思其实是“ Yet Another Markup Language"（仍是一种标记语言）。该配置文件类型主要强调这种语言是主要侧重于匹配值数据，而不是以标记为中心， 以标记为中心的主要是XML语言。

YAML 可读性高，容易理解，用来表达数据序列化的格式。它的语法与其他高级语言很像，对于各种数据类型有属于自己较为简单的表达方式，它使用空白符缩进，层次分明，对于需要表达或者编辑的数据结构和各种配置文件等使用yaml文件非常方便。

YAML 配置文件后缀为**.yml**，例如application.yml、bootstrap.yml。

# 基本语法：

- 采用key :value格式，kv之间需要用空格间隔
- ‘#’表示注释
- 字符串值不使用引号，如果要加引号，” “会转义字符串中的特殊字符(例如`\n`)，‘ ’不会转义字符串中的特殊字符。
- 使用缩进来表示配置之间的层级关系
- 缩进的空格数没有具体要求，只需满足相同层级的元素左对齐
- 该文件对配置中的大小写很敏感

# 数据类型：

标量：不可再分的值，int、filoat、boolean、string、date、null

```yaml
k: v
```

对象：键值对的集合，有map、hash、set、object 

```yaml
#行内写法  
k: {k1:v1,k2:v2,k3:v3}
#分行写法
k: 
  k1: v1
  k2: v2
  k3: v3
```

数组：一组按次序排列的值，有array、list、queue

```yaml
#行内写法 
k: [v1,v2,v3]
#分行写法
k:
 - v1
 - v2
 - v3
```

配置类型

```java
@Data
public class Student {
	private String userName;
	private Boolean committee;
	private Date birth;
	private Integer age;
	private Teacher teacher;
	private String[] subjects;
	private List<String> aside;
	private Map<String, Double> score;
	private Map<String, List<Teacher>> allteachers;
}

@Data
public class Teacher {
	private String name;
	private Integer age;
}
```

yaml表示以上对象

```yaml
student:
  userName: zhangsan
  boss: false
  birth: 2022/08/9 11:11:11
  age: 18
  teacher: 
    name: xiaohong
    weight: 35
  subjects: [语文,数学,英语]
  aside: 
    - lisi
    - wangwu
  score:
    english: 66
    math: 66.5
    chinese: 66
  allteachers:
    woman:
      - {name: xioahuang}
      - {name: xiaoli,age: 47}
    male: [{name: xiaochen,age: 47}]
```

# 文本块(如果想引入多行的文本块，则使用**|**符号)

```yaml
title: |
   Hello jack!!
   I am good!
   Thanks! 
```

# 引用

用到`&`锚点和`*`星号，`&`用来建立锚点，`<<`表示合并到当前数据，`*`用来引用锚点

```yaml
zhangsan: &zhangsan
  name: 张三
  age: 18

committee: 
  position: LifeMember
  <<: *zhangsan
```

上面最终相当于如下：

```yaml
zhangsan: &zhangsan
  name: 张三
  age: 18

committee: 
  position: LifeMember
  name: 张三
  age: 18
```

还有一种文件内引用，引用已经定义好的变量，如下：

```yaml
host: http://yaoqi.com
path: ${host}/person/add  
# 最终值为 http://yaoqi.com/person/add  
```

本次分享到此结束啦~~~~
