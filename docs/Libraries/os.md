# OS

***
##  Documentation.
### OS is a library that allows you to interact with the operating system.
###

***
## Index.

* [fonction Chown(path string, uid int, gid int)](#fonction-chown)
* [fonction ClearEnv()](#fonction-clearenv)
* [fonction Create(path string)](#fonction-create)
* [fonction Getegid() int](#fonction-gategid)
* [fonction GetEnv(key string) string](#fonction-getenv)
* [fonction Geteuid() int](#fonction-geteuid)
* [fonction Getgid() int](#fonction-getgid)
* [fonction GetHostName() string](#fonction-gethostname)
* [fonction Getpid() int](#fonction-getpid)
* [fonction Getppid() int](#fonction-getppid)
* [fonction Getuid() int](#fonction-getuid)
* [fonction GetUserHomeDir() string](#fonction-getuserhomedir)
* [fonction Getwd() string](#fonction-getwd)
* [fonction Mkdir(path string)](#fonction-mkdir)
* [fonction ReadDir(path string) []string](#fonction-readdir)
* [fonction ReadFile(path string) string](#fonction-readfile)
* [fonction Remove(path string)](#fonction-remove)
* [fonction RemoveAll(path string)](#fonction-removeall)
* [fonction SetEnv(key string, value string)](#fonction-setenv)
* [fonction SetEnvByFile(path string)](#fonction-setenvbyfile)
* [fonction UnsetEnv(key string)](#fonction-unsetenv)
* [fonction WriteFile(path string, data string)](#fonction-writefile)
##
### Fonction Chrown
```
function chown(path string, uid int, gid int)
```
Chown changes the owner and group of the file
### Fonction ClearEnv
```
function clearEnv()
```
Clears all environment variables
### Fonction Create
```
function reate(path string)
```
Creates a file and returns the file object
### Fonction Getegid
```
function getegid() int
```
Gets the effective group ID of the calling process
### Fonction GetEnv
```
function getEnv(key string) string
```
Returns an environment variable
### Fonction Geteuid
```
function geteuid() int
```
Gets the effective user ID of the calling process
### Fonction Getgid
```
function getgid() int
```
Gets the group ID of the calling process
### Fonction GetHostName
```
function getHostName() string
```
Gets the hostname of the machine
### Fonction Getpid
```
function getpid() int
```
Gets the process ID of the calling process
### Fonction Getppid
```
function getppid() int
```
Gets the process ID of the parent of the calling process
### Fonction Getuid
```
function getuid() int
```
Gets the user ID of the calling process
### Fonction GetUserHomeDir
```
function getUserHomeDir() string
```
Gets the home directory of the current user
### Fonction Getwd
```
function getwd() string
```
Gets the current working directory
### Fonction Mkdir
```
function mkdir(path string)
```
Creates a new directory
### Fonction ReadDir
```
function readDir(path string) []string
```
Reads a directory and returns the names of the files and directories
### Fonction ReadFile
```
function readFile(path string) string
```
Returns the content of a file
### Fonction Remove
```
function remove(path string)
```
Removes a file
### Fonction RemoveAll
```
function removeAll(path string)
```
Removes a directory and all its contents
### Fonction SetEnv
```
function setEnv(key string, value string)
```
Sets an environment variable
### Fonction SetEnvByFile
```
function setEnvByFile(path string)
```
Sets an environment variable by reading a file
### Fonction UnsetEnv
```
function unsetEnv(key string)
```
Unsets an environment variable
### Fonction WriteFile
```
function writeFile(path string, data string)
```
Writes a file with the given content
##