
# GROUPIE-TRACKER-geolocalization
* `ayessenb` 
* `roshakba` 



#### Groupie-tracker-geolocation consists on making your site :
We followed the same principles as the first subjects.
    * Groupie Tracker Geolocalization consists on mapping the different concerts locations of a certain artist/band given by the Client.
    * We used a process of converting addresses (ex: Germany Mainz) into geographic coordinates (ex: 49,59380 8,15052) which you must use to place markers for the concerts locations of a certain artist/band on a map.
    * We used the Map API that we found more appropriate.


#### Description:

* We used text/template pakage to recieve and send GET requests.
* GET /: Sends HTML response, the main page.
* and using the JSON file which contains the data that need to be displayed to pages
* Linked external css file to make good design 
* For realising the dynamnic changes of the bands and list of suggestions,
 JS was used for the implementation on the front ent.



## Usage/Examples
Cloning storage to your host
```CMD/Terminal 
git clone git@github.com:compygirl/groupie-tracker-geolocalization.git
```
Go to the downloaded repository:

```CMD/Terminal 
cd groupie-tracker-geolocalization
```
Run a Server:
```CMD/Terminal 
go run main.go
```

Follow the link on the terminal:
```CMD/Terminal 
Starting server got testing... http://127.0.0.1:8080 
```
you can play with the page by choosing the music band's image and so on.
As well as searching the strings with filters
Zoom in the map



## HTTP status codes
* OK (200), if everything went without errors.
* Not Found, if the wrong URL was provided.
* Bad Request, for incorrect requests. for exaple, the id is out of range.
* Internal Server Error, for unhandled errors.
* If string is not a substring of any items which were mentioned in the first paragraph, 
then the page will be displaying "NOT FOUND"



## Feedback

If you liked our project, we would be grateful if you could add `Star` to the repository.

Alem Students @ayyssenb @roshakba