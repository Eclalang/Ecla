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
- [GetHostname() string](#gethostname)
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

### Example :
```ecla
import "os";

function testChown() {
    os.chown("test.txt", 0, 0);
}
```

## ClearEnv
```
function clearEnv()
```
Clears all environment variables

### Example :
```ecla
import "os";

function testClearEnv() {
    os.clearEnv();
}
```

## Create
```
function create(name string)
```
Creates a file and returns the file object

### Example :
```ecla
import "os";

function testCreate() {
    os.create("test.txt");
}
```

## Getegid
```
function getegid() int
```
Gets the effective group ID of the calling process

### Example :
```ecla
import "console";
import "os";

function testGetegid() {
    var egid int = os.getegid();
    console.println(egid);
}
```

## GetEnv
```
function getEnv(key string) string
```
Returns an environment variable

### Example :
```ecla
import "console";
import "os";

function testGetEnv() {
    var env string = os.getEnv("API_KEY");
    console.println(env);
}
```

## Geteuid
```
function geteuid() int
```
Gets the effective user ID of the calling process

### Example :
```ecla
import "console";
import "os";

function testGeteuid() {
    var euid int = os.geteuid();
    console.println(euid);
}
```

## Getgid
```
function getgid() int
```
Gets the group ID of the calling process

### Example :
```ecla
import "console";
import "os";

function testGetgid() {
    var gid int = os.getgid();
    console.println(gid);
}
```

## GetHostname
```
function getHostname() string
```
Gets the hostname of the machine

### Example :
```ecla
import "console";
import "os";

function testGetHostname() {
    var hostname string = os.getHostname();
    console.println(hostname);
}
```

## Getpid
```
function getpid() int
```
Gets the process ID of the calling process

### Example :
```ecla
import "console";
import "os";

function testGetpid() {
    var pid int = os.getpid();
    console.println(pid);
}
```

## Getppid
```
function getppid() int
```
Gets the process ID of the parent of the calling process

### Example :
```ecla
import "console";
import "os";

function testGetppid() {
    var ppid int = os.getppid();
    console.println(ppid);
}
```

## Getuid
```
function getuid() int
```
Gets the user ID of the calling process

### Example :
```ecla
import "console";
import "os";

function testGetuid() {
    var uid int = os.getuid();
    console.println(uid);
}
```

## GetUserHomeDir
```
function getUserHomeDir() string
```
Gets the home directory of the current user

### Example :
```ecla
import "console";
import "os";

function testGetUserHomeDir() {
    var homeDir string = os.getUserHomeDir();
    console.println(homeDir);
}
```

## Getwd
```
function getwd() string
```
Gets the current working directory

### Example :
```ecla
import "console";
import "os";

function testGetwd() {
    var wd string = os.getwd();
    console.println(wd);
}
```

## Mkdir
```
function mkdir(name string)
```
Creates a new directory

### Example :
```ecla
import "os";

function testMkdir() {
    os.mkdir("testDir");
}
```

## ReadDir
```
function readDir(name string) []string
```
Reads a directory and returns the names of the files and directories

### Example :
```ecla
import "console";
import "os";

function testReadDir() {
    var files []string = os.readDir("testDir");
    console.println(files);
}
```

## ReadFile
```
function readFile(str string) string
```
Returns the content of a file

### Example :
```ecla
import "console";
import "os";

function testReadFile() {
    var content string = os.readFile("test.txt");
    console.println(content);
}
```

## Remove
```
function remove(filename string)
```
Removes a file

### Example :
```ecla
import "os";

function testRemove() {
    os.remove("test.txt");
}
```

## RemoveAll
```
function removeAll(path string)
```
Removes a directory and all its contents

### Example :
```ecla
import "os";

function testRemoveAll() {
    os.removeAll("testDir");
}
```

## SetEnv
```
function setEnv(key string, value string)
```
Sets an environment variable

### Example :
```ecla
import "os";

function testSetEnv() {
    os.setEnv("API_KEY", "123456789");
}
```

## SetEnvByFile
```
function setEnvByFile(filename string)
```
Sets an environment variable by reading a file

### Example :
```ecla
import "os";

function testSetEnvByFile() {
    os.setEnvByFile("test.txt");
}
```

## UnsetEnv
```
function unsetEnv(key string)
```
Unsets an environment variable

### Example :
```ecla
import "os";

function testUnsetEnv() {
    os.unsetEnv("API_KEY");
}
```

## WriteFile
```
function writeFile(filename string, content string)
```
Writes a file with the given content

### Example :
```ecla
import "os";

function testWriteFile() {
    os.writeFile("test.txt", "Hello World!");
}
```