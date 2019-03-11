streamstatus
==============================

Streamstatus is a small cli tool to check the status of your favorite twitch streamers.

## Installation

- If you are running (archlinux)[https://www.archlinux.org/] use the (AUR-package)[https://aur.archlinux.org/packages/streamstatus/]

- Clone this repo and build with `go build`

## Usage

Add `.streamstatus` to your home directory. Then add streamer and category.

Example:
```
~ cat .streamstatus
demuslim sc2
basetradetv sc2 
TaKeTV sc2 
Lamboking sc2 
Adam13531 code
```

## Run

Simply run with `streamstatus` to see all streamers.

```
~ streamstatus
Your streams are currently :
[ 0] demuslim         sc2    online
[ 1] basetradetv      sc2    offline
[ 2] TaKeTV           sc2    online
[ 3] Lamboking        sc2    offline
[ 4] Adam13531        code   online
```

Run `streamstatus` with a category eg. `streamstatus code`to see all streamers in this category.

```
~ streamstatus
Your streams are currently :
[ 0] Adam13531        code   online
```

