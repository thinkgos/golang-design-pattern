# Go 语言设计模式

[![Build Status](https://travis-ci.org/thinkgos/golang-design-pattern.svg?branch=master)](https://travis-ci.org/thinkgos/golang-design-pattern)
[![codecov](https://codecov.io/gh/thinkgos/golang-design-pattern/branch/master/graph/badge.svg)](https://codecov.io/gh/thinkgos/golang-design-pattern)
![Action Status](https://github.com/thinkgos/golang-design-pattern/workflows/Go/badge.svg)

设计模式 Golang实现－《研磨设计模式》读书笔记  
Go 语言设计模式的实例代码

- #### 创建型模式
  - [x] 01.简单工厂模式 (Simple Factory)
  - [x] 02.工厂方法模式 (Factory Method)
  - [x] 03.抽象工厂模式 (Abstract Factory)
  - [x] 04.创建者模式 (Builder)
  - [x] 05.原型模式 (Prototype)
  - [x] 06.单例模式 (Singleton)

- #### 结构型模式

  - [x] 01.外观模式 (Facade)
  - [x] 02.适配器模式 (Adapter)
  - [x] 03.代理模式 (Proxy)
  - [ ] 04.组合模式 (Composite)
  - [x] 05.享元模式 (Flyweight)
  - [x] 06.装饰模式 (Decorator)
  - [x] 07.桥接模式 (Bridge)

- #### 行为型模式

  - [ ] 01.中介者模式 (Mediator)
  - [x] 02.观察者模式 (Observer)
  - [x] 03.命令模式 (Command)
  - [ ] 04.迭代器模式 (Iterator)
  - [x] 05.模板方法模式 (Template Method)
  - [x] 06.策略模式 (Strategy)
  - [ ] 07.状态模式 (State)
  - [ ] 08.备忘录模式 (Memento)
  - [x] 09.解释器模式 (Interpreter)
  - [ ] 10.职责链模式 (Chain of Responsibility)
  - [ ] 11.访问者模式 (Visitor)

## 六大基本原则
   
   - 单一职责原则: 指出每一个模块或类应当拥有的一个部分责任的功能由提供软件，而责任应完全封装在类
   - 开闭原则: 说对扩展开放,对修改关闭
   - 里式替换原则: 子类可以实现父类的抽象方法,但是不能覆盖父类的非抽象方法在我们做系统设计时,经常会设计接口或抽象类,然后由子类来实现抽象方法
   - 接口隔离原则: 接口应该尽量细化,一个接口对应一个功能模块,同时接口里面的方法应该尽可能的少,使接口更加轻便灵活。
   - 依赖反转原则: 是指一种特定的解耦（传统的依赖关系创建在高层次上，而具体的策略设置则应用在低层次的模块上）形式，使得高层次的模块不依赖于低层次的模块的实现细节，依赖关系被颠倒（反转），从而使得低层次模块依赖于高层次模块的需求抽象。
   - 迪米特法则: 一个实体应当尽量少的与其他实体之间发生相互作用,使得系统功能模块相对独立
