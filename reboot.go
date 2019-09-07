package main
import (
    "syscall"
)

/*
syscall: 底层的外部包,提供了操作系统底层调用的基本接口
下面程序可以让Linux系统重启
*/

const LINUX_REBOOT_MAGIC1 uintptr = 0xfee1dead
const LINUX_REBOOT_MAGIC2 uintptr = 672274793
const LINUX_REBOOT_CMD_RESTART uintptr = 0x1234567

func main() {
    syscall.Syscall(syscall.SYS_REBOOT,
        LINUX_REBOOT_MAGIC1,
        LINUX_REBOOT_MAGIC2,
        LINUX_REBOOT_CMD_RESTART)
}