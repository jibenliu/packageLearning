package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"time"
)

const HeadSize = 4

type Buffer struct {
	reader    io.Reader
	header    string
	buf       []byte
	bufLength int
	start     int
	end       int
}

func NewBuffer(reader io.Reader, header string, len int) *Buffer {
	buf := make([]byte, len)
	return &Buffer{reader, header, buf, len, 0, 0}
}

func (b *Buffer) grow() {
	if b.start == 0 {
		return
	}
	copy(b.buf, b.buf[b.start:b.end])
	b.end -= b.start
	b.start = 0
}

func (b *Buffer) len() int {
	return b.end - b.start
}

func (b *Buffer) seek(n int) ([]byte, error) {
	if b.end-b.start >= n {
		buf := b.buf[b.start : b.start+n]
		return buf, nil
	}
	return nil, errors.New("not enough")
}

func (b *Buffer) read(offset, n int) []byte {
	b.start += offset
	buf := b.buf[b.start : b.start+n]
	b.start += n
	return buf
}

func (b *Buffer) readFomReader() error {
	if b.end == b.bufLength {
		return errors.New(fmt.Sprintf("一个完整的数据包太长已经超过你定义的example.BufferLength(%d)\n", b.bufLength))
	}
	n, err := b.reader.Read(b.buf[b.end:])
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	b.end += n
	return nil
}

func (b *Buffer) checkMsg(msg chan string) error {
	HeaderLength := HeadSize + len(b.header)
	headBuf, err1 := b.seek(HeaderLength)
	if err1 != nil {
		return nil
	}
	if string(headBuf[:len(b.header)]) == b.header {

	} else {
		return errors.New("消息头部不正确")
	}

	contentSize := int(binary.BigEndian.Uint32(headBuf[len(b.header):]))
	if b.len() >= contentSize-HeaderLength {
		contentBuf := b.read(HeaderLength, contentSize)
		msg <- string(contentBuf)
		err3 := b.checkMsg(msg)
		if err3 != nil {
			return err3
		}
	}
	return nil
}

func (b *Buffer) Read(msg chan string) error {
	for {
		b.grow()
		err1 := b.readFomReader()
		if err1 != nil {
			return err1
		}

		err2 := b.checkMsg(msg)
		if err2 != nil {
			return err2
		}
	}
}
