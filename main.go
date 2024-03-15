package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os/exec"
	"time"
)

/*
Usage of gomonitor.exe:

	-interval int
	      Interval in seconds between connect attempts (default 10)
	-ip string
	      IP address to connect (default "10.45.1.20")
	-port string
	      Port to use for connect (default "443")
	-protocol string
	      Protocol to use for connect (tcp/udp) (default "tcp")
	-sound string
	      Path to the sound file (default "C:\\Windows\\Media\\ringout.wav")
	-timeout int
	      Timeout in seconds after which the alarm sound starts (default 60)

Usage example:
gomonitor -ip=192.168.0.1 -port=80 -protocol=udp -sound=/path/to/sound/file.wav
or
.\gomonitor.exe -ip=10.45.1.1 -port=22 -sound=D:\goMonitor\Notification.wav
*/
func main() {
	// Определение аргументов командной строки с значениями по умолчанию
	defaultIP := "10.45.1.20"
	defaultPort := "443"
	defaultProtocol := "tcp"
	defaultSoundFile := "C:\\Windows\\Media\\ringout.wav"
	defaultTimeout := 60  // Установка значения по умолчанию для таймаута в секундах
	defaultInterval := 10 // Установка значения по умолчанию для интервала в секундах

	ip := flag.String("ip", defaultIP, "IP address to connect")
	port := flag.String("port", defaultPort, "Port to use for connect")
	protocol := flag.String("protocol", defaultProtocol, "Protocol to use for connect (tcp/udp)")
	soundFile := flag.String("sound", defaultSoundFile, "Path to the sound file")
	timeout := flag.Int("timeout", defaultTimeout, "Timeout in seconds after which the alarm sound starts")
	interval := flag.Int("interval", defaultInterval, "Interval in seconds between connect attempts")

	flag.Parse() // Разбор аргументов командной строки

	alarmPlaying := false
	elapsed := 0

	for {
		success := connect(*ip, *port, *protocol)
		if !success {
			elapsed += *interval
			if !alarmPlaying && elapsed >= *timeout {
				fmt.Printf("Connect to %s:%s (%s) failed after %d seconds. Starting alarm and playing sound...\n", *ip, *port, *protocol, elapsed)
				// Включаем звуковую тревогу
				go playSound(*soundFile, &alarmPlaying)
			}
		} else {
			if alarmPlaying {
				fmt.Printf("Connect to %s:%s (%s) succeeded. Stopconnect alarm and playing sound...\n", *ip, *port, *protocol)
				// Останавливаем звуковую тревогу
				stopSound(&alarmPlaying)
			} else {
				fmt.Printf("Connect to %s:%s (%s) succeeded.\n", *ip, *port, *protocol)
			}
			elapsed = 0 // Сброс счетчика времени после успешного пинга
		}
		time.Sleep(time.Duration(*interval) * time.Second)
	}
}

func connect(ip string, port string, protocol string) bool {
	address := fmt.Sprintf("%s:%s", ip, port)
	conn, err := net.DialTimeout(protocol, address, time.Second)
	if err != nil {
		fmt.Println("Connect failed:", err)
		return false
	}
	conn.Close()
	return true
}

func stopSound(alarmPlaying *bool) {
	*alarmPlaying = false
}

// When compiling for Windows:
func playSound(soundFile string, alarmPlaying *bool) {
	*alarmPlaying = true

	cmd := exec.Command("C:\\Program Files (x86)\\Windows Media Player\\wmplayer.exe", soundFile) // Путь к файлу для воспроизведения передается как аргумент функции playSound
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Windows Media Player has been started with %s!", soundFile)

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
