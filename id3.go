package acousticid
import "github.com/mikkyang/id3-go"

type Id3 struct {
	Artist string
	Title string
	Album string
}

func (i *Id3) set(file string) {
	mp3File, err := id3.Open(file)
	defer mp3File.Close()
	if err != nil {
		panic(err)
	}
	mp3File.SetArtist(i.Artist)
	mp3File.SetTitle(i.Title)
	mp3File.SetAlbum(i.Album)
}