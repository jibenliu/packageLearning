linux内核Inotify接口简介

inotify中主要涉及3个接口。分别是inotify_init, inotify_add_watch,read。具体如下：

| 接口名                                | 作用                     |
|------------------------------------|------------------------|
| int fd = inotify_init()            | 创建inotify实例，返回对应的文件描述符 |
| inotify_add_watch (fd, path, mask) | 注册被监视目录或文件的事件          |
| read (fd,buf, BUF_LEN)             | 读取监听到的文件事件             |

Inotify可以监听的文件系统事件列表：

| 事件名称             | 事件说明                                              |
|------------------|---------------------------------------------------|
| IN_ACCESS        | 文件被访问                                             |
| IN_MODIFY        | 文件被 write                                         |
| IN_CLOSE_WRITE   | 可写文件被 close                                       |
| IN_OPEN          | 文件被 open                                          |
| IN_MOVED_TO      | 文件被移来，如 mv、cp                                     |
| IN_CREATE        | 创建新文件                                             |
| IN_DELETE        | 文件被删除，如 rm                                        |
| IN_DELETE_SELF   | 自删除，即一个可执行文件在执行时删除自己                              |
| IN_MOVE_SELF     | 自移动，即一个可执行文件在执行时移动自己                              |
| IN_ATTRIB        | 文件属性被修改，如 chmod、chown、touch 等                     |
| IN_CLOSE_NOWRITE | 不可写文件被 close                                      |
| IN_MOVED_FROM    | 文件被移走,如 mv                                        |
| IN_UNMOUNT       | 宿主文件系统被 umount                                    |
| IN_CLOSE         | 文件被关闭，等同于(IN_CLOSE_WRITE &#124; IN_CLOSE_NOWRITE) |
| IN_MOVE          | 文件被移动，等同于(IN_MOVED_FROM &#124; IN_MOVED_TO)       |
