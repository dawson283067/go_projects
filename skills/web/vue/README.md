# 初始化一个vue项目

```js
npm create vue@latest

D:\Development\go_projects\skills\web\vue>npm create vue@latest
Need to install the following packages:
create-vue@3.9.2
Ok to proceed? (y) y

Vue.js - The Progressive JavaScript Framework

√ Project name: ... vue-project
√ Add TypeScript? ... No
√ Add JSX Support? ... Yes
√ Add Vue Router for Single Page Application development? ... Yes
√ Add Pinia for state management? ... Yes
√ Add Vitest for Unit Testing? ... No
√ Add an End-to-End Testing Solution? » No
√ Add ESLint for code quality? ... Yes
√ Add Prettier for code formatting? ... Yes

Scaffolding project in D:\Development\go_projects\skills\web\vue\vue-project...

Done. Now run:

  cd vue-project
  npm install
  npm run format
  npm run dev
```

yrm ls 列出源
yrm use 使用某个源

```js
D:\Development\go_projects\skills\web\vue>dir
 驱动器 D 中的卷没有标签。
 卷的序列号是 E6B1-BF25

 D:\Development\go_projects\skills\web\vue 的目录

2024-02-28  04:46 PM    <DIR>          .
2024-02-28  04:46 PM    <DIR>          ..
2024-02-28  04:47 PM               843 README.md
2024-02-28  04:46 PM    <DIR>          vue-project
               1 个文件            843 字节
               3 个目录 1,880,253,165,568 可用字节
```

```js
PS D:\Development\go_projects\skills\web\vue\vue-project> npm install

added 221 packages, and audited 222 packages in 12s

42 packages are looking for funding
  run `npm fund` for details

found 0 vulnerabilities
```

```sh
PS D:\Development\go_projects\skills\web\vue\vue-project> npm run format

> vue-project@0.0.0 format
> prettier --write src/   

src/App.vue 106ms (unchanged)
src/assets/base.css 10ms (unchanged)
src/assets/main.css 5ms (unchanged)
src/components/HelloWorld.vue 19ms (unchanged)
src/components/icons/IconCommunity.vue 4ms (unchanged)
src/components/icons/IconDocumentation.vue 2ms (unchanged)
src/components/icons/IconEcosystem.vue 3ms (unchanged)
src/components/icons/IconSupport.vue 2ms (unchanged)
src/components/icons/IconTooling.vue 3ms (unchanged)
src/components/TheWelcome.vue 13ms (unchanged)
src/components/WelcomeItem.vue 12ms (unchanged)
src/main.js 6ms (unchanged)
src/router/index.js 9ms (unchanged)
src/stores/counter.js 7ms (unchanged)
src/views/AboutView.vue 2ms (unchanged)
src/views/HomeView.vue 3ms (unchanged)
```

```sh
PS D:\Development\go_projects\skills\web\vue\vue-project> npm run dev

> vue-project@0.0.0 dev
> vite


  VITE v5.1.4  ready in 506 ms

  ➜  Local:   http://localhost:5173/
  ➜  Network: use --host to expose
  ➜  press h + enter to show help
h

  Shortcuts
  press r + enter to restart the server
  press u + enter to show server url   
  press o + enter to open in browser   
  press c + enter to clear console     
  press q + enter to quit
```





