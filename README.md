# Card Games

Implementierung von Kartenspielen für die Konsole.
Es sind einige Dateien in Packages vorgegeben, die ein Grundgerüst für ein Spiel bilden.

## Package `cards`

Dieses Package enthält Datentypen für Spielkarten, Kartenstapel und Handkarten
von Spielern.
Außerdem sind diverse Methoden und Funktionen implementiert,
die in Spielen verwendet werdenkönnen.

Die Typen in diesem Spiel sind generisch, d.h. nicht auf ein bestimmtes Spiel ausgelegt.
Es geht um allgemeine Funktionalitäten wie z.B. das Ziehen von Karten, die allen
(oder sehr vielen) Kartenspielen gemein sind.

## Package `gamestate`

Dieses Package enthält eine Struktur für die Verwaltung des Spiel-Zustands.
Es enthält eine Liste von Handkarten, einen Karten- und einen Ablagestapel
und es kennt den Spieler, der an der Reihe ist.

Auch dieses Struct ist generisch und nicht für ein bestimmtes Spiel gemacht.
Es ist noch sehr rudimentär, es fehlen Tests und man könnte auch noch
Convenience-Methoden z.B. für den Spielerwechsel oder die Abfrage des aktuellen
oder eines bestimmten Spielers hinzufügen.

## Package `maumau`

Hier ist ein Grundgerüst für ein Mau-Mau-Spiel vorgegeben.
Näheres siehe dortige [README.md](maumau/README.md).
