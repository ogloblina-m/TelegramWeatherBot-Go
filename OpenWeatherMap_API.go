type SearchResults struct {
    ready   bool
    Query   string
    Results []Result
}
type Result struct {
    Name, Description, URL string
}