package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base32"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	maxClockDrift = 1
	maxAttempts   = 3
)

type TOTPConfig struct {
	UserID string
	Secret string
	Step   int64
}

var rdb *redis.Client
var ctx = context.Background()

func generateTOTP(config *TOTPConfig, currentTime time.Time) (string, error) {
	interval := currentTime.Unix() / config.Step
	return generateCodeFromInterval(config.Secret, interval)
}

func generateCodeFromInterval(secret string, interval int64) (string, error) {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(interval))

	h := hmac.New(sha512.New, []byte(secret))
	_, err := h.Write(b[:])
	if err != nil {
		return "", err
	}
	hashed := h.Sum(nil)

	offset := hashed[len(hashed)-1] & 0xF
	value := int(hashed[offset]&0x7F)<<24 |
		int(hashed[offset+1]&0xFF)<<16 |
		int(hashed[offset+2]&0xFF)<<8 |
		int(hashed[offset+3]&0xFF)

	code := value % 1000000
	return fmt.Sprintf("%06d", code), nil
}

func validateTOTP(providedCode string, config *TOTPConfig, currentTime time.Time) (bool, error) {
	lastUsedCode, err := rdb.Get(ctx, "lastUsedCode:"+config.UserID).Result()
	if err == nil && lastUsedCode == providedCode {
		return false, errors.New("code is used a second time - replay attack detected")
	}

	currentInterval := currentTime.Unix() / config.Step

	for drift := -maxClockDrift; drift <= maxClockDrift; drift++ {
		code, err := generateCodeFromInterval(config.Secret, currentInterval+int64(drift))
		if err != nil {
			return false, err
		}
		if code == providedCode {
			rdb.Set(ctx, "lastUsedCode:"+config.UserID, providedCode, time.Duration(config.Step)*time.Second)
			rdb.Del(ctx, "failedAttempts:"+config.UserID)
			return true, nil
		}
	}

	failedAttempts, _ := rdb.Incr(ctx, "failedAttempts:"+config.UserID).Result()
	if failedAttempts > maxAttempts {
		return false, errors.New("rate limit exceeded")
	}

	return false, nil
}

func main() {

	fmt.Println("")
	fmt.Println("# 1764Snippets")
	fmt.Println("## Crypto Time-based One-Time Password (TOTP)")

	/*
		This Go code implements the Time-based One-Time Password (TOTP) authentication mechanism. It uses HMAC with SHA512 for code generation and integrates with Redis for rate limiting and replay attack protection. The code provides functions to generate and validate TOTP codes, and it uses Docker (through a Makefile) to ensure a Redis instance is running during execution.

		This snippet is a modified version of the 42Snippets "Golang Time-based One-Time Password (TOTP)".
	*/

	// Initialize Redis client
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// This secret should be securely generated and shared with the user
	// For the purpose of this demo, it's hardcoded
	secret := "JBSWY3DPEHPK3PXP"
	userID := "user12345"

	decodedSecret, err := base32.StdEncoding.DecodeString(strings.ToUpper(secret))
	if err != nil {
		fmt.Println("Error decoding secret:", err)
		return
	}

	config := &TOTPConfig{
		UserID: userID,
		Secret: string(decodedSecret),
		Step:   30,
	}

	totp, err := generateTOTP(config, time.Now())
	if err != nil {
		fmt.Println("Error generating TOTP:", err)
		return
	}

	fmt.Println("Generated TOTP: ", totp)

	fmt.Print("Using the TOTP: ")
	isValid, err := validateTOTP(totp, config, time.Now())
	if err != nil {
		fmt.Println("Error validating TOTP:", err)
		return
	}
	if isValid {
		fmt.Println("TOTP is valid!")
	} else {
		fmt.Println("TOTP is invalid.")
	}

	fmt.Print("Using 111111 as TOTP: ")
	isNotValid, err := validateTOTP("111111", config, time.Now())
	if err != nil {
		fmt.Println("Error validating TOTP: ", err)
		return
	}
	if isNotValid {
		fmt.Println("TOTP is valid!")
	} else {
		fmt.Println("TOTP is invalid.")
	}

	// To demonstrate that it works only one time
	fmt.Println("Using the TOTP again:")
	isValid2, err := validateTOTP(totp, config, time.Now())
	if err != nil {
		fmt.Println("Error validating TOTP: ", err)
		return
	}
	if isValid2 {
		fmt.Println("TOTP is valid!")
	} else {
		fmt.Println("TOTP is invalid.")
	}
}
