# Guide de Démarrage Rapide du Langage de Programmation Ecla

Bienvenue dans Ecla ! Ce guide vous aidera à démarrer avec les bases du langage de programmation Ecla.

## Installation

Avant de pouvoir commencer à coder en Ecla, vous devrez installer le langage sur votre système. Veuillez suivre le [guide d'installation](https://github.com/Eclalang/Ecla/blob/main/INSTALL.md) pour des instructions détaillées.

## Création d'un Simple Programme Bonjour le Monde

Plongeons dans la création d'un simple programme "Bonjour, le Monde !" en Ecla. Nous allons décomposer chaque étape :

### Importation des Bibliothèques

```ecla
import "console";
```

En Ecla, vous pouvez importer des bibliothèques pour étendre les fonctionnalités de vos programmes. Ici, nous importons la bibliothèque `console` pour permettre les opérations d'entrée et de sortie de base.

### Déclaration de Variables

```ecla
var hello string = "Hello";
world := "World !";
```

En Ecla, vous pouvez déclarer des variables en utilisant le mot-clé `var` suivi du nom de la variable et de son type, éventuellement en l'initialisant avec une valeur. Alternativement, vous pouvez utiliser le raccourci `:=` pour une déclaration et une initialisation de variable implicite.

### Concaténation de Chaînes de Caractères

```ecla
var helloWorld string;
helloWorld = hello + ", " + world;
```

Ici, nous concaténons les chaînes "Hello" et "World !" dans une nouvelle variable `helloWorld` en utilisant l'opérateur `+`.

### Création d'une Fonction

```ecla
function printHelloWorld() {
    console.println(helloWorld);
}
```

En Ecla, vous pouvez définir des fonctions en utilisant le mot-clé `function` suivi du nom de la fonction et de ses paramètres (s'il y en a). Les fonctions encapsulent des blocs de code réutilisables. Ici, nous définissons une fonction `printHelloWorld` qui affiche la valeur de `helloWorld` dans la console.

### Appel d'une Fonction

```ecla
printHelloWorld();
```

Pour exécuter le code à l'intérieur d'une fonction, vous devez l'appeler. Ici, nous appelons la fonction `printHelloWorld` pour afficher "Hello, World !" dans la console.

### Code Complet

```ecla
import "console";

var hello string = "Hello";
world := "World !";

var helloWorld string;
helloWorld = hello + ", " + world;

function printHelloWorld() {
    console.println(helloWorld);
}

printHelloWorld();
```

Félicitations ! Vous venez de créer votre premier programme en Ecla.

### Exécutez votre Programme

Pour exécuter votre programme, enregistrez le code dans un fichier avec l'extension `.ecla` (par exemple, `hello.ecla`). Ensuite, ouvrez un terminal et naviguez vers le répertoire où le fichier est enregistré. Si vous avez installé Ecla correctement, vous pouvez exécuter le programme en utilisant la commande suivante :

```bash
ecla hello.ecla
```

Quand vous exécutez cette commande, vous devriez voir "Hello, World !" s'afficher dans la console.

Vous êtes maintenant prêt à vous lancer dans votre voyage dans le monde de la programmation avec Ecla. Si vous souhaitez en savoir plus sur le langage, consultez la [documentation officielle](https://github.com/Eclalang/LearnEcla/blob/main/README.md)
> la documentation est encore en cours de construction, mais vous pouvez toujours y trouver des informations utiles.