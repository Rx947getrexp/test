你可以使用以下命令在不停止服务的情况下清理日志文件：

```bash
> /path/to/error.log
```

这个命令将会清空文件内容，但不会删除文件。这样，正在写入这个文件的进程可以继续写入。

然后，你可以设置日志轮转来自动管理日志文件的大小。这可以使用logrotate工具完成，它可以自动轮转、压缩、删除和邮件日志文件。每个日志文件可以通过创建一个配置文件来进行设置。

例如，一个简单的logrotate配置文件可能如下所示：

```bash
/path/to/error.log {
    daily
    rotate 7
    compress
    missingok
    notifempty
    create 0640 root adm
}
```

这个配置文件的意思是：

- daily：每天轮转一次日志文件。
- rotate 7：保存7个旧的日志文件。
- compress：轮转后的文件进行压缩。
- missingok：如果日志文件不存在，不报错继续下一个。
- notifempty：如果日志文件为空，不进行轮转。
- create 0640 root adm：轮转后新建日志文件的权限和所有者。

你可以根据自己的需求调整这个配置。然后，你需要将这个配置文件放到/etc/logrotate.d/目录下，logrotate会自动读取这个目录下的配置文件。

最后，你可以通过crontab设置定时任务，定时运行logrotate命令：

```bash
0 0 * * * /usr/sbin/logrotate /etc/logrotate.d/errorlog
```

这个命令的意思是每天0点0分运行logrotate命令，轮转/etc/logrotate.d/errorlog配置文件中指定的日志文件。