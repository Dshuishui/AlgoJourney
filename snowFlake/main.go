package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	epoch         int64 = 1609459200000 // 起始时间戳（2021-01-01）
	machineBits   uint8 = 10            // 机器ID所占位数
	sequenceBits  uint8 = 12            // 序列号所占位数
	maxMachineID        = -1 ^ (-1 << machineBits)
	maxSequence         = -1 ^ (-1 << sequenceBits)

	timeShift     = machineBits + sequenceBits
	machineShift  = sequenceBits
)

type Snowflake struct {
	mu        sync.Mutex
	timestamp int64
	machineID int64
	sequence  int64
}

func NewSnowflake(machineID int64) *Snowflake {
	if machineID < 0 || machineID > maxMachineID {
		panic("machine ID out of range")
	}
	return &Snowflake{machineID: machineID}
}

func (s *Snowflake) NextID() (int64, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    now := time.Now().UnixMilli()

    if now < s.timestamp {
        // 时钟回拨：等待短时间或返回错误（这里选择等待）
        // 等待回到上次时间或超时
        waitUntil := s.timestamp
        for now < waitUntil {
            time.Sleep(time.Millisecond) // 减少 CPU 消耗
            now = time.Now().UnixMilli()
            // 可添加超时保护，避免无限等待
        }
    }

    if now == s.timestamp {
        s.sequence = (s.sequence + 1) & maxSequence  // 当时间戳相同的时候，自增序列号即可，如果当前时间戳的序列号都被分配完，就把将其划分到写一个时间戳去
        if s.sequence == 0 {
            // 序列用完，等待下一毫秒
            for now <= s.timestamp {
                time.Sleep(time.Microsecond * 50) // 比忙等省CPU
                now = time.Now().UnixMilli()
            }
        }
    } else {
        s.sequence = 0
    }

    s.timestamp = now
    id := ((now - epoch) << timeShift) | (s.machineID << machineShift) | s.sequence
    return id, nil
}

func main() {
	node := NewSnowflake(1) // 节点ID = 1
	for range 10 {
		fmt.Println(node.NextID())
	}
}
