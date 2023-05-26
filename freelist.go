package main

// metaPage is the maximum pgnum that is used by the db for its own purposes. For now, only page 0 is used as the
// header page. It means all other page numbers can be used.
const metaPage = 0

type freelist struct {
	maxPage       pgnum   // Holds the maximum page allocated. maxPage*PageSize = fileSize
	releasedPages []pgnum // Pages that were previouslly allocated but are now free
}

func newFreelist() *freelist {
	return &freelist{
		maxPage:       metaPage,
		releasedPages: []pgnum{},
	}
}

func (fr *freelist) getNextPage() pgnum {
	// If possible, fetch pages first from the released pages.
	// Else, increase the maximum page
	if len(fr.releasedPages) != 0 {
		pageID := fr.releasedPages[len(fr.releasedPages)-1]
		fr.releasedPages = fr.releasedPages[:len(fr.releasedPages)-1]
		return pageID
	}
	fr.maxPage += 1
	return fr.maxPage
}

func (fr *freelist) releasePage(page pgnum) {
	fr.releasedPages = append(fr.releasedPages, page)
}
