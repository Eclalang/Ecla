# Os

OS library allows you to interact with the operating system.

## Index

- [Chown(path string, uid int, gid int)](#chown)
- [ClearEnv()](#clearenv)
- [Create(name string)](#create)
- [Getegid() int](#getegid)
- [GetEnv(key string) string](#getenv)
- [Geteuid() int](#geteuid)
- [Getgid() int](#getgid)
- [GetHostName() string](#gethostname)
- [Getpid() int](#getpid)
- [Getppid() int](#getppid)
- [Getuid() int](#getuid)
- [GetUserHomeDir() string](#getuserhomedir)
- [Getwd() string](#getwd)
- [Mkdir(name string)](#mkdir)
- [ReadDir(name string) []string](#readdir)
- [ReadFile(str string) string](#readfile)
- [Remove(filename string)](#remove)
- [RemoveAll(path string)](#removeall)
- [SetEnv(key string, value string)](#setenv)
- [SetEnvByFile(filename string)](#setenvbyfile)
- [UnsetEnv(key string)](#unsetenv)
- [WriteFile(filename string, content string)](#writefile)

## Chown
```
function chown(path string, uid int, gid int)
```
Chown changes the owner and group of the file

## ClearEnv
```
function clearEnv()
```
Clears all environment variables

## Create
```
function create(name string)
```
Creates a file and returns the file object

## Getegid
```
function getegid() int
```
Gets the effective group ID of the calling process

## GetEnv
```
function getEnv(key string) string
```
Returns an environment variable

## Geteuid
```
function geteuid() int
```
Gets the effective user ID of the calling process

## Getgid
```
function getgid() int
```
Gets the group ID of the calling process

## GetHostName
```
function getHostName() string
```
Gets the hostname of the machine

## Getpid
```
function getpid() int
```
Gets the process ID of the calling process

## Getppid
```
function getppid() int
```
Gets the process ID of the parent of the calling process

## Getuid
```
function getuid() int
```
Gets the user ID of the calling process

## GetUserHomeDir
```
function getUserHomeDir() string
```
Gets the home directory of the current user

## Getwd
```
function getwd() string
```
Gets the current working directory

## Mkdir
```
function mkdir(name string)
```
Creates a new directory

## ReadDir
```
function readDir(name string) []string
```
Reads a directory and returns the names of the files and directories

## ReadFile
```
function readFile(str string) string
```
Returns the content of a file

## Remove
```
function remove(filename string)
```
Removes a file

## RemoveAll
```
function removeAll(path string)
```
Removes a directory and all its contents

## SetEnv
```
function setEnv(key string, value string)
```
Sets an environment variable

## SetEnvByFile
```
function setEnvByFile(filename string)
```
Sets an environment variable by reading a file

## UnsetEnv
```
function unsetEnv(key string)
```
Unsets an environment variable

## WriteFile
```
function writeFile(filename string, content string)
```
Writes a file with the given content