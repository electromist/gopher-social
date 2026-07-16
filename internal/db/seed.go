package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/electromist/gopher-social.git/internal/store"
)

var usernames = []string{
	"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi",
	"ivan", "judy", "mallory", "niaj", "oscar", "peggy", "quinn", "rudy",
	"sybil", "trent", "ursula", "victor", "wendy", "xander", "yvonne", "zach",
}

var titles = []string{
	"The Silent Voyager", "Echoes of Tomorrow", "Master of Shadows",
	"The Quantum Blueprint", "Whispers of the Cosmos", "The Last Cipher",
	"Architect of Dreams", "The Fractal Path", "Stormbound Horizons",
}

var contents = []string{
	"Exploring the hidden patterns in everyday data reveals how small decisions shape larger outcomes.",
	"A journey through the cosmos of imagination allows us to explore infinite possibilities.",
	"Unveiling the secrets behind great innovations is like peeling back the layers of history.",
	"When technology meets creativity, magic happens in ways that transform the world.",
	"The art of simplifying complex problems lies in breaking them into smaller, manageable steps.",
}

var tags = []string{"go", "backend", "web", "database", "api"}

var comments = []string{
	"Great read!", "Very informative.", "I learned a lot from this.",
	"Could you explain more about the second point?", "Awesome perspective.",
}

// Seed function accept karega Store interface aur DB connection
func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	// 100 Users
	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)
	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			log.Println("Error creating user:", err)
			_ = tx.Rollback()
			return
		}
	}
	tx.Commit()

	// 200 Posts
	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	// 500 Comments
	comms := generateComments(500, users, posts)
	for _, comment := range comms {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		// Example: alice231@example.com (jaisa transcript me bataya tha)
		user := &store.User{
			Username: usernames[rand.Intn(len(usernames))] + fmt.Sprintf("%d", i),
			Email:    usernames[rand.Intn(len(usernames))] + fmt.Sprintf("%d@example.com", i),
			Role: store.Role{
				Name: "user",
			},
		}
		user.Password.Set("123123")
		users[i] = user
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))], // Always 2 tags
			},
		}
	}
	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		post := posts[rand.Intn(len(posts))]
		cms[i] = &store.Comment{
			UserID:  user.ID,
			PostID:  post.ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
