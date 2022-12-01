# Mau Mau

In diesem Package ist ein erster Wurf für eine Mau-Mau-Spiellogik vorgegeben.
Ihre Aufgabe besteht darin, das Spiel zu vervollständigen bzw. zu erweitern.

## Funktion für die Abfrage eines Spielzuges

In der Vorgabe wird bisher hartcodiert ein Spieler nach seinem Spielzug gefragt.
Verlagern Sie diesen Code in eine Funktion, die den `state` und eine Spielernummer
erwartet und die den Zug mit diesem Spieler durchführt.

*Hinweis:* Diese Funktion ist eine Vorbereitung für die folgende Aufgabe.
Durch diese Funktion werden Sie flexibel, um beliebige Spieler nach ihren Zügen fragen
zu können.

## Vollständige Spielrunden

In der Vorgabe wird bisher nur ein Spieler nach seinem Zug gefragt.
Erweitern Sie das Spiel, so dass es bis zu Ende läuft:

* Das Spiel soll reihum alle Spieler nach ihren Zügen fragt und diese durchführen.
* Das Spiel soll so lange laufen, wie noch mehr als ein Spieler Karten hat.
* Es sollen nur solche Spieler nach Karten gefragt werden, die auch wirklich Karten haben.

## Basis-Spielregeln umsetzen

Bei Mau Mau ist gefordert, dass die gelegte Spielkarte zur
Karte auf dem Ablagestapel passen muss.
Dies ist die Basis-Spielregel, nach der gezogen wird.
Setzen Sie diese Zug-Regel um:

* Schreiben Sie eine Funktion `match()`, die zwei Spielkarten erwartet
  und `true` oder `false` liefert, je nachdem, ob die Karten passen.
* Verwenden Sie die Funktion `match()`, um beim Durchführen eines Zuges
  zu prüfen, ob der Zug erlaub ist und ihn nur dann durchführen.

## Ermöglichen, dass der Spieler eine Karte zieht

Neben dem Ausspielen einer Karte muss der Spieler auch die Wahl haben,
eine Karte vom Stapel zu ziehen.
Insbesondere muss er eine Karte ziehen, falls er keine legen kann.

## Fehlerbehandlung

Die vorgegebenen Funktionen enthalten bisher keinerlei Fehlerbehandlung.
Auch bei den vorherigen Aufgaben war eine Fehlerbehandlung noch nicht gefragt.
D.h. es wird bspw. beim Ziehen von Karten aus einem `Deck` oder einer `Hand`
nicht geprüft, ob überhaupt noch Karten vorhanden sind.
Auch beim Ausspielen von Karten im Mau-Mau-Gerüst wird nicht geprüft,
ob die Auswahl überhaupt gültig war.

Ergänzen Sie an sinnvollen Stellen Fehlerbehandlungen.

* Beim Durchführen eines Zuges soll geprüft werden,
  ob die Wahl des Spielers eine Karte aus seiner Hand ist.
* Beim Ziehen von Karten soll vorher geprüft werden,
  ob es überhaupt noch Karten auf auf dem Stapel gibt.
  Falls nicht, sollen alle Karten auf dem Ablagestapel außer der obersten neu gemischt
  und auf das Deck gelegt werden.
* Falls Ihnen weitere Fehler oder sinnvolle Prüfungen auffallen,
  sollten Sie natürlich auch diese berücksichtigen.

## Sonderregeln

Es gibt bei Mau Mau diverse Sonderregeln, die je nach Spielrunde unterschiedlich
streng ausgelegt werden:

* Wird eine "7" gelegt, muss der nächste Spieler zwei Karten ziehen,
  * Falls möglich, kann der Spieler stattdessen selbst eine "7" legen.
  * Dann muss der übernächste Spieler vier Karten ziehen.
  * Analog für die nächsten beiden "7".
* Wird eine "8" gelegt, muss der nächste Spieler aussetzen.
  * ... es sei denn, er kann selbst eine "8" legen.
* Wird ein Bube gelegt, darf der Spieler sich eine Farbe wünschen.
* Wird ein Ass gelegt, darf/muss der Spieler gleich noch einmal legen.

Setzen Sie diese Sonderregeln um.
Sie können auch zu Beginn des Spiels wählbar machen, welche Sonderregeln gelten sollen.

*Hinweis:* Versuchen Sie nicht gleich, alles auf einmal umzusetzen.
Es bietet sich an, mit einer einfachen Regel wie der "8"er-Regel anzufangen und
auch das "Gegenlegen" nicht sofort umzusetzen. Gleiches gilt für die "7"er-Regel.

## Bot

Setzen Sie einen Bot um, der das Spiel spielt.

*Hinweis:* Auch hier müssen Sie nicht gleich eine perfekte KI bauen,
sondern sie sollten sich langsam steigern:

* Starten Sie mit einem Bot, der zumindest korrekte Züge macht, falls dies möglich ist.
* Als nächstes könnten Sie dafür sorgen, dass der Bot Farben bevorzugt,
  von denen er viele Karten hat.
* Der nächste Schritt ist ein Sonderregel-optimiertes Spiel, bei dem der Bot
  z.B. "7" oder "8" bevorzugt.
* Noch besser wird es, wenn der Bot sich merkt, welche Karten andere Spieler gespielt
  oder sich gewünscht haben und versucht, diese zu vermeiden.
