# server
The API KEY for onestepGps should be put into a file in server/keys.ini

To run use

```
go run server.go
```

The backend should be available at http://localhost:8080

# client
google map key should go into client/mapkey.ts as and exported value

```
client yarn dev
```

```
eg export default 'GOOGLE_API_KEY'
```
The site should then run at http://localhost:3000


#Interesting things

In the popup that appears when clicking on a map marker, I created a reusable component that can handle both notes and phone numbers. Rules can be passed to the component to check validity.


The navigation drawers are responsive and move to the bottom of the page at less than 800px.
A better check would be to use the mobile property of useDisplay; however, this page would be difficult to load on a mobile device.


While thinking about why a user would want to make a marker invisible on the map, I guessed that they would want to recenter around those markers. When the user marks a marker as invisible, I provided a button to 
recenter the map around the markers.


I also assumed that positions with latitude and longitude 0 were offline.


Thanks for looking!
