
Use io in golang to know file descriptor.

```go
// goのos.Fileの構造体
type file struct {
	pfd         poll.FD
	name        string
	dirinfo     *dirInfo // nil unless directory being read
	nonblock    bool     // whether we set nonblocking mode
	stdoutOrErr bool     // whether this is stdout or stderr
	appendMode  bool     // whether file is opened for appending
}

```

poll.FDはfile descriptorを表している

Linuxカーネルプロセル内部ではopen舌ファイルに識別子をつけている。（非負整数）
これをfile descriptorと呼び、goではpoll.FD型がこれを表している

```go
// poll.FDの型
type FD struct {
	// Lock sysfd and serialize access to Read and Write methods.
	fdmu fdMutex

	// System file descriptor. Immutable until Close.
	Sysfd int

	// I/O poller.
	pd pollDesc

	// Writev cache.
	iovecs *[]syscall.Iovec

	// Semaphore signaled when file is closed.
	csema uint32

	// Non-zero if this file has been set to blocking mode.
	isBlocking uint32

	// Whether this is a streaming descriptor, as opposed to a
	// packet-based descriptor like a UDP socket. Immutable.
	IsStream bool

	// Whether a zero byte read indicates EOF. This is false for a
	// message based socket connection.
	ZeroReadIsEOF bool

	// Whether this is a file rather than a network socket.
	isFile bool
}

```

Sysfdがまさにfile descriptorの識別子を表している。
コメントではCloseされるまで変更されないとなっている

ちなみにOpenしてないものにも一応識別子は存在しており、それはinode番号と呼ばれている
fdとはその番号とは関係ないので、同じファイルだろうがOpenした時々で異なる番号がつけられる可能性がある

