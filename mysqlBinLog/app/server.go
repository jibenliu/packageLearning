package app

import (
	"bufio"
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/go-mysql-org/go-mysql/replication"
	"io"
	"net"
	"os"
	"time"
)

const (
	MinProtocolVersion byte = 10

	OkHeader  byte = 0x00
	ErrHeader byte = 0xff
)

const MaxPayloadLength = 1<<24 - 1

type Server struct {
	Cfg             *Config
	Ctx             context.Context
	conn            net.Conn
	io              *PacketIo
	registerSuccess bool
}

func (s *Server) Run() {
	defer func() {
		s.Quit()
	}()

	s.dump()
}

func (s *Server) dump() {
	err := s.handshake()
	if err != nil {
		panic(err)
	}
	s.invalidChecksum()
	fmt.Println("dump ...")
	s.register()
	s.writeDumpCommand()
	parser := replication.NewBinlogParser()
	for {
		//time.Sleep(2 * time.Second)
		//s.query("select 1")

		data, err := s.io.readPacket()
		if err != nil || len(data) == 0 {
			continue
		}

		//s.Quit()

		if data[0] == OkHeader {
			//skip ok
			data = data[1:]
			if e, err := parser.Parse(data); err == nil {
				e.Dump(os.Stdout)
			} else {
				fmt.Println(err)
			}
		} else {
			s.io.HandleError(data)
		}
	}
}

func (s *Server) invalidChecksum() {
	sql := `SET @master_binlog_checksum='NONE'`
	if err := s.query(sql); err != nil {
		fmt.Println(err)
	}
	//must-read from tcp connection , either will be blocked
	_, _ = s.io.readPacket()
}

func (s *Server) handshake() error {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", s.Cfg.Host, s.Cfg.Port), 10*time.Second)
	if err != nil {
		return err
	}

	tc := conn.(*net.TCPConn)
	_ = tc.SetKeepAlive(true)
	_ = tc.SetNoDelay(true)
	s.conn = tc

	s.io = &PacketIo{}
	s.io.r = bufio.NewReaderSize(s.conn, 16*1024)
	s.io.w = tc

	data, err := s.io.readPacket()
	if err != nil {
		return err
	}

	if data[0] == ErrHeader {
		return errors.New("error packet")
	}

	if data[0] < MinProtocolVersion {
		return fmt.Errorf("version is too lower, current:%d", data[0])
	}

	pos := 1 + bytes.IndexByte(data[1:], 0x00) + 1
	connId := uint32(binary.LittleEndian.Uint32(data[pos : pos+4]))
	pos += 4
	salt := data[pos : pos+8]

	pos += 8 + 1
	capability := uint32(binary.LittleEndian.Uint16(data[pos : pos+2]))

	pos += 2

	var status uint16
	var pluginName string
	if len(data) > pos {
		//skip charset
		pos++
		status = binary.LittleEndian.Uint16(data[pos : pos+2])
		pos += 2
		capability = uint32(binary.LittleEndian.Uint16(data[pos:pos+2]))<<16 | capability
		pos += 2

		pos += 10 + 1
		salt = append(salt, data[pos:pos+12]...)
		pos += 13

		if end := bytes.IndexByte(data[pos:], 0x00); end != -1 {
			pluginName = string(data[pos : pos+end])
		} else {
			pluginName = string(data[pos:])
		}
	}

	fmt.Printf("conn_id:%v, status:%d, plugin:%v\n", connId, status, pluginName)

	//write
	capability = 500357
	length := 4 + 4 + 1 + 23
	length += len(s.Cfg.User) + 1

	pass := []byte(s.Cfg.Pass)
	auth := calPassword(salt[:20], pass)
	length += 1 + len(auth)
	data = make([]byte, length+4)

	data[4] = byte(capability)
	data[5] = byte(capability >> 8)
	data[6] = byte(capability >> 16)
	data[7] = byte(capability >> 24)

	//utf8
	data[12] = byte(33)
	pos = 13 + 23
	if len(s.Cfg.User) > 0 {
		pos += copy(data[pos:], s.Cfg.User)
	}

	pos++
	data[pos] = byte(len(auth))
	pos += 1 + copy(data[pos+1:], auth)

	err = s.io.writePacket(data)
	if err != nil {
		return fmt.Errorf("write auth packet error")
	}

	pk, err := s.io.readPacket()
	if err != nil {
		return err
	}

	if pk[0] == OkHeader {
		fmt.Println("handshake ok ")
		return nil
	} else if pk[0] == ErrHeader {
		s.io.HandleError(pk)
		return errors.New("handshake error ")
	}

	return nil
}

func (s *Server) writeDumpCommand() {
	s.io.seq = 0
	data := make([]byte, 4+1+4+2+4+len(s.Cfg.LogFile))
	pos := 4
	data[pos] = 18 //dump binlog
	pos++
	binary.LittleEndian.PutUint32(data[pos:], uint32(s.Cfg.Position))
	pos += 4

	//dump command flag
	binary.LittleEndian.PutUint16(data[pos:], 0)
	pos += 2

	binary.LittleEndian.PutUint32(data[pos:], uint32(s.Cfg.ServerId))
	pos += 4

	copy(data[pos:], s.Cfg.LogFile)

	_ = s.io.writePacket(data)
	//ok
	res, _ := s.io.readPacket()
	if res[0] == OkHeader {
		fmt.Println("send dump command return ok.")
	} else {
		s.io.HandleError(res)
	}
}

func (s *Server) register() {
	s.io.seq = 0
	hostname, _ := os.Hostname()
	data := make([]byte, 4+1+4+1+len(hostname)+1+len(s.Cfg.User)+1+len(s.Cfg.Pass)+2+4+4)
	pos := 4
	data[pos] = 21 //register slave  command
	pos++
	binary.LittleEndian.PutUint32(data[pos:], uint32(s.Cfg.ServerId))
	pos += 4

	data[pos] = uint8(len(hostname))
	pos++
	n := copy(data[pos:], hostname)
	pos += n

	data[pos] = uint8(len(s.Cfg.User))
	pos++
	n = copy(data[pos:], s.Cfg.User)
	pos += n

	data[pos] = uint8(len(s.Cfg.Pass))
	pos++
	n = copy(data[pos:], s.Cfg.Pass)
	pos += n

	binary.LittleEndian.PutUint16(data[pos:], uint16(s.Cfg.Port))
	pos += 2

	binary.LittleEndian.PutUint32(data[pos:], 0)
	pos += 4

	//master id = 0
	binary.LittleEndian.PutUint32(data[pos:], 0)

	_ = s.io.writePacket(data)

	//ok
	res, _ := s.io.readPacket()
	if res[0] == OkHeader {
		fmt.Println("register success.")
		s.registerSuccess = true
	} else {
		s.io.HandleError(data)
	}
}

func (s *Server) writeCommand(command byte) {
	s.io.seq = 0
	_ = s.io.writePacket([]byte{
		0x01, //1 byte long
		0x00,
		0x00,
		0x00, //seq
		command,
	})
}

func (s *Server) query(q string) error {
	s.io.seq = 0
	length := len(q) + 1
	data := make([]byte, length+4)
	data[4] = 3
	copy(data[5:], q)
	return s.io.writePacket(data)
}

func (s *Server) Quit() {
	//quit
	s.writeCommand(byte(1))
	//maybe only close
	if err := s.conn.Close(); nil != err {
		fmt.Printf("error in close :%v\n", err)
	}
}

type PacketIo struct {
	r   *bufio.Reader
	w   io.Writer
	seq uint8
}

func (p *PacketIo) readPacket() ([]byte, error) {
	//to read header
	header := []byte{0, 0, 0, 0}
	if _, err := io.ReadFull(p.r, header); err != nil {
		return nil, err
	}

	length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)
	if length == 0 {
		p.seq++
		return []byte{}, nil
	}

	if length == 1 {
		return nil, fmt.Errorf("invalid payload")
	}

	seq := uint8(header[3])
	if p.seq != seq {
		return nil, fmt.Errorf("invalid seq %d", seq)
	}

	p.seq++
	data := make([]byte, length)
	if _, err := io.ReadFull(p.r, data); err != nil {
		return nil, err
	} else {
		if length < MaxPayloadLength {
			return data, nil
		}
		var buf []byte
		buf, err = p.readPacket()
		if err != nil {
			return nil, err
		}
		if len(buf) == 0 {
			return data, nil
		} else {
			return append(data, buf...), nil
		}
	}
}

func (p *PacketIo) writePacket(data []byte) error {
	length := len(data) - 4
	if length >= MaxPayloadLength {
		data[0] = 0xff
		data[1] = 0xff
		data[2] = 0xff
		data[3] = p.seq

		if n, err := p.w.Write(data[:4+MaxPayloadLength]); err != nil {
			return fmt.Errorf("write find error")
		} else if n != 4+MaxPayloadLength {
			return fmt.Errorf("not equal max pay load length")
		} else {
			p.seq++
			length -= MaxPayloadLength
			data = data[MaxPayloadLength:]
		}
	}

	data[0] = byte(length)
	data[1] = byte(length >> 8)
	data[2] = byte(length >> 16)
	data[3] = p.seq

	if n, err := p.w.Write(data); err != nil {
		return errors.New("write find error")
	} else if n != len(data) {
		return errors.New("not equal length")
	} else {
		p.seq++
		return nil
	}
}

func calPassword(scramble, password []byte) []byte {
	crypt := sha1.New()
	crypt.Write(password)
	stage1 := crypt.Sum(nil)

	crypt.Reset()
	crypt.Write(stage1)
	hash := crypt.Sum(nil)

	crypt.Reset()
	crypt.Write(scramble)
	crypt.Write(hash)
	scramble = crypt.Sum(nil)

	for i := range scramble {
		scramble[i] ^= stage1[i]
	}

	return scramble
}

func (p *PacketIo) HandleError(data []byte) {
	pos := 1
	code := binary.LittleEndian.Uint16(data[pos:])
	pos += 2
	pos++
	state := string(data[pos : pos+5])
	pos += 5
	msg := string(data[pos:])
	fmt.Printf("code:%d, state:%s, msg:%s\n", code, state, msg)
}
