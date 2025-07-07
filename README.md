# AuraNest
🎼 AuraNest: Purposeful Music Organization for Serious Listeners | AuraNest is a structured music recommendation backend built for listeners who don’t scroll aimlessly. Every song is placed at the intersection of a specific mood (vibe) and time of day — nothing random, nothing accidental.
Pre-defined list of available moods (Vibes) and Time Slots of the day.
---------------------
##🎧 Vibes Table##
| ID | Label         | Emoji |
| -- | ------------- | ----- |
| 1  | Sunny Side Up | 🌞   |
| 2  | Workout       | 🏋️   |
| 3  | Cooking       | 🍳   |
| 4  | Cleaning      | 🧼   |
| 5  | Driving       | 🚗   |
| 6  | Rainy Day     | 🌧️   |
| 7  | Focus Mode    | 🎯   |
| 8  | Rage Release  | 🔥   |

--------------------------
##⏰ Time Slots Table

| ID | Label         | Start Hour | End Hour | Emoji |
| -- | ------------- | ---------- | -------- | ----- |
| 1  | Early Morning | 5          | 8        | 🌄    |
| 2  | Morning       | 8          | 12       | ☀️    |
| 3  | Afternoon     | 12         | 16       | 🌤️   |
| 4  | Evening       | 16         | 19       | 🌆    |
| 5  | Night         | 19         | 23       | 🌙    |
| 6  | Late Night    | 23         | 5        | 🌌    |
--------------------------

## 🧪 How to Test the API Locally

```bash
# Step 1: Run the Go server
go run main.go

# Step 2: Use these endpoints with Postman or browser:

# Get all vibes
GET http://localhost:8080/vibes

# Get all time slots
GET http://localhost:8080/timeslots

# Get one song by vibe and time
GET http://localhost:8080/songs?vibe=Workout&time=Morning
GET http://localhost:8080/songs?vibe=Focus Mode&time=Afternoon
e.g. GET http://localhost:8080/songs?vibe=Sunny Side Up&time=Early Morning
```

