# Go代码

+ skills: 项目开发中的一些基础技能
+ vblog: Vblog项目（单体服务）
+ devcloud-mini: Devcloud-mini（微服务项目）
+ devcloud: 业务项目（Devcloud完整版）

## 基础过度到项目的重点

+ 基础篇：基础技能（用法+原理）
+ 项目篇：软件工程（系统性解决问题的方法，软件设计）

## 项目课体量带来的问题

1. 项目中用到的工具 都是及时更新，20%-30%技术更新
2. 基础篇（单文件，单包开发）；项目课 代码量至少上千，几十个文件或者包共同组成一个工程，项目课代码量多，逻辑复杂
3. 程序跑不起来是正常情况，写的代码一遍能跑起来是惊喜。需要学会自己debug
4. 项目课程中代码多，流程复杂，代码是现场写的，是有几率出现问题，需要现场debug

## 软件开发生命周期

1. 老板的商业洞察，立项（BP），使用对象，市场前景，商业模型（怎么赚钱），开发的方向
2. 产品设计：规划产品
    1. 需求收集：用户的痛点，用户需要的是什么，给谁用的。
    2. 系统设计：前端负责人/后端负责热 测试负责人/运维负责人
       + 概要设计：基本流程跑通，关键字段定义请求，只有一个想清楚业务流程
         + 流程设计：使用流程说明，可以参考竞品
         + 用户界面设计：UI/交互，可以参考竞品
         + 软件架构设计：采用单体服务还是微服务。哪些模块。数据库，是否需要缓存。技术栈和架构设计是一体的。
       + 详细设计：
         + 每个模块，每个具体功能的定义是什么，具体到能出页面，能定义接口，需要评审
         + 产品原型，高保真：使用专门的产品工具：蓝图、axure、figma
    3. 产品研发：
       + 前端：根据原型开发界面，软件对接后端，一起发布进行测试
       + 后端：后端研发负责人，有参与系统架构设计，任务分解
         + 整体规划：服务交互流程，接口定义：
         + 服务A：Vblog
         + 服务B：cmdb
         + 服务C：用户中心
         + 开发环境联合测试
    4. 产品测试 等着版本进行提测
       + 集成测试：功能测试/回归测试/安全测试/性能测试
    5. 运维发布：上线 配置域名 开放给 公网用户使用
       + 灰度验证：单元使用一个业务账号来进行线上测试，内部账号进行测试
       + 灰度发布：0% ~ 100% 到新版本
       + A/B 

产品迭代: 需求 ---> 开发 ---> 验证 ---> 上线

