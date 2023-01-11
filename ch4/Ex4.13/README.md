# Postic 📸
Postic is a tool made in **GO** to download poster movie images from the [OMDb API](https://omdbapi.com/).  

# Setup
First [get your OMDb API KEY](https://omdbapi.com/apikey.aspx), then create a `.env` file in the root directory and fill it with the following content:  
```
omdb_API_Key=<YOUR API KEY>
```

# How to use it 🤔?
In the command line type the binary and the movie name:
```
$ go run . Casablanca
Movie
Title: Casablanca
Poster URL:https://m.media-amazon.com/images/M/MV5BY2IzZGY2YmEtYzljNS00NTM5LTgwMzUtMzM1NjQ4NGI0OTk0XkEyXkFqcGdeQXVyNDYyMDk5MTU@._V1_SX300.jpg
Poster downloaded!
Filename created: "Casablanca.jpg"
```

## Thanks & Resources 😄  
Thanks to me (XD) and internet for starting this journey with this great programming language. I coded this little challenge from the book [The Go Programming Language](https://www.gopl.io/), hope you liked the little project.
