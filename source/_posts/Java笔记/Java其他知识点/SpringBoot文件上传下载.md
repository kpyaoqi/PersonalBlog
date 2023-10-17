---
title: SpringBoot文件上传下载

date: 2022-03-12	

categories: Java其他知识点	

tags: [Java笔记,Java其他知识点]
---	

# SpringBoot文件上传下载

## 后台代码

```java
@RestController
@RequestMapping("/file")
public class FileController {
    
    private static String filePath = "上传及下载文件的路径";

    //文件上传
    @PostMapping ("/upload")
    public String upload(@RequestParam("file") MultipartFile file)  {
        try {
            if (file.isEmpty()) {
                return "文件内容为空";
            }
            //获取文件名
            String filename = file.getOriginalFilename();
            //若需要修改上传的文件名，则获取后缀名，再拼接到新的filename前缀名称
            //获取后缀名
            //String suffixName = filename.substring(filename.lastIndexOf("."));
            //设置文件路径
            File dest = new File(filePath+filename);
            //检测是否存在目录,若不存在则新建
            if (!dest.getParentFile().exists()) {
                dest.getParentFile().mkdirs();
            }
            //文件写入
            file.transferTo(dest);
            return "上传成功";
        } catch (IOException e) {
            e.printStackTrace();
        }
        return "上传失败";
    }

    //文件下载
    @GetMapping("/download/{fileName}")
    public String download(HttpServletResponse response,@PathVariable("fileName") String fileName) throws IOException{
        File file = new File(filePath+fileName);
        if (file.exists()){
            // 设置下载的文件名
            response.addHeader("Content-Disposition", "attachment;fileName=" + fileName);
            byte[] buffer = new byte[1024];
            BufferedInputStream bin = new BufferedInputStream(new FileInputStream(file));
            BufferedOutputStream bout = new BufferedOutputStream(response.getOutputStream());
            int i = bis.read(buffer);
            while (i != -1) {
                bout.write(buffer, 0, i);
                i = bin.read(buffer);
                //释放资源
                if (bin != null) {
                    bin.close();
                }
                if (bout != null) {
                    bout.close();
                }
            }
            return "下载成功";
        }
        return "下载失败";
    }
}
```

## 前端代码

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>springboot文件上传与下载</title>
</head>
<body>
<div>文件上传</div>
<!--enctype="multipart/form-data":指定传输类型为二进制类型，一般上传文件时使用，上传方法也必须是“POST”-->
<form action="upload" method="POST" enctype="multipart/form-data">
    请选择文件： <input type="file" name="file"/>
    <input type="submit"/>
</form>
<hr/>
<a href="download/文件名">下载文件</a>
</body>
</html>
```

注意：springboot对上传的文件大小有默认大小不能1MB，超过1MB会出现这个错误：org.springframework.web.multipart.MultipartException。

解决方法：配置第一行是设置单个文件的大小，第二行是设置单次请求的文件的总大小，配置方法根据SpringBoot版本不同而不同。

1.Spring Boot 1.3 或之前的版本，配置:

```yaml
multipart.maxFileSize = 100Mb
multipart.maxRequestSize = 200Mb
```

2.Spring Boot 1.4 版本后配置:

```yaml
spring.http.multipart.maxFileSize = 100Mb
spring.http.multipart.maxRequestSize = 200Mb
```

3.Spring Boot 2.0 之后的版本配置修改为: 

```yaml
#单位由Mb改为MB了
spring.servlet.multipart.max-file-size = 100MB
spring.servlet.multipart.max-request-size = 200MB
```
如果想要不限制文件上传的大小，那么就把两个值都设置为-1
