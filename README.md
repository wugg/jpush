# jpush-go

极光推送 Go 语言客户端 SDK，支持 Push、Schedule、Device、Report 全套 API。

## 功能特性

- 推送（Push）：支持全平台、标签、别名、RegistrationID 等多种推送方式
- 定时任务（Schedule）：支持单次和周期定时推送
- 设备管理（Device）：设备标签、别名管理
- 统计报告（Report）：送达统计、消息统计、用户统计
- 批量推送：支持按 RegistrationID 和别名批量推送
- 推送撤销：支持撤销正在发送或排队的消息

## 安装

```bash
go get github.com/yourusername/jpush-go
```

## 快速开始

```go
package main

import (
    "fmt"
    jpush "github.com/yourusername/jpush-go"
)

func main() {
    // 创建客户端
    client := jpush.NewClient("your_app_key", "your_master_secret")

    // 构建推送
    result, err := client.Push().
        SetPlatformAll().
        SetAudienceAll().
        SetNotificationAlert("Hello JPush!").
        Send()

    if err != nil {
        fmt.Printf("推送失败: %v\n", err)
        return
    }

    fmt.Printf("推送成功，消息ID: %s\n", result.MsgID)
}
```

## 推送

### 基础推送

```go
client.Push().
    SetPlatformAll().
    SetAudienceAll().
    SetNotificationAlert("通知内容").
    Send()
```

### 分平台推送

```go
android := &jpush.AndroidNotification{
    Alert:     "Android通知",
    Title:     "标题",
    ChannelID: "default",
    Extras:    map[string]interface{}{"key": "value"},
}

ios := &jpush.IOSNotification{
    Alert:  "iOS通知",
    Sound:  "default",
    Badge:  1,
    Extras: map[string]interface{}{"key": "value"},
}

client.Push().
    SetPlatforms("ios", "android").
    AddTag("vip", "active").
    SetAndroidNotification(android).
    SetIosNotification(ios).
    SetApnsProduction(true).
    Send()
```

### 按别名推送

```go
client.Push().
    SetPlatformAll().
    SetAudienceWithMap(map[string][]string{
        "alias": {"user_alias_1", "user_alias_2"},
    }).
    SetNotificationAlert("发送给指定用户").
    Send()
```

### 按 RegistrationID 推送

```go
client.Push().
    SetPlatformAll().
    AddRegistrationID("registration_id_1", "registration_id_2").
    SetNotificationAlert("指定设备推送").
    Send()
```

### 自定义消息

```go
message := &jpush.Message{
    MsgContent:  "自定义消息内容",
    Title:       "消息标题",
    ContentType: "text",
    Extras:      map[string]interface{}{"from": "jpush"},
}

client.Push().
    SetPlatformAll().
    SetAudienceAll().
    SetMessage(message).
    Send()
```

### 推送选项

```go
options := &jpush.Options{
    TimeToLive:      3600,      // 离线消息保留时间（秒）
    ApnsProduction:  true,       // iOS 生产环境
    BigPushDuration: 0,         // 定速推送时长
}

client.Push().
    SetPlatformAll().
    SetAudienceAll().
    SetNotificationAlert("通知").
    SetOptions(options).
    Send()
```

## 定时任务

### 单次定时推送

```go
// 创建推送载荷
push := client.Push().
    SetPlatformAll().
    SetAudienceAll().
    SetNotificationAlert("定时推送")

pushPayload := map[string]interface{}{}
data, _ := json.Marshal(push)
json.Unmarshal(data, &pushPayload)

// 创建定时任务
schedule := client.Schedule().
    SetName("每日提醒").
    SetEnabled(true).
    SetPush(pushPayload).
    SetSingleTrigger("2024-12-01 12:00:00")

result, err := schedule.Create()
```

### 周期定时推送

```go
// 每天定时推送
trigger := jpush.NewDailyTrigger(
    "2024-01-01 00:00:00", // 开始时间
    "10:00:00",            // 执行时间
    1,                     // 每1天
)

// 每周定时推送
trigger := jpush.NewWeeklyTrigger(
    "2024-01-01 00:00:00",
    "10:00:00",
    []string{"MON", "WED", "FRI"}, // 周一、周三、周五
    1,
)

// 每月定时推送
trigger := jpush.NewMonthlyTrigger(
    "2024-01-01 00:00:00",
    "10:00:00",
    []string{"01", "15"}, // 每月1日和15日
    1,
)

schedule := client.Schedule().
    SetName("定时推送").
    SetEnabled(true).
    SetPush(pushPayload).
    SetPeriodicalTrigger(trigger)

result, err := schedule.Create()
```

### 定时任务管理

```go
// 获取详情
detail, _ := client.Schedule().GetByID(scheduleID)

// 获取列表
list, _ := client.Schedule().GetList(1)

// 删除任务
client.Schedule().Delete(scheduleID)
```

## 设备管理

### 设备信息

```go
device := client.Device()

// 获取设备信息
info, _ := device.GetDeviceInfo("registration_id_xxx")

// 更新别名
device.UpdateDeviceAlias("registration_id_xxx", "new_alias")

// 添加设备标签
device.AddDeviceTag("registration_id_xxx", "vip", "active")
```

### 标签管理

```go
// 获取所有标签
tags, _ := device.GetTagList()

// 添加标签用户
device.AddTagUsers("vip", "reg_id_1", "reg_id_2")

// 检查用户是否在标签中
check, _ := device.CheckTagUserExist("vip", "reg_id_1")

// 删除标签
device.DeleteTag("vip", "android")
```

### 别名管理

```go
// 获取别名用户
users, _ := device.GetAliasUsers("user_alias", "android")

// 删除别名
device.DeleteAlias("user_alias", "ios")
```

## 统计报告

### 送达统计

```go
report := client.Report()

// 获取送达统计
received, _ := report.GetReceived("msg_id_xxx")

// 获取送达详情
detail, _ := report.GetReceivedDetail("msg_id_xxx")
```

### 消息统计

```go
// 获取消息统计
messages, _ := report.GetMessages("msg_id_1,msg_id_2")

// 获取消息详情
detail, _ := report.GetMessagesDetail("msg_id_xxx")
```

### 用户统计

```go
// 每日用户统计
users, _ := report.GetUsersDaily("2024-01-01", 7)

// 每小时用户统计
hourly, _ := report.GetUsersHourly("2024-01-01", 24)
```

## 批量推送

### 按 RegistrationID 批量推送

```go
pushList := []map[string]interface{}{
    {
        "platform": "all",
        "audience": map[string][]string{
            "registration_id": {"reg_id_1"},
        },
        "notification": map[string]interface{}{
            "alert": "推送内容1",
        },
    },
    // ...更多推送
}

result, _ := client.BatchPushByRegID(pushList)
```

## 推送撤销

```go
// 撤销正在发送或排队的消息
result, err := client.RevokePush(msgID)
```

## 区域配置

默认使用极光国内服务器，可配置使用北京机房：

```go
client := jpush.NewClient(
    "your_app_key",
    "your_master_secret",
    jpush.WithZone(jpush.ZoneBJ),
)
```

## 错误处理

SDK 定义了多种错误类型：

```go
// 认证失败
_, err := client.Push().Send()
if _, ok := err.(*jpush.Unauthorized); ok {
    // 检查 AppKey 和 MasterSecret
}

// API 错误
if failure, ok := err.(*jpush.JPushFailure); ok {
    fmt.Printf("错误码: %d, 错误信息: %s\n", failure.ErrorCode, failure.ErrorMessage)
}

// 连接错误
if _, ok := err.(*jpush.APIConnectionException); ok {
    // 检查网络连接
}
```

## License

MIT License
