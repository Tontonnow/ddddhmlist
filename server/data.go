package server

func (x *Data) Len() int {
	return len(x.VideoList)
}
func (x *Data) Less(i, j int) bool {
	if x.VideoList[i].Episode == x.VideoList[j].Episode {
		return x.VideoList[i].EpisodeId < x.VideoList[j].EpisodeId
	}
	return x.VideoList[i].Episode < x.VideoList[j].Episode
}
func (x *Data) Swap(i, j int) {
	x.VideoList[i], x.VideoList[j] = x.VideoList[j], x.VideoList[i]
}
