# 桥接模式

see [桥接模式](https://www.runoob.com/design-pattern/bridge-pattern.html)

桥接（Bridge）是用于把抽象化与实现化解耦，使得二者可以独立变化,使得两部分独立扩展。    
这种类型的设计模式属于结构型模式，它通过提供抽象化和实现化之间的桥接结构，来实现二者的解耦。

桥接模式类似于策略模式，区别在于策略模式封装一系列算法使得算法可以互相替换。  
而桥接属多对多关系. 定义二者可独立变化的接口

桥接模式,策略模式,适配器模式的对比:
- 适配器模式,多对一,
- 策略模式,类似适配器模式,多对一,提供策略,在运行时可以动态替换
- 桥接模式,多对多,抽象化与实现解耦,使得各自部份可以独立变化.
