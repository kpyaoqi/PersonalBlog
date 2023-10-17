---
title: 2-SpringMvc

date: 2022-05-20	

categories: SpringMvc	

tags: [Java笔记,SpringMvc]
---	

## @RequestMapping注解:将请求和控制器关联起来

标识一个类时：设置请求的初始信息

标识一个方法时：设置请求的具体信息

这个注解里面还有很多属性：

- value:设置多个请求映射的地址
- method：设置能通过的请求方式（springmvc提供了扩展注解：@GetMapping、@PostMapping、@PutMapping、@DeleteMapping）
- params：设置能通过的请求必须携带的请求参数
- headers：设置能通过的请求必须携带的请求头参数

### 请求过来了那当然就是如何获取参数了？

1. 通过ServletAPI获取：将HttpServletRequest作为控制器的形参，会封装放弃那请求的请求参数

2. 通过控制器方法获取：当控制器的形参名和请求参数名相同时，DispatcherServlet会将请求参数自动赋值给形参

3. @RequestParam：将请求参数与控制器形参创建映射关系，标注在形参前面，获取请求参数名与形参名相同的，有三个属性（value：指定形参赋值的参数名、required：设置请求参数是否必须、defaultValue：默认值为参数名）

4. @RequestHeader：将请求头信息与控制器形参创建映射关系，用法与@RequestParam

5. @CookieValue：将Cookie数据与控制器形参创建映射关系，用法与@RequestParam
6. 通过POJO获取：控制器形参设置一个实体类类型，当请求参数中的参数与该实体类中的属性一一对应，则会给其赋值

### 域对象的数据共享

1. 使用ServletAPI向request域对象共享数据:控制器设置形参”HttpServletRequest request“，调用方法request.setAttribute("AB", "data");
2. 使用ModelAndView向request域对象共享数据：创建一个新的ModelAndView()对象，调用方法：mod.addObject("AB", "data"); 设置视图，实现页面跳转 ：mod.setViewName("success"); 返回ModelAndView()对象：return mav; 
3. 使用Model向request域对象共享数据:控制器设置形参”Model model“，调用方法model.addAttribute("AB", "data");
4. 使用map向request域对象共享数据::控制器设置形参”Map<String, Object> map“,调用方法map.put("AB", "data");
5. 使用ModelMap向request域对象共享数据:控制器设置形参”ModelMap modelMap“，调用方法modelMap.addAttribute("AB", "data");
6. 向session域共享数据:控制器设置形参”HttpSession session”，调用方法session.setAttribute("AB", "data");
7. 向application域共享数据:控制器设置形参”HttpSession session”，创建session.getServletContext(); 对象，调用方法application.setAttribute("AB", "data");

### SpringMvc视图的渲染

- ThymeleafView：当控制器所请求的视图没有前缀的时候，会默认被配置文件中的视图解析器根据前缀和后缀拼接解析，然后转发

- 转发视图：当控制器的请求的视图含有前缀“forword:”时，不会经过视图解析器，而是将前缀去掉，转发到剩下的路径

- 重定向视图：当控制器的请求的视图含有前缀“redirect:”时，不会经过视图解析器，而是将前缀去掉，重定向到剩下的路径

  （当控制器方法中只需要实现页面跳转，则可以设置view controller标签）

```xml
<!--
path：设置处理的请求地址 
view-name：设置请求地址所对应的视图名称 
--> 
<mvc:view-controller path="/router" view-name="success"></mvc:view-controller>
<!--开启注解驱动-->
<mvc:annotation-driven />
```

### 请求体的转换

1. @RequestBody:标注在控制器形参位置，请求体会赋值给形参
2. RequestEntity：是一种封装请求报文的类型，标注在形参位置，请求报文会赋值给形参，可以通过getHeaders()获取请求头信息，getBody()获取请求体信息等
3. @ResponseBody：标注在控制器方法，将该方法的返回体直接响应在浏览器上
4. ResponseEntity：用于控制器方法的返回值类型，该控制器方法的返回值就是响应到浏览器的响应报文
5. @RestController：标识在控制器的类上，就相当于为类添加了@Controller注解，并且为其中的每个方法添加了@ResponseBody注解

### 拦截器和异常处理器

在配置文件中配置拦截器

```xml
<bean class="com.yaoqqi.interceptor.FirstInterceptor"></bean> 
<ref bean="firstInterceptor"></ref> 
<!-- 以上两种配置方式都是对DispatcherServlet所处理的所有的请求进行拦截 --> 
<mvc:interceptor> 
    <mvc:mapping path="/**"/> 
    <mvc:exclude-mapping path="/testRequestEntity"/> 
    <ref bean="firstInterceptor"></ref> 
</mvc:interceptor> 
<!--
    以上配置方式可以通过ref或bean标签设置拦截器，通过mvc:mapping设置需要拦截的请求，通过 
    mvc:exclude-mapping设置需要排除的请求，即不需要拦截的请求 
-->
```

拦截器有三个抽象方法：

- preHandle：控制器方法执行之前执行preHandle()，返回值boolean类型表示是否令其通过，返回true为通过并调用控制器方法；返回false表示拦截，即不调用控制器方法
- postHandle：控制器方法执行之后执行postHandle()
- afterComplation：处理完视图和模型数据，渲染视图完毕之后执行afterComplation()

### 异常处理器

```java
//@ControllerAdvice将当前类标识为异常处理的组件 
@ControllerAdvice 
public class ExceptionController { 
    //@ExceptionHandler用于设置所标识方法处理的异常 
    @ExceptionHandler(ArithmeticException.class) 
    //ex表示当前请求处理中出现的异常对象 
    public String handleArithmeticException(Exception ex, Model model){ 
        model.addAttribute("ex", ex); 
        return "error"; 
    } 
}
```

