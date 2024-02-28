
// JS变量作用域
/*
{
    var a = 100

    {
        console.log(a)
    }

    console.log(a)
}
*/

// 解构赋值

a1 = [1,2,3]
x = a1[0]
y = a1[1]
z = a1[2]

// 变量数组 <-map-> value数组
[x, [y, z]] = [1, [2, 3]]
console.log(x,y,z)

// 对象的解构赋值
b1 = {name: '张三', age: 10}
// console.log({name, age})

var {name, age} = b1
// name = b1[name]
// age = b1[age]
console.log(name, age)


// go fmt （fmt.printf("你好！ %s, 我是%s", "张三", "李四")）
console.log('你好！' + name + '，我是李四')
console.log(`你好！${name}，我是李四`)

// 错误处理
var a = 'test'
// 中间逻辑，不小心把a 修改位null 或者undefine了

// 通过try catch 捕获代码片段异常
try {
    a = null
    console.log(a[0])
} catch (error) {
    // 这里没有把具体的报错信息打印
    // console.log('error')
    console.log(error)
} finally {
    console.log("finally logic")
}


// 手动抛出业务异常
try {
    // 1. new exception
    // Error 是一种类型
    var e = new Error('自定义异常')
    // 2. 再抛出 exception
    throw e
} catch (error) {
    console.log(error)
}


// 函数
function abs(x) {
    if (x >= 0) {
        return x
    } else {
        return -x
    }
}
console.log(abs(-10))
// console.log(abs(true))


// 方法：绑定在对象上的函数
var person = {name: "张三", age: 10}
person.greet = function() {
    // this 指代当前绑定的对象
    console.log(`你好！我是${this.name}`)
}

// 这里的greet就是方法
person.greet()


// 箭头函数
// person.greet = () => {console.log(`你好！我是${this.name}`)}  这里不能这么些，this是空

// function sqt(x) {
//     return x * x;
// }
// console.log(sqt(10))

// x => x * x
// 定义函数的参数和返回值，并没有定义函数名称
var sqt = x => x * x
// 完整版 var sqt = (x) => { return x * x }

console.log(sqt(10))










// js 导入与导出

import {firstName, lastName, year} from './profile.mjs';

console.log(firstName, lastName, year)


// map解构赋值
// { firstName, lastName, year } = {firstName: firstName, lastName: lastName, year: year}

// var firstName = 'xxx'
// 如果出现变量冲突，可以通过别名的方式，对导入的变量进行命名
// import {firstName as profileFirstName, lastName, year} from './profile.mjs'
// console.log(firstName, profoleFirstName, year)

// 能不能通过包名来应用包里面的
// pkg.firstName
// pkg.lastName
// pkg.year
// 引入另外一个概念，默认导出
// 所有的变量绑定在一个变量上导出



// import { MYAPP } from "./profile.mjs"
// console.log(firstName, MYAPP.firstName, MYAPP.lastName, MYAPP.year)
// 这是早期解决冲突的方式，很接近于package的概念了。但是现在还是需要有MYAPP的变量。
// 还是需要知道有MYAPP，pkg.xxx

// 引出新的概念，默认导出
// { firstName, lastName, year } = {firstName: firstName, lastName: lastName, year: year}
// default as pkg  导出为一个default空间，重命名为pkg。

// 这里把default改名为profile
// import profile from './profile.mjs'
// console.log(profile.firstName, profile.lastName, profile.year)
