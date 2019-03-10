streamstatus
==============================

Streamstatus is a small cli tool to check the status of your favorite twitch streamers.

## Build

Build with `go build -o streamstatus`

## Usage

Add a file `.streamstatus` in your home directory. Add streamer and add category.

Example:
```
~ cat .streamstatus
demuslim sc2
TaKeTV sc2
sentdex python
jakenbakelive travel
```

## Output
```
~ streamstatus
Your streams are currently :
[ 0] demuslim         sc2    online
[ 1] basetradetv      sc2    offline
[ 2] TaKeTV           sc2    online
[ 3] Lamboking        sc2    offline
```

