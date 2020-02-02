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
| **[Booleans](#Booleans)** | Booleanske verdier | Verdier som kun kan være `true` eller `false`. |

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

**Spesielle typer heltall:** `uint8`, `uint16`, `uint32`, `uint64`, `int8`, `int16`, `int32` og `in64`.  
Tallene bak forteller hvor mange bits hver type bruker.  

> **Eksempel:** `uint8` forventer 8 positive bits.

### Floating-Point Numbers

1. Desimaltall i Go er unøyaktige.
2. Desimaltall i Go har en definert størrelse (32 eller 64 bit). Jo høyere, jo mer nøyaktig.
3. I tillegg til nummer, kan Floating-Point Numbers vise andre verdier:
    * NaN (Not A Number): ting som 0 av 0
    * Positiv og Negativ uendelig

Go bruker to forskjellige typer desimaltall, `float32` og `float64`.  
Disse er ofte nevnt som `single precision` og `double precision`.  

> _Standarden er å jobbe med `float64` så langt det er mulig._

I tillegg er det to ekstra typer for å representere komplekse nummere (imaginære nummere):  
`complex64` og `complex128`.

### Strings

Streng konstanter kan defineres med både "hermetegn" og gravistegn ( ` ) [engelsk: *backticks*].

Forskjellen mellom disse er at tekst mellom hermetegn ikke kan gå over flere linjer, og tillater spesielle rømningssekvenser. Eksempler på dette er:

* "\n" som blir byttet ut med en ny linje
* "\t" som blir byttet ut med en tab

**Andre ting å bemerke om strenger:**  

* Alle tegn inne i en streng regnes som en del av strengen. Dette gjelder også mellomrom.
* Strenger starter fra `0`, ***ikke*** `1`, når man teller antall tegn.
* Om du ber om ett spesifikt tegn inne i en streng, så vil dette returneres som en [`int`](#Integers) med byte koden til tegnet.
* Konkatenering (å sette sammen), er hva vi gjør for å slå sammen to strenger. Bare tenk på det som vanlig addisjon, men med bokstaver.

**Funksjoner ofte benyttet i strings:**  
| Funksjon | Hva den gjør |
|-|-|
| `len` | Finner lengden på en string |
| `[nummer]` | Går til et tegn inne i strengen. Starter ***ikke*** å telle fra `1`, men fra `0`. |
| `+`| Slår sammen to strenger. Se [Mattematikk](#Mattematikk) for mer utdypende informasjon. |

### Booleans

Booleanske verdier kan bare være `true` eller `false`. Egentlig er det en 1-bit integer verdi, men vi bruker det som en egen datatype.  
Vi bruker tre forskjellige logiske operatører i sammenheng med Booleanske verdier.

| Type | Leses som | Forklaring |
|------|-----------|------------|
| `&&` | and | Begge verdier må være like for at dette skal gå opp |
| `||` | or | Verdien er enten det ene eller det andre |
| `!` | not | Ikke denne verdien |

**Oversikt:** `and`
| Uttrykk | Verdi |
|---------|-------|
| `true && true` | `true` |
| `true && false` | `false` |
| `false && true` | `false` |
| `false && false` | `false` |

**Oversikt:** `or`
| Uttrykk | Verdi |
|---------|-------|
| `true || true` | `true` |
| `true || false` | `true` |
| `false || true` | `true` |
| `false || false` | `false` |

**Oversikt:** `not`
| Uttrykk | Verdi |
|---------|-------|
| `!true` | `false` |
| `!false` | `true`|
