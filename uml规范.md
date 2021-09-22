<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [如何画UML](#%E5%A6%82%E4%BD%95%E7%94%BBuml)
  - [前言](#%E5%89%8D%E8%A8%80)
  - [UML](#uml)
  - [类](#%E7%B1%BB)
  - [类的关系](#%E7%B1%BB%E7%9A%84%E5%85%B3%E7%B3%BB)
    - [1、依赖关系](#1%E4%BE%9D%E8%B5%96%E5%85%B3%E7%B3%BB)
    - [2、继承关系](#2%E7%BB%A7%E6%89%BF%E5%85%B3%E7%B3%BB)
    - [3、实现关系](#3%E5%AE%9E%E7%8E%B0%E5%85%B3%E7%B3%BB)
    - [4、关联关系](#4%E5%85%B3%E8%81%94%E5%85%B3%E7%B3%BB)
    - [5、聚合关系](#5%E8%81%9A%E5%90%88%E5%85%B3%E7%B3%BB)
    - [6、组合关系](#6%E7%BB%84%E5%90%88%E5%85%B3%E7%B3%BB)
  - [总结](#%E6%80%BB%E7%BB%93)
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

- "+" 表示public
- "-" 表示private
- "#" 表示protected

### 类的关系

#### 1、依赖关系

依赖关系，就是构造这个类的时候需要依赖其他的类，比如：动物，动物有新陈代谢，新陈代谢就需要依赖水，食物，氧气，所以动物依赖水和氧气，他们之间及时依赖关系。  

<img src="/img/uml-rely.png"  alt="uml" align=center />

#### 2、继承关系

继承（泛化）关系，它指定了子类如何特化父类的所有特征和行为。用带空心三角形的实线表示。   

图中的动物，鸟，鸭，唐老鸭之前的关系就是继承的关系   

<img src="/img/uml-inherit.png"  alt="uml" align=center />

#### 3、实现关系

一种类与接口的关系，表示类是接口所有特征和行为的实现。  

用带空心三角形的虚线表示  

例如：大雁实现了飞行的接口   

<img src="/img/uml-interface.png"  alt="uml" align=center />

棒棒糖表示法   

接口还有另一种的表示方法俗称棒棒糖表示法  

<img src="/img/uml-sugar.png"  alt="uml" align=center />

#### 4、关联关系

所谓关联关系，就是这个类有一个属性是其他类。  

用实箭线表示  

例子：比如企鹅，在每年特定的季节才会下蛋，所以需要知道气候的变化。  

<img src="/img/uml-connect.png"  alt="uml" align=center />

#### 5、聚合关系  

聚合关系表示的是一种弱的'拥有'关系，是强的关联关系;  

用带空心菱形的实线表示  

特点： 部分对象的生命周期并不由整体对象来管理。也就是说，当整体对象已经不存在的时候，部分的对象还是可能继续存在的。比如：一只大雁脱离了雁群，依然是可以继续存活的。   

<img src="/img/uml-agg.png"  alt="uml" align=center />

#### 6、组合关系

组合关系是一种强的'拥有'关系，体现了严格的部分和整体的关系，部分和整体的生命周期一样。  

用带实心菱形的实线表示，线头的两端会有数字1和2，这被成为基数。表明这一端的类可以有几个实例。比如鸟有两个翅膀。  

<img src="/img/uml-have.png"  alt="uml" align=center />

### 总结

这里通过动物这个例子，对 uml 中几种经常用到的模型做了简单的分析，总体看下来也不是很难  

这里主要是参考【大话设计模式】，有时间的话建议花时间阅读下   

### 参考

【大话设计模式】一本关于设计模式不错的书籍   