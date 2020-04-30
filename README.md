# Jackbox Scanner

A little CLI tool written to randomly scan [Jackbox](https://jackbox.tv) Servers for active rooms.

Features:
* Filter only rooms which can be joined as players. (`-playable` flag)
* Filter only rooms with no password protection (default. Turn this off with `-passwords` flag)
* Concurrent requests for fast performance. Usually you can find a room within 10-20 seconds. (Modify concurrency level with `-workers` flag)
* Manually check specific 4-letter room codes by passing them as positional arguments.

## Installation

You'll need a `go` compiler: just `go get` it:

```
$ go get github.com/kklash/jackbox_scanner
$ jackbox_scanner  -h
Usage of ./jackbox_scanner:
  -first
      Terminate upon finding first valid room code.
  -passwords
      Show games which are password-protected.
  -playable
      Only count games open to joining as players.
  -verbose
      Enable verbose logging of each request.
  -workers int
      Number of parallel workers checking rooms. (default 5)
```

## Usage

To scan for any available and unlocked rooms, just run the program with no arguments:

```
$ jackbox_scanner
```

It will start printing the number of room codes it has checked thus far, and will dump the room info on any active ones it finds:

```
Room codes checked: 47
FOUND ROOM CODE: LQJE
{
  "apptag": "quiplash2",
  "audienceEnabled": true,
  "numAudience": 0,
  "error": "",
  "joinAs": "audience",
  "requiresPassword": false,
  "roomid": "LQJE",
  "server": "ecast.jackboxgames.com"
}
```

If you wish to scan specific room codes, pass them as positional arguments, like so:

```
$ jackbox_scanner abbb jjsz nquf
ABBB: EMPTY
NQUF: **** FOUND ROOM ****
JJSZ: EMPTY
```
