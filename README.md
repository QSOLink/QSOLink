# qsolink
JSON/REST based Ham Radio Logging Platform

This project and repo will serve as the server and database for the operator to log contacts.

QSO Link (*pronounced Q S O Link*), Named after the [Ham Radio Q Codes](https://hamradioprep.com/ham-radio-q-codes/), is intended to be a cross platform ham radio contact logging package.  It comes from a series of frustrations over the many years I've been a ham, where replacing a computer resulted in lost log databases.  Or needing to come up with solutions to sync logs when using a laptop and a desktop, and possibly even a phone depending on the context of how I was operating.  My goal is to use a JSON/REST API server, that can be accessed through a number of different front ends.

Future goals will include writing a Terminal User Interface front end, inspired by [Fabian Kurz, DJ5CW](https://fkurz.net/ham/) logging package [yfklog](https://fkurz.net/ham/yfklog.html).  In addition to a GUI interface, inspired by [N3FPJ's Amateur Contact Log](https://www.n3fjp.com/aclog.html).  Ideally both of these clients will be cross platform, working on Linux, FreeBSD, OSX, and Windows.  Finally I think having a web interface that interacts with the API would be great as well.

I will be working on this project as part of #100DaysOfCode officially starting on September 18, 2023.

### Dependencies Installation
ensure you have docker and docker compose installed and working.

### Running QSLlink for development
```
$ docker compose up
```

### Future Integrations
I think integrating with a callsign lookup database like [HamDB.org](https://hamdb.org/api) or [callbook.info](https://callook.info/api_reference.php) would be fantastic.
I should also ensure that QSOLink can post our logbook to the [QRZ Logbook](https://www.qrz.com/docs/logbook30/api) which has an API as well.


