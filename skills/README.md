# 项目技能

+ 单元测试： unit-test


## vscode运行go test显示打印日志

```
File --> Preferences --> Settings --> (User)Workbench --> Settings Editor --> Edit in settings.json

最后添加一行："go.testFlags": ["-count=1","-v"]

{
    "hediet.vscode-drawio.resizeImages": null,
    "diffEditor.ignoreTrimWhitespace": false,
    "explorer.confirmDelete": false,
    "workbench.settings.applyToAllProfiles": [     

    ],
    "go.testFlags": ["-count=1","-v"]
}
````