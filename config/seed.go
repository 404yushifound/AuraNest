package config

import (
	"AuraNest/models"
	"log"
)

// 🔍 Helper function to get Vibe ID by label
func getVibeID(label string) uint {
	var vibe models.Vibe
	DB.Where("label = ?", label).First(&vibe)
	return vibe.ID
}

// 🔍 Helper function to get TimeSlot ID by label
func getTimeSlotID(label string) uint {
	var timeSlot models.TimeSlot
	DB.Where("label = ?", label).First(&timeSlot)
	return timeSlot.ID
}

// 🌱 SeedData populates Vibes, TimeSlots, and Songs
func SeedData() {
	vibes := []models.Vibe{
		{Label: "Sunny Side Up", Emoji: "🌞"},
		{Label: "Workout", Emoji: "🏋️"},
		{Label: "Cooking", Emoji: "🍳"},
		{Label: "Cleaning", Emoji: "🧼"},
		{Label: "Driving", Emoji: "🚗"},
		{Label: "Rainy Day", Emoji: "🌧️"},
		{Label: "Focus Mode", Emoji: "🎯"},
		{Label: "Rage Release", Emoji: "🔥"},
	}

	timeSlots := []models.TimeSlot{
		{Label: "Early Morning", StartHour: 5, EndHour: 8, Emoji: "🌄"},
		{Label: "Morning", StartHour: 8, EndHour: 12, Emoji: "☀️"},
		{Label: "Afternoon", StartHour: 12, EndHour: 16, Emoji: "🌤️"},
		{Label: "Evening", StartHour: 16, EndHour: 19, Emoji: "🌆"},
		{Label: "Night", StartHour: 19, EndHour: 23, Emoji: "🌙"},
		{Label: "Late Night", StartHour: 23, EndHour: 5, Emoji: "🌌"},
	}

	for _, vibe := range vibes {
		if err := DB.FirstOrCreate(&vibe, models.Vibe{Label: vibe.Label}).Error; err != nil {
			log.Println("Error seeding vibe:", vibe.Label)
		}
	}

	for _, slot := range timeSlots {
		if err := DB.FirstOrCreate(&slot, models.TimeSlot{Label: slot.Label}).Error; err != nil {
			log.Println("Error seeding time slot:", slot.Label)
		}
	}

	songs := []models.Song{
	// 🌞 Sunny Side Up
	{Title: "Gayatri Mantra (Piano)", Artist: "Traditional", YouTubeLink: "https://youtu.be/dtfF_Q6o0gc?si=oeHGdcA39O-U7UgL", VibeID: getVibeID("Sunny Side Up"), TimeSlotID: getTimeSlotID("Early Morning")},
	{Title: "Serendipity", Artist: "Jimin", YouTubeLink: "https://youtu.be/BEMaH9Sm3lQ", VibeID: getVibeID("Sunny Side Up"), TimeSlotID: getTimeSlotID("Morning")},
	{Title: "Still With You", Artist: "Jungkook", YouTubeLink: "https://youtu.be/dtfF_Q6o0gc", VibeID: getVibeID("Sunny Side Up"), TimeSlotID: getTimeSlotID("Afternoon")},
	{Title: "Rain", Artist: "BTS", YouTubeLink: "https://youtu.be/0lapF4DQPKQ", VibeID: getVibeID("Sunny Side Up"), TimeSlotID: getTimeSlotID("Evening")},
	{Title: "One of the Girls (Electric Guitar Ver.)", Artist: "The Weeknd", YouTubeLink: "https://youtu.be/0Y3Dvb1KkBQ", VibeID: getVibeID("Sunny Side Up"), TimeSlotID: getTimeSlotID("Night")},
	{Title: "Watermelon Sugar", Artist: "Harry Styles", YouTubeLink: "https://youtu.be/E07s5ZYygMg", VibeID: getVibeID("Sunny Side Up"), TimeSlotID: getTimeSlotID("Late Night")},

	// 🏋️ Workout
	{Title: "Blood Sweat & Tears", Artist: "BTS", YouTubeLink: "https://youtu.be/hmE9f-TEutc", VibeID: getVibeID("Workout"), TimeSlotID: getTimeSlotID("Early Morning")},
	{Title: "BOOMBAYAH", Artist: "BLACKPINK", YouTubeLink: "https://youtu.be/bwmSjveL3Lc", VibeID: getVibeID("Workout"), TimeSlotID: getTimeSlotID("Morning")},
	{Title: "As If It’s Your Last", Artist: "BLACKPINK", YouTubeLink: "https://youtu.be/Amq-qlqbjYA", VibeID: getVibeID("Workout"), TimeSlotID: getTimeSlotID("Afternoon")},
	{Title: "Playing With Fire", Artist: "BLACKPINK", YouTubeLink: "https://youtu.be/9pdj4iJD08s", VibeID: getVibeID("Workout"), TimeSlotID: getTimeSlotID("Evening")},
	{Title: "End Game", Artist: "Taylor Swift ft. Ed Sheeran", YouTubeLink: "https://youtu.be/dfnCAmr569k", VibeID: getVibeID("Workout"), TimeSlotID: getTimeSlotID("Night")},
	{Title: "TakeDown", Artist: "Kpop Demon Hunters", YouTubeLink: "https://youtu.be/m_vw2kzV3M0", VibeID: getVibeID("Workout"), TimeSlotID: getTimeSlotID("Late Night")},

	// 🍳 Cooking
	{Title: "Your Idol", Artist: "K-pop Demon Hunters", YouTubeLink: "https://youtu.be/-V5uYuTjxlY", VibeID: getVibeID("Cooking"), TimeSlotID: getTimeSlotID("Early Morning")},
	{Title: "Psycho", Artist: "Red Velvet", YouTubeLink: "https://youtu.be/uR8Mrt1IpXg", VibeID: getVibeID("Cooking"), TimeSlotID: getTimeSlotID("Morning")},
	{Title: "Egotistic", Artist: "Mamamoo", YouTubeLink: "https://youtu.be/pHtxTSiPh5I", VibeID: getVibeID("Cooking"), TimeSlotID: getTimeSlotID("Afternoon")},
	{Title: "Icky", Artist: "KARD", YouTubeLink: "https://youtu.be/Zc3ZBq4DnKo", VibeID: getVibeID("Cooking"), TimeSlotID: getTimeSlotID("Evening")},
	{Title: "Starboy", Artist: "The Weeknd", YouTubeLink: "https://youtu.be/34Na4j8AVgA", VibeID: getVibeID("Cooking"), TimeSlotID: getTimeSlotID("Night")},
	{Title: "Wake Up Call", Artist: "Maroon 5", YouTubeLink: "https://youtu.be/KRaWnd3LJfs", VibeID: getVibeID("Cooking"), TimeSlotID: getTimeSlotID("Late Night")},

	// 🧼 Cleaning
	{Title: "Egotistic", Artist: "Mamamoo", YouTubeLink: "https://youtu.be/pHtxTSiPh5I", VibeID: getVibeID("Cleaning"), TimeSlotID: getTimeSlotID("Early Morning")},
	{Title: "Icky", Artist: "KARD", YouTubeLink: "https://youtu.be/Zc3ZBq4DnKo", VibeID: getVibeID("Cleaning"), TimeSlotID: getTimeSlotID("Morning")},
	{Title: "Psycho", Artist: "Red Velvet", YouTubeLink: "https://youtu.be/uR8Mrt1IpXg", VibeID: getVibeID("Cleaning"), TimeSlotID: getTimeSlotID("Afternoon")},
	{Title: "Your Idol", Artist: "K-pop Demon Hunters", YouTubeLink: "https://youtu.be/-V5uYuTjxlY", VibeID: getVibeID("Cleaning"), TimeSlotID: getTimeSlotID("Evening")},
	{Title: "As If It’s Your Last", Artist: "BLACKPINK", YouTubeLink: "https://youtu.be/Amq-qlqbjYA", VibeID: getVibeID("Cleaning"), TimeSlotID: getTimeSlotID("Night")},
	{Title: "End Game", Artist: "Taylor Swift ft. Ed Sheeran", YouTubeLink: "https://youtu.be/dfnCAmr569k", VibeID: getVibeID("Cleaning"), TimeSlotID: getTimeSlotID("Late Night")},

	// 🚗 Driving
	{Title: "Handlebars", Artist: "Jennie", YouTubeLink: "https://youtu.be/m4CJhL9Mfzc", VibeID: getVibeID("Driving"), TimeSlotID: getTimeSlotID("Early Morning")},
	{Title: "House of Cards", Artist: "BTS", YouTubeLink: "https://youtu.be/fikbwrWFy3g", VibeID: getVibeID("Driving"), TimeSlotID: getTimeSlotID("Morning")},
	{Title: "Wake Up Call", Artist: "Maroon 5", YouTubeLink: "https://youtu.be/KRaWnd3LJfs", VibeID: getVibeID("Driving"), TimeSlotID: getTimeSlotID("Afternoon")},
	{Title: "Kalank (Title Track)", Artist: "Arijit Singh", YouTubeLink: "https://youtu.be/m1DTB1sQo1U", VibeID: getVibeID("Driving"), TimeSlotID: getTimeSlotID("Evening")},
	{Title: "Jhol", Artist: "Seedhe Maut", YouTubeLink: "https://youtu.be/1KCE_WCRomY", VibeID: getVibeID("Driving"), TimeSlotID: getTimeSlotID("Night")},
	{Title: "House of Cards", Artist: "BTS", YouTubeLink: "https://youtu.be/fikbwrWFy3g", VibeID: getVibeID("Driving"), TimeSlotID: getTimeSlotID("Late Night")},

	// 🌧️ Rainy Day
	{Title: "Still With You", Artist: "Jungkook", YouTubeLink: "https://youtu.be/dtfF_Q6o0gc", VibeID: getVibeID("Rainy Day"), TimeSlotID: getTimeSlotID("Early Morning")},
	{Title: "Rain", Artist: "BTS", YouTubeLink: "https://youtu.be/0lapF4DQPKQ", VibeID: getVibeID("Rainy Day"), TimeSlotID: getTimeSlotID("Morning")},
	{Title: "Luka Chhupi", Artist: "Lata Mangeshkar", YouTubeLink: "https://youtu.be/gvDfXbWj4yo", VibeID: getVibeID("Rainy Day"), TimeSlotID: getTimeSlotID("Afternoon")},
	{Title: "House of Cards", Artist: "BTS", YouTubeLink: "https://youtu.be/fikbwrWFy3g", VibeID: getVibeID("Rainy Day"), TimeSlotID: getTimeSlotID("Evening")},
	{Title: "Serendipity", Artist: "Jimin", YouTubeLink: "https://youtu.be/BEMaH9Sm3lQ", VibeID: getVibeID("Rainy Day"), TimeSlotID: getTimeSlotID("Night")},
	{Title: "Rain", Artist: "BTS", YouTubeLink: "https://youtu.be/0lapF4DQPKQ", VibeID: getVibeID("Rainy Day"), TimeSlotID: getTimeSlotID("Late Night")},

	// 🎯 Focus Mode
	{Title: "One of the Girls (Electric Guitar Ver.)", Artist: "The Weeknd", YouTubeLink: "https://youtu.be/0Y3Dvb1KkBQ", VibeID: getVibeID("Focus Mode"), TimeSlotID: getTimeSlotID("Early Morning")},
	{Title: "My Heart Will Go On", Artist: "Celine Dion", YouTubeLink: "https://youtu.be/FHG2oizTlpY", VibeID: getVibeID("Focus Mode"), TimeSlotID: getTimeSlotID("Morning")},
	{Title: "When We Were Young", Artist: "Adele", YouTubeLink: "https://youtu.be/DDWKuo3gXMQ", VibeID: getVibeID("Focus Mode"), TimeSlotID: getTimeSlotID("Afternoon")},
	{Title: "Starboy", Artist: "The Weeknd", YouTubeLink: "https://youtu.be/34Na4j8AVgA", VibeID: getVibeID("Focus Mode"), TimeSlotID: getTimeSlotID("Evening")},
	{Title: "As It Was", Artist: "Harry Styles", YouTubeLink: "https://youtu.be/H5v3kku4y6Q", VibeID: getVibeID("Focus Mode"), TimeSlotID: getTimeSlotID("Night")},
	{Title: "Watermelon Sugar", Artist: "Harry Styles", YouTubeLink: "https://youtu.be/E07s5ZYygMg", VibeID: getVibeID("Focus Mode"), TimeSlotID: getTimeSlotID("Late Night")},

	// 🔥 Rage Release
	{Title: "TakeDown", Artist: "Kpop Demon Hunters", YouTubeLink: "https://youtu.be/m_vw2kzV3M0", VibeID: getVibeID("Rage Release"), TimeSlotID: getTimeSlotID("Early Morning")},
	{Title: "Black Widow", Artist: "Rita Ora", YouTubeLink: "https://youtu.be/u3u22OYqFGo", VibeID: getVibeID("Rage Release"), TimeSlotID: getTimeSlotID("Morning")},
	{Title: "Labour", Artist: "Paris Paloma", YouTubeLink: "https://youtu.be/8bHoR4Nqk4U", VibeID: getVibeID("Rage Release"), TimeSlotID: getTimeSlotID("Afternoon")},
	{Title: "Get Back", Artist: "Demi Lovato", YouTubeLink: "https://youtu.be/2qDoLhP65cQ", VibeID: getVibeID("Rage Release"), TimeSlotID: getTimeSlotID("Evening")},
	{Title: "Heart Attack", Artist: "Demi Lovato", YouTubeLink: "https://youtu.be/HyHDKwwJb68", VibeID: getVibeID("Rage Release"), TimeSlotID: getTimeSlotID("Night")},
	{Title: "Labour", Artist: "Paris Paloma", YouTubeLink: "https://youtu.be/8bHoR4Nqk4U", VibeID: getVibeID("Rage Release"), TimeSlotID: getTimeSlotID("Late Night")},
}

	for _, song := range songs {
		if err := DB.FirstOrCreate(&song, models.Song{Title: song.Title, VibeID: song.VibeID, TimeSlotID: song.TimeSlotID}).Error; err != nil {
			log.Println("Error seeding song:", song.Title)
		}
	}

	log.Println("✅ Seeded vibes, time slots, and songs!")
}

