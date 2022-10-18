# exileum - A BitTorrent CLI written in Go
Named after the lead developer of [TorrentPier](https://en.wikipedia.org/wiki/TorrentPier), the open-source BitTorrent tracker that 
powers [the best](https://rutracker.org/forum/index.php) torrenting community on Earth (RuTracker :))

### roadmap

1. Snowball (core)

| BEP # | Title | 
| --- | ----------- |
| [3](https://www.bittorrent.org/beps/bep_0003.html) | BitTorrent protocol	 |
| [5](https://www.bittorrent.org/beps/bep_0005.html) | DHT protocol |

2. Boxer (features I)

| BEP # | Title | 
| --- | ----------- |
| [53](https://www.bittorrent.org/beps/bep_0053.html) | Magnet link file selection	 |
| [11](https://www.bittorrent.org/beps/bep_0011.html) | Peer exchange |
|[16](https://www.bittorrent.org/beps/bep_0016.html) |Super-seeding	|

3. Napoleon (features II)
tbd:
- GUI?

## Contributing
Welcomed! 


## Notes

1. connect to tracker
    - GET announce key in metainfo
    - RESPONSE: a list of peers 



functions:
- main
- function that makes initial annoucne request response