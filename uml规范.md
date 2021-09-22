<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [如何画UML](#%E5%A6%82%E4%BD%95%E7%94%BBuml)
  - [前言](#%E5%89%8D%E8%A8%80)
  - [UML](#uml)
  - [类](#%E7%B1%BB)
  - [类的关系](#%E7%B1%BB%E7%9A%84%E5%85%B3%E7%B3%BB)
    - [1、依赖关系](#1%E4%BE%9D%E8%B5%96%E5%85%B3%E7%B3%BB)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 如何画UML

### 前言

最近在学习设计模式，其中不免涉及到 UML,这里来复习下 UML 是如何画的。  

### UML 

这里根据【大话设计模式中】中动物和鸟关系的例子，重新画了这个关系的 uml 类图    

<img src="/img/uml-demo.png"  alt="uml" align=center />

这里根据上面的这个例子，我们一一来展开分析  

### 类

类是具有相似结构、行为和关系的一组对象的描述符，是面向对象系统中最重要的构造块  

<img src="/img/uml-class.png"  alt="uml" align=center />

上面的图片，从第一格往下面分析  

- 第一层 显示类的名称，如果是抽象类就用斜体表示  

- 第二层 类的特性，通常是字段和属性  

- 第三层 类的操作，通常是方法和行为

前面的符号需要我们特殊注意  

- + 表示public
- - 表示private
- # 表示protected

### 类的关系

#### 1、依赖关系

依赖关系，就是构造这个类的时候需要依赖其他的类，比如：动物，动物有新陈代谢，新陈代谢就需要依赖水，食物，氧气，所以动物依赖水和氧气，他们之间及时依赖关系。  




### 参考

【大话设计模式】一本关于设计模式不错的书籍   