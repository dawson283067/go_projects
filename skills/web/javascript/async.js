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