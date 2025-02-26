package db

import (
	"context"
	"fmt"
	"log"
	"maps"
	"math/rand"
	"slices"

	"github.com/willystw/golang-simple-social/internal/store"
)

var usernames = []string{
	"EchoVibe", "FrostByte", "SolarWave", "QuantumBlaze", "MysticWhisper", "VortexNova", "NebulaShifter", "LunarSpecter", "ArcaneFusion", "CosmicStreak", "PixelDrift", "BlazeQuest", "NovaWanderer", "ZenEcho", "FrostVanguard", "ThunderZion", "CyberNebula", "StellarWhirlwind", "ShadowSpark", "AstralSurge", "EchoFlare", "NebulaWraith", "CelestialEcho", "GalacticForge", "LunarVortex", "ZenithBlaze", "MidnightNova", "StellarPulse", "SolarShadow", "PhantomGlimmer", "TitanStorm", "DreamFusion", "CelestialFlare", "QuantumGlitch", "VortexDynamo", "NebulaDrift",
}

var titles = []string{"Mastering Go Arrays", "Go Slices vs Arrays", "Understanding Go Pointers", "Effective Go Concurrency", "Go Maps Explained", "Go Structs: A Deep Dive", "Error Handling in Go", "Go Best Practices for Beginners", "Introduction to Go Channels", "Go Functions and Closures", "Exploring Go's Garbage Collector", "Go and Web Development", "Optimizing Go Code", "Testing in Go: A Guide", "Go Performance Tuning"}

var content string = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur fringilla volutpat sapien et pellentesque. Mauris egestas lorem sit amet mi fermentum sollicitudin. In commodo euismod justo eu mollis. Maecenas dignissim lorem ut sapien laoreet dapibus vel nec magna. Mauris aliquam justo ullamcorper ipsum aliquet dignissim. Mauris ac commodo lorem. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus lacus ligula, efficitur ac erat vitae, pretium ultricies purus. Fusce sagittis dui elementum tellus pellentesque, eu auctor sem pharetra. Phasellus massa magna, tincidunt eu condimentum eget, consequat et urna. Mauris at massa vitae nibh iaculis elementum. Pellentesque eu sagittis leo."

var tags = []string{
	"tech", "art", "fashion", "food", "music", "travel", "fitness", "nature", "photography", "lifestyle", "gaming", "sports", "health", "adventure", "design", "movies", "books", "writing", "DIY", "business",
}

var usercomments = []string{
	"Great work!", "Love this!", "So inspiring!", "Amazing content!", "Keep it up!", "Nice job!", "Very creative!", "Well done!", "So cool!", "This is awesome!", "Nice one!", "I love it!", "Looking great!", "So impressive!", "This is fantastic!", "Great idea!", "So unique!", "Well executed!", "Incredible!", "Totally awesome!",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, nil, user); err != nil {
			log.Println("error creating user:", err)
			return

		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)

	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("error creating post:", err)
			return
		}
	}

	log.Println("seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		username := usernames[i%len(usernames)] + fmt.Sprintf("%d", i)

		users[i] = &store.User{
			Username: username,
			Email:    username + "@example.com",
		}
		users[i].SetPassword("123123")

	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		tags := map[string]int{
			tags[rand.Intn(len(tags))]: 1,
			tags[rand.Intn(len(tags))]: 1,
		}
		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: content,
			Tags:    slices.Collect(maps.Keys(tags)),
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	comments := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		post := posts[rand.Intn(len(posts))]
		usercomment := usercomments[rand.Intn(len(usercomments))]

		comments[i] = &store.Comment{
			PostID:  post.ID,
			UserID:  user.ID,
			Content: usercomment,
		}
	}

	return comments
}
