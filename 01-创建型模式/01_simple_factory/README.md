# 简单工厂模式

see [简单工厂模式](https://www.runoob.com/design-pattern/factory-pattern.html)

Go语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类.  
NewXXX函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。

在这个simple factory包中只有API 接口和NewAPI函数为包外可见,可根据不同的条件产生不同的功能,封装了实现细节。
