## Go 常用设计模式

- 单例模式包含饿汉式和懒汉式两种实现
- 工厂模式包含简单工厂、工厂方法、抽象工厂、DI 容器
- 代理模式包含静态代理、动态代理（采用 go generate 模拟）
- 观察者模式包含观察者模式、eventbus

|  **类型**  |               **设计模式（Github）**                | **常用** |                      **博客**                      |
| :--------: | :-------------------------------------------------: | :------: | :------------------------------------------------: |
| **创建型** |         单例模式(Singleton Design Pattern)          |    ✅     |              Go 设计模式 01-单例模式               |
|            |          工厂模式(Factory Design Pattern)           |    ✅     |          Go 设计模式 02-工厂模式&DI 容器           |
|            |         建造者模式(Builder Design Pattern)          |    ✅     |             Go 设计模式 03-建造者模式              |
|            |         原型模式(Prototype Design Pattern)          |    ❌     |              Go 设计模式 04-原型模式               |
| **结构型** |           代理模式(Proxy Design Pattern)            |    ✅     | Go 设计模式 06-代理模式(generate 实现类似动态代理) |
|            |           桥接模式(Bridge Design Pattern)           |    ✅     |              Go 设计模式 07-桥接模式               |
|            |        装饰器模式(Decorator Design Pattern)         |    ✅     |             Go 设计模式 08-装饰器模式              |
|            |         适配器模式(Adapter Design Pattern)          |    ✅     |             Go 设计模式 09-适配器模式              |
|            |           门面模式(Facade Design Pattern)           |    ❌     |              Go 设计模式 10-门面模式               |
|            |         组合模式(Composite Design Pattern)          |    ❌     |              Go 设计模式 11-组合模式               |
|            |         享元模式(Flyweight Design Pattern)          |    ❌     |              Go 设计模式 12-享元模式               |
| **行为型** |         观察者模式(Observer Design Pattern)         |    ✅     |   Go 设计模式 13-观察者模式(实现简单的 EventBus)   |
|            |      模板模式(Template Method Design Pattern)       |    ✅     |              Go 模板模式 14-模板模式               |
|            |      策略模式(Strategy Method Design Pattern)       |    ✅     |              Go 设计模式 15-策略模式               |
|            | 职责链模式(Chain Of Responsibility Design Pattern)] |    ✅     |    Go 设计模式 16-职责链模式(Gin 的中间件实现)     |
|            |           状态模式(State Design Pattern)            |    ✅     |              Go 设计模式 17-状态模式               |
|            |         迭代器模式(Iterator Design Pattern)         |    ✅     |             Go 设计模式 18-迭代器模式              |
|            |         访问者模式(Visitor Design Pattern)          |    ❌     |             Go 设计模式 19-访问者模式              |
|            |         备忘录模式(Memento Design Pattern)          |    ❌     |             Go 设计模式 20-备忘录模式              |
|            |          命令模式(Command Design Pattern)           |    ❌     |              Go 设计模式 21-命令模式               |
|            |       解释器模式(Interpreter Design Pattern)        |    ❌     |             Go 设计模式 22-解释器模式              |
|            |          中介模式(Mediator Design Pattern)          |    ❌     |              Go 设计模式 23-中介模式               |