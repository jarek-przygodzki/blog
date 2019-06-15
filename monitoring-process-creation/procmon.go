package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

/*
#include <linux/netlink.h>
#include <linux/connector.h>
#include <linux/cn_proc.h>
*/
import "C"

func main() {
	sock, err := unix.Socket(unix.AF_NETLINK, unix.SOCK_DGRAM, unix.NETLINK_CONNECTOR)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	addr := &unix.SockaddrNetlink{Family: unix.AF_NETLINK, Groups: C.CN_IDX_PROC, Pid: uint32(os.Getpid())}
	err = unix.Bind(sock, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bind: %v\n", err)
		return
	}
	defer func() {
		send(sock, C.PROC_CN_MCAST_IGNORE)
		unix.Close(sock)
	}()

	err = send(sock, C.PROC_CN_MCAST_LISTEN)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	for {
		p := make([]byte, 1024)
		nbytes, from, err := unix.Recvfrom(sock, p, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
		pid := from.(*unix.SockaddrNetlink).Pid
		if pid != 0 {
			fmt.Fprintf(os.Stderr, "sender was not kernel (PID %d)", pid)
			return
		}
		nlmessages, err := syscall.ParseNetlinkMessage(p[:nbytes])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}

		for _, m := range nlmessages {
			if m.Header.Type == unix.NLMSG_DONE {
				cnhdr := (*C.struct_cn_msg)(unsafe.Pointer(&m.Data[0]))
				ptr := uintptr(unsafe.Pointer(cnhdr))
				ptr += unsafe.Sizeof(*cnhdr)
				pe := (*C.struct_proc_event)(unsafe.Pointer(ptr))
				switch pe.what {
				case C.PROC_EVENT_EXEC:
					e := (*C.struct_exec_proc_event)(unsafe.Pointer(&pe.event_data))
					fmt.Printf("Process started: PID %d\n", e.process_pid)
				case C.PROC_EVENT_EXIT:
					e := (*C.struct_exit_proc_event)(unsafe.Pointer(&pe.event_data))
					fmt.Printf("Process exited: PID %d\n", e.process_pid)
				}
			}
		}
	}
}

func send(sockfd int, op C.enum_proc_cn_mcast_op) error {
	cnhdr := &C.struct_cn_msg{
		id:  C.struct_cb_id{idx: C.CN_IDX_PROC, val: C.CN_VAL_PROC},
		len: C.__u16(C.sizeof_enum_proc_cn_mcast_op),
	}
	header := unix.NlMsghdr{
		Len:   uint32(C.sizeof_struct_nlmsghdr + C.sizeof_struct_cn_msg + C.sizeof_enum_proc_cn_mcast_op),
		Type:  uint16(unix.NLMSG_DONE),
		Flags: 0,
		Seq:   0,
		Pid:   uint32(os.Getpid()),
	}

	buf := bytes.NewBuffer(make([]byte, 0, header.Len))

	binary.Write(buf, binary.LittleEndian, header)
	binary.Write(buf, binary.LittleEndian, cnhdr)
	binary.Write(buf, binary.LittleEndian, op)

	destAddr := &unix.SockaddrNetlink{Family: unix.AF_NETLINK, Groups: C.CN_IDX_PROC, Pid: 0} // 0 is the kernel
	return unix.Sendto(sockfd, buf.Bytes(), 0, destAddr)

}