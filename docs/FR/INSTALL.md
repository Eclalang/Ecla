# Comment installer Ecla

## Télécharger une release

Vous pouvez télécharger l'interpréteur Ecla en tant que binaire autonome pré-construit pour votre système depuis la dernière release sur https://github.com/Eclalang/Ecla/releases

Renommez le fichier exécutable en `ecla` ou `ecla.exe` selon votre système d'exploitation, et ajoutez le chemin du fichier à votre variable d'environnement PATH.

Redémarrez votre terminal pour appliquer les changements.

## Construire depuis les sources

Alternativement, vous pouvez construire l'interpréteur Ecla depuis les sources

## Prérequis

Installer Go 1.21.1 ou plus.

### Windows:

- Télécharger l'installateur Go pour Windows sur le site officiel de Go : https://go.dev/dl/

- Ouvrez le fichier d'installation et suivez les instructions pour installer Go.

- Après l'installation, ouvrez l'invite de commande de Windows (cmd) et écrivez "go version". Cela devrait afficher la version de Go que vous venez tout juste d'installer.

### Linux:
- Utilisez les commandes suivantes pour installer Go:

```bash
wget https://dl.google.com/go/go1.21.1.linux-amd64.tar.gz
sudo tar -xvf go1.21.1.linux-amd64.tar.gz
sudo mv go /usr/local
```

- Ajoutez les lignes suivantes à votre fichier ~/.bashrc pour ajouter Go a votre variable d'environnement PATH:

```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

- Sauvegardez et fermez le fichier. Ensuite, rechargez votre terminal et écrivez "go version". Cela devrait afficher la version de Go que vous venez d'installer.

## Construire l'interpréteur d'Ecla

- Clonez le répertoire:

```bash
git clone https://github.com/Eclalang/Ecla.git
cd Ecla
```

- Exécutez la commande suivante pour créer le fichier exécutable:

```bash
go build -o .
```

> Si vous le voulez vous pouvez aussi exporter l'executable a vos variables d'environnements.

## Vous êtes prêt à utiliser Ecla!
