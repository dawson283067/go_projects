// profile.mjs
export var firstName = 'Michael';
export var lastName = 'Jackson';
export var year = 1958;

// 1. 通过把变量组装成一个map 导出
// 等于：{firstName;firstName, lastName:lastName, year:year}
// export {firstName, lastName, year}

// 2. 唯一的全局变量MYAPP
// export var MYAPP = {
//     firstName,
//     lastName,
//     year
// }


// 3. 默认导出
// default = {firstName: firstName, lastName: lastName, year: year}
// export default {
//     firstName, 
//     lastName, 
//     year
// }


// 以上就是具名导出和匿名导出。

// 这里是ES6的 import/export 导入导出

// 