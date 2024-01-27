# 开发文档

## 功能
- 项目及构建
- [x] 项目 CRUD 
- [ ] 项目解析
- [ ] 构建计划 CRUD
- [ ] 构建计划配置
- [ ] 执行构建任务
- [ ] 构建历史
- 环境
- [ ] 工具初始化
- [x] 工具 CRUD

## 乱七八糟的想法
- [ ] 项目仓库类型区分 GIT、DIR（本地目录）、SVN 三种，GIT、SVN 支持用户名密码导入和 SSH 私钥导入，这些信息以 json 格式存储在 config 属性中
- [ ] 服务器管理，以 SSH 方式配置 IP、PORT，可以通过 SSH 协议访问服务器
- [ ] SSH Key 的管理，支持生成 SSH 私钥文件（可设置密码），可以用来访问服务器
- [ ] 有个叫做 Guacamole 的项目似乎可以远程控制服务器，尝试融合进来
- [ ] 

## 功能拆解流程
- 项目解析
- [x] 区分项目仓库类型：GIT、DIR、SVN
    - GIT 项目解析
    - [x] 1. 判断项目目录下是否存在 .git 目录，若存在表示已经导入过仓库，执行第 3 步，若不存在执行第 2 步
    - [x] 2. 从 config 属性中解析仓库信息，支持用户名密码导入和 SSH 私钥导入，执行 git clone xxx
    - [ ] 3. 遍历整个目录树，解析指定文件（比如 package.json、pom.xml），汇总到 parsedInfo 属性，存储到数据库
    - DIR 项目解析
    - [x] 判断目录是否存在，若存在继续执行
    - [ ] 遍历整个目录树，解析指定文件（比如 package.json、pom.xml），汇总到 parsedInfo 属性，存储到数据库
    - SVN 项目解析
    - [ ] 1. 判断项目目录下是否存在 .svn 目录，若存在表示已经导入过仓库，执行第 3 步，若不存在执行第 2 步
    - [ ] 2. 从 config 属性中解析仓库信息，支持用户名密码导入和 SSH 私钥导入，把 SVN 项目克隆到本地
    - [ ] 3. 遍历整个目录树，解析指定文件（比如 package.json、pom.xml），汇总到 parsedInfo 属性，存储到数据库
- GIT 仓库导入
- [x] 支持三种方式认证：用户名密码、SSH 私钥、AccessToken
- [ ] 访问 GIT 仓库两种协议：HTTP 和 SSH，HTTP 支持 Credentials 和 AccessToken，SSH 支持 SSH 私钥访问