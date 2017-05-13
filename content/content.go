package content

/**
 * struct for the comic page
 */

type Page struct {
	Id, PageNumber                     int
	Title, Path, ImgUrl, DisqusId, Act string
}
