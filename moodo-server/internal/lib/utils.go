package lib

var MoodToGenreMapping = map[string][]int{
	"uplifting":  {35, 18, 10751},    // Comedy, Drama, Family (feel-good and inspiring)
	"vibey":      {10402, 35, 10749}, // Music, Comedy, Romance (chill, aesthetic vibes)
	"nostalgic":  {10751, 16, 36},    // Family, Animation, History (throwbacks and childhood memories)
	"thrilling":  {53, 27, 9648},     // Thriller, Horror, Mystery (exciting, keeps you on edge)
	"wholesome":  {10751, 18, 35},    // Family, Drama, Comedy (heartwarming and feel-good content)
	"futuristic": {878, 12, 14},      // Science Fiction, Adventure, Fantasy (futuristic or surreal)
	"dark":       {27, 53, 9648},     // Horror, Thriller, Mystery (moody and intense)
	"cringe":     {35, 10770},        // Comedy, TV Movie (awkward or "so-bad-it's-good" content)
	"retro":      {36, 10752, 10402}, // History, War, Music (vintage, throwback vibes)
	"brainy":     {99, 18},           // Documentary, Drama (thought-provoking content)
	"rom-com":    {10749, 35},        // Romance, Comedy (light-hearted romantic comedy)
	"chaotic":    {28, 35, 53},       // Action, Comedy, Thriller (unpredictable and high-energy)
}
