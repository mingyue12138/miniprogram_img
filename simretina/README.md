# simretina
## 当前进度和完成情况
### 进度
1.视频片段及结构化信息打包上报、结构化算法结果信息上报、图片及结构化信息打包上报未实现
2.其余接口均有实现

### 完成情况
1.大部分接口返回值和postman测试中相同,大量接口msg返回“该设备尚未建立连接”
2.视频上传和图片上传两个接口返回Required request body is missing
3.token值均未添加


## vscode调试
直接在cmd/main.go按F5后，程序使用的启动路径是cmd，导致不能正确读取相对路径。  
配置launch.json可以一劳永逸的解决，还可以在launch.json里面设置env进行自测。

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "golang",
            "type": "go",    
            "request": "launch", 
            "mode": "auto", 
            "env": { 
                "common.db_conn": "titan/xianwei@172.16.5.22:3306@titandb"
            },
            "program": "${workspaceFolder}/cmd" ,
            "cwd": "${workspaceFolder}"
        }
    ]
}
```

