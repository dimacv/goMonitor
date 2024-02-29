package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os/exec"
	"time"
)

func main() {
	// Определение аргументов командной строки с значениями по умолчанию
	defaultIP := "10.45.1.1"
	defaultSoundFile := "Start-Process 'C:\\Program Files (x86)\\Windows Media Player\\wmplayer.exe'" // -ArgumentList 'C:\\Windows\\Media\\ringout.wav' -PassThru | Out-File 'C:\\ARHIV\\pingmonitor\\log.txt'" // Убедитесь, что у вас есть этот файл в директории с программой
	//// When compiling for Linux:
	//defaultSoundFile := "./Notification.wav"

	ip := flag.String("ip", defaultIP, "IP address to ping")
	soundFile := flag.String("sound", defaultSoundFile, "Path to the sound file")

	flag.Parse() // Разбор аргументов командной строки

	interval := 2 * time.Second // Интервал между попытками пинга
	alarmPlaying := false

	for {
		success := ping(*ip)
		if !success {
			if !alarmPlaying {
				fmt.Printf("Ping to %s failed! Starting alarm...\n", *ip)
				// Включаем звуковую тревогу
				go playSound(soundFile, &alarmPlaying)
			}
		} else {
			if alarmPlaying {
				fmt.Printf("Ping to %s succeeded! Stopping alarm...\n", *ip)
				// Останавливаем звуковую тревогу
				stopSound(&alarmPlaying)
			} else {
				fmt.Printf("Ping to %s succeeded!\n", *ip)
			}
		}
		time.Sleep(interval)
	}
}

func ping(ip string) bool {
	conn, err := net.DialTimeout("tcp", ip+":80", time.Second)
	if err != nil {
		fmt.Println("Ping failed:", err)
		return false
	}
	conn.Close()
	return true
}

func stopSound(alarmPlaying *bool) {

	*alarmPlaying = false
}

func playSound(soundFile *string, alarmPlaying *bool) {
	*alarmPlaying = true

	cmd := exec.Command("C:\\Program Files (x86)\\Windows Media Player\\wmplayer.exe", "C:\\Windows\\Media\\ringout.wav")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Windows Media Player has been started with ringout.wav!")

	*alarmPlaying = false
}

//// When compiling for Linux:
//func playSound(soundFile *string, alarmPlaying *bool) {
//	*alarmPlaying = true
//	cmd := exec.Command("aplay", *soundFile)
//	err := cmd.Start()
//	if err != nil {
//		fmt.Println("Failed to start sound:", err)
//		*alarmPlaying = false
//		return
//	}
//	cmd.Wait()
//	*alarmPlaying = false
//}
