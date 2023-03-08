package tagcloud

// TagCloud aggregates statistics about used tags
type TagCloud struct {
	tags            map[string]*TagStat
	tagsPrioritized []TagStat
}

// TagStat represents statistics regarding single tag
type TagStat struct {
	Tag             string
	OccurrenceCount int
}

// New should create a valid TagCloud instance
func New() *TagCloud {
	return &TagCloud{tags: make(map[string]*TagStat), tagsPrioritized: make([]TagStat, 0)}
}

// AddTag should add a tag to the cloud if it wasn't present and increase tag occurrence count
// thread-safety is not needed
func (tagCloud *TagCloud) AddTag(tag string) {
	_, ok := tagCloud.tags[tag]
	if !ok {
		tagCloud.tags[tag] = &TagStat{Tag: tag, OccurrenceCount: 0}
	}
	tagCloud.tags[tag].OccurrenceCount += 1

	tagCloud.tagsPrioritized = delPrev(tagCloud.tagsPrioritized, *tagCloud.tags[tag])
	tagCloud.tagsPrioritized = add(tagCloud.tagsPrioritized, *tagCloud.tags[tag])
}

func binSearch(a []TagStat, x int) int {
	l, r := -1, len(a)
	for r-l != 1 {
		m := (l + r) / 2
		if a[m].OccurrenceCount >= x {
			l = m
		} else {
			r = m
		}
	}
	return r
}

func delPrev(a []TagStat, tagStat TagStat) []TagStat {
	prevOccurrenceCount := tagStat.OccurrenceCount - 1
	for i := binSearch(a, prevOccurrenceCount) - 1; i >= 0 && a[i].OccurrenceCount == prevOccurrenceCount; i-- {
		if a[i].Tag != tagStat.Tag {
			continue
		}
		for j := i; j < len(a)-1; j++ {
			a[j] = a[j+1]
		}
		return a[:len(a)-1]
	}
	return a
}

func add(a []TagStat, tagStat TagStat) []TagStat {
	id := binSearch(a, tagStat.OccurrenceCount)
	a = append(a, tagStat)
	for i := len(a) - 1; i >= id && i > 0; i-- {
		a[i] = a[i-1]
	}
	a[id] = tagStat
	return a
}

// TopN should return top N most frequent tags ordered in descending order by occurrence count
// if there are multiple tags with the same occurrence count then the order is defined by implementation
// if n is greater that TagCloud size then all elements should be returned
// thread-safety is not needed
// there are no restrictions on time complexity
func (tagCloud *TagCloud) TopN(n int) []TagStat {
	return tagCloud.tagsPrioritized[0:min(len(tagCloud.tags), n)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
