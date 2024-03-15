## This is a small monitor of the availability of a network resource using ping to it.  
## When the ping stops, a sound begins to be made.  
## When the ping resumes, the sound stops.    
 
Usage of gomonitor.exe:
  -interval int
        Interval in seconds between connect attempts (default 10)
  -ip string
        IP address to connect (default "10.45.1.20")
  -port string
  -port string
        Port to use for connect (default "443")
  -protocol string
        Protocol to use for connect (tcp/udp) (default "tcp")
  -sound string
        Path to the sound file (default "C:\\Windows\\Media\\ringout.wav")
  -timeout int
        Timeout in seconds after which the alarm sound starts (default 60)
______________________________________________________________________________________
  
Usage example:

usage when compiling and running:  

go run main.go -ip "8.8.8.8" -sound "path/to/soundfile.wav"  
  
    
use when running an already compiled program:  
  
goMonitor -ip "8.8.8.8" -sound "path/to/soundfile.wav" 


