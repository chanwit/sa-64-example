package entity

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Video{},
		&User{},
		&Playlist{},
		&Resolution{},
		&WatchVideo{},
	)

	db = database

	db.Model(&User{}).Create(&User{
		Name:  "Chanwit",
		Email: "chanwit@gmail.com",
	})
	db.Model(&User{}).Create(&User{
		Name:  "Name",
		Email: "name@example.com",
	})

	var chanwit User
	var name User
	db.Raw("SELECT * FROM users WHERE email = ?", "chanwit@gmail.com").Scan(&chanwit)
	db.Raw("SELECT * FROM users WHERE email = ?", "name@example.com").Scan(&name)

	// --- Video Data
	saLecture4 := Video{
		Name:  "SA Lecture 4",
		Url:   "https://youtu.be/123",
		Owner: chanwit,
	}
	db.Model(&Video{}).Create(&saLecture4)

	howTo := Video{
		Name:  "How to ...",
		Url:   "https://youtu.be/456",
		Owner: chanwit,
	}
	db.Model(&Video{}).Create(&howTo)

	helloWorld := Video{
		Name:  "Hello World with C",
		Url:   "https://youtu.be/789",
		Owner: name,
	}
	db.Model(&Video{}).Create(&helloWorld)

	// Resolution Data
	res360p := Resolution{
		Value: "360p",
	}
	db.Model(&Resolution{}).Create(&res360p)

	res480p := Resolution{
		Value: "480p",
	}
	db.Model(&Resolution{}).Create(&res480p)

	res720p := Resolution{
		Value: "720p",
	}
	db.Model(&Resolution{}).Create(&res720p)

	watchedPlayListOfChanwit := Playlist{
		Title: "Watched",
		Owner: chanwit,
	}
	db.Model(&Playlist{}).Create(&watchedPlayListOfChanwit)

	musicPlayListOfChanwit := Playlist{
		Title: "Music",
		Owner: chanwit,
	}
	db.Model(&Playlist{}).Create(&musicPlayListOfChanwit)

	watchedPlayListOfName := Playlist{
		Title: "Watched",
		Owner: name,
	}
	db.Model(&Playlist{}).Create(&watchedPlayListOfName)

	// watch 1
	db.Model(&WatchVideo{}).Create(&WatchVideo{
		Playlist:    watchedPlayListOfChanwit,
		Video:       saLecture4,
		WatchedTime: time.Now(),
		Resolution:  res720p,
	})
	// watch 2
	db.Model(&WatchVideo{}).Create(&WatchVideo{
		Playlist:    watchedPlayListOfName,
		Video:       helloWorld,
		WatchedTime: time.Now(),
		Resolution:  res480p,
	})
	// watch 3
	db.Model(&WatchVideo{}).Create(&WatchVideo{
		Playlist:    watchedPlayListOfChanwit,
		Video:       helloWorld,
		WatchedTime: time.Now(),
		Resolution:  res720p,
	})

	//
	// === Query
	//

	var target User
	db.Model(&User{}).Find(&target, db.Where("email = ?", "chanwit@gmail.com"))

	var watchedPlaylist Playlist
	db.Model(&Playlist{}).Find(&watchedPlaylist, db.Where("title = ? and owner_id = ?", "Watched", target.ID))

	var watchedList []*WatchVideo
	db.Model(&WatchVideo{}).
		Joins("Playlist").
		Joins("Resolution").
		Joins("Video").
		Find(&watchedList, db.Where("playlist_id = ?", watchedPlaylist.ID))

	for _, wl := range watchedList {
		fmt.Printf("Watch Video: %v\n", wl.ID)
		fmt.Printf("%v\n", wl.Playlist.Title)
		fmt.Printf("%v\n", wl.Resolution.Value)
		fmt.Printf("%v\n", wl.Video.Name)
		fmt.Println("====")
	}
}
