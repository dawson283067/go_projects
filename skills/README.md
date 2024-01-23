# 项目技能

+ 单元测试： unit-test


## vscode运行go test显示打印日志

```
File --> Preferences --> Settings --> (User)Workbench --> Settings Editor --> Edit in settings.json

最后添加一行："go.testFlags": ["-count=1","-v"]

vs 配置 go test flag配置：go.testFlags
+ -count=1: (cached)，单元测试没有修改也想强制运行一次
+ -v: 单元测试内部想打印 大小信息 print, t.Log()，需要加上该参数

{
    "hediet.vscode-drawio.resizeImages": null,
    "diffEditor.ignoreTrimWhitespace": false,
    "explorer.confirmDelete": false,
    "workbench.settings.applyToAllProfiles": [     

    ],
    "go.testFlags": ["-count=1","-v"]
}
````