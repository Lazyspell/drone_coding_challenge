# drone_coding_challenge

PIZZA DRONE CHALLENGE Here is the challenge
Create a backend service to calculate routes for a small pizza chain that recently purchased one drone to deliver their pizzas. Drone range is 25 miles and must come back to the store for a new battery before it runs out of energy.
The drone will reach out to your backend service to receive its next destination in GPS coordinates.
The drone will only request the next destination:
● once it is ready to leave the pizza store
● once it has reached a previously commanded destination and made a delivery.
You may use any programming language you prefer. Your code should be syntactically correct, but it is not required to be functional.
Expected functionality:

1. Given a dynamically sized CSV file with order time and delivery address, calculate the optimal path (you are free to define what “optimal path” means to you).

example inbound CSV: <use your own size/data>
order_time | address -------------------------------------------
8/11/2022 12:16 | 1016 Huntingdon Dr, San Jose, CA 95129 8/11/2022 12:22 | 1580 Benton St, Sunnyvale, CA 94087 8/11/2022 12:41 | 10239 E Estates Dr, Cupertino, CA 95014 8/11/2022 12:53 | 2112 Rockhurst Ct, Santa Clara, CA 95051 ...

2.  Provide a HTTP GET endpoint that will send the drone its next destination in GPS coordinates
    upon request. Remember, the drone will only reach out to your backend service
    ● when it’s ready to leave the pizza store.
    ● to get the next destination once it has completed its previous destination or

3.  The drone will reach out to your backend service with secure credentials. Your backend service
    must authenticate the drone and provide it with a token that it can use to authorize all subsequent requests to your REST API. This token must be validated before providing the drone any data. The token must expire 1 hour after being issued and the drone must request to refresh the token or obtain a new one.
    Summary
    You are welcome to implement any additional functionality you think is important for this scenario. Time box, and don’t chase any deep rabbit holes. Invest as if it were the first ProofOfConcept version.
    And most important... HAVE FUN and ENJOY WRITING/SOLVING IT !!
