Implementiere bitte in Go das Composite Pattern für das Fenster-Beispiel aus der Vorlesung. Es gibt ein Interface Item mit Price() und Display(level int). Glass und Frame sind die Blätter und haben jeweils einen Namen und einen Preis. Window ist das Composite, enthält 1..* Items (also auch wieder Windows) und hat eine Add-Methode. Der Preis eines Windows ist die Summe der Preise seiner Items, Display gibt den Baum eingerückt aus.

In main bitte ein Fenster mit Rahmen und 2x2 Scheiben nachbauen, wobei die Scheibe unten rechts selbst wieder ein kleines Fenster mit eigenem Rahmen und 2x2 Glas ist. Dann Display aufrufen und den Gesamtpreis ausgeben.

Code einfach halten, keine externen Packages, mit go run . ausführbar.
