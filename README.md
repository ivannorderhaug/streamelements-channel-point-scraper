# Streamelements Channel Points Scraper
The purpose of this application is to retrieve all the usernames and points from the Streamelements API and save it in a simple readable and reuseable format. This application is custom made for the channel owner of https://www.twitch.tv/roshtein and will only work with that channel. 

## Installation
Simply install it from the releases tab or by clicking on this link: [download](https://github.com/ivannorderhaug/streamelements-channel-point-scraper/releases/tag/Release)

For those of you who want to build it yourself, download Go version 19.3 from [https://go.dev/doc/install](https://go.dev/doc/install).

Once Go is installed, open a terminal and clone the repository with:
```bash
git clone https://github.com/ivannorderhaug/streamelements-channel-point-scraper.git
```
Then head into project folder with:
```bash
cd streamelements-channel-point-scraper
```
Finally, build it with:
```bash
go build .\cmd\main.go
```
## Usage
Run the executeable and it will do it's magic. 

The data retrieved will be saved to a file called 'users.json' and will be located in the same directory as the executeable.

An example of how the data will look is as follows:

```json
{
 "users": [
  {
   "username": "roshtein",
   "points": 1627489
  },
  {
   "username": "intelmage",
   "points": 1613390
  },
  {
   "username": "fabeyy",
   "points": 1590041
  },
  {
   "username": "gomamisok",
   "points": 1583725
  },
  {
   "username": "mrsatchels",
   "points": 1582991
  },
  {
   "username": "commanderroot",
   "points": 1578715
  },
  {
   "username": "l332k",
   "points": 1576136
  },
  {
   "username": "jasonkyan",
   "points": 1569061
  }
  ]
}

```


## License

[MIT](https://choosealicense.com/licenses/mit/)
