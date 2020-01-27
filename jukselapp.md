# Jukselapp om Go på Norsk (Oversatt av meg)

## Vokabular og ordliste

Liste over forskjellige begreper og hva de kan oversettes til. Jeg er ikke en offisiell oversetter, eller en programmerer enda, men dette er erfaringene jeg har gjort meg opp så langt i læringsløpet.

---

### Datatyper

| Engelsk | Oversettelse | Definisjon |
|---------|--------------|------------|
| **[Integers](#Integers)** | Heltall | Alle tall som kan skrives uten komma. |
| **[Floating-Point Numbers](#Floating-Point-Numbers)** | Desimaltall | Alle tall som ikke er heltall. |
| **[Strings](#Strings)** | Streng | En liste med tegn med en bestemt lengde. |

### Mattematikk

| Tegn | Funksjon |
|------|----------|
| `+` | Addisjon |
| `-` | Substraksjon |
| `*` | Multiplikasjon |
| `/` | Divisjon |
| `%` | Rest |
Mattematikk i programmering er ikke begrenset til bare tall, så å plusse to strenger sammen er helt innenfor.  
**Eksempel:** `"Matte" + "matikk" = Mattematikk`

---

## Utdypende

### Integers

Forskjellige typer integers/heltall:

| Engelsk | Norsk |Egenskaper |
|------|----------|-----------|
| `uint` | Usignert Heltall | Kan bare være positive tall, eller 0. |
| `int` | Signert Heltall | Kan være både positive og negative tall. |
| `byte` | byte | Det samme som `uint8`, altså 8 positive bit. |
| `rune` | rune | Det samme som `int32`, altså 64 bit positive eller negative. |

> **Spesielle typer heltall:** `uint8`, `uint16`, `uint32`, `uint64`, `int8`, `int16`, `int32` og `in64`.  
> Tallene bak forteller hvor mange bits hver type bruker.  
>
> **Eksempel:** `uint8` forventer 8 positive bits.

### Floating-Point Numbers

1. Desimaltall i Go er unøyaktige.
2. Desimaltall i Go har en definert størrelse (32 eller 64 bit). Jo høyere, jo mer nøyaktig.
3. I tillegg til nummer, kan Floating-Point Numbers vise andre verdier:
    * NaN (Not A Number): ting som 0 av 0
    * Positiv og Negativ uendelig

> Go bruker to forskjellige typer desimaltall, `float32` og `float64`.  
Disse er ofte nevnt som `single precision` og `double precision`.  

_Standarden er å jobbe med `float64` så langt det er mulig._
> I tillegg er det to ekstra typer for å representere komplekse nummere (imaginære nummere): `complex64` og `complex128`.

### Strings

Streng konstanter kan defineres med både "hermetegn" og gravistegn ( ` ) [engelsk: *backticks*].

Forskjellen mellom disse er at tekst mellom hermetegn ikke kan gå over flere linjer, og tillater spesielle rømningssekvenser. Eksempler på dette er:

* "\n" som blir byttet ut med en ny linje
* "\t" som blir byttet ut med en tab
