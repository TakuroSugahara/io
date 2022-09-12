
Use io in golang to know file descriptor.

# file descriptorについて

### os.Fileの型と、fdとは何か

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

Linuxカーネルプロセル内部ではopenファイルに識別子をつけている。（非負整数）
これをfile descriptorと呼び、goではpoll.FD型がこれを表している

### FDの型とinode番号

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

### os.Openの処理とfdの取得方法

`os.Open` の内部では `syscall.Open` を呼び出している。
つまり、`syscall.Open` を使うことで、fdを取得している。

### file.Readと低レイヤーの繋がり

`file.Read` では いくつかの非公開関数を呼び出しながら最終的に `syscall.Read` をfdを渡して呼び出している
処理の順番は以下

1. os.File型のReadメソッド
1. os.Fileのreadメソッド
1. poll.FD型のReadメソッド
1. syscall.Readメソッド
1. osカーネルのシステムコールの読み込み処理

Wirteも大体同様の処理になっている
