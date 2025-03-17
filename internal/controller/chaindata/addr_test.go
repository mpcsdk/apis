package chaindata

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shengdoushi/base58"
)

func TestAddress_Scan(t *testing.T) {
	// 1. 移除现有 0x 前缀
	evmAddr := "0xECa9bC828A3005B9a3b909f2cc5c2a54794DE05F"
	evmAddr = strings.ToLower(evmAddr)
	addr := strings.TrimPrefix(evmAddr, "0x")

	// 2. 添加 0x41 前缀
	addr = "41" + addr
	addrBytes := common.Hex2Bytes(addr)
	// 3. 计算双重 SHA256 哈希
	hash1 := sha256.Sum256(addrBytes)
	hash2 := sha256.Sum256(hash1[:])
	checksum := hash2[0:4]

	addchecksum := append(addrBytes, checksum...)

	// 10. Base58 编码
	tronAddr := base58.Encode(addchecksum, TronAlphabet)
	fmt.Println(tronAddr)
	// TXYZopYRdj2D9XRtbG411XZZ3kM5VkAeBf
}
