# Comment installer Ecla

## Prérequis

Installer Go 1.19 ou plus.

### Windows:

- Télécharger l'installateur Go pour Windows sur le site officiel de Go : https://go.dev/dl/

- Ouvrez le fichier d'installation et suivez les instructions pour installer Go.

- Après l'installation, ouvrez l'invite de commande de Windows (cmd) et écrivez "go version". Cela devrait afficher la version de Go que vous venez tout juste d'installer.

### Linux:
- Utilisez les commandes suivantes pour installer Go:

```bash
wget https://dl.google.com/go/go1.19.linux-amd64.tar.gz
sudo tar -xvf go1.19.linux-amd64.tar.gz
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