# go-workerpool
Worker pool implementation. Target is `Make it be a high performance golang module, used in any way`
##Introduction
Inspiration by below link, want to implement a common library to support mass throughput system arch.

http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/

##Prototype

Producer -> Generate Data

Consumer -> Processing Data

Worker   -> Do the real job
