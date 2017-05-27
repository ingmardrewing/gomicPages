package content

/**
 * struct for the comic page
 */

type Page struct {
	Id, PageNumber                                  int
	Title, Description, Path, ImgUrl, DisqusId, Act string
}

func EmptyPage() Page {
	return Page{0, 0, "", "", "", "", "", ""}
}
