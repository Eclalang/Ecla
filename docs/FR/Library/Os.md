# Os

La librairie OS permet d'interagir avec le système d'exploitation.

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
Chown change le propriétaire et le groupe du fichier

### Exemple :
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
Efface toutes les variables d'environnement

### Exemple :
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
Crée un fichier et retourne l'objet du fichier

### Exemple :
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
Récupère l'ID effectif du groupe du processus appelant


### Exemple :
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
Retourne une variable d'environnement

### Exemple :
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
Récupère l'ID effectif de l'utilisateur appelant le processus

### Exemple :
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
Récupère l'ID de groupe du processus appelant

### Exemple :
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
Récupère le nom d'hôte de la machine

### Exemple :
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
Récupère l'ID du processus appelant

### Exemple :
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
Récupère l'ID du processus parent du processus appelant

### Exemple :
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
Récupère l'ID utilisateur du processus appelant

### Exemple :
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
Récupère le répertoire utilisateur de l'utilisateur actuel

### Exemple :
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
Récupère le répertoire courant

### Exemple :
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
Crée un nouveau répertoire

### Exemple :
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
Lit le répertoire et retourne le nom des fichiers et des répertoires

### Exemple :
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
Retourne le contenu du fichier

### Exemple :
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
Supprime un fichier

### Exemple :
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
Supprime un répertoire et son contenu

### Exemple :
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
Définit une variable d'environnement

### Exemple :
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
Définit une variable d'environnement en lisant un fichier

### Exemple :
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
Supprime une variable d'environnement

### Exemple :
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
Écrit un fichier avec le contenu donné

### Exemple :
```ecla
import "os";

function testWriteFile() {
    os.writeFile("test.txt", "Hello World!");
}
```