var a = (fn) => {
    console.log('网络请求a处理中')

    // 失败后就直接返回

    // 成功就执行fn
    fn()
}

var b = (fn) => {
    console.log('网络请求b处理中')

    // 失败后就直接返回

    // 成功就执行fn
    fn()
}

var c = (fn) => {
    console.log('网络请求c处理中')

    // 失败后就直接返回

    // 成功就执行fn
    fn()
}

// 同时调用， a,b,c同时调用
// curl url1,url2,url3 相当于同时打开3个网站
// a()
// b()
// c()


// js里面使用同步编程： a --> b --> c
a(() => {
    // a 执行完成后 调用函数fn
    console.log('网络请求a处理完成')
    b(() => {
        // b 执行完成后 调用函数fn
        console.log('网络请求b处理完成')
        c(() => {
            // c 执行完成后 调用函数fn
            console.log('网络请求c处理完成')
        })
    })
})


// 没有Promise情况下，调用vblog的 blog list api
// 成功：[item ...]
// 失败：提示用户
function BlogList(success, failed) {
    // io 网络请求
    // 通过http status code

    // 成功 调用success处理返回的数据
    var items = []
    success(items)

    // 失败
    var err = new Error('请求失败')
    failed(err)  
}

// 为了规范异步编程，专门发明一个概念叫Promise对象，把你这个函数构造成一个Promise对象
var po = new Promise(BlogList)
po.then((data) => {
    console.log(data)
}).catch((err) => {
    console.log(err)
}).finally(() => {
    console.log("final")
})

// 以上是从异步编程到Promise的第一步进化

// 写的是同步的，执行逻辑是异步的。有没有这种更好的方式呢？
// async/await，执行逻辑和go里面是一样的
// await 等待promise对象的返回
async function f1() {
    try {
        const p = new Promise(BlogList)
        // await(p)
        var resp = await p
        console.log("x", resp)
    } catch (error) {
        console.log(error)
    }
}

f1()

// 调用API的具体例子：axios 这个库做说明













