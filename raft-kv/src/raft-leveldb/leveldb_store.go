package raft_leveldb

import (
	"bytes"
	"encoding/binary"
	"github.com/hashicorp/go-msgpack/codec"
	"github.com/hashicorp/raft"
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	DB_PATH = "/Users/didi/Documents/Go/src/github.com/turingkv/raft-kv/src/raft-leveldb/db1"
)

// Decode reverses the encode operation on a byte slice input
func decodeMsgPack(buf []byte, out interface{}) error {
	r := bytes.NewBuffer(buf)
	hd := codec.MsgpackHandle{}
	dec := codec.NewDecoder(r, &hd)
	return dec.Decode(out)
}

// Encode writes an encoded object to a new bytes buffer
func encodeMsgPack(in interface{}) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(nil)
	hd := codec.MsgpackHandle{}
	enc := codec.NewEncoder(buf, &hd)
	err := enc.Encode(in)
	return buf, err
}

// Converts bytes to an integer
func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

// Converts a uint to a byte slice
func uint64ToBytes(u uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, u)
	return buf
}

type LeveldbStore struct {
	db   *leveldb.DB
	path string
}

type Options struct {
	Path string
}

func NewLeveldbStore(path string) (*LeveldbStore, error) {
	return New(Options{Path: path})
}

func New(options Options) (*LeveldbStore, error) {
	db, err := leveldb.OpenFile(options.Path, nil)

	if err != nil {
		return nil, err
	}

	store := &LeveldbStore{
		db:   db,
		path: options.Path,
	}

	return store, nil
}

func (l *LeveldbStore) Close() error {
	return l.db.Close()
}

func (l *LeveldbStore) Set(k, v []byte) error {

	if err := l.db.Put(k, v, nil); err != nil {
		return err
	}
	return nil

}

func (l *LeveldbStore) Get(k []byte) ([]byte, error) {

	val, err := l.db.Get(k, nil)
	if err != nil {
		return nil, err
	}

	return append([]byte(nil), val...), nil
}

func (l *LeveldbStore) SetUint64(key []byte, val uint64) error {
	return l.Set(key, uint64ToBytes(val))
}

func (l *LeveldbStore) GetUint64(key []byte) (uint64, error) {
	val, err := l.Get(key)
	if err != nil {
		return 0, err
	}
	return bytesToUint64(val), nil
}

/*
	用来存储raft的日志
*/
func (l *LeveldbStore) StoreLogs(logs []*raft.Log) error {

	for _, log := range logs {
		key := uint64ToBytes(log.Index)
		val, err := encodeMsgPack(log)
		if err != nil {
			return err
		}
		if err := l.db.Put(key, val.Bytes(), nil); err != nil {
			return err
		}
	}

	return nil
}

func (l *LeveldbStore) StoreLog(log *raft.Log) error {
	return l.StoreLogs([]*raft.Log{log})
}

func (l *LeveldbStore) GetLog(idx uint64, log *raft.Log) error {

	val, err := l.db.Get(uint64ToBytes(idx), nil)
	if err != nil {
		return err
	}

	if val == nil {
		return raft.ErrLogNotFound
	}
	return decodeMsgPack(val, log)
}

func (l *LeveldbStore) FirstIndex() (uint64, error) {

	iter := l.db.NewIterator(nil, nil)

	if iter.First() {
		return bytesToUint64(iter.Key()), nil
	} else {
		return 0, nil
	}

}

func (l *LeveldbStore) LastIndex() (uint64, error) {
	iter := l.db.NewIterator(nil, nil)

	if iter.Last() {
		return bytesToUint64(iter.Key()), nil
	} else {
		return 0, nil
	}
}

func (l *LeveldbStore) DeleteRange(min, max uint64) error {

	minKey := uint64ToBytes(min)
	iter := l.db.NewIterator(nil, nil)

	if iter.Seek(minKey) {
		if err := l.db.Delete(iter.Key(), nil); err != nil {
			return err
		}
		for iter.Next() {
			if bytesToUint64(iter.Key()) > max {
				break
			}
			if err := l.db.Delete(iter.Key(), nil); err != nil {
				return err
			}
		}
	}
	return nil
}
