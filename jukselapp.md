# Jukselapp om Go på Norsk (Oversatt av meg)

## Vokabular og ordliste

Liste over forskjellige begreper og hva de kan oversettes til. Jeg er ikke en offisiell oversetter, eller en programmerer enda, men dette er erfaringene jeg har gjort meg opp så langt i læringsløpet.

---

## Datatyper

| Engelsk | Oversettelse | Definisjon |
|---------|--------------|------------|
| Integers | Heltall | Alle tall som kan skrives uten komma. |
| Floating Point Numbers | Desimaltall | Alle tall som ikke er heltall. |
| [Strings](#Strings) | Streng | En liste med tegn med en bestemt lengde. |

---

### Utdypende

#### Strings

Streng konstanter kan defineres med både "hermetegn" og gravistegn ( ` ) [engelsk: *backticks*].

Forskjellen mellom disse er at tekst mellom hermetegn ikke kan gå over flere linjer, og tillater spesielle rømningssekvenser. Eksempler på dette er:

* "\n" som blir byttet ut med en ny linje
* "\t" som blir byttet ut med en tab
