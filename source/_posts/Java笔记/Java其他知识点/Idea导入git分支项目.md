---
title: Idea导入git分支项目

date: 2022-03-07	

categories: Java其他知识点	

tags: [Java笔记,Java其他知识点]
---	

# Idea导入git分支项目

因为直接通过git的clone命令默认是直接导入远程仓库的master主支的代码，无法导入分支代码

## 如何导入分支代码？

首先创建个空文件夹，打开文件夹右击出现git命令行，执行以下代码：

![image-20220727110600522](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Java笔记/Java其他知识点/img/image-20220727110600522.png) 

打开idea，在idea里打开刚才新建的项目文件夹（File->open）

右键刚才打开的项目，从远程获取代码库

![image-20220727110923223](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Java笔记/Java其他知识点/img/image-20220727110923223.png) 

选择所需要的分支代码并从远程获取代码并合并本地的版本

![image-20220727111037749](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Java笔记/Java其他知识点/img/image-20220727111037749.png) 

![image-20220727111202072](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Java笔记/Java其他知识点/img/image-20220727111202072.png) 

等待Idea加载代码，完成！

![image-20220727111659418](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Java笔记/Java其他知识点/img/image-20220727111659418.png) 

若需要新项目联合maven，只需右键文件夹，选择Add Frameworks Support—>Maven

![image-20220727112412339](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Java笔记/Java其他知识点/img/image-20220727112412339.png) 

![image-20220727112304055](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Java笔记/Java其他知识点/img/image-20220727112304055.png) 