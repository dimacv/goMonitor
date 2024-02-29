## This is a small monitor of the availability of a network resource using ping to it.  
## When the ping stops, a sound begins to be made.  
## When the ping resumes, the sound stops.    
  
    
usage when compiling and running:  

go run main.go -ip "8.8.8.8" -sound "path/to/soundfile.wav"  
  
    
use when running an already compiled program:  
  
goMonitor -ip "8.8.8.8" -sound "path/to/soundfile.wav"  
