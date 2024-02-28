// for of 遍历array

arr1 = [1, 2, 3]
for (var x of arr1) {
    console.log(x);
}

// for of 遍历map
var obj = {
    name: 'Jack',
    age: 20,
    city: 'Beijing'
};
// Object.keys, Aray/Object/Error
// 工具方法
for (var key of Object.keys(obj)) {
    console.log(obj[key]);
}


// for each
// 使用 array对象提供一个迭代器，传递一个函数作为参数
// fn = element => {}
// 它不能中断循环
arr1.forEach(element => {
    console.log(element)
});

