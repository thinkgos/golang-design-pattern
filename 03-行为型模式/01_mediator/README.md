# 中介者模式

see [中介者模式](https://www.runoob.com/design-pattern/mediator-pattern.html)

中介者模式（Mediator Pattern）是用来降低多个对象和类之间的通信复杂性。  
这种模式提供了一个中介类，该类通常处理不同类之间的通信，并支持松耦合，使代码易于维护。  
中介者模式属于行为型模式。

中介者模式封装对象之间互交，使依赖变的简单，并且使复杂互交简单化，封装在中介者中。

例子中的中介者使用单例模式生成中介者。
中介者的change使用switch判断类型。
