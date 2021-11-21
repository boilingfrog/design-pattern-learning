<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [职责链模式](#%E8%81%8C%E8%B4%A3%E9%93%BE%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 职责链模式

### 定义

职责链模式(Chain Of Responsibility):使多个对象都有机会处理请求，从而避免请求的发送者和接收者之间的藕合关系。将这个对象连城一条链，并沿着这条链
传递该请求，知道一个对象处理完为止。   

是一种处理请求的模式，它让多个处理器都有机会处理该请求，直到其中某个处理成功为止。责任链模式把多个处理器串成链，然后让请求在链上传递。   

举个栗子  

<img src="/img/pattern-processor.png" alt="responsibility" />  


### 优点

### 缺点

### 适用范围

### 代码实现

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/职责链模式
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001   