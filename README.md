## Golf Logger – Desktop Klient/Server Portfolio Projekt

### Oversigt

Golf Logger er en desktop-applikation til at administrere din golftaske og holde styr på personlige slaglængder.

Projektet er bygget som en praktisk demonstration af moderne client-server arkitektur.

---

### Funktioner

**Brugerautentificering:**

* Sikker JWT-baseret registrering og login

**Taskeadministration:**

* Se, tilføj eller fjern køller fra din personlige taske
* Indstil og rediger afstande for hver kølle med en brugervenlig GUI

**Fuld Client-Server Adskillelse:**

* Go backend (REST API med SQLite-lagring)
* Go Fyne desktop frontend

**Rent UI/UX:**

* Responsiv, cross-platform desktop-app
* Adgangskoder er altid skjulte; fejl- og succesmeddelelser håndteres elegant

**Fremtidssikret:**

* “Kommer snart” placeholder til Golf Range Map-visualisering (viser, at projektet er klar til videreudvikling)

---

### Teknologier & Frameworks

**Sprog:**

* Go: bruges både til backend (API-server) og frontend (desktop-GUI)

**Backend:**

* Egen REST API (Go net/http)
* SQLite til bruger- og kølledata
* Adgangskoder hashes med Go’s standardbiblioteker

**Frontend:**

* Fyne – cross-platform Go desktop GUI toolkit

**API-integration:**

* Strukturen er designet til nem tilføjelse af eksterne Golfbane-API'er

---

### Sådan Kører Du Projektet

1. Clone repository
2. Installer Go (version 1.18 eller højere)
3. Kør backend API’en:

   ```bash
   go run main.go  
   ```
4. Kør GUI-applikationen:

   ```bash
   cd gui  
   go run .  
   ```
   
5. Opret en konto og log ind via desktop-klientens brugerflade!

---

### Hvorfor Dette Projekt?

Kodebasen demonstrerer:

* Design og implementering af ægte klient-server-funktionalitet
* Sikker og professionel GUI-samtidighed i Go
* Ren opdeling af logik, vedligeholdelsesvenlig struktur og produktionsklare udvidelsespunkter

---

### Ambitioner / Fremtidige Udvidelser

* Tilføj Golf Range Map-visualisering med baner og skud-overlays
* Integrér live-bane-API’er for detaljerede bane- og huldata
* Tilføj skud-logging og faner for performance-analyse/statistik

---
